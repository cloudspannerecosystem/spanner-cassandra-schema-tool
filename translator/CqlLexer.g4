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
 * Lexical definitions for CQL.
 */
lexer grammar CqlLexer;

options {
  caseInsensitive = true;
}

fragment ALPHA: [A-Z];

// Keywords
K_CREATE   : 'CREATE';
K_TABLE    : 'TABLE';
K_IF       : 'IF';
K_NOT      : 'NOT';
K_EXISTS   : 'EXISTS';
K_PRIMARY  : 'PRIMARY';
K_KEY      : 'KEY';
K_WITH     : 'WITH';
K_KEYSPACE : 'KEYSPACE';

K_ASCII    : 'ASCII';
K_BIGINT   : 'BIGINT';
K_BLOB     : 'BLOB';
K_BOOLEAN  : 'BOOLEAN';
K_COUNTER  : 'COUNTER';
K_DATE     : 'DATE';
K_DECIMAL  : 'DECIMAL';
K_DOUBLE   : 'DOUBLE';
K_FLOAT    : 'FLOAT';
K_INET     : 'INET';
K_INT      : 'INT';
K_SMALLINT : 'SMALLINT';
K_TEXT     : 'TEXT';
K_TIME     : 'TIME';
K_TIMESTAMP: 'TIMESTAMP';
K_TIMEUUID : 'TIMEUUID';
K_TINYINT  : 'TINYINT';
K_UUID     : 'UUID';
K_VARCHAR  : 'VARCHAR';
K_VARINT   : 'VARINT';
K_MAP      : 'MAP';
K_SET      : 'SET';
K_LIST     : 'LIST';


// Punctuation
SEMICOLON       : ';';
DQUOTE          : '"';
DOT             : '.';
COMMA           : ',';
L_PAREN         : '(';
R_PAREN         : ')';
L_ANGLE_BRACKET : '<';
R_ANGLE_BRACKET : '>';

// Identifiers
IDENTIFIER: ALPHA[A-Z0-9_]*;
IDENTIFIER_WITH_HYPHEN: ALPHA [A-Z0-9_-]* '-' [A-Z0-9_-]*;

// Hidden
WS
    : [ \t\r\n]+ -> channel (HIDDEN)
    ;

COMMENT
    : ('-- ' | '//') .*? ('\n' | '\r') -> channel (HIDDEN)
    ;

MULTILINE_COMMENT
    : '/*' .*? '*/' -> channel (HIDDEN)
    ;

UNKNOWN
    : .
    ;