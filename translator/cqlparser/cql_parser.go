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

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CqlParser struct {
	*antlr.BaseParser
}

var CqlParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func cqlparserParserInit() {
	staticData := &CqlParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'CREATE'", "'TABLE'", "'IF'", "'NOT'", "'EXISTS'", "'PRIMARY'",
		"'KEY'", "'WITH'", "'KEYSPACE'", "'ASCII'", "'BIGINT'", "'BLOB'", "'BOOLEAN'",
		"'COUNTER'", "'DATE'", "'DECIMAL'", "'DOUBLE'", "'FLOAT'", "'INET'",
		"'INT'", "'SMALLINT'", "'TEXT'", "'TIME'", "'TIMESTAMP'", "'TIMEUUID'",
		"'TINYINT'", "'UUID'", "'VARCHAR'", "'VARINT'", "'MAP'", "'SET'", "'LIST'",
		"';'", "'\"'", "'.'", "','", "'('", "')'", "'<'", "'>'",
	}
	staticData.SymbolicNames = []string{
		"", "K_CREATE", "K_TABLE", "K_IF", "K_NOT", "K_EXISTS", "K_PRIMARY",
		"K_KEY", "K_WITH", "K_KEYSPACE", "K_ASCII", "K_BIGINT", "K_BLOB", "K_BOOLEAN",
		"K_COUNTER", "K_DATE", "K_DECIMAL", "K_DOUBLE", "K_FLOAT", "K_INET",
		"K_INT", "K_SMALLINT", "K_TEXT", "K_TIME", "K_TIMESTAMP", "K_TIMEUUID",
		"K_TINYINT", "K_UUID", "K_VARCHAR", "K_VARINT", "K_MAP", "K_SET", "K_LIST",
		"SEMICOLON", "DQUOTE", "DOT", "COMMA", "L_PAREN", "R_PAREN", "L_ANGLE_BRACKET",
		"R_ANGLE_BRACKET", "IDENTIFIER", "IDENTIFIER_WITH_HYPHEN", "WS", "COMMENT",
		"MULTILINE_COMMENT", "UNKNOWN",
	}
	staticData.RuleNames = []string{
		"root", "cqlStatement", "createTable", "createKeyspace", "columnDefinitionList",
		"columnDefinition", "primaryKeyClause", "primaryKey", "partitionKey",
		"clusteringColumns", "columnName", "tableName", "keyspaceName", "generalIdentifier",
		"cqlType", "primaryKeyKeywords", "ifNotExist", "cqlNativeType", "cqlCollectionType",
		"wihtTableOptions", "nonSemicolonToken",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 176, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 3, 0, 45, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 51, 8, 1, 1, 2,
		1, 2, 1, 2, 3, 2, 56, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 63, 8,
		2, 1, 3, 1, 3, 1, 3, 3, 3, 68, 8, 3, 1, 3, 1, 3, 1, 3, 5, 3, 73, 8, 3,
		10, 3, 12, 3, 76, 9, 3, 1, 4, 1, 4, 1, 4, 5, 4, 81, 8, 4, 10, 4, 12, 4,
		84, 9, 4, 1, 4, 1, 4, 3, 4, 88, 8, 4, 1, 5, 1, 5, 1, 5, 3, 5, 93, 8, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 103, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 110, 8, 8, 10, 8, 12, 8, 113, 9, 8, 1, 8,
		1, 8, 3, 8, 117, 8, 8, 1, 9, 1, 9, 1, 9, 5, 9, 122, 8, 9, 10, 9, 12, 9,
		125, 9, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 132, 8, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 3, 14, 142, 8, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 165, 8, 18, 1, 19, 1, 19, 5, 19, 169, 8, 19, 10, 19, 12, 19, 172, 9,
		19, 1, 20, 1, 20, 1, 20, 0, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
		22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 0, 4, 1, 0, 41, 42, 1, 0, 10, 29,
		1, 0, 31, 32, 3, 0, 1, 32, 34, 42, 46, 46, 171, 0, 42, 1, 0, 0, 0, 2, 50,
		1, 0, 0, 0, 4, 52, 1, 0, 0, 0, 6, 64, 1, 0, 0, 0, 8, 77, 1, 0, 0, 0, 10,
		89, 1, 0, 0, 0, 12, 94, 1, 0, 0, 0, 14, 99, 1, 0, 0, 0, 16, 116, 1, 0,
		0, 0, 18, 118, 1, 0, 0, 0, 20, 126, 1, 0, 0, 0, 22, 131, 1, 0, 0, 0, 24,
		135, 1, 0, 0, 0, 26, 137, 1, 0, 0, 0, 28, 141, 1, 0, 0, 0, 30, 143, 1,
		0, 0, 0, 32, 146, 1, 0, 0, 0, 34, 150, 1, 0, 0, 0, 36, 164, 1, 0, 0, 0,
		38, 166, 1, 0, 0, 0, 40, 173, 1, 0, 0, 0, 42, 44, 3, 2, 1, 0, 43, 45, 5,
		33, 0, 0, 44, 43, 1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46,
		47, 5, 0, 0, 1, 47, 1, 1, 0, 0, 0, 48, 51, 3, 4, 2, 0, 49, 51, 3, 6, 3,
		0, 50, 48, 1, 0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 3, 1, 0, 0, 0, 52, 53, 5,
		1, 0, 0, 53, 55, 5, 2, 0, 0, 54, 56, 3, 32, 16, 0, 55, 54, 1, 0, 0, 0,
		55, 56, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 3, 22, 11, 0, 58, 59, 5,
		37, 0, 0, 59, 60, 3, 8, 4, 0, 60, 62, 5, 38, 0, 0, 61, 63, 3, 38, 19, 0,
		62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 5, 1, 0, 0, 0, 64, 65, 5, 1,
		0, 0, 65, 67, 5, 9, 0, 0, 66, 68, 3, 32, 16, 0, 67, 66, 1, 0, 0, 0, 67,
		68, 1, 0, 0, 0, 68, 69, 1, 0, 0, 0, 69, 70, 3, 24, 12, 0, 70, 74, 5, 8,
		0, 0, 71, 73, 3, 40, 20, 0, 72, 71, 1, 0, 0, 0, 73, 76, 1, 0, 0, 0, 74,
		72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 7, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 77, 82, 3, 10, 5, 0, 78, 79, 5, 36, 0, 0, 79, 81, 3, 10, 5, 0, 80, 78,
		1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0,
		83, 87, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 36, 0, 0, 86, 88, 3,
		12, 6, 0, 87, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 9, 1, 0, 0, 0, 89,
		90, 3, 20, 10, 0, 90, 92, 3, 28, 14, 0, 91, 93, 3, 30, 15, 0, 92, 91, 1,
		0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 11, 1, 0, 0, 0, 94, 95, 3, 30, 15, 0,
		95, 96, 5, 37, 0, 0, 96, 97, 3, 14, 7, 0, 97, 98, 5, 38, 0, 0, 98, 13,
		1, 0, 0, 0, 99, 102, 3, 16, 8, 0, 100, 101, 5, 36, 0, 0, 101, 103, 3, 18,
		9, 0, 102, 100, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 15, 1, 0, 0, 0,
		104, 117, 3, 20, 10, 0, 105, 106, 5, 37, 0, 0, 106, 111, 3, 20, 10, 0,
		107, 108, 5, 36, 0, 0, 108, 110, 3, 20, 10, 0, 109, 107, 1, 0, 0, 0, 110,
		113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 114,
		1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 5, 38, 0, 0, 115, 117, 1, 0,
		0, 0, 116, 104, 1, 0, 0, 0, 116, 105, 1, 0, 0, 0, 117, 17, 1, 0, 0, 0,
		118, 123, 3, 20, 10, 0, 119, 120, 5, 36, 0, 0, 120, 122, 3, 20, 10, 0,
		121, 119, 1, 0, 0, 0, 122, 125, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 123,
		124, 1, 0, 0, 0, 124, 19, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 126, 127, 5,
		41, 0, 0, 127, 21, 1, 0, 0, 0, 128, 129, 3, 24, 12, 0, 129, 130, 5, 35,
		0, 0, 130, 132, 1, 0, 0, 0, 131, 128, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0,
		132, 133, 1, 0, 0, 0, 133, 134, 5, 41, 0, 0, 134, 23, 1, 0, 0, 0, 135,
		136, 3, 26, 13, 0, 136, 25, 1, 0, 0, 0, 137, 138, 7, 0, 0, 0, 138, 27,
		1, 0, 0, 0, 139, 142, 3, 34, 17, 0, 140, 142, 3, 36, 18, 0, 141, 139, 1,
		0, 0, 0, 141, 140, 1, 0, 0, 0, 142, 29, 1, 0, 0, 0, 143, 144, 5, 6, 0,
		0, 144, 145, 5, 7, 0, 0, 145, 31, 1, 0, 0, 0, 146, 147, 5, 3, 0, 0, 147,
		148, 5, 4, 0, 0, 148, 149, 5, 5, 0, 0, 149, 33, 1, 0, 0, 0, 150, 151, 7,
		1, 0, 0, 151, 35, 1, 0, 0, 0, 152, 153, 5, 30, 0, 0, 153, 154, 5, 39, 0,
		0, 154, 155, 3, 34, 17, 0, 155, 156, 5, 36, 0, 0, 156, 157, 3, 34, 17,
		0, 157, 158, 5, 40, 0, 0, 158, 165, 1, 0, 0, 0, 159, 160, 7, 2, 0, 0, 160,
		161, 5, 39, 0, 0, 161, 162, 3, 34, 17, 0, 162, 163, 5, 40, 0, 0, 163, 165,
		1, 0, 0, 0, 164, 152, 1, 0, 0, 0, 164, 159, 1, 0, 0, 0, 165, 37, 1, 0,
		0, 0, 166, 170, 5, 8, 0, 0, 167, 169, 3, 40, 20, 0, 168, 167, 1, 0, 0,
		0, 169, 172, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171,
		39, 1, 0, 0, 0, 172, 170, 1, 0, 0, 0, 173, 174, 7, 3, 0, 0, 174, 41, 1,
		0, 0, 0, 17, 44, 50, 55, 62, 67, 74, 82, 87, 92, 102, 111, 116, 123, 131,
		141, 164, 170,
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

// CqlParserInit initializes any static state used to implement CqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CqlParserInit() {
	staticData := &CqlParserParserStaticData
	staticData.once.Do(cqlparserParserInit)
}

// NewCqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCqlParser(input antlr.TokenStream) *CqlParser {
	CqlParserInit()
	this := new(CqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CqlParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CqlParser.g4"

	return this
}

// CqlParser tokens.
const (
	CqlParserEOF                    = antlr.TokenEOF
	CqlParserK_CREATE               = 1
	CqlParserK_TABLE                = 2
	CqlParserK_IF                   = 3
	CqlParserK_NOT                  = 4
	CqlParserK_EXISTS               = 5
	CqlParserK_PRIMARY              = 6
	CqlParserK_KEY                  = 7
	CqlParserK_WITH                 = 8
	CqlParserK_KEYSPACE             = 9
	CqlParserK_ASCII                = 10
	CqlParserK_BIGINT               = 11
	CqlParserK_BLOB                 = 12
	CqlParserK_BOOLEAN              = 13
	CqlParserK_COUNTER              = 14
	CqlParserK_DATE                 = 15
	CqlParserK_DECIMAL              = 16
	CqlParserK_DOUBLE               = 17
	CqlParserK_FLOAT                = 18
	CqlParserK_INET                 = 19
	CqlParserK_INT                  = 20
	CqlParserK_SMALLINT             = 21
	CqlParserK_TEXT                 = 22
	CqlParserK_TIME                 = 23
	CqlParserK_TIMESTAMP            = 24
	CqlParserK_TIMEUUID             = 25
	CqlParserK_TINYINT              = 26
	CqlParserK_UUID                 = 27
	CqlParserK_VARCHAR              = 28
	CqlParserK_VARINT               = 29
	CqlParserK_MAP                  = 30
	CqlParserK_SET                  = 31
	CqlParserK_LIST                 = 32
	CqlParserSEMICOLON              = 33
	CqlParserDQUOTE                 = 34
	CqlParserDOT                    = 35
	CqlParserCOMMA                  = 36
	CqlParserL_PAREN                = 37
	CqlParserR_PAREN                = 38
	CqlParserL_ANGLE_BRACKET        = 39
	CqlParserR_ANGLE_BRACKET        = 40
	CqlParserIDENTIFIER             = 41
	CqlParserIDENTIFIER_WITH_HYPHEN = 42
	CqlParserWS                     = 43
	CqlParserCOMMENT                = 44
	CqlParserMULTILINE_COMMENT      = 45
	CqlParserUNKNOWN                = 46
)

// CqlParser rules.
const (
	CqlParserRULE_root                 = 0
	CqlParserRULE_cqlStatement         = 1
	CqlParserRULE_createTable          = 2
	CqlParserRULE_createKeyspace       = 3
	CqlParserRULE_columnDefinitionList = 4
	CqlParserRULE_columnDefinition     = 5
	CqlParserRULE_primaryKeyClause     = 6
	CqlParserRULE_primaryKey           = 7
	CqlParserRULE_partitionKey         = 8
	CqlParserRULE_clusteringColumns    = 9
	CqlParserRULE_columnName           = 10
	CqlParserRULE_tableName            = 11
	CqlParserRULE_keyspaceName         = 12
	CqlParserRULE_generalIdentifier    = 13
	CqlParserRULE_cqlType              = 14
	CqlParserRULE_primaryKeyKeywords   = 15
	CqlParserRULE_ifNotExist           = 16
	CqlParserRULE_cqlNativeType        = 17
	CqlParserRULE_cqlCollectionType    = 18
	CqlParserRULE_wihtTableOptions     = 19
	CqlParserRULE_nonSemicolonToken    = 20
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CqlStatement() ICqlStatementContext
	EOF() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) CqlStatement() ICqlStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlStatementContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(CqlParserEOF, 0)
}

