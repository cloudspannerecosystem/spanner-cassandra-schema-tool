// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated from CqlParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // CqlParser

import "github.com/antlr4-go/antlr/v4"

// CqlParserListener is a complete listener for a parse tree produced by CqlParser.
type CqlParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterCqlStatement is called when entering the cqlStatement production.
	EnterCqlStatement(c *CqlStatementContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterCreateKeyspace is called when entering the createKeyspace production.
	EnterCreateKeyspace(c *CreateKeyspaceContext)

	// EnterColumnDefinitionList is called when entering the columnDefinitionList production.
	EnterColumnDefinitionList(c *ColumnDefinitionListContext)

	// EnterColumnDefinition is called when entering the columnDefinition production.
	EnterColumnDefinition(c *ColumnDefinitionContext)

	// EnterPrimaryKeyClause is called when entering the primaryKeyClause production.
	EnterPrimaryKeyClause(c *PrimaryKeyClauseContext)

	// EnterPrimaryKey is called when entering the primaryKey production.
	EnterPrimaryKey(c *PrimaryKeyContext)

	// EnterPartitionKey is called when entering the partitionKey production.
	EnterPartitionKey(c *PartitionKeyContext)

	// EnterClusteringColumns is called when entering the clusteringColumns production.
	EnterClusteringColumns(c *ClusteringColumnsContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterKeyspaceName is called when entering the keyspaceName production.
	EnterKeyspaceName(c *KeyspaceNameContext)

	// EnterGeneralIdentifier is called when entering the generalIdentifier production.
	EnterGeneralIdentifier(c *GeneralIdentifierContext)

	// EnterCqlType is called when entering the cqlType production.
	EnterCqlType(c *CqlTypeContext)

	// EnterPrimaryKeyKeywords is called when entering the primaryKeyKeywords production.
	EnterPrimaryKeyKeywords(c *PrimaryKeyKeywordsContext)

	// EnterIfNotExist is called when entering the ifNotExist production.
	EnterIfNotExist(c *IfNotExistContext)

	// EnterCqlNativeType is called when entering the cqlNativeType production.
	EnterCqlNativeType(c *CqlNativeTypeContext)

	// EnterCqlCollectionType is called when entering the cqlCollectionType production.
	EnterCqlCollectionType(c *CqlCollectionTypeContext)

	// EnterWihtTableOptions is called when entering the wihtTableOptions production.
	EnterWihtTableOptions(c *WihtTableOptionsContext)

	// EnterNonSemicolonToken is called when entering the nonSemicolonToken production.
	EnterNonSemicolonToken(c *NonSemicolonTokenContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitCqlStatement is called when exiting the cqlStatement production.
	ExitCqlStatement(c *CqlStatementContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitCreateKeyspace is called when exiting the createKeyspace production.
	ExitCreateKeyspace(c *CreateKeyspaceContext)

	// ExitColumnDefinitionList is called when exiting the columnDefinitionList production.
	ExitColumnDefinitionList(c *ColumnDefinitionListContext)

	// ExitColumnDefinition is called when exiting the columnDefinition production.
	ExitColumnDefinition(c *ColumnDefinitionContext)

	// ExitPrimaryKeyClause is called when exiting the primaryKeyClause production.
	ExitPrimaryKeyClause(c *PrimaryKeyClauseContext)

	// ExitPrimaryKey is called when exiting the primaryKey production.
	ExitPrimaryKey(c *PrimaryKeyContext)

	// ExitPartitionKey is called when exiting the partitionKey production.
	ExitPartitionKey(c *PartitionKeyContext)

	// ExitClusteringColumns is called when exiting the clusteringColumns production.
	ExitClusteringColumns(c *ClusteringColumnsContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitKeyspaceName is called when exiting the keyspaceName production.
	ExitKeyspaceName(c *KeyspaceNameContext)

	// ExitGeneralIdentifier is called when exiting the generalIdentifier production.
	ExitGeneralIdentifier(c *GeneralIdentifierContext)

	// ExitCqlType is called when exiting the cqlType production.
	ExitCqlType(c *CqlTypeContext)

	// ExitPrimaryKeyKeywords is called when exiting the primaryKeyKeywords production.
	ExitPrimaryKeyKeywords(c *PrimaryKeyKeywordsContext)

	// ExitIfNotExist is called when exiting the ifNotExist production.
	ExitIfNotExist(c *IfNotExistContext)

	// ExitCqlNativeType is called when exiting the cqlNativeType production.
	ExitCqlNativeType(c *CqlNativeTypeContext)

	// ExitCqlCollectionType is called when exiting the cqlCollectionType production.
	ExitCqlCollectionType(c *CqlCollectionTypeContext)

	// ExitWihtTableOptions is called when exiting the wihtTableOptions production.
	ExitWihtTableOptions(c *WihtTableOptionsContext)

	// ExitNonSemicolonToken is called when exiting the nonSemicolonToken production.
	ExitNonSemicolonToken(c *NonSemicolonTokenContext)
}
