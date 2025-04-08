/*
 * Copyright (C) 2025 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package translator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	cql "github.com/cloudspannerecosystem/spanner-cassandra-schema-tool/translator/cqlparser"
)

func getCqlParseTreeRoot(cqlQuery string) (root cql.IRootContext, err error) {
	// The CqlParserErrorListener fires a panic for SyntaxError.
	defer func() {
		if r := recover(); r != nil {
			root = nil
			err = fmt.Errorf("%v", r)
		}
	}()

	lexer := cql.NewCqlLexer(antlr.NewInputStream(cqlQuery))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := cql.NewCqlParser(stream)
	if parser == nil {
		return nil, fmt.Errorf("error while creating parser object")
	}

	parser.RemoveErrorListeners()
	errorListener := new(CqlParserErrorListener)
	parser.AddErrorListener(errorListener)

	return parser.Root(), nil
}

func mapCqlNativeTypeToSpannerType(cqlNativeType cql.ICqlNativeTypeContext) (string, error) {
	if (cqlNativeType.K_BIGINT() != nil) ||
		(cqlNativeType.K_COUNTER() != nil) ||
		(cqlNativeType.K_INT() != nil) ||
		(cqlNativeType.K_SMALLINT() != nil) ||
		(cqlNativeType.K_TIME() != nil) ||
		(cqlNativeType.K_TINYINT() != nil) {
		return "INT64", nil
	}
	if cqlNativeType.K_BOOLEAN() != nil {
		return "BOOL", nil
	}
	if cqlNativeType.K_TIMESTAMP() != nil {
		return "TIMESTAMP", nil
	}
	if cqlNativeType.K_DATE() != nil {
		return "DATE", nil
	}
	if cqlNativeType.K_FLOAT() != nil {
		return "FLOAT32", nil
	}
	if cqlNativeType.K_DOUBLE() != nil {
		return "FLOAT64", nil
	}
	if (cqlNativeType.K_DECIMAL() != nil) ||
		(cqlNativeType.K_VARINT() != nil) {
		return "NUMERIC", nil
	}
	if cqlNativeType.K_BLOB() != nil {
		return "BYTES(MAX)", nil
	}
	if (cqlNativeType.K_ASCII() != nil) ||
		(cqlNativeType.K_INET() != nil) ||
		(cqlNativeType.K_TEXT() != nil) ||
		(cqlNativeType.K_VARCHAR() != nil) ||
		(cqlNativeType.K_TIMEUUID() != nil) ||
		(cqlNativeType.K_UUID() != nil) {
		return "STRING(MAX)", nil
	}
	return "", fmt.Errorf("Unknown type %s.", cqlNativeType.GetText())
}

func mapCqlCollectionTypeToSpannerType(cqlCollectionType cql.ICqlCollectionTypeContext) (string, error) {
	if cqlCollectionType.K_MAP() != nil {
		if len(cqlCollectionType.AllCqlNativeType()) != 2 {
			return "", fmt.Errorf("Unknown type %s.", cqlCollectionType.GetText())
		}
		for i := 0; i < 2; i++ {
			if cqlCollectionType.CqlNativeType(i).K_COUNTER() != nil {
				return "", fmt.Errorf("Counters are not allowed inside collections: %s", cqlCollectionType.GetText())
			}
			if _, err := mapCqlNativeTypeToSpannerType(cqlCollectionType.CqlNativeType(i)); err != nil {
				return "", fmt.Errorf("Unknown type %s.", cqlCollectionType.GetText())
			}
		}
		return "JSON", nil
	}
	if (cqlCollectionType.K_SET() != nil) || cqlCollectionType.K_LIST() != nil {
		if len(cqlCollectionType.AllCqlNativeType()) != 1 {
			return "", fmt.Errorf("Unknown type %s.", cqlCollectionType.GetText())
		}
		if cqlCollectionType.CqlNativeType(0).K_COUNTER() != nil {
			return "", fmt.Errorf("Counters are not allowed inside collections: %s", cqlCollectionType.GetText())
		}
		innerSpannerType, err := mapCqlNativeTypeToSpannerType(cqlCollectionType.CqlNativeType(0))
		if err != nil {
			return "", fmt.Errorf("Unknown type %s.", cqlCollectionType.GetText())
		}
		return fmt.Sprintf("ARRAY<%s>", innerSpannerType), nil
	}
	return "", fmt.Errorf("Unknown type %s.", cqlCollectionType.GetText())
}

func mapCqlTypeToSpannerType(cqlType cql.ICqlTypeContext) (string, error) {
	if cqlType.CqlNativeType() != nil {
		spannerType, err := mapCqlNativeTypeToSpannerType(cqlType.CqlNativeType())
		if err != nil {
			return "", err
		}
		return spannerType, nil
	}
	if cqlType.CqlCollectionType() != nil {
		spannerType, err := mapCqlCollectionTypeToSpannerType(cqlType.CqlCollectionType())
		if err != nil {
			return "", err
		}
		return spannerType, nil
	}
	return "", fmt.Errorf("Unknown type %s.", cqlType.GetText())
}

func validatePrimaryKeyType(cqlType cql.ICqlTypeContext, columnName string) error {
	if cqlType.CqlCollectionType() != nil {
		return fmt.Errorf("Invalid non-frozen collection type %s for PRIMARY KEY column '%s'", cqlType.CqlCollectionType().GetText(), columnName)
	}
	if cqlType.CqlNativeType() != nil && cqlType.CqlNativeType().K_COUNTER() != nil {
		return fmt.Errorf("counter type is not supported for PRIMARY KEY column '%s'", columnName)
	}
	return nil
}

func extractPrimaryKeyColumns(primaryKeyClause cql.IPrimaryKeyClauseContext) []string {
	primaryKey := primaryKeyClause.PrimaryKey()
	var primaryKeyColumns []string

	// Collects all columns in the partition key.
	for _, columnName := range primaryKey.PartitionKey().AllColumnName() {
		primaryKeyColumns = append(primaryKeyColumns, columnName.GetText())
	}

	// Collects all columns in the clustering key.
	if primaryKey.ClusteringColumns() != nil {
		for _, columnName := range primaryKey.ClusteringColumns().AllColumnName() {
			primaryKeyColumns = append(primaryKeyColumns, columnName.GetText())
		}
	}

	return primaryKeyColumns
}

func isInTableLevelPrimaryKey(tableLevelPrimaryKey []string, column string) bool {
	column = strings.ToLower(column)
	for _, pkColumn := range tableLevelPrimaryKey {
		// Case-insensitive comparison.
		if strings.ToLower(pkColumn) == column {
			return true
		}
	}
	return false
}

// ToSpannerStmt attempts to translate a CQL statement into a equivalent Spanner statement.
//
// Input:
//
//	cqlStmt      string: The CQL statement.
//	databaseName string: The expected database (keyspace) name for validation.
//
// Output:
//
//	spannerStmt string: The translated SQL (empty if ignored).
//	ignored	 	bool:   True if the statement was not a CREATE TABLE.
//	err     	error:  Error during translation or keyspace validation (nil if ignored due to statement type).
func ToSpannerStmt(cqlStmt, databaseName string) (string, bool, error) {
	root, err := getCqlParseTreeRoot(cqlStmt)
	if err != nil {
		return "", false, err // Return error if parsing fails
	}

	if root.CqlStatement().CreateTable() != nil {
		stmt, err := toSpannerCreateTableStmt(root.CqlStatement().CreateTable(), databaseName)
		return stmt, false, err
	}

	// The stmt is ignored if it is not a CREATE TABLE stmt.
	return "", true, nil
}

func toSpannerCreateTableStmt(createTable cql.ICreateTableContext, databaseName string) (string, error) {
	// Verify the Cassandra keyspace name is the Spanner database name.
	if createTable.TableName().KeyspaceName() != nil {
		keyspace := createTable.TableName().KeyspaceName().GetText()
		if !strings.EqualFold(keyspace, databaseName) {
			return "", fmt.Errorf("Keyspace '%s' does not match the Spanner database '%s'", keyspace, databaseName)
		}
	}

	var spannerDdlBuilder strings.Builder
	tableName := createTable.TableName().IDENTIFIER().GetText()
	if createTable.IfNotExist() != nil {
		spannerDdlBuilder.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", tableName))
	} else {
		spannerDdlBuilder.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", tableName))
	}

	// Primary key columns are required for translating the column definition to enforce the following constraints:
	// 1. Primary key columns cannot be of counter or collection (list, set, map) types.
	// 2. Primary key columns must be non-nullable. (Note: In Spanner, a primary key column can be nullable if `NOT NULL` is not explicitly defined in the column definition).
	columnDefinitionList := createTable.ColumnDefinitionList()
	var tableLevelPrimaryKey []string
	if columnDefinitionList.PrimaryKeyClause() != nil {
		tableLevelPrimaryKey = extractPrimaryKeyColumns(columnDefinitionList.PrimaryKeyClause())
	}

	var columnLevelPrimaryKey string
	hasNonCounterColumn := false
	hasCounterColumn := false
	for _, columnDefinition := range columnDefinitionList.AllColumnDefinition() {
		// TODO: Column name can not be duplicated.
		columnName := columnDefinition.ColumnName().GetText()

		isColumnLevelPk := columnDefinition.PrimaryKeyKeywords() != nil
		if isColumnLevelPk {
			if len(tableLevelPrimaryKey) > 0 || len(columnLevelPrimaryKey) > 0 {
				return "", fmt.Errorf("Multiple PRIMARY KEY specified for table '%s' (exactly one required)", tableName)
			}
			columnLevelPrimaryKey = columnName
		}

		isNotNullRequired := false
		cqlType := columnDefinition.CqlType()
		if isInTableLevelPrimaryKey(tableLevelPrimaryKey, columnName) || isColumnLevelPk {
			if err := validatePrimaryKeyType(cqlType, columnName); err != nil {
				return "", err
			}
			isNotNullRequired = true
		} else {
			isCounterColumn := (cqlType.CqlNativeType() != nil) && (cqlType.CqlNativeType().K_COUNTER() != nil)
			if isCounterColumn {
				hasCounterColumn = true
				isNotNullRequired = true
			} else {
				hasNonCounterColumn = true
			}
			if hasCounterColumn && hasNonCounterColumn {
				return "", errors.New("Cannot mix counter and non counter columns in the same table")
			}
		}

		spannerType, err := mapCqlTypeToSpannerType(cqlType)
		if err != nil {
			return "", err
		}

		// The lowercase value is required for the 'cassandra_type' option since Spanner currently only accepts lowercase values.
		if isNotNullRequired {
			spannerDdlBuilder.WriteString(fmt.Sprintf(" %s %s NOT NULL OPTIONS (cassandra_type = '%s'),\n", columnName, spannerType, strings.ToLower(cqlType.GetText())))
		} else {
			spannerDdlBuilder.WriteString(fmt.Sprintf(" %s %s OPTIONS (cassandra_type = '%s'),\n", columnName, spannerType, strings.ToLower(cqlType.GetText())))
		}
	}
	spannerDdlBuilder.WriteString(")")

	// Generate the Spanner table-level primary key.
	if len(tableLevelPrimaryKey) > 0 && len(columnLevelPrimaryKey) > 0 {
		return "", fmt.Errorf("Multiple PRIMARY KEY specified for table '%s' (exactly one required)", tableName)
	} else if len(tableLevelPrimaryKey) > 0 {
		// Cassandra (composite) partition key and clustering columns are combined to form the composite primary key in Spanner
		// Example: PRIMARY KEY ((a,b,),c,d) -> PRIMARY KEY(a,b,c,d)
		spannerDdlBuilder.WriteString(fmt.Sprintf(" PRIMARY KEY (%s)", strings.Join(tableLevelPrimaryKey, ", ")))
	} else if len(columnLevelPrimaryKey) > 0 {
		spannerDdlBuilder.WriteString(fmt.Sprintf(" PRIMARY KEY (%s)", columnLevelPrimaryKey))
	} else {
		return "", fmt.Errorf("No PRIMARY KEY specified for table '%s' (exactly one required)", tableName)
	}

	return spannerDdlBuilder.String(), nil
}