func (s *RootContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(CqlParserSEMICOLON, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *CqlParser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CqlParserRULE_root)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.CqlStatement()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserSEMICOLON {
		{
			p.SetState(43)
			p.Match(CqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(46)
		p.Match(CqlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlStatementContext is an interface to support dynamic dispatch.
type ICqlStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CreateTable() ICreateTableContext
	CreateKeyspace() ICreateKeyspaceContext

	// IsCqlStatementContext differentiates from other interfaces.
	IsCqlStatementContext()
}

type CqlStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlStatementContext() *CqlStatementContext {
	var p = new(CqlStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlStatement
	return p
}

func InitEmptyCqlStatementContext(p *CqlStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlStatement
}

func (*CqlStatementContext) IsCqlStatementContext() {}

func NewCqlStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlStatementContext {
	var p = new(CqlStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlStatement

	return p
}

func (s *CqlStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlStatementContext) CreateTable() ICreateTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTableContext)
}

func (s *CqlStatementContext) CreateKeyspace() ICreateKeyspaceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateKeyspaceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateKeyspaceContext)
}

func (s *CqlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlStatement(s)
	}
}

func (s *CqlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlStatement(s)
	}
}

func (p *CqlParser) CqlStatement() (localctx ICqlStatementContext) {
	localctx = NewCqlStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CqlParserRULE_cqlStatement)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.CreateTable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.CreateKeyspace()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateTableContext is an interface to support dynamic dispatch.
type ICreateTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_CREATE() antlr.TerminalNode
	K_TABLE() antlr.TerminalNode
	TableName() ITableNameContext
	L_PAREN() antlr.TerminalNode
	ColumnDefinitionList() IColumnDefinitionListContext
	R_PAREN() antlr.TerminalNode
	IfNotExist() IIfNotExistContext
	WihtTableOptions() IWihtTableOptionsContext

	// IsCreateTableContext differentiates from other interfaces.
	IsCreateTableContext()
}

type CreateTableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableContext() *CreateTableContext {
	var p = new(CreateTableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_createTable
	return p
}

func InitEmptyCreateTableContext(p *CreateTableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_createTable
}

func (*CreateTableContext) IsCreateTableContext() {}

func NewCreateTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableContext {
	var p = new(CreateTableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_createTable

	return p
}

func (s *CreateTableContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *CreateTableContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLE, 0)
}

func (s *CreateTableContext) TableName() ITableNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableNameContext)
}

func (s *CreateTableContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *CreateTableContext) ColumnDefinitionList() IColumnDefinitionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnDefinitionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionListContext)
}

func (s *CreateTableContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *CreateTableContext) IfNotExist() IIfNotExistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistContext)
}

func (s *CreateTableContext) WihtTableOptions() IWihtTableOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWihtTableOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWihtTableOptionsContext)
}

func (s *CreateTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCreateTable(s)
	}
}

func (s *CreateTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCreateTable(s)
	}
}

func (p *CqlParser) CreateTable() (localctx ICreateTableContext) {
	localctx = NewCreateTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CqlParserRULE_createTable)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(CqlParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Match(CqlParserK_TABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_IF {
		{
			p.SetState(54)
			p.IfNotExist()
		}

	}
	{
		p.SetState(57)
		p.TableName()
	}
	{
		p.SetState(58)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.ColumnDefinitionList()
	}
	{
		p.SetState(60)
		p.Match(CqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_WITH {
		{
			p.SetState(61)
			p.WihtTableOptions()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICreateKeyspaceContext is an interface to support dynamic dispatch.
type ICreateKeyspaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_CREATE() antlr.TerminalNode
	K_KEYSPACE() antlr.TerminalNode
	KeyspaceName() IKeyspaceNameContext
	K_WITH() antlr.TerminalNode
	IfNotExist() IIfNotExistContext
	AllNonSemicolonToken() []INonSemicolonTokenContext
	NonSemicolonToken(i int) INonSemicolonTokenContext

	// IsCreateKeyspaceContext differentiates from other interfaces.
	IsCreateKeyspaceContext()
}

type CreateKeyspaceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateKeyspaceContext() *CreateKeyspaceContext {
	var p = new(CreateKeyspaceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_createKeyspace
	return p
}

func InitEmptyCreateKeyspaceContext(p *CreateKeyspaceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_createKeyspace
}

func (*CreateKeyspaceContext) IsCreateKeyspaceContext() {}

func NewCreateKeyspaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateKeyspaceContext {
	var p = new(CreateKeyspaceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_createKeyspace

	return p
}

func (s *CreateKeyspaceContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateKeyspaceContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *CreateKeyspaceContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEYSPACE, 0)
}

func (s *CreateKeyspaceContext) KeyspaceName() IKeyspaceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyspaceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyspaceNameContext)
}

func (s *CreateKeyspaceContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *CreateKeyspaceContext) IfNotExist() IIfNotExistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistContext)
}

func (s *CreateKeyspaceContext) AllNonSemicolonToken() []INonSemicolonTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INonSemicolonTokenContext); ok {
			len++
		}
	}

	tst := make([]INonSemicolonTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INonSemicolonTokenContext); ok {
			tst[i] = t.(INonSemicolonTokenContext)
			i++
		}
	}

	return tst
}

func (s *CreateKeyspaceContext) NonSemicolonToken(i int) INonSemicolonTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonSemicolonTokenContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonSemicolonTokenContext)
}

func (s *CreateKeyspaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateKeyspaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateKeyspaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCreateKeyspace(s)
	}
}

func (s *CreateKeyspaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCreateKeyspace(s)
	}
}

