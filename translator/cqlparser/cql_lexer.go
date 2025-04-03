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

// Code generated from CqlLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CqlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cqllexerLexerInit() {
	staticData := &CqlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'CREATE'", "'TABLE'", "'IF'", "'NOT'", "'EXISTS'", "'PRIMARY'",
		"'KEY'", "'ASCII'", "'BIGINT'", "'BLOB'", "'BOOLEAN'", "'COUNTER'",
		"'DATE'", "'DECIMAL'", "'DOUBLE'", "'FLOAT'", "'INET'", "'INT'", "'SMALLINT'",
		"'TEXT'", "'TIME'", "'TIMESTAMP'", "'TIMEUUID'", "'TINYINT'", "'UUID'",
		"'VARCHAR'", "'VARINT'", "'MAP'", "'SET'", "'LIST'", "';'", "'\"'",
		"'.'", "','", "'('", "')'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "K_CREATE", "K_TABLE", "K_IF", "K_NOT", "K_EXISTS", "K_PRIMARY",
		"K_KEY", "K_ASCII", "K_BIGINT", "K_BLOB", "K_BOOLEAN", "K_COUNTER",
		"K_DATE", "K_DECIMAL", "K_DOUBLE", "K_FLOAT", "K_INET", "K_INT", "K_SMALLINT",
		"K_TEXT", "K_TIME", "K_TIMESTAMP", "K_TIMEUUID", "K_TINYINT", "K_UUID",
		"K_VARCHAR", "K_VARINT", "K_MAP", "K_SET", "K_LIST", "SEMICOLON", "DQUOTE",
		"DOT", "COMMA", "L_PAREN", "R_PAREN", "L_ANGLE_BRACKET", "R_ANGLE_BRACKET",
		"IDENTIFIER", "IDENTIFIER_WITH_HYPHEN", "WS", "COMMENT", "MULTILINE_COMMENT",
		"UNKNOWN",
	}
	staticData.RuleNames = []string{
		"ALPHA", "K_CREATE", "K_TABLE", "K_IF", "K_NOT", "K_EXISTS", "K_PRIMARY",
		"K_KEY", "K_ASCII", "K_BIGINT", "K_BLOB", "K_BOOLEAN", "K_COUNTER",
		"K_DATE", "K_DECIMAL", "K_DOUBLE", "K_FLOAT", "K_INET", "K_INT", "K_SMALLINT",
		"K_TEXT", "K_TIME", "K_TIMESTAMP", "K_TIMEUUID", "K_TINYINT", "K_UUID",
		"K_VARCHAR", "K_VARINT", "K_MAP", "K_SET", "K_LIST", "SEMICOLON", "DQUOTE",
		"DOT", "COMMA", "L_PAREN", "R_PAREN", "L_ANGLE_BRACKET", "R_ANGLE_BRACKET",
		"IDENTIFIER", "IDENTIFIER_WITH_HYPHEN", "WS", "COMMENT", "MULTILINE_COMMENT",
		"UNKNOWN",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 357, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 5, 39, 299, 8, 39,
		10, 39, 12, 39, 302, 9, 39, 1, 40, 1, 40, 5, 40, 306, 8, 40, 10, 40, 12,
		40, 309, 9, 40, 1, 40, 1, 40, 5, 40, 313, 8, 40, 10, 40, 12, 40, 316, 9,
		40, 1, 41, 4, 41, 319, 8, 41, 11, 41, 12, 41, 320, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 330, 8, 42, 1, 42, 5, 42, 333, 8, 42,
		10, 42, 12, 42, 336, 9, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1,
		43, 1, 43, 5, 43, 346, 8, 43, 10, 43, 12, 43, 349, 9, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 2, 334, 347, 0, 45, 1, 0, 3, 1, 5, 2,
		7, 3, 9, 4, 11, 5, 13, 6, 15, 7, 17, 8, 19, 9, 21, 10, 23, 11, 25, 12,
		27, 13, 29, 14, 31, 15, 33, 16, 35, 17, 37, 18, 39, 19, 41, 20, 43, 21,
		45, 22, 47, 23, 49, 24, 51, 25, 53, 26, 55, 27, 57, 28, 59, 29, 61, 30,
		63, 31, 65, 32, 67, 33, 69, 34, 71, 35, 73, 36, 75, 37, 77, 38, 79, 39,
		81, 40, 83, 41, 85, 42, 87, 43, 89, 44, 1, 0, 27, 2, 0, 65, 90, 97, 122,
		2, 0, 67, 67, 99, 99, 2, 0, 82, 82, 114, 114, 2, 0, 69, 69, 101, 101, 2,
		0, 65, 65, 97, 97, 2, 0, 84, 84, 116, 116, 2, 0, 66, 66, 98, 98, 2, 0,
		76, 76, 108, 108, 2, 0, 73, 73, 105, 105, 2, 0, 70, 70, 102, 102, 2, 0,
		78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 88, 88, 120, 120, 2, 0,
		83, 83, 115, 115, 2, 0, 80, 80, 112, 112, 2, 0, 77, 77, 109, 109, 2, 0,
		89, 89, 121, 121, 2, 0, 75, 75, 107, 107, 2, 0, 71, 71, 103, 103, 2, 0,
		85, 85, 117, 117, 2, 0, 68, 68, 100, 100, 2, 0, 86, 86, 118, 118, 2, 0,
		72, 72, 104, 104, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 5, 0, 45, 45,
		48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 2, 0, 10,
		10, 13, 13, 362, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0,
		0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0,
		0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0,
		0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1,
		0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39,
		1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0,
		47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0,
		0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0,
		0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1,
		0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85,
		1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 1, 91, 1, 0, 0, 0, 3,
		93, 1, 0, 0, 0, 5, 100, 1, 0, 0, 0, 7, 106, 1, 0, 0, 0, 9, 109, 1, 0, 0,
		0, 11, 113, 1, 0, 0, 0, 13, 120, 1, 0, 0, 0, 15, 128, 1, 0, 0, 0, 17, 132,
		1, 0, 0, 0, 19, 138, 1, 0, 0, 0, 21, 145, 1, 0, 0, 0, 23, 150, 1, 0, 0,
		0, 25, 158, 1, 0, 0, 0, 27, 166, 1, 0, 0, 0, 29, 171, 1, 0, 0, 0, 31, 179,
		1, 0, 0, 0, 33, 186, 1, 0, 0, 0, 35, 192, 1, 0, 0, 0, 37, 197, 1, 0, 0,
		0, 39, 201, 1, 0, 0, 0, 41, 210, 1, 0, 0, 0, 43, 215, 1, 0, 0, 0, 45, 220,
		1, 0, 0, 0, 47, 230, 1, 0, 0, 0, 49, 239, 1, 0, 0, 0, 51, 247, 1, 0, 0,
		0, 53, 252, 1, 0, 0, 0, 55, 260, 1, 0, 0, 0, 57, 267, 1, 0, 0, 0, 59, 271,
		1, 0, 0, 0, 61, 275, 1, 0, 0, 0, 63, 280, 1, 0, 0, 0, 65, 282, 1, 0, 0,
		0, 67, 284, 1, 0, 0, 0, 69, 286, 1, 0, 0, 0, 71, 288, 1, 0, 0, 0, 73, 290,
		1, 0, 0, 0, 75, 292, 1, 0, 0, 0, 77, 294, 1, 0, 0, 0, 79, 296, 1, 0, 0,
		0, 81, 303, 1, 0, 0, 0, 83, 318, 1, 0, 0, 0, 85, 329, 1, 0, 0, 0, 87, 341,
		1, 0, 0, 0, 89, 355, 1, 0, 0, 0, 91, 92, 7, 0, 0, 0, 92, 2, 1, 0, 0, 0,
		93, 94, 7, 1, 0, 0, 94, 95, 7, 2, 0, 0, 95, 96, 7, 3, 0, 0, 96, 97, 7,
		4, 0, 0, 97, 98, 7, 5, 0, 0, 98, 99, 7, 3, 0, 0, 99, 4, 1, 0, 0, 0, 100,
		101, 7, 5, 0, 0, 101, 102, 7, 4, 0, 0, 102, 103, 7, 6, 0, 0, 103, 104,
		7, 7, 0, 0, 104, 105, 7, 3, 0, 0, 105, 6, 1, 0, 0, 0, 106, 107, 7, 8, 0,
		0, 107, 108, 7, 9, 0, 0, 108, 8, 1, 0, 0, 0, 109, 110, 7, 10, 0, 0, 110,
		111, 7, 11, 0, 0, 111, 112, 7, 5, 0, 0, 112, 10, 1, 0, 0, 0, 113, 114,
		7, 3, 0, 0, 114, 115, 7, 12, 0, 0, 115, 116, 7, 8, 0, 0, 116, 117, 7, 13,
		0, 0, 117, 118, 7, 5, 0, 0, 118, 119, 7, 13, 0, 0, 119, 12, 1, 0, 0, 0,
		120, 121, 7, 14, 0, 0, 121, 122, 7, 2, 0, 0, 122, 123, 7, 8, 0, 0, 123,
		124, 7, 15, 0, 0, 124, 125, 7, 4, 0, 0, 125, 126, 7, 2, 0, 0, 126, 127,
		7, 16, 0, 0, 127, 14, 1, 0, 0, 0, 128, 129, 7, 17, 0, 0, 129, 130, 7, 3,
		0, 0, 130, 131, 7, 16, 0, 0, 131, 16, 1, 0, 0, 0, 132, 133, 7, 4, 0, 0,
		133, 134, 7, 13, 0, 0, 134, 135, 7, 1, 0, 0, 135, 136, 7, 8, 0, 0, 136,
		137, 7, 8, 0, 0, 137, 18, 1, 0, 0, 0, 138, 139, 7, 6, 0, 0, 139, 140, 7,
		8, 0, 0, 140, 141, 7, 18, 0, 0, 141, 142, 7, 8, 0, 0, 142, 143, 7, 10,
		0, 0, 143, 144, 7, 5, 0, 0, 144, 20, 1, 0, 0, 0, 145, 146, 7, 6, 0, 0,
		146, 147, 7, 7, 0, 0, 147, 148, 7, 11, 0, 0, 148, 149, 7, 6, 0, 0, 149,
		22, 1, 0, 0, 0, 150, 151, 7, 6, 0, 0, 151, 152, 7, 11, 0, 0, 152, 153,
		7, 11, 0, 0, 153, 154, 7, 7, 0, 0, 154, 155, 7, 3, 0, 0, 155, 156, 7, 4,
		0, 0, 156, 157, 7, 10, 0, 0, 157, 24, 1, 0, 0, 0, 158, 159, 7, 1, 0, 0,
		159, 160, 7, 11, 0, 0, 160, 161, 7, 19, 0, 0, 161, 162, 7, 10, 0, 0, 162,
		163, 7, 5, 0, 0, 163, 164, 7, 3, 0, 0, 164, 165, 7, 2, 0, 0, 165, 26, 1,
		0, 0, 0, 166, 167, 7, 20, 0, 0, 167, 168, 7, 4, 0, 0, 168, 169, 7, 5, 0,
		0, 169, 170, 7, 3, 0, 0, 170, 28, 1, 0, 0, 0, 171, 172, 7, 20, 0, 0, 172,
		173, 7, 3, 0, 0, 173, 174, 7, 1, 0, 0, 174, 175, 7, 8, 0, 0, 175, 176,
		7, 15, 0, 0, 176, 177, 7, 4, 0, 0, 177, 178, 7, 7, 0, 0, 178, 30, 1, 0,
		0, 0, 179, 180, 7, 20, 0, 0, 180, 181, 7, 11, 0, 0, 181, 182, 7, 19, 0,
		0, 182, 183, 7, 6, 0, 0, 183, 184, 7, 7, 0, 0, 184, 185, 7, 3, 0, 0, 185,
		32, 1, 0, 0, 0, 186, 187, 7, 9, 0, 0, 187, 188, 7, 7, 0, 0, 188, 189, 7,
		11, 0, 0, 189, 190, 7, 4, 0, 0, 190, 191, 7, 5, 0, 0, 191, 34, 1, 0, 0,
		0, 192, 193, 7, 8, 0, 0, 193, 194, 7, 10, 0, 0, 194, 195, 7, 3, 0, 0, 195,
		196, 7, 5, 0, 0, 196, 36, 1, 0, 0, 0, 197, 198, 7, 8, 0, 0, 198, 199, 7,
		10, 0, 0, 199, 200, 7, 5, 0, 0, 200, 38, 1, 0, 0, 0, 201, 202, 7, 13, 0,
		0, 202, 203, 7, 15, 0, 0, 203, 204, 7, 4, 0, 0, 204, 205, 7, 7, 0, 0, 205,
		206, 7, 7, 0, 0, 206, 207, 7, 8, 0, 0, 207, 208, 7, 10, 0, 0, 208, 209,
		7, 5, 0, 0, 209, 40, 1, 0, 0, 0, 210, 211, 7, 5, 0, 0, 211, 212, 7, 3,
		0, 0, 212, 213, 7, 12, 0, 0, 213, 214, 7, 5, 0, 0, 214, 42, 1, 0, 0, 0,
		215, 216, 7, 5, 0, 0, 216, 217, 7, 8, 0, 0, 217, 218, 7, 15, 0, 0, 218,
		219, 7, 3, 0, 0, 219, 44, 1, 0, 0, 0, 220, 221, 7, 5, 0, 0, 221, 222, 7,
		8, 0, 0, 222, 223, 7, 15, 0, 0, 223, 224, 7, 3, 0, 0, 224, 225, 7, 13,
		0, 0, 225, 226, 7, 5, 0, 0, 226, 227, 7, 4, 0, 0, 227, 228, 7, 15, 0, 0,
		228, 229, 7, 14, 0, 0, 229, 46, 1, 0, 0, 0, 230, 231, 7, 5, 0, 0, 231,
		232, 7, 8, 0, 0, 232, 233, 7, 15, 0, 0, 233, 234, 7, 3, 0, 0, 234, 235,
		7, 19, 0, 0, 235, 236, 7, 19, 0, 0, 236, 237, 7, 8, 0, 0, 237, 238, 7,
		20, 0, 0, 238, 48, 1, 0, 0, 0, 239, 240, 7, 5, 0, 0, 240, 241, 7, 8, 0,
		0, 241, 242, 7, 10, 0, 0, 242, 243, 7, 16, 0, 0, 243, 244, 7, 8, 0, 0,
		244, 245, 7, 10, 0, 0, 245, 246, 7, 5, 0, 0, 246, 50, 1, 0, 0, 0, 247,
		248, 7, 19, 0, 0, 248, 249, 7, 19, 0, 0, 249, 250, 7, 8, 0, 0, 250, 251,
		7, 20, 0, 0, 251, 52, 1, 0, 0, 0, 252, 253, 7, 21, 0, 0, 253, 254, 7, 4,
		0, 0, 254, 255, 7, 2, 0, 0, 255, 256, 7, 1, 0, 0, 256, 257, 7, 22, 0, 0,
		257, 258, 7, 4, 0, 0, 258, 259, 7, 2, 0, 0, 259, 54, 1, 0, 0, 0, 260, 261,
		7, 21, 0, 0, 261, 262, 7, 4, 0, 0, 262, 263, 7, 2, 0, 0, 263, 264, 7, 8,
		0, 0, 264, 265, 7, 10, 0, 0, 265, 266, 7, 5, 0, 0, 266, 56, 1, 0, 0, 0,
		267, 268, 7, 15, 0, 0, 268, 269, 7, 4, 0, 0, 269, 270, 7, 14, 0, 0, 270,
		58, 1, 0, 0, 0, 271, 272, 7, 13, 0, 0, 272, 273, 7, 3, 0, 0, 273, 274,
		7, 5, 0, 0, 274, 60, 1, 0, 0, 0, 275, 276, 7, 7, 0, 0, 276, 277, 7, 8,
		0, 0, 277, 278, 7, 13, 0, 0, 278, 279, 7, 5, 0, 0, 279, 62, 1, 0, 0, 0,
		280, 281, 5, 59, 0, 0, 281, 64, 1, 0, 0, 0, 282, 283, 5, 34, 0, 0, 283,
		66, 1, 0, 0, 0, 284, 285, 5, 46, 0, 0, 285, 68, 1, 0, 0, 0, 286, 287, 5,
		44, 0, 0, 287, 70, 1, 0, 0, 0, 288, 289, 5, 40, 0, 0, 289, 72, 1, 0, 0,
		0, 290, 291, 5, 41, 0, 0, 291, 74, 1, 0, 0, 0, 292, 293, 5, 60, 0, 0, 293,
		76, 1, 0, 0, 0, 294, 295, 5, 62, 0, 0, 295, 78, 1, 0, 0, 0, 296, 300, 3,
		1, 0, 0, 297, 299, 7, 23, 0, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0,
		0, 300, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 80, 1, 0, 0, 0, 302,
		300, 1, 0, 0, 0, 303, 307, 3, 1, 0, 0, 304, 306, 7, 24, 0, 0, 305, 304,
		1, 0, 0, 0, 306, 309, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 307, 308, 1, 0,
		0, 0, 308, 310, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 310, 314, 5, 45, 0, 0,
		311, 313, 7, 24, 0, 0, 312, 311, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314,
		312, 1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 82, 1, 0, 0, 0, 316, 314, 1,
		0, 0, 0, 317, 319, 7, 25, 0, 0, 318, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0,
		0, 320, 318, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322,
		323, 6, 41, 0, 0, 323, 84, 1, 0, 0, 0, 324, 325, 5, 45, 0, 0, 325, 326,
		5, 45, 0, 0, 326, 330, 5, 32, 0, 0, 327, 328, 5, 47, 0, 0, 328, 330, 5,
		47, 0, 0, 329, 324, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0, 330, 334, 1, 0, 0,
		0, 331, 333, 9, 0, 0, 0, 332, 331, 1, 0, 0, 0, 333, 336, 1, 0, 0, 0, 334,
		335, 1, 0, 0, 0, 334, 332, 1, 0, 0, 0, 335, 337, 1, 0, 0, 0, 336, 334,
		1, 0, 0, 0, 337, 338, 7, 26, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 6, 42,
		0, 0, 340, 86, 1, 0, 0, 0, 341, 342, 5, 47, 0, 0, 342, 343, 5, 42, 0, 0,
		343, 347, 1, 0, 0, 0, 344, 346, 9, 0, 0, 0, 345, 344, 1, 0, 0, 0, 346,
		349, 1, 0, 0, 0, 347, 348, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 350,
		1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 350, 351, 5, 42, 0, 0, 351, 352, 5, 47,
		0, 0, 352, 353, 1, 0, 0, 0, 353, 354, 6, 43, 0, 0, 354, 88, 1, 0, 0, 0,
		355, 356, 9, 0, 0, 0, 356, 90, 1, 0, 0, 0, 8, 0, 300, 307, 314, 320, 329,
		334, 347, 1, 0, 1, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CqlLexerInit initializes any static state used to implement CqlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCqlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CqlLexerInit() {
	staticData := &CqlLexerLexerStaticData
	staticData.once.Do(cqllexerLexerInit)
}

// NewCqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCqlLexer(input antlr.CharStream) *CqlLexer {
	CqlLexerInit()
	l := new(CqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CqlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CqlLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CqlLexer tokens.
const (
	CqlLexerK_CREATE               = 1
	CqlLexerK_TABLE                = 2
	CqlLexerK_IF                   = 3
	CqlLexerK_NOT                  = 4
	CqlLexerK_EXISTS               = 5
	CqlLexerK_PRIMARY              = 6
	CqlLexerK_KEY                  = 7
	CqlLexerK_ASCII                = 8
	CqlLexerK_BIGINT               = 9
	CqlLexerK_BLOB                 = 10
	CqlLexerK_BOOLEAN              = 11
	CqlLexerK_COUNTER              = 12
	CqlLexerK_DATE                 = 13
	CqlLexerK_DECIMAL              = 14
	CqlLexerK_DOUBLE               = 15
	CqlLexerK_FLOAT                = 16
	CqlLexerK_INET                 = 17
	CqlLexerK_INT                  = 18
	CqlLexerK_SMALLINT             = 19
	CqlLexerK_TEXT                 = 20
	CqlLexerK_TIME                 = 21
	CqlLexerK_TIMESTAMP            = 22
	CqlLexerK_TIMEUUID             = 23
	CqlLexerK_TINYINT              = 24
	CqlLexerK_UUID                 = 25
	CqlLexerK_VARCHAR              = 26
	CqlLexerK_VARINT               = 27
	CqlLexerK_MAP                  = 28
	CqlLexerK_SET                  = 29
	CqlLexerK_LIST                 = 30
	CqlLexerSEMICOLON              = 31
	CqlLexerDQUOTE                 = 32
	CqlLexerDOT                    = 33
	CqlLexerCOMMA                  = 34
	CqlLexerL_PAREN                = 35
	CqlLexerR_PAREN                = 36
	CqlLexerL_ANGLE_BRACKET        = 37
	CqlLexerR_ANGLE_BRACKET        = 38
	CqlLexerIDENTIFIER             = 39
	CqlLexerIDENTIFIER_WITH_HYPHEN = 40
	CqlLexerWS                     = 41
	CqlLexerCOMMENT                = 42
	CqlLexerMULTILINE_COMMENT      = 43
	CqlLexerUNKNOWN                = 44
)
