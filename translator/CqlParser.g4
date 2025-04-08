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
 
/**
 * Grammar for CQL.
 */
parser grammar CqlParser;

options {
  tokenVocab = CqlLexer;
}

root
    : cqlStatement SEMICOLON? EOF
    ;

cqlStatement
    : createTable
    | createKeyspace
    ;

createTable
   : K_CREATE K_TABLE ifNotExist? tableName L_PAREN columnDefinitionList R_PAREN wihtTableOptions?
   ;

// TODO: Consider using a more precise syntax for options.
createKeyspace
   : K_CREATE K_KEYSPACE ifNotExist? keyspaceName K_WITH nonSemicolonToken*
   ;

columnDefinitionList
   : columnDefinition (COMMA columnDefinition)* (COMMA primaryKeyClause)?
   ;

columnDefinition
   : columnName cqlType primaryKeyKeywords?
   ;

// TODO: Merge the rule into columnDefinition.
primaryKeyClause
   : primaryKeyKeywords L_PAREN primaryKey R_PAREN
   ;

primaryKey
    : partitionKey (COMMA clusteringColumns)?
    ;

partitionKey
    : columnName
    | L_PAREN columnName (COMMA columnName)* R_PAREN
    ;

clusteringColumns
    : columnName (COMMA columnName)*
    ;

columnName
    : IDENTIFIER
    ;

tableName
    : (keyspaceName DOT)? IDENTIFIER
    ;

// A Cassandra keyspace name is conventionally defined as an IDENTIFIER. Spanner
// database names, however, permit hyphens. To ensure compatibility when mapping
// Spanner databases to Cassandra keyspaces, the parsing rule is relaxed to
// allow hyphens in keyspace names.
keyspaceName
    : generalIdentifier
    ;

generalIdentifier
    : IDENTIFIER_WITH_HYPHEN
    | IDENTIFIER
    ;

// TODO: Add support on frozen, duration, and tuple.
cqlType
    : cqlNativeType
    | cqlCollectionType
    ;

primaryKeyKeywords
    : K_PRIMARY K_KEY
    ;

ifNotExist
   : K_IF K_NOT K_EXISTS
   ;

cqlNativeType
    : K_ASCII
    | K_BIGINT
    | K_BLOB
    | K_BOOLEAN
    | K_COUNTER
    | K_DATE
    | K_DECIMAL
    | K_DOUBLE
    | K_FLOAT
    | K_INET
    | K_INT
    | K_SMALLINT
    | K_TEXT
    | K_TIME
    | K_TIMESTAMP
    | K_TIMEUUID
    | K_TINYINT
    | K_UUID
    | K_VARCHAR
    | K_VARINT
    ;

cqlCollectionType
    : K_MAP L_ANGLE_BRACKET cqlNativeType COMMA cqlNativeType R_ANGLE_BRACKET
    | (K_SET | K_LIST) L_ANGLE_BRACKET cqlNativeType R_ANGLE_BRACKET
    ;

// TODO: Consider using a more precise syntax.
wihtTableOptions
    : K_WITH nonSemicolonToken*
    ;

nonSemicolonToken
    : K_CREATE
    | K_TABLE
    | K_IF
    | K_NOT
    | K_EXISTS
    | K_PRIMARY
    | K_KEY
    | K_WITH
    | K_KEYSPACE
    | K_ASCII
    | K_BIGINT
    | K_BLOB
    | K_BOOLEAN
    | K_COUNTER
    | K_DATE
    | K_DECIMAL
    | K_DOUBLE
    | K_FLOAT
    | K_INET
    | K_INT
    | K_SMALLINT
    | K_TEXT
    | K_TIME
    | K_TIMESTAMP
    | K_TIMEUUID
    | K_TINYINT
    | K_UUID
    | K_VARCHAR
    | K_VARINT
    | K_MAP
    | K_SET
    | K_LIST
    | DQUOTE
    | DOT
    | COMMA
    | L_PAREN
    | R_PAREN
    | L_ANGLE_BRACKET
    | R_ANGLE_BRACKET
    | IDENTIFIER
    | IDENTIFIER_WITH_HYPHEN
    | UNKNOWN
    ;