func (p *CqlParser) CreateKeyspace() (localctx ICreateKeyspaceContext) {
	localctx = NewCreateKeyspaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CqlParserRULE_createKeyspace)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(64)
		p.Match(CqlParserK_CREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(65)
		p.Match(CqlParserK_KEYSPACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_IF {
		{
			p.SetState(66)
			p.IfNotExist()
		}

	}
	{
		p.SetState(69)
		p.KeyspaceName()
	}
	{
		p.SetState(70)
		p.Match(CqlParserK_WITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&79156247265278) != 0 {
		{
			p.SetState(71)
			p.NonSemicolonToken()
		}

		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnDefinitionListContext is an interface to support dynamic dispatch.
type IColumnDefinitionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnDefinition() []IColumnDefinitionContext
	ColumnDefinition(i int) IColumnDefinitionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	PrimaryKeyClause() IPrimaryKeyClauseContext

	// IsColumnDefinitionListContext differentiates from other interfaces.
	IsColumnDefinitionListContext()
}

type ColumnDefinitionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionListContext() *ColumnDefinitionListContext {
	var p = new(ColumnDefinitionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinitionList
	return p
}

func InitEmptyColumnDefinitionListContext(p *ColumnDefinitionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinitionList
}

func (*ColumnDefinitionListContext) IsColumnDefinitionListContext() {}

func NewColumnDefinitionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionListContext {
	var p = new(ColumnDefinitionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_columnDefinitionList

	return p
}

func (s *ColumnDefinitionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionListContext) AllColumnDefinition() []IColumnDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IColumnDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnDefinitionContext); ok {
			tst[i] = t.(IColumnDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ColumnDefinitionListContext) ColumnDefinition(i int) IColumnDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *ColumnDefinitionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CqlParserCOMMA)
}

func (s *ColumnDefinitionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, i)
}

func (s *ColumnDefinitionListContext) PrimaryKeyClause() IPrimaryKeyClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyClauseContext)
}

func (s *ColumnDefinitionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterColumnDefinitionList(s)
	}
}

func (s *ColumnDefinitionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitColumnDefinitionList(s)
	}
}

func (p *CqlParser) ColumnDefinitionList() (localctx IColumnDefinitionListContext) {
	localctx = NewColumnDefinitionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CqlParserRULE_columnDefinitionList)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.ColumnDefinition()
	}
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(78)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(79)
				p.ColumnDefinition()
			}

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(85)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(86)
			p.PrimaryKeyClause()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnDefinitionContext is an interface to support dynamic dispatch.
type IColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ColumnName() IColumnNameContext
	CqlType() ICqlTypeContext
	PrimaryKeyKeywords() IPrimaryKeyKeywordsContext

	// IsColumnDefinitionContext differentiates from other interfaces.
	IsColumnDefinitionContext()
}

type ColumnDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionContext() *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinition
	return p
}

func InitEmptyColumnDefinitionContext(p *ColumnDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnDefinition
}

func (*ColumnDefinitionContext) IsColumnDefinitionContext() {}

func NewColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_columnDefinition

	return p
}

func (s *ColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionContext) ColumnName() IColumnNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnDefinitionContext) CqlType() ICqlTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlTypeContext)
}

func (s *ColumnDefinitionContext) PrimaryKeyKeywords() IPrimaryKeyKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyKeywordsContext)
}

func (s *ColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterColumnDefinition(s)
	}
}

func (s *ColumnDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitColumnDefinition(s)
	}
}

func (p *CqlParser) ColumnDefinition() (localctx IColumnDefinitionContext) {
	localctx = NewColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CqlParserRULE_columnDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.ColumnName()
	}
	{
		p.SetState(90)
		p.CqlType()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserK_PRIMARY {
		{
			p.SetState(91)
			p.PrimaryKeyKeywords()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryKeyClauseContext is an interface to support dynamic dispatch.
type IPrimaryKeyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryKeyKeywords() IPrimaryKeyKeywordsContext
	L_PAREN() antlr.TerminalNode
	PrimaryKey() IPrimaryKeyContext
	R_PAREN() antlr.TerminalNode

	// IsPrimaryKeyClauseContext differentiates from other interfaces.
	IsPrimaryKeyClauseContext()
}

type PrimaryKeyClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyClauseContext() *PrimaryKeyClauseContext {
	var p = new(PrimaryKeyClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyClause
	return p
}

func InitEmptyPrimaryKeyClauseContext(p *PrimaryKeyClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyClause
}

func (*PrimaryKeyClauseContext) IsPrimaryKeyClauseContext() {}

func NewPrimaryKeyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyClauseContext {
	var p = new(PrimaryKeyClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_primaryKeyClause

	return p
}

func (s *PrimaryKeyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyClauseContext) PrimaryKeyKeywords() IPrimaryKeyKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyKeywordsContext)
}

func (s *PrimaryKeyClauseContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *PrimaryKeyClauseContext) PrimaryKey() IPrimaryKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryKeyContext)
}

func (s *PrimaryKeyClauseContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *PrimaryKeyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPrimaryKeyClause(s)
	}
}

func (s *PrimaryKeyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPrimaryKeyClause(s)
	}
}

func (p *CqlParser) PrimaryKeyClause() (localctx IPrimaryKeyClauseContext) {
	localctx = NewPrimaryKeyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CqlParserRULE_primaryKeyClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.PrimaryKeyKeywords()
	}
	{
		p.SetState(95)
		p.Match(CqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(96)
		p.PrimaryKey()
	}
	{
		p.SetState(97)
		p.Match(CqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryKeyContext is an interface to support dynamic dispatch.
type IPrimaryKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PartitionKey() IPartitionKeyContext
	COMMA() antlr.TerminalNode
	ClusteringColumns() IClusteringColumnsContext

	// IsPrimaryKeyContext differentiates from other interfaces.
	IsPrimaryKeyContext()
}

type PrimaryKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyContext() *PrimaryKeyContext {
	var p = new(PrimaryKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKey
	return p
}

func InitEmptyPrimaryKeyContext(p *PrimaryKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKey
}

func (*PrimaryKeyContext) IsPrimaryKeyContext() {}

func NewPrimaryKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyContext {
	var p = new(PrimaryKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_primaryKey

	return p
}

func (s *PrimaryKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyContext) PartitionKey() IPartitionKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPartitionKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPartitionKeyContext)
}

func (s *PrimaryKeyContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, 0)
}

func (s *PrimaryKeyContext) ClusteringColumns() IClusteringColumnsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClusteringColumnsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClusteringColumnsContext)
}

func (s *PrimaryKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPrimaryKey(s)
	}
}

func (s *PrimaryKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPrimaryKey(s)
	}
}

func (p *CqlParser) PrimaryKey() (localctx IPrimaryKeyContext) {
	localctx = NewPrimaryKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CqlParserRULE_primaryKey)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.PartitionKey()
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CqlParserCOMMA {
		{
			p.SetState(100)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.ClusteringColumns()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPartitionKeyContext is an interface to support dynamic dispatch.
type IPartitionKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnName() []IColumnNameContext
	ColumnName(i int) IColumnNameContext
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPartitionKeyContext differentiates from other interfaces.
	IsPartitionKeyContext()
}

type PartitionKeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartitionKeyContext() *PartitionKeyContext {
	var p = new(PartitionKeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_partitionKey
	return p
}

func InitEmptyPartitionKeyContext(p *PartitionKeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_partitionKey
}

func (*PartitionKeyContext) IsPartitionKeyContext() {}

func NewPartitionKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartitionKeyContext {
	var p = new(PartitionKeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_partitionKey

	return p
}

func (s *PartitionKeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PartitionKeyContext) AllColumnName() []IColumnNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnNameContext); ok {
			len++
		}
	}

	tst := make([]IColumnNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnNameContext); ok {
			tst[i] = t.(IColumnNameContext)
			i++
		}
	}

	return tst
}

