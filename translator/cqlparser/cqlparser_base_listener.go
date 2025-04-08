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

// BaseCqlParserListener is a complete listener for a parse tree produced by CqlParser.
type BaseCqlParserListener struct{}

var _ CqlParserListener = &BaseCqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCqlParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCqlParserListener) ExitRoot(ctx *RootContext) {}

// EnterCqlStatement is called when production cqlStatement is entered.
func (s *BaseCqlParserListener) EnterCqlStatement(ctx *CqlStatementContext) {}

// ExitCqlStatement is called when production cqlStatement is exited.
func (s *BaseCqlParserListener) ExitCqlStatement(ctx *CqlStatementContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseCqlParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseCqlParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterCreateKeyspace is called when production createKeyspace is entered.
func (s *BaseCqlParserListener) EnterCreateKeyspace(ctx *CreateKeyspaceContext) {}

// ExitCreateKeyspace is called when production createKeyspace is exited.
func (s *BaseCqlParserListener) ExitCreateKeyspace(ctx *CreateKeyspaceContext) {}

// EnterColumnDefinitionList is called when production columnDefinitionList is entered.
func (s *BaseCqlParserListener) EnterColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// ExitColumnDefinitionList is called when production columnDefinitionList is exited.
func (s *BaseCqlParserListener) ExitColumnDefinitionList(ctx *ColumnDefinitionListContext) {}

// EnterColumnDefinition is called when production columnDefinition is entered.
func (s *BaseCqlParserListener) EnterColumnDefinition(ctx *ColumnDefinitionContext) {}

// ExitColumnDefinition is called when production columnDefinition is exited.
func (s *BaseCqlParserListener) ExitColumnDefinition(ctx *ColumnDefinitionContext) {}

// EnterPrimaryKeyClause is called when production primaryKeyClause is entered.
func (s *BaseCqlParserListener) EnterPrimaryKeyClause(ctx *PrimaryKeyClauseContext) {}

// ExitPrimaryKeyClause is called when production primaryKeyClause is exited.
func (s *BaseCqlParserListener) ExitPrimaryKeyClause(ctx *PrimaryKeyClauseContext) {}

// EnterPrimaryKey is called when production primaryKey is entered.
func (s *BaseCqlParserListener) EnterPrimaryKey(ctx *PrimaryKeyContext) {}

// ExitPrimaryKey is called when production primaryKey is exited.
func (s *BaseCqlParserListener) ExitPrimaryKey(ctx *PrimaryKeyContext) {}

// EnterPartitionKey is called when production partitionKey is entered.
func (s *BaseCqlParserListener) EnterPartitionKey(ctx *PartitionKeyContext) {}

// ExitPartitionKey is called when production partitionKey is exited.
func (s *BaseCqlParserListener) ExitPartitionKey(ctx *PartitionKeyContext) {}

// EnterClusteringColumns is called when production clusteringColumns is entered.
func (s *BaseCqlParserListener) EnterClusteringColumns(ctx *ClusteringColumnsContext) {}

// ExitClusteringColumns is called when production clusteringColumns is exited.
func (s *BaseCqlParserListener) ExitClusteringColumns(ctx *ClusteringColumnsContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseCqlParserListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseCqlParserListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseCqlParserListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseCqlParserListener) ExitTableName(ctx *TableNameContext) {}

// EnterKeyspaceName is called when production keyspaceName is entered.
func (s *BaseCqlParserListener) EnterKeyspaceName(ctx *KeyspaceNameContext) {}

// ExitKeyspaceName is called when production keyspaceName is exited.
func (s *BaseCqlParserListener) ExitKeyspaceName(ctx *KeyspaceNameContext) {}

// EnterGeneralIdentifier is called when production generalIdentifier is entered.
func (s *BaseCqlParserListener) EnterGeneralIdentifier(ctx *GeneralIdentifierContext) {}

// ExitGeneralIdentifier is called when production generalIdentifier is exited.
func (s *BaseCqlParserListener) ExitGeneralIdentifier(ctx *GeneralIdentifierContext) {}

// EnterCqlType is called when production cqlType is entered.
func (s *BaseCqlParserListener) EnterCqlType(ctx *CqlTypeContext) {}

// ExitCqlType is called when production cqlType is exited.
func (s *BaseCqlParserListener) ExitCqlType(ctx *CqlTypeContext) {}

// EnterPrimaryKeyKeywords is called when production primaryKeyKeywords is entered.
func (s *BaseCqlParserListener) EnterPrimaryKeyKeywords(ctx *PrimaryKeyKeywordsContext) {}

// ExitPrimaryKeyKeywords is called when production primaryKeyKeywords is exited.
func (s *BaseCqlParserListener) ExitPrimaryKeyKeywords(ctx *PrimaryKeyKeywordsContext) {}

// EnterIfNotExist is called when production ifNotExist is entered.
func (s *BaseCqlParserListener) EnterIfNotExist(ctx *IfNotExistContext) {}

// ExitIfNotExist is called when production ifNotExist is exited.
func (s *BaseCqlParserListener) ExitIfNotExist(ctx *IfNotExistContext) {}

// EnterCqlNativeType is called when production cqlNativeType is entered.
func (s *BaseCqlParserListener) EnterCqlNativeType(ctx *CqlNativeTypeContext) {}

// ExitCqlNativeType is called when production cqlNativeType is exited.
func (s *BaseCqlParserListener) ExitCqlNativeType(ctx *CqlNativeTypeContext) {}

// EnterCqlCollectionType is called when production cqlCollectionType is entered.
func (s *BaseCqlParserListener) EnterCqlCollectionType(ctx *CqlCollectionTypeContext) {}

// ExitCqlCollectionType is called when production cqlCollectionType is exited.
func (s *BaseCqlParserListener) ExitCqlCollectionType(ctx *CqlCollectionTypeContext) {}

// EnterWihtTableOptions is called when production wihtTableOptions is entered.
func (s *BaseCqlParserListener) EnterWihtTableOptions(ctx *WihtTableOptionsContext) {}

// ExitWihtTableOptions is called when production wihtTableOptions is exited.
func (s *BaseCqlParserListener) ExitWihtTableOptions(ctx *WihtTableOptionsContext) {}

// EnterNonSemicolonToken is called when production nonSemicolonToken is entered.
func (s *BaseCqlParserListener) EnterNonSemicolonToken(ctx *NonSemicolonTokenContext) {}

// ExitNonSemicolonToken is called when production nonSemicolonToken is exited.
func (s *BaseCqlParserListener) ExitNonSemicolonToken(ctx *NonSemicolonTokenContext) {}