func (s *PartitionKeyContext) ColumnName(i int) IColumnNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *PartitionKeyContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *PartitionKeyContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *PartitionKeyContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CqlParserCOMMA)
}

func (s *PartitionKeyContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, i)
}

func (s *PartitionKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartitionKeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartitionKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPartitionKey(s)
	}
}

func (s *PartitionKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPartitionKey(s)
	}
}

func (p *CqlParser) PartitionKey() (localctx IPartitionKeyContext) {
	localctx = NewPartitionKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CqlParserRULE_partitionKey)
	var _la int

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.ColumnName()
		}

	case CqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.Match(CqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.ColumnName()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == CqlParserCOMMA {
			{
				p.SetState(107)
				p.Match(CqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(108)
				p.ColumnName()
			}

			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(114)
			p.Match(CqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClusteringColumnsContext is an interface to support dynamic dispatch.
type IClusteringColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumnName() []IColumnNameContext
	ColumnName(i int) IColumnNameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsClusteringColumnsContext differentiates from other interfaces.
	IsClusteringColumnsContext()
}

type ClusteringColumnsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClusteringColumnsContext() *ClusteringColumnsContext {
	var p = new(ClusteringColumnsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_clusteringColumns
	return p
}

func InitEmptyClusteringColumnsContext(p *ClusteringColumnsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_clusteringColumns
}

func (*ClusteringColumnsContext) IsClusteringColumnsContext() {}

func NewClusteringColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClusteringColumnsContext {
	var p = new(ClusteringColumnsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_clusteringColumns

	return p
}

func (s *ClusteringColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ClusteringColumnsContext) AllColumnName() []IColumnNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnNameContext); ok {
			len++
		}
	}

	tst := make([]IColumnNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnNameContext); ok {
			tst[i] = t.(IColumnNameContext)
			i++
		}
	}

	return tst
}

func (s *ClusteringColumnsContext) ColumnName(i int) IColumnNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ClusteringColumnsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CqlParserCOMMA)
}

func (s *ClusteringColumnsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, i)
}

func (s *ClusteringColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClusteringColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClusteringColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterClusteringColumns(s)
	}
}

func (s *ClusteringColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitClusteringColumns(s)
	}
}

func (p *CqlParser) ClusteringColumns() (localctx IClusteringColumnsContext) {
	localctx = NewClusteringColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CqlParserRULE_clusteringColumns)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.ColumnName()
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CqlParserCOMMA {
		{
			p.SetState(119)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.ColumnName()
		}

		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnName
	return p
}

func InitEmptyColumnNameContext(p *ColumnNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_columnName
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (p *CqlParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CqlParserRULE_columnName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(CqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITableNameContext is an interface to support dynamic dispatch.
type ITableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	KeyspaceName() IKeyspaceNameContext
	DOT() antlr.TerminalNode

	// IsTableNameContext differentiates from other interfaces.
	IsTableNameContext()
}

type TableNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableNameContext() *TableNameContext {
	var p = new(TableNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_tableName
	return p
}

func InitEmptyTableNameContext(p *TableNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_tableName
}

func (*TableNameContext) IsTableNameContext() {}

func NewTableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableNameContext {
	var p = new(TableNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_tableName

	return p
}

func (s *TableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *TableNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
}

func (s *TableNameContext) KeyspaceName() IKeyspaceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyspaceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyspaceNameContext)
}

func (s *TableNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(CqlParserDOT, 0)
}

func (s *TableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterTableName(s)
	}
}

func (s *TableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitTableName(s)
	}
}

func (p *CqlParser) TableName() (localctx ITableNameContext) {
	localctx = NewTableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CqlParserRULE_tableName)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(128)
			p.KeyspaceName()
		}
		{
			p.SetState(129)
			p.Match(CqlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(133)
		p.Match(CqlParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyspaceNameContext is an interface to support dynamic dispatch.
type IKeyspaceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GeneralIdentifier() IGeneralIdentifierContext

	// IsKeyspaceNameContext differentiates from other interfaces.
	IsKeyspaceNameContext()
}

type KeyspaceNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyspaceNameContext() *KeyspaceNameContext {
	var p = new(KeyspaceNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_keyspaceName
	return p
}

func InitEmptyKeyspaceNameContext(p *KeyspaceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_keyspaceName
}

func (*KeyspaceNameContext) IsKeyspaceNameContext() {}

func NewKeyspaceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyspaceNameContext {
	var p = new(KeyspaceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_keyspaceName

	return p
}

func (s *KeyspaceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyspaceNameContext) GeneralIdentifier() IGeneralIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeneralIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeneralIdentifierContext)
}

func (s *KeyspaceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyspaceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyspaceNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterKeyspaceName(s)
	}
}

func (s *KeyspaceNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitKeyspaceName(s)
	}
}

func (p *CqlParser) KeyspaceName() (localctx IKeyspaceNameContext) {
	localctx = NewKeyspaceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CqlParserRULE_keyspaceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.GeneralIdentifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGeneralIdentifierContext is an interface to support dynamic dispatch.
type IGeneralIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsGeneralIdentifierContext differentiates from other interfaces.
	IsGeneralIdentifierContext()
}

type GeneralIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeneralIdentifierContext() *GeneralIdentifierContext {
	var p = new(GeneralIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_generalIdentifier
	return p
}

func InitEmptyGeneralIdentifierContext(p *GeneralIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_generalIdentifier
}

func (*GeneralIdentifierContext) IsGeneralIdentifierContext() {}

func NewGeneralIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeneralIdentifierContext {
	var p = new(GeneralIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_generalIdentifier

	return p
}

func (s *GeneralIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *GeneralIdentifierContext) IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER_WITH_HYPHEN, 0)
}

func (s *GeneralIdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
}

func (s *GeneralIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeneralIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeneralIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterGeneralIdentifier(s)
	}
}

func (s *GeneralIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitGeneralIdentifier(s)
	}
}

func (p *CqlParser) GeneralIdentifier() (localctx IGeneralIdentifierContext) {
	localctx = NewGeneralIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CqlParserRULE_generalIdentifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(137)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CqlParserIDENTIFIER || _la == CqlParserIDENTIFIER_WITH_HYPHEN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlTypeContext is an interface to support dynamic dispatch.
type ICqlTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CqlNativeType() ICqlNativeTypeContext
	CqlCollectionType() ICqlCollectionTypeContext

	// IsCqlTypeContext differentiates from other interfaces.
	IsCqlTypeContext()
}

type CqlTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlTypeContext() *CqlTypeContext {
	var p = new(CqlTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlType
	return p
}

func InitEmptyCqlTypeContext(p *CqlTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlType
}

func (*CqlTypeContext) IsCqlTypeContext() {}

func NewCqlTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlTypeContext {
	var p = new(CqlTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlType

	return p
}

func (s *CqlTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlTypeContext) CqlNativeType() ICqlNativeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlNativeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlNativeTypeContext)
}

func (s *CqlTypeContext) CqlCollectionType() ICqlCollectionTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlCollectionTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlCollectionTypeContext)
}

func (s *CqlTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlType(s)
	}
}

func (s *CqlTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlType(s)
	}
}

func (p *CqlParser) CqlType() (localctx ICqlTypeContext) {
	localctx = NewCqlTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CqlParserRULE_cqlType)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_ASCII, CqlParserK_BIGINT, CqlParserK_BLOB, CqlParserK_BOOLEAN, CqlParserK_COUNTER, CqlParserK_DATE, CqlParserK_DECIMAL, CqlParserK_DOUBLE, CqlParserK_FLOAT, CqlParserK_INET, CqlParserK_INT, CqlParserK_SMALLINT, CqlParserK_TEXT, CqlParserK_TIME, CqlParserK_TIMESTAMP, CqlParserK_TIMEUUID, CqlParserK_TINYINT, CqlParserK_UUID, CqlParserK_VARCHAR, CqlParserK_VARINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.CqlNativeType()
		}

	case CqlParserK_MAP, CqlParserK_SET, CqlParserK_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)
			p.CqlCollectionType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryKeyKeywordsContext is an interface to support dynamic dispatch.
type IPrimaryKeyKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_PRIMARY() antlr.TerminalNode
	K_KEY() antlr.TerminalNode

	// IsPrimaryKeyKeywordsContext differentiates from other interfaces.
	IsPrimaryKeyKeywordsContext()
}

type PrimaryKeyKeywordsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryKeyKeywordsContext() *PrimaryKeyKeywordsContext {
	var p = new(PrimaryKeyKeywordsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyKeywords
	return p
}

func InitEmptyPrimaryKeyKeywordsContext(p *PrimaryKeyKeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_primaryKeyKeywords
}

func (*PrimaryKeyKeywordsContext) IsPrimaryKeyKeywordsContext() {}

func NewPrimaryKeyKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryKeyKeywordsContext {
	var p = new(PrimaryKeyKeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_primaryKeyKeywords

	return p
}

func (s *PrimaryKeyKeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryKeyKeywordsContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PRIMARY, 0)
}

func (s *PrimaryKeyKeywordsContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEY, 0)
}

func (s *PrimaryKeyKeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryKeyKeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryKeyKeywordsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterPrimaryKeyKeywords(s)
	}
}

func (s *PrimaryKeyKeywordsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitPrimaryKeyKeywords(s)
	}
}

func (p *CqlParser) PrimaryKeyKeywords() (localctx IPrimaryKeyKeywordsContext) {
	localctx = NewPrimaryKeyKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CqlParserRULE_primaryKeyKeywords)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(CqlParserK_PRIMARY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(CqlParserK_KEY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfNotExistContext is an interface to support dynamic dispatch.
type IIfNotExistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_IF() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode

	// IsIfNotExistContext differentiates from other interfaces.
	IsIfNotExistContext()
}

type IfNotExistContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistContext() *IfNotExistContext {
	var p = new(IfNotExistContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifNotExist
	return p
}

func InitEmptyIfNotExistContext(p *IfNotExistContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_ifNotExist
}

func (*IfNotExistContext) IsIfNotExistContext() {}

func NewIfNotExistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistContext {
	var p = new(IfNotExistContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_ifNotExist

	return p
}

func (s *IfNotExistContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *IfNotExistContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOT, 0)
}

func (s *IfNotExistContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *IfNotExistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterIfNotExist(s)
	}
}

func (s *IfNotExistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitIfNotExist(s)
	}
}

func (p *CqlParser) IfNotExist() (localctx IIfNotExistContext) {
	localctx = NewIfNotExistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CqlParserRULE_ifNotExist)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(CqlParserK_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(147)
		p.Match(CqlParserK_NOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Match(CqlParserK_EXISTS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlNativeTypeContext is an interface to support dynamic dispatch.
type ICqlNativeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_ASCII() antlr.TerminalNode
	K_BIGINT() antlr.TerminalNode
	K_BLOB() antlr.TerminalNode
	K_BOOLEAN() antlr.TerminalNode
	K_COUNTER() antlr.TerminalNode
	K_DATE() antlr.TerminalNode
	K_DECIMAL() antlr.TerminalNode
	K_DOUBLE() antlr.TerminalNode
	K_FLOAT() antlr.TerminalNode
	K_INET() antlr.TerminalNode
	K_INT() antlr.TerminalNode
	K_SMALLINT() antlr.TerminalNode
	K_TEXT() antlr.TerminalNode
	K_TIME() antlr.TerminalNode
	K_TIMESTAMP() antlr.TerminalNode
	K_TIMEUUID() antlr.TerminalNode
	K_TINYINT() antlr.TerminalNode
	K_UUID() antlr.TerminalNode
	K_VARCHAR() antlr.TerminalNode
	K_VARINT() antlr.TerminalNode

	// IsCqlNativeTypeContext differentiates from other interfaces.
	IsCqlNativeTypeContext()
}

type CqlNativeTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlNativeTypeContext() *CqlNativeTypeContext {
	var p = new(CqlNativeTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlNativeType
	return p
}

func InitEmptyCqlNativeTypeContext(p *CqlNativeTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlNativeType
}

func (*CqlNativeTypeContext) IsCqlNativeTypeContext() {}

func NewCqlNativeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlNativeTypeContext {
	var p = new(CqlNativeTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlNativeType

	return p
}

func (s *CqlNativeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlNativeTypeContext) K_ASCII() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASCII, 0)
}

func (s *CqlNativeTypeContext) K_BIGINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BIGINT, 0)
}

func (s *CqlNativeTypeContext) K_BLOB() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BLOB, 0)
}

func (s *CqlNativeTypeContext) K_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BOOLEAN, 0)
}

func (s *CqlNativeTypeContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COUNTER, 0)
}

func (s *CqlNativeTypeContext) K_DATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DATE, 0)
}

func (s *CqlNativeTypeContext) K_DECIMAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DECIMAL, 0)
}

func (s *CqlNativeTypeContext) K_DOUBLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DOUBLE, 0)
}

func (s *CqlNativeTypeContext) K_FLOAT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FLOAT, 0)
}

func (s *CqlNativeTypeContext) K_INET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INET, 0)
}

func (s *CqlNativeTypeContext) K_INT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INT, 0)
}

func (s *CqlNativeTypeContext) K_SMALLINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SMALLINT, 0)
}

func (s *CqlNativeTypeContext) K_TEXT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TEXT, 0)
}

func (s *CqlNativeTypeContext) K_TIME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIME, 0)
}

func (s *CqlNativeTypeContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMESTAMP, 0)
}

func (s *CqlNativeTypeContext) K_TIMEUUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMEUUID, 0)
}

func (s *CqlNativeTypeContext) K_TINYINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TINYINT, 0)
}

func (s *CqlNativeTypeContext) K_UUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UUID, 0)
}

func (s *CqlNativeTypeContext) K_VARCHAR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARCHAR, 0)
}

func (s *CqlNativeTypeContext) K_VARINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARINT, 0)
}

func (s *CqlNativeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlNativeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlNativeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlNativeType(s)
	}
}

func (s *CqlNativeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlNativeType(s)
	}
}

func (p *CqlParser) CqlNativeType() (localctx ICqlNativeTypeContext) {
	localctx = NewCqlNativeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CqlParserRULE_cqlNativeType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1073740800) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICqlCollectionTypeContext is an interface to support dynamic dispatch.
type ICqlCollectionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_MAP() antlr.TerminalNode
	L_ANGLE_BRACKET() antlr.TerminalNode
	AllCqlNativeType() []ICqlNativeTypeContext
	CqlNativeType(i int) ICqlNativeTypeContext
	COMMA() antlr.TerminalNode
	R_ANGLE_BRACKET() antlr.TerminalNode
	K_SET() antlr.TerminalNode
	K_LIST() antlr.TerminalNode

	// IsCqlCollectionTypeContext differentiates from other interfaces.
	IsCqlCollectionTypeContext()
}

type CqlCollectionTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCqlCollectionTypeContext() *CqlCollectionTypeContext {
	var p = new(CqlCollectionTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlCollectionType
	return p
}

func InitEmptyCqlCollectionTypeContext(p *CqlCollectionTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_cqlCollectionType
}

func (*CqlCollectionTypeContext) IsCqlCollectionTypeContext() {}

func NewCqlCollectionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CqlCollectionTypeContext {
	var p = new(CqlCollectionTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_cqlCollectionType

	return p
}

func (s *CqlCollectionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *CqlCollectionTypeContext) K_MAP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MAP, 0)
}

func (s *CqlCollectionTypeContext) L_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserL_ANGLE_BRACKET, 0)
}

func (s *CqlCollectionTypeContext) AllCqlNativeType() []ICqlNativeTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICqlNativeTypeContext); ok {
			len++
		}
	}

	tst := make([]ICqlNativeTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICqlNativeTypeContext); ok {
			tst[i] = t.(ICqlNativeTypeContext)
			i++
		}
	}

	return tst
}

func (s *CqlCollectionTypeContext) CqlNativeType(i int) ICqlNativeTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICqlNativeTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICqlNativeTypeContext)
}

func (s *CqlCollectionTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, 0)
}

func (s *CqlCollectionTypeContext) R_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserR_ANGLE_BRACKET, 0)
}

func (s *CqlCollectionTypeContext) K_SET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SET, 0)
}

func (s *CqlCollectionTypeContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIST, 0)
}

func (s *CqlCollectionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CqlCollectionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CqlCollectionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterCqlCollectionType(s)
	}
}

func (s *CqlCollectionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitCqlCollectionType(s)
	}
}

func (p *CqlParser) CqlCollectionType() (localctx ICqlCollectionTypeContext) {
	localctx = NewCqlCollectionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CqlParserRULE_cqlCollectionType)
	var _la int

	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CqlParserK_MAP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(CqlParserK_MAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.CqlNativeType()
		}
		{
			p.SetState(155)
			p.Match(CqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.CqlNativeType()
		}
		{
			p.SetState(157)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CqlParserK_SET, CqlParserK_LIST:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			_la = p.GetTokenStream().LA(1)

			if !(_la == CqlParserK_SET || _la == CqlParserK_LIST) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(160)
			p.Match(CqlParserL_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.CqlNativeType()
		}
		{
			p.SetState(162)
			p.Match(CqlParserR_ANGLE_BRACKET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWihtTableOptionsContext is an interface to support dynamic dispatch.
type IWihtTableOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_WITH() antlr.TerminalNode
	AllNonSemicolonToken() []INonSemicolonTokenContext
	NonSemicolonToken(i int) INonSemicolonTokenContext

	// IsWihtTableOptionsContext differentiates from other interfaces.
	IsWihtTableOptionsContext()
}

type WihtTableOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWihtTableOptionsContext() *WihtTableOptionsContext {
	var p = new(WihtTableOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_wihtTableOptions
	return p
}

func InitEmptyWihtTableOptionsContext(p *WihtTableOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_wihtTableOptions
}

func (*WihtTableOptionsContext) IsWihtTableOptionsContext() {}

func NewWihtTableOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WihtTableOptionsContext {
	var p = new(WihtTableOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_wihtTableOptions

	return p
}

func (s *WihtTableOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *WihtTableOptionsContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *WihtTableOptionsContext) AllNonSemicolonToken() []INonSemicolonTokenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INonSemicolonTokenContext); ok {
			len++
		}
	}

	tst := make([]INonSemicolonTokenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INonSemicolonTokenContext); ok {
			tst[i] = t.(INonSemicolonTokenContext)
			i++
		}
	}

	return tst
}

func (s *WihtTableOptionsContext) NonSemicolonToken(i int) INonSemicolonTokenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INonSemicolonTokenContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INonSemicolonTokenContext)
}

func (s *WihtTableOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WihtTableOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WihtTableOptionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterWihtTableOptions(s)
	}
}

func (s *WihtTableOptionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitWihtTableOptions(s)
	}
}

func (p *CqlParser) WihtTableOptions() (localctx IWihtTableOptionsContext) {
	localctx = NewWihtTableOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CqlParserRULE_wihtTableOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(CqlParserK_WITH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&79156247265278) != 0 {
		{
			p.SetState(167)
			p.NonSemicolonToken()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INonSemicolonTokenContext is an interface to support dynamic dispatch.
type INonSemicolonTokenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_CREATE() antlr.TerminalNode
	K_TABLE() antlr.TerminalNode
	K_IF() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_EXISTS() antlr.TerminalNode
	K_PRIMARY() antlr.TerminalNode
	K_KEY() antlr.TerminalNode
	K_WITH() antlr.TerminalNode
	K_KEYSPACE() antlr.TerminalNode
	K_ASCII() antlr.TerminalNode
	K_BIGINT() antlr.TerminalNode
	K_BLOB() antlr.TerminalNode
	K_BOOLEAN() antlr.TerminalNode
	K_COUNTER() antlr.TerminalNode
	K_DATE() antlr.TerminalNode
	K_DECIMAL() antlr.TerminalNode
	K_DOUBLE() antlr.TerminalNode
	K_FLOAT() antlr.TerminalNode
	K_INET() antlr.TerminalNode
	K_INT() antlr.TerminalNode
	K_SMALLINT() antlr.TerminalNode
	K_TEXT() antlr.TerminalNode
	K_TIME() antlr.TerminalNode
	K_TIMESTAMP() antlr.TerminalNode
	K_TIMEUUID() antlr.TerminalNode
	K_TINYINT() antlr.TerminalNode
	K_UUID() antlr.TerminalNode
	K_VARCHAR() antlr.TerminalNode
	K_VARINT() antlr.TerminalNode
	K_MAP() antlr.TerminalNode
	K_SET() antlr.TerminalNode
	K_LIST() antlr.TerminalNode
	DQUOTE() antlr.TerminalNode
	DOT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	L_ANGLE_BRACKET() antlr.TerminalNode
	R_ANGLE_BRACKET() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode
	UNKNOWN() antlr.TerminalNode

	// IsNonSemicolonTokenContext differentiates from other interfaces.
	IsNonSemicolonTokenContext()
}

type NonSemicolonTokenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonSemicolonTokenContext() *NonSemicolonTokenContext {
	var p = new(NonSemicolonTokenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_nonSemicolonToken
	return p
}

func InitEmptyNonSemicolonTokenContext(p *NonSemicolonTokenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CqlParserRULE_nonSemicolonToken
}

func (*NonSemicolonTokenContext) IsNonSemicolonTokenContext() {}

func NewNonSemicolonTokenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonSemicolonTokenContext {
	var p = new(NonSemicolonTokenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CqlParserRULE_nonSemicolonToken

	return p
}

func (s *NonSemicolonTokenContext) GetParser() antlr.Parser { return s.parser }

func (s *NonSemicolonTokenContext) K_CREATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_CREATE, 0)
}

func (s *NonSemicolonTokenContext) K_TABLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TABLE, 0)
}

func (s *NonSemicolonTokenContext) K_IF() antlr.TerminalNode {
	return s.GetToken(CqlParserK_IF, 0)
}

func (s *NonSemicolonTokenContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_NOT, 0)
}

func (s *NonSemicolonTokenContext) K_EXISTS() antlr.TerminalNode {
	return s.GetToken(CqlParserK_EXISTS, 0)
}

func (s *NonSemicolonTokenContext) K_PRIMARY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_PRIMARY, 0)
}

func (s *NonSemicolonTokenContext) K_KEY() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEY, 0)
}

func (s *NonSemicolonTokenContext) K_WITH() antlr.TerminalNode {
	return s.GetToken(CqlParserK_WITH, 0)
}

func (s *NonSemicolonTokenContext) K_KEYSPACE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_KEYSPACE, 0)
}

func (s *NonSemicolonTokenContext) K_ASCII() antlr.TerminalNode {
	return s.GetToken(CqlParserK_ASCII, 0)
}

func (s *NonSemicolonTokenContext) K_BIGINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BIGINT, 0)
}

func (s *NonSemicolonTokenContext) K_BLOB() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BLOB, 0)
}

func (s *NonSemicolonTokenContext) K_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CqlParserK_BOOLEAN, 0)
}

func (s *NonSemicolonTokenContext) K_COUNTER() antlr.TerminalNode {
	return s.GetToken(CqlParserK_COUNTER, 0)
}

func (s *NonSemicolonTokenContext) K_DATE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DATE, 0)
}

func (s *NonSemicolonTokenContext) K_DECIMAL() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DECIMAL, 0)
}

func (s *NonSemicolonTokenContext) K_DOUBLE() antlr.TerminalNode {
	return s.GetToken(CqlParserK_DOUBLE, 0)
}

func (s *NonSemicolonTokenContext) K_FLOAT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_FLOAT, 0)
}

func (s *NonSemicolonTokenContext) K_INET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INET, 0)
}

func (s *NonSemicolonTokenContext) K_INT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_INT, 0)
}

func (s *NonSemicolonTokenContext) K_SMALLINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SMALLINT, 0)
}

func (s *NonSemicolonTokenContext) K_TEXT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TEXT, 0)
}

func (s *NonSemicolonTokenContext) K_TIME() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIME, 0)
}

func (s *NonSemicolonTokenContext) K_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMESTAMP, 0)
}

func (s *NonSemicolonTokenContext) K_TIMEUUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TIMEUUID, 0)
}

func (s *NonSemicolonTokenContext) K_TINYINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_TINYINT, 0)
}

func (s *NonSemicolonTokenContext) K_UUID() antlr.TerminalNode {
	return s.GetToken(CqlParserK_UUID, 0)
}

func (s *NonSemicolonTokenContext) K_VARCHAR() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARCHAR, 0)
}

func (s *NonSemicolonTokenContext) K_VARINT() antlr.TerminalNode {
	return s.GetToken(CqlParserK_VARINT, 0)
}

func (s *NonSemicolonTokenContext) K_MAP() antlr.TerminalNode {
	return s.GetToken(CqlParserK_MAP, 0)
}

func (s *NonSemicolonTokenContext) K_SET() antlr.TerminalNode {
	return s.GetToken(CqlParserK_SET, 0)
}

func (s *NonSemicolonTokenContext) K_LIST() antlr.TerminalNode {
	return s.GetToken(CqlParserK_LIST, 0)
}

func (s *NonSemicolonTokenContext) DQUOTE() antlr.TerminalNode {
	return s.GetToken(CqlParserDQUOTE, 0)
}

func (s *NonSemicolonTokenContext) DOT() antlr.TerminalNode {
	return s.GetToken(CqlParserDOT, 0)
}

func (s *NonSemicolonTokenContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CqlParserCOMMA, 0)
}

func (s *NonSemicolonTokenContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserL_PAREN, 0)
}

func (s *NonSemicolonTokenContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CqlParserR_PAREN, 0)
}

func (s *NonSemicolonTokenContext) L_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserL_ANGLE_BRACKET, 0)
}

func (s *NonSemicolonTokenContext) R_ANGLE_BRACKET() antlr.TerminalNode {
	return s.GetToken(CqlParserR_ANGLE_BRACKET, 0)
}

func (s *NonSemicolonTokenContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER, 0)
}

func (s *NonSemicolonTokenContext) IDENTIFIER_WITH_HYPHEN() antlr.TerminalNode {
	return s.GetToken(CqlParserIDENTIFIER_WITH_HYPHEN, 0)
}

func (s *NonSemicolonTokenContext) UNKNOWN() antlr.TerminalNode {
	return s.GetToken(CqlParserUNKNOWN, 0)
}

func (s *NonSemicolonTokenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonSemicolonTokenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonSemicolonTokenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.EnterNonSemicolonToken(s)
	}
}

func (s *NonSemicolonTokenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CqlParserListener); ok {
		listenerT.ExitNonSemicolonToken(s)
	}
}

func (p *CqlParser) NonSemicolonToken() (localctx INonSemicolonTokenContext) {
	localctx = NewNonSemicolonTokenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CqlParserRULE_nonSemicolonToken)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&79156247265278) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
