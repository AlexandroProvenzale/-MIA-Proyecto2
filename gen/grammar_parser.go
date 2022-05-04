// Code generated from Grammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grammar

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Proyecto2/Program"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 187,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 6, 2, 32, 10, 2, 13, 2, 14, 2,
	33, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 58,
	10, 3, 3, 4, 3, 4, 6, 4, 62, 10, 4, 13, 4, 14, 4, 63, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 84, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 6, 7, 94, 10, 7, 13, 7, 14, 7, 95, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 124,
	10, 8, 3, 9, 3, 9, 6, 9, 128, 10, 9, 13, 9, 14, 9, 129, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 142, 10, 10,
	3, 11, 3, 11, 6, 11, 146, 10, 11, 13, 11, 14, 11, 147, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 160, 10, 12,
	3, 13, 3, 13, 6, 13, 164, 10, 13, 13, 13, 14, 13, 165, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 5, 14, 182, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 2, 2, 16, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 2, 3, 4, 2, 35, 35, 37, 39,
	2, 197, 2, 31, 3, 2, 2, 2, 4, 57, 3, 2, 2, 2, 6, 59, 3, 2, 2, 2, 8, 83,
	3, 2, 2, 2, 10, 85, 3, 2, 2, 2, 12, 91, 3, 2, 2, 2, 14, 123, 3, 2, 2, 2,
	16, 125, 3, 2, 2, 2, 18, 141, 3, 2, 2, 2, 20, 143, 3, 2, 2, 2, 22, 159,
	3, 2, 2, 2, 24, 161, 3, 2, 2, 2, 26, 181, 3, 2, 2, 2, 28, 183, 3, 2, 2,
	2, 30, 32, 5, 4, 3, 2, 31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 3, 3, 2, 2, 2, 35, 36, 5, 6, 4, 2,
	36, 37, 7, 40, 2, 2, 37, 58, 3, 2, 2, 2, 38, 39, 5, 10, 6, 2, 39, 40, 7,
	40, 2, 2, 40, 58, 3, 2, 2, 2, 41, 42, 5, 12, 7, 2, 42, 43, 7, 40, 2, 2,
	43, 58, 3, 2, 2, 2, 44, 45, 5, 16, 9, 2, 45, 46, 7, 40, 2, 2, 46, 58, 3,
	2, 2, 2, 47, 48, 5, 20, 11, 2, 48, 49, 7, 40, 2, 2, 49, 58, 3, 2, 2, 2,
	50, 51, 5, 24, 13, 2, 51, 52, 7, 40, 2, 2, 52, 58, 3, 2, 2, 2, 53, 54,
	5, 28, 15, 2, 54, 55, 7, 40, 2, 2, 55, 58, 3, 2, 2, 2, 56, 58, 7, 40, 2,
	2, 57, 35, 3, 2, 2, 2, 57, 38, 3, 2, 2, 2, 57, 41, 3, 2, 2, 2, 57, 44,
	3, 2, 2, 2, 57, 47, 3, 2, 2, 2, 57, 50, 3, 2, 2, 2, 57, 53, 3, 2, 2, 2,
	57, 56, 3, 2, 2, 2, 58, 5, 3, 2, 2, 2, 59, 61, 7, 6, 2, 2, 60, 62, 5, 8,
	5, 2, 61, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 8, 4, 1, 2, 66, 7, 3, 2, 2, 2,
	67, 68, 7, 15, 2, 2, 68, 69, 7, 34, 2, 2, 69, 70, 7, 35, 2, 2, 70, 84,
	8, 5, 1, 2, 71, 72, 7, 18, 2, 2, 72, 73, 7, 34, 2, 2, 73, 74, 7, 30, 2,
	2, 74, 84, 8, 5, 1, 2, 75, 76, 7, 16, 2, 2, 76, 77, 7, 34, 2, 2, 77, 78,
	7, 27, 2, 2, 78, 84, 8, 5, 1, 2, 79, 80, 7, 17, 2, 2, 80, 81, 7, 34, 2,
	2, 81, 82, 7, 28, 2, 2, 82, 84, 8, 5, 1, 2, 83, 67, 3, 2, 2, 2, 83, 71,
	3, 2, 2, 2, 83, 75, 3, 2, 2, 2, 83, 79, 3, 2, 2, 2, 84, 9, 3, 2, 2, 2,
	85, 86, 7, 7, 2, 2, 86, 87, 7, 18, 2, 2, 87, 88, 7, 34, 2, 2, 88, 89, 7,
	30, 2, 2, 89, 90, 8, 6, 1, 2, 90, 11, 3, 2, 2, 2, 91, 93, 7, 8, 2, 2, 92,
	94, 5, 14, 8, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 93, 3, 2,
	2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 8, 7, 1, 2, 98, 13,
	3, 2, 2, 2, 99, 100, 7, 15, 2, 2, 100, 101, 7, 34, 2, 2, 101, 102, 7, 35,
	2, 2, 102, 124, 8, 8, 1, 2, 103, 104, 7, 18, 2, 2, 104, 105, 7, 34, 2,
	2, 105, 106, 7, 30, 2, 2, 106, 124, 8, 8, 1, 2, 107, 108, 7, 16, 2, 2,
	108, 109, 7, 34, 2, 2, 109, 110, 7, 27, 2, 2, 110, 124, 8, 8, 1, 2, 111,
	112, 7, 17, 2, 2, 112, 113, 7, 34, 2, 2, 113, 114, 7, 28, 2, 2, 114, 124,
	8, 8, 1, 2, 115, 116, 7, 19, 2, 2, 116, 117, 7, 34, 2, 2, 117, 118, 7,
	29, 2, 2, 118, 124, 8, 8, 1, 2, 119, 120, 7, 21, 2, 2, 120, 121, 7, 34,
	2, 2, 121, 122, 7, 37, 2, 2, 122, 124, 8, 8, 1, 2, 123, 99, 3, 2, 2, 2,
	123, 103, 3, 2, 2, 2, 123, 107, 3, 2, 2, 2, 123, 111, 3, 2, 2, 2, 123,
	115, 3, 2, 2, 2, 123, 119, 3, 2, 2, 2, 124, 15, 3, 2, 2, 2, 125, 127, 7,
	9, 2, 2, 126, 128, 5, 18, 10, 2, 127, 126, 3, 2, 2, 2, 128, 129, 3, 2,
	2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2,
	131, 132, 8, 9, 1, 2, 132, 17, 3, 2, 2, 2, 133, 134, 7, 18, 2, 2, 134,
	135, 7, 34, 2, 2, 135, 136, 7, 30, 2, 2, 136, 142, 8, 10, 1, 2, 137, 138,
	7, 21, 2, 2, 138, 139, 7, 34, 2, 2, 139, 140, 7, 37, 2, 2, 140, 142, 8,
	10, 1, 2, 141, 133, 3, 2, 2, 2, 141, 137, 3, 2, 2, 2, 142, 19, 3, 2, 2,
	2, 143, 145, 7, 11, 2, 2, 144, 146, 5, 22, 12, 2, 145, 144, 3, 2, 2, 2,
	146, 147, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148,
	149, 3, 2, 2, 2, 149, 150, 8, 11, 1, 2, 150, 21, 3, 2, 2, 2, 151, 152,
	7, 23, 2, 2, 152, 153, 7, 34, 2, 2, 153, 154, 7, 31, 2, 2, 154, 160, 8,
	12, 1, 2, 155, 156, 7, 19, 2, 2, 156, 157, 7, 34, 2, 2, 157, 158, 7, 29,
	2, 2, 158, 160, 8, 12, 1, 2, 159, 151, 3, 2, 2, 2, 159, 155, 3, 2, 2, 2,
	160, 23, 3, 2, 2, 2, 161, 163, 7, 12, 2, 2, 162, 164, 5, 26, 14, 2, 163,
	162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166,
	3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 8, 13, 1, 2, 168, 25, 3, 2,
	2, 2, 169, 170, 7, 25, 2, 2, 170, 171, 7, 34, 2, 2, 171, 172, 9, 2, 2,
	2, 172, 182, 8, 14, 1, 2, 173, 174, 7, 26, 2, 2, 174, 175, 7, 34, 2, 2,
	175, 176, 9, 2, 2, 2, 176, 182, 8, 14, 1, 2, 177, 178, 7, 23, 2, 2, 178,
	179, 7, 34, 2, 2, 179, 180, 7, 31, 2, 2, 180, 182, 8, 14, 1, 2, 181, 169,
	3, 2, 2, 2, 181, 173, 3, 2, 2, 2, 181, 177, 3, 2, 2, 2, 182, 27, 3, 2,
	2, 2, 183, 184, 7, 13, 2, 2, 184, 185, 8, 15, 1, 2, 185, 29, 3, 2, 2, 2,
	14, 33, 57, 63, 83, 95, 123, 129, 141, 147, 159, 165, 181,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "'='",
}
var symbolicNames = []string{
	"", "PAUSES", "EXEC", "REP", "MKDISK", "RMDISK", "FDISK", "MOUNT", "UNMOUNT",
	"MKFS", "LOGIN", "LOGOUT", "MKGRP", "SIZE", "FIT", "UNIT", "PATH", "TYPE",
	"DELETEP", "NAME", "ADD", "ID", "FS", "USR", "PASSW", "E_FIT", "E_UNIT",
	"E_TYPE", "E_PATH", "E_ID", "PATH1", "PATH2", "IGUAL", "ENTERO", "NEGATIVO",
	"IDENTIFICADOR", "COMPLEMENTO", "E_USRS", "NEWLINE", "WHITESPACE", "COMENTARIO",
}

var ruleNames = []string{
	"start", "comando", "mkdisk_f", "mkparam", "rmdisk_f", "fdisk_f", "fdiskparam",
	"mount_f", "mountparam", "mkfs_f", "mkfsparam", "login_f", "loginparam",
	"logout_f",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GrammarParser struct {
	*antlr.BaseParser
}

func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	this := new(GrammarParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

var info_MKDISK Program.InfoMKDisk
var info_FDISK Program.InfoFDisk
var info_MOUNT Program.InfoMount
var info_MKFS Program.InfoMkfs
var info_LOGIN Program.InfoLogin

func initializeMKDISK(MKDISK *Program.InfoMKDisk) {
	MKDISK.Path = ""
	MKDISK.Fit = ""
	MKDISK.Size = 0
	MKDISK.Unit = ""
}

func initializeFDISK(FDISK *Program.InfoFDisk) {
	FDISK.Path = ""
	FDISK.Fit = ""
	FDISK.Size = 0
	FDISK.Unit = ""
	FDISK.Type = ""
	FDISK.Name = ""
}

func initializeMOUNT(MOUNT *Program.InfoMount) {
	MOUNT.Path = ""
	MOUNT.Name = ""
}

func initializeMKFS(MKFS *Program.InfoMkfs) {
	MKFS.Id = ""
	MKFS.Type = ""
}

func initializeLOGIN(LOGIN *Program.InfoLogin) {
	LOGIN.User = ""
	LOGIN.Pass = ""
	LOGIN.Id = ""
}

// GrammarParser tokens.
const (
	GrammarParserEOF           = antlr.TokenEOF
	GrammarParserPAUSES        = 1
	GrammarParserEXEC          = 2
	GrammarParserREP           = 3
	GrammarParserMKDISK        = 4
	GrammarParserRMDISK        = 5
	GrammarParserFDISK         = 6
	GrammarParserMOUNT         = 7
	GrammarParserUNMOUNT       = 8
	GrammarParserMKFS          = 9
	GrammarParserLOGIN         = 10
	GrammarParserLOGOUT        = 11
	GrammarParserMKGRP         = 12
	GrammarParserSIZE          = 13
	GrammarParserFIT           = 14
	GrammarParserUNIT          = 15
	GrammarParserPATH          = 16
	GrammarParserTYPE          = 17
	GrammarParserDELETEP       = 18
	GrammarParserNAME          = 19
	GrammarParserADD           = 20
	GrammarParserID            = 21
	GrammarParserFS            = 22
	GrammarParserUSR           = 23
	GrammarParserPASSW         = 24
	GrammarParserE_FIT         = 25
	GrammarParserE_UNIT        = 26
	GrammarParserE_TYPE        = 27
	GrammarParserE_PATH        = 28
	GrammarParserE_ID          = 29
	GrammarParserPATH1         = 30
	GrammarParserPATH2         = 31
	GrammarParserIGUAL         = 32
	GrammarParserENTERO        = 33
	GrammarParserNEGATIVO      = 34
	GrammarParserIDENTIFICADOR = 35
	GrammarParserCOMPLEMENTO   = 36
	GrammarParserE_USRS        = 37
	GrammarParserNEWLINE       = 38
	GrammarParserWHITESPACE    = 39
	GrammarParserCOMENTARIO    = 40
)

// GrammarParser rules.
const (
	GrammarParserRULE_start      = 0
	GrammarParserRULE_comando    = 1
	GrammarParserRULE_mkdisk_f   = 2
	GrammarParserRULE_mkparam    = 3
	GrammarParserRULE_rmdisk_f   = 4
	GrammarParserRULE_fdisk_f    = 5
	GrammarParserRULE_fdiskparam = 6
	GrammarParserRULE_mount_f    = 7
	GrammarParserRULE_mountparam = 8
	GrammarParserRULE_mkfs_f     = 9
	GrammarParserRULE_mkfsparam  = 10
	GrammarParserRULE_login_f    = 11
	GrammarParserRULE_loginparam = 12
	GrammarParserRULE_logout_f   = 13
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) AllComando() []IComandoContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComandoContext)(nil)).Elem())
	var tst = make([]IComandoContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComandoContext)
		}
	}

	return tst
}

func (s *StartContext) Comando(i int) IComandoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComandoContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComandoContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GrammarParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserMKDISK)|(1<<GrammarParserRMDISK)|(1<<GrammarParserFDISK)|(1<<GrammarParserMOUNT)|(1<<GrammarParserMKFS)|(1<<GrammarParserLOGIN)|(1<<GrammarParserLOGOUT))) != 0) || _la == GrammarParserNEWLINE {
		{
			p.SetState(28)
			p.Comando()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IComandoContext is an interface to support dynamic dispatch.
type IComandoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComandoContext differentiates from other interfaces.
	IsComandoContext()
}

type ComandoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComandoContext() *ComandoContext {
	var p = new(ComandoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_comando
	return p
}

func (*ComandoContext) IsComandoContext() {}

func NewComandoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComandoContext {
	var p = new(ComandoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_comando

	return p
}

func (s *ComandoContext) GetParser() antlr.Parser { return s.parser }

func (s *ComandoContext) Mkdisk_f() IMkdisk_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkdisk_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkdisk_fContext)
}

func (s *ComandoContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GrammarParserNEWLINE, 0)
}

func (s *ComandoContext) Rmdisk_f() IRmdisk_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRmdisk_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRmdisk_fContext)
}

func (s *ComandoContext) Fdisk_f() IFdisk_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFdisk_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFdisk_fContext)
}

func (s *ComandoContext) Mount_f() IMount_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMount_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMount_fContext)
}

func (s *ComandoContext) Mkfs_f() IMkfs_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkfs_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkfs_fContext)
}

func (s *ComandoContext) Login_f() ILogin_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogin_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogin_fContext)
}

func (s *ComandoContext) Logout_f() ILogout_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogout_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogout_fContext)
}

func (s *ComandoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComandoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComandoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterComando(s)
	}
}

func (s *ComandoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitComando(s)
	}
}

func (p *GrammarParser) Comando() (localctx IComandoContext) {
	localctx = NewComandoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_comando)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserMKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Mkdisk_f()
		}
		{
			p.SetState(34)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserRMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)
			p.Rmdisk_f()
		}
		{
			p.SetState(37)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserFDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.Fdisk_f()
		}
		{
			p.SetState(40)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(42)
			p.Mount_f()
		}
		{
			p.SetState(43)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKFS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(45)
			p.Mkfs_f()
		}
		{
			p.SetState(46)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(48)
			p.Login_f()
		}
		{
			p.SetState(49)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(51)
			p.Logout_f()
		}
		{
			p.SetState(52)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserNEWLINE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(54)
			p.Match(GrammarParserNEWLINE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMkdisk_fContext is an interface to support dynamic dispatch.
type IMkdisk_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMkdisk_fContext differentiates from other interfaces.
	IsMkdisk_fContext()
}

type Mkdisk_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdisk_fContext() *Mkdisk_fContext {
	var p = new(Mkdisk_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkdisk_f
	return p
}

func (*Mkdisk_fContext) IsMkdisk_fContext() {}

func NewMkdisk_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkdisk_fContext {
	var p = new(Mkdisk_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkdisk_f

	return p
}

func (s *Mkdisk_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkdisk_fContext) MKDISK() antlr.TerminalNode {
	return s.GetToken(GrammarParserMKDISK, 0)
}

func (s *Mkdisk_fContext) AllMkparam() []IMkparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMkparamContext)(nil)).Elem())
	var tst = make([]IMkparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMkparamContext)
		}
	}

	return tst
}

func (s *Mkdisk_fContext) Mkparam(i int) IMkparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMkparamContext)
}

func (s *Mkdisk_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkdisk_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkdisk_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkdisk_f(s)
	}
}

func (s *Mkdisk_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkdisk_f(s)
	}
}

func (p *GrammarParser) Mkdisk_f() (localctx IMkdisk_fContext) {
	localctx = NewMkdisk_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_mkdisk_f)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(GrammarParserMKDISK)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH))) != 0) {
		{
			p.SetState(58)
			p.Mkparam()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.CreateDisk(info_MKDISK)
	initializeMKDISK(&info_MKDISK)

	return localctx
}

// IMkparamContext is an interface to support dynamic dispatch.
type IMkparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_E_PATH returns the _E_PATH token.
	Get_E_PATH() antlr.Token

	// Get_E_FIT returns the _E_FIT token.
	Get_E_FIT() antlr.Token

	// Get_E_UNIT returns the _E_UNIT token.
	Get_E_UNIT() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_E_PATH sets the _E_PATH token.
	Set_E_PATH(antlr.Token)

	// Set_E_FIT sets the _E_FIT token.
	Set_E_FIT(antlr.Token)

	// Set_E_UNIT sets the _E_UNIT token.
	Set_E_UNIT(antlr.Token)

	// IsMkparamContext differentiates from other interfaces.
	IsMkparamContext()
}

type MkparamContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_ENTERO antlr.Token
	_E_PATH antlr.Token
	_E_FIT  antlr.Token
	_E_UNIT antlr.Token
}

func NewEmptyMkparamContext() *MkparamContext {
	var p = new(MkparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkparam
	return p
}

func (*MkparamContext) IsMkparamContext() {}

func NewMkparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkparamContext {
	var p = new(MkparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkparam

	return p
}

func (s *MkparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkparamContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *MkparamContext) Get_E_PATH() antlr.Token { return s._E_PATH }

func (s *MkparamContext) Get_E_FIT() antlr.Token { return s._E_FIT }

func (s *MkparamContext) Get_E_UNIT() antlr.Token { return s._E_UNIT }

func (s *MkparamContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *MkparamContext) Set_E_PATH(v antlr.Token) { s._E_PATH = v }

func (s *MkparamContext) Set_E_FIT(v antlr.Token) { s._E_FIT = v }

func (s *MkparamContext) Set_E_UNIT(v antlr.Token) { s._E_UNIT = v }

func (s *MkparamContext) SIZE() antlr.TerminalNode {
	return s.GetToken(GrammarParserSIZE, 0)
}

func (s *MkparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *MkparamContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *MkparamContext) PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserPATH, 0)
}

func (s *MkparamContext) E_PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_PATH, 0)
}

func (s *MkparamContext) FIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserFIT, 0)
}

func (s *MkparamContext) E_FIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_FIT, 0)
}

func (s *MkparamContext) UNIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserUNIT, 0)
}

func (s *MkparamContext) E_UNIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_UNIT, 0)
}

func (s *MkparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkparam(s)
	}
}

func (s *MkparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkparam(s)
	}
}

func (p *GrammarParser) Mkparam() (localctx IMkparamContext) {
	localctx = NewMkparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_mkparam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(66)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(67)

			var _m = p.Match(GrammarParserENTERO)

			localctx.(*MkparamContext)._ENTERO = _m
		}
		info_MKDISK.Size = (func() int {
			if localctx.(*MkparamContext).Get_ENTERO() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*MkparamContext).Get_ENTERO().GetText())
				return i
			}
		}())

	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(70)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(71)

			var _m = p.Match(GrammarParserE_PATH)

			localctx.(*MkparamContext)._E_PATH = _m
		}
		info_MKDISK.Path = strings.ReplaceAll((func() string {
			if localctx.(*MkparamContext).Get_E_PATH() == nil {
				return ""
			} else {
				return localctx.(*MkparamContext).Get_E_PATH().GetText()
			}
		}()), "\"", "")

	case GrammarParserFIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(73)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(74)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(75)

			var _m = p.Match(GrammarParserE_FIT)

			localctx.(*MkparamContext)._E_FIT = _m
		}
		info_MKDISK.Fit = (func() string {
			if localctx.(*MkparamContext).Get_E_FIT() == nil {
				return ""
			} else {
				return localctx.(*MkparamContext).Get_E_FIT().GetText()
			}
		}())

	case GrammarParserUNIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(77)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(78)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(79)

			var _m = p.Match(GrammarParserE_UNIT)

			localctx.(*MkparamContext)._E_UNIT = _m
		}
		info_MKDISK.Unit = (func() string {
			if localctx.(*MkparamContext).Get_E_UNIT() == nil {
				return ""
			} else {
				return localctx.(*MkparamContext).Get_E_UNIT().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRmdisk_fContext is an interface to support dynamic dispatch.
type IRmdisk_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_E_PATH returns the _E_PATH token.
	Get_E_PATH() antlr.Token

	// Set_E_PATH sets the _E_PATH token.
	Set_E_PATH(antlr.Token)

	// IsRmdisk_fContext differentiates from other interfaces.
	IsRmdisk_fContext()
}

type Rmdisk_fContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_E_PATH antlr.Token
}

func NewEmptyRmdisk_fContext() *Rmdisk_fContext {
	var p = new(Rmdisk_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_rmdisk_f
	return p
}

func (*Rmdisk_fContext) IsRmdisk_fContext() {}

func NewRmdisk_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rmdisk_fContext {
	var p = new(Rmdisk_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_rmdisk_f

	return p
}

func (s *Rmdisk_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Rmdisk_fContext) Get_E_PATH() antlr.Token { return s._E_PATH }

func (s *Rmdisk_fContext) Set_E_PATH(v antlr.Token) { s._E_PATH = v }

func (s *Rmdisk_fContext) RMDISK() antlr.TerminalNode {
	return s.GetToken(GrammarParserRMDISK, 0)
}

func (s *Rmdisk_fContext) PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserPATH, 0)
}

func (s *Rmdisk_fContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *Rmdisk_fContext) E_PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_PATH, 0)
}

func (s *Rmdisk_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rmdisk_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rmdisk_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRmdisk_f(s)
	}
}

func (s *Rmdisk_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRmdisk_f(s)
	}
}

func (p *GrammarParser) Rmdisk_f() (localctx IRmdisk_fContext) {
	localctx = NewRmdisk_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_rmdisk_f)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(GrammarParserRMDISK)
	}
	{
		p.SetState(84)
		p.Match(GrammarParserPATH)
	}
	{
		p.SetState(85)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(86)

		var _m = p.Match(GrammarParserE_PATH)

		localctx.(*Rmdisk_fContext)._E_PATH = _m
	}
	Program.RemoveDisk(strings.ReplaceAll((func() string {
		if localctx.(*Rmdisk_fContext).Get_E_PATH() == nil {
			return ""
		} else {
			return localctx.(*Rmdisk_fContext).Get_E_PATH().GetText()
		}
	}()), "\"", ""))

	return localctx
}

// IFdisk_fContext is an interface to support dynamic dispatch.
type IFdisk_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFdisk_fContext differentiates from other interfaces.
	IsFdisk_fContext()
}

type Fdisk_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFdisk_fContext() *Fdisk_fContext {
	var p = new(Fdisk_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_fdisk_f
	return p
}

func (*Fdisk_fContext) IsFdisk_fContext() {}

func NewFdisk_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fdisk_fContext {
	var p = new(Fdisk_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_fdisk_f

	return p
}

func (s *Fdisk_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Fdisk_fContext) FDISK() antlr.TerminalNode {
	return s.GetToken(GrammarParserFDISK, 0)
}

func (s *Fdisk_fContext) AllFdiskparam() []IFdiskparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFdiskparamContext)(nil)).Elem())
	var tst = make([]IFdiskparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFdiskparamContext)
		}
	}

	return tst
}

func (s *Fdisk_fContext) Fdiskparam(i int) IFdiskparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFdiskparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFdiskparamContext)
}

func (s *Fdisk_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fdisk_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fdisk_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFdisk_f(s)
	}
}

func (s *Fdisk_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFdisk_f(s)
	}
}

func (p *GrammarParser) Fdisk_f() (localctx IFdisk_fContext) {
	localctx = NewFdisk_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_fdisk_f)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(GrammarParserFDISK)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH)|(1<<GrammarParserTYPE)|(1<<GrammarParserNAME))) != 0) {
		{
			p.SetState(90)
			p.Fdiskparam()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.CrearParticion(info_FDISK)
	initializeFDISK(&info_FDISK)

	return localctx
}

// IFdiskparamContext is an interface to support dynamic dispatch.
type IFdiskparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Get_E_PATH returns the _E_PATH token.
	Get_E_PATH() antlr.Token

	// Get_E_FIT returns the _E_FIT token.
	Get_E_FIT() antlr.Token

	// Get_E_UNIT returns the _E_UNIT token.
	Get_E_UNIT() antlr.Token

	// Get_E_TYPE returns the _E_TYPE token.
	Get_E_TYPE() antlr.Token

	// GetE_name returns the e_name token.
	GetE_name() antlr.Token

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// Set_E_PATH sets the _E_PATH token.
	Set_E_PATH(antlr.Token)

	// Set_E_FIT sets the _E_FIT token.
	Set_E_FIT(antlr.Token)

	// Set_E_UNIT sets the _E_UNIT token.
	Set_E_UNIT(antlr.Token)

	// Set_E_TYPE sets the _E_TYPE token.
	Set_E_TYPE(antlr.Token)

	// SetE_name sets the e_name token.
	SetE_name(antlr.Token)

	// IsFdiskparamContext differentiates from other interfaces.
	IsFdiskparamContext()
}

type FdiskparamContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_ENTERO antlr.Token
	_E_PATH antlr.Token
	_E_FIT  antlr.Token
	_E_UNIT antlr.Token
	_E_TYPE antlr.Token
	e_name  antlr.Token
}

func NewEmptyFdiskparamContext() *FdiskparamContext {
	var p = new(FdiskparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_fdiskparam
	return p
}

func (*FdiskparamContext) IsFdiskparamContext() {}

func NewFdiskparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FdiskparamContext {
	var p = new(FdiskparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_fdiskparam

	return p
}

func (s *FdiskparamContext) GetParser() antlr.Parser { return s.parser }

func (s *FdiskparamContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *FdiskparamContext) Get_E_PATH() antlr.Token { return s._E_PATH }

func (s *FdiskparamContext) Get_E_FIT() antlr.Token { return s._E_FIT }

func (s *FdiskparamContext) Get_E_UNIT() antlr.Token { return s._E_UNIT }

func (s *FdiskparamContext) Get_E_TYPE() antlr.Token { return s._E_TYPE }

func (s *FdiskparamContext) GetE_name() antlr.Token { return s.e_name }

func (s *FdiskparamContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *FdiskparamContext) Set_E_PATH(v antlr.Token) { s._E_PATH = v }

func (s *FdiskparamContext) Set_E_FIT(v antlr.Token) { s._E_FIT = v }

func (s *FdiskparamContext) Set_E_UNIT(v antlr.Token) { s._E_UNIT = v }

func (s *FdiskparamContext) Set_E_TYPE(v antlr.Token) { s._E_TYPE = v }

func (s *FdiskparamContext) SetE_name(v antlr.Token) { s.e_name = v }

func (s *FdiskparamContext) SIZE() antlr.TerminalNode {
	return s.GetToken(GrammarParserSIZE, 0)
}

func (s *FdiskparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *FdiskparamContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *FdiskparamContext) PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserPATH, 0)
}

func (s *FdiskparamContext) E_PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_PATH, 0)
}

func (s *FdiskparamContext) FIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserFIT, 0)
}

func (s *FdiskparamContext) E_FIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_FIT, 0)
}

func (s *FdiskparamContext) UNIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserUNIT, 0)
}

func (s *FdiskparamContext) E_UNIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_UNIT, 0)
}

func (s *FdiskparamContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GrammarParserTYPE, 0)
}

func (s *FdiskparamContext) E_TYPE() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_TYPE, 0)
}

func (s *FdiskparamContext) NAME() antlr.TerminalNode {
	return s.GetToken(GrammarParserNAME, 0)
}

func (s *FdiskparamContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFICADOR, 0)
}

func (s *FdiskparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FdiskparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FdiskparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterFdiskparam(s)
	}
}

func (s *FdiskparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitFdiskparam(s)
	}
}

func (p *GrammarParser) Fdiskparam() (localctx IFdiskparamContext) {
	localctx = NewFdiskparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_fdiskparam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(98)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(99)

			var _m = p.Match(GrammarParserENTERO)

			localctx.(*FdiskparamContext)._ENTERO = _m
		}
		info_FDISK.Size = (func() int {
			if localctx.(*FdiskparamContext).Get_ENTERO() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*FdiskparamContext).Get_ENTERO().GetText())
				return i
			}
		}())

	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(102)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(103)

			var _m = p.Match(GrammarParserE_PATH)

			localctx.(*FdiskparamContext)._E_PATH = _m
		}
		info_FDISK.Path = strings.ReplaceAll((func() string {
			if localctx.(*FdiskparamContext).Get_E_PATH() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).Get_E_PATH().GetText()
			}
		}()), "\"", "")

	case GrammarParserFIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(106)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(107)

			var _m = p.Match(GrammarParserE_FIT)

			localctx.(*FdiskparamContext)._E_FIT = _m
		}
		info_FDISK.Fit = (func() string {
			if localctx.(*FdiskparamContext).Get_E_FIT() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).Get_E_FIT().GetText()
			}
		}())

	case GrammarParserUNIT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(110)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(111)

			var _m = p.Match(GrammarParserE_UNIT)

			localctx.(*FdiskparamContext)._E_UNIT = _m
		}
		info_FDISK.Unit = (func() string {
			if localctx.(*FdiskparamContext).Get_E_UNIT() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).Get_E_UNIT().GetText()
			}
		}())

	case GrammarParserTYPE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(113)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(114)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(115)

			var _m = p.Match(GrammarParserE_TYPE)

			localctx.(*FdiskparamContext)._E_TYPE = _m
		}
		info_FDISK.Type = (func() string {
			if localctx.(*FdiskparamContext).Get_E_TYPE() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).Get_E_TYPE().GetText()
			}
		}())

	case GrammarParserNAME:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(117)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(118)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(119)

			var _m = p.Match(GrammarParserIDENTIFICADOR)

			localctx.(*FdiskparamContext).e_name = _m
		}
		info_FDISK.Name = (func() string {
			if localctx.(*FdiskparamContext).GetE_name() == nil {
				return ""
			} else {
				return localctx.(*FdiskparamContext).GetE_name().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMount_fContext is an interface to support dynamic dispatch.
type IMount_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMount_fContext differentiates from other interfaces.
	IsMount_fContext()
}

type Mount_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMount_fContext() *Mount_fContext {
	var p = new(Mount_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mount_f
	return p
}

func (*Mount_fContext) IsMount_fContext() {}

func NewMount_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mount_fContext {
	var p = new(Mount_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mount_f

	return p
}

func (s *Mount_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Mount_fContext) MOUNT() antlr.TerminalNode {
	return s.GetToken(GrammarParserMOUNT, 0)
}

func (s *Mount_fContext) AllMountparam() []IMountparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMountparamContext)(nil)).Elem())
	var tst = make([]IMountparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMountparamContext)
		}
	}

	return tst
}

func (s *Mount_fContext) Mountparam(i int) IMountparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMountparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMountparamContext)
}

func (s *Mount_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mount_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mount_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMount_f(s)
	}
}

func (s *Mount_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMount_f(s)
	}
}

func (p *GrammarParser) Mount_f() (localctx IMount_fContext) {
	localctx = NewMount_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_mount_f)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(GrammarParserMOUNT)
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserPATH || _la == GrammarParserNAME {
		{
			p.SetState(124)
			p.Mountparam()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.MontarDisco(info_MOUNT)
	initializeMOUNT(&info_MOUNT)

	return localctx
}

// IMountparamContext is an interface to support dynamic dispatch.
type IMountparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_E_PATH returns the _E_PATH token.
	Get_E_PATH() antlr.Token

	// GetE_name returns the e_name token.
	GetE_name() antlr.Token

	// Set_E_PATH sets the _E_PATH token.
	Set_E_PATH(antlr.Token)

	// SetE_name sets the e_name token.
	SetE_name(antlr.Token)

	// IsMountparamContext differentiates from other interfaces.
	IsMountparamContext()
}

type MountparamContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_E_PATH antlr.Token
	e_name  antlr.Token
}

func NewEmptyMountparamContext() *MountparamContext {
	var p = new(MountparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mountparam
	return p
}

func (*MountparamContext) IsMountparamContext() {}

func NewMountparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MountparamContext {
	var p = new(MountparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mountparam

	return p
}

func (s *MountparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MountparamContext) Get_E_PATH() antlr.Token { return s._E_PATH }

func (s *MountparamContext) GetE_name() antlr.Token { return s.e_name }

func (s *MountparamContext) Set_E_PATH(v antlr.Token) { s._E_PATH = v }

func (s *MountparamContext) SetE_name(v antlr.Token) { s.e_name = v }

func (s *MountparamContext) PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserPATH, 0)
}

func (s *MountparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *MountparamContext) E_PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_PATH, 0)
}

func (s *MountparamContext) NAME() antlr.TerminalNode {
	return s.GetToken(GrammarParserNAME, 0)
}

func (s *MountparamContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFICADOR, 0)
}

func (s *MountparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MountparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MountparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMountparam(s)
	}
}

func (s *MountparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMountparam(s)
	}
}

func (p *GrammarParser) Mountparam() (localctx IMountparamContext) {
	localctx = NewMountparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_mountparam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(132)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(133)

			var _m = p.Match(GrammarParserE_PATH)

			localctx.(*MountparamContext)._E_PATH = _m
		}
		info_MOUNT.Path = strings.ReplaceAll((func() string {
			if localctx.(*MountparamContext).Get_E_PATH() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).Get_E_PATH().GetText()
			}
		}()), "\"", "")

	case GrammarParserNAME:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(135)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(136)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(137)

			var _m = p.Match(GrammarParserIDENTIFICADOR)

			localctx.(*MountparamContext).e_name = _m
		}
		info_MOUNT.Name = (func() string {
			if localctx.(*MountparamContext).GetE_name() == nil {
				return ""
			} else {
				return localctx.(*MountparamContext).GetE_name().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMkfs_fContext is an interface to support dynamic dispatch.
type IMkfs_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMkfs_fContext differentiates from other interfaces.
	IsMkfs_fContext()
}

type Mkfs_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfs_fContext() *Mkfs_fContext {
	var p = new(Mkfs_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkfs_f
	return p
}

func (*Mkfs_fContext) IsMkfs_fContext() {}

func NewMkfs_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkfs_fContext {
	var p = new(Mkfs_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkfs_f

	return p
}

func (s *Mkfs_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkfs_fContext) MKFS() antlr.TerminalNode {
	return s.GetToken(GrammarParserMKFS, 0)
}

func (s *Mkfs_fContext) AllMkfsparam() []IMkfsparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMkfsparamContext)(nil)).Elem())
	var tst = make([]IMkfsparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMkfsparamContext)
		}
	}

	return tst
}

func (s *Mkfs_fContext) Mkfsparam(i int) IMkfsparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkfsparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMkfsparamContext)
}

func (s *Mkfs_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkfs_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkfs_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkfs_f(s)
	}
}

func (s *Mkfs_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkfs_f(s)
	}
}

func (p *GrammarParser) Mkfs_f() (localctx IMkfs_fContext) {
	localctx = NewMkfs_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GrammarParserRULE_mkfs_f)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(GrammarParserMKFS)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserTYPE || _la == GrammarParserID {
		{
			p.SetState(142)
			p.Mkfsparam()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.Formatear(info_MKFS)
	initializeMKFS(&info_MKFS)

	return localctx
}

// IMkfsparamContext is an interface to support dynamic dispatch.
type IMkfsparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_E_ID returns the _E_ID token.
	Get_E_ID() antlr.Token

	// Get_E_TYPE returns the _E_TYPE token.
	Get_E_TYPE() antlr.Token

	// Set_E_ID sets the _E_ID token.
	Set_E_ID(antlr.Token)

	// Set_E_TYPE sets the _E_TYPE token.
	Set_E_TYPE(antlr.Token)

	// IsMkfsparamContext differentiates from other interfaces.
	IsMkfsparamContext()
}

type MkfsparamContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_E_ID   antlr.Token
	_E_TYPE antlr.Token
}

func NewEmptyMkfsparamContext() *MkfsparamContext {
	var p = new(MkfsparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkfsparam
	return p
}

func (*MkfsparamContext) IsMkfsparamContext() {}

func NewMkfsparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfsparamContext {
	var p = new(MkfsparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkfsparam

	return p
}

func (s *MkfsparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfsparamContext) Get_E_ID() antlr.Token { return s._E_ID }

func (s *MkfsparamContext) Get_E_TYPE() antlr.Token { return s._E_TYPE }

func (s *MkfsparamContext) Set_E_ID(v antlr.Token) { s._E_ID = v }

func (s *MkfsparamContext) Set_E_TYPE(v antlr.Token) { s._E_TYPE = v }

func (s *MkfsparamContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *MkfsparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *MkfsparamContext) E_ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_ID, 0)
}

func (s *MkfsparamContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GrammarParserTYPE, 0)
}

func (s *MkfsparamContext) E_TYPE() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_TYPE, 0)
}

func (s *MkfsparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfsparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfsparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkfsparam(s)
	}
}

func (s *MkfsparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkfsparam(s)
	}
}

func (p *GrammarParser) Mkfsparam() (localctx IMkfsparamContext) {
	localctx = NewMkfsparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GrammarParserRULE_mkfsparam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(157)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(150)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(151)

			var _m = p.Match(GrammarParserE_ID)

			localctx.(*MkfsparamContext)._E_ID = _m
		}
		info_MKFS.Id = (func() string {
			if localctx.(*MkfsparamContext).Get_E_ID() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).Get_E_ID().GetText()
			}
		}())

	case GrammarParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(153)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(154)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(155)

			var _m = p.Match(GrammarParserE_TYPE)

			localctx.(*MkfsparamContext)._E_TYPE = _m
		}
		info_MKFS.Type = (func() string {
			if localctx.(*MkfsparamContext).Get_E_TYPE() == nil {
				return ""
			} else {
				return localctx.(*MkfsparamContext).Get_E_TYPE().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILogin_fContext is an interface to support dynamic dispatch.
type ILogin_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogin_fContext differentiates from other interfaces.
	IsLogin_fContext()
}

type Login_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogin_fContext() *Login_fContext {
	var p = new(Login_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_login_f
	return p
}

func (*Login_fContext) IsLogin_fContext() {}

func NewLogin_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Login_fContext {
	var p = new(Login_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_login_f

	return p
}

func (s *Login_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Login_fContext) LOGIN() antlr.TerminalNode {
	return s.GetToken(GrammarParserLOGIN, 0)
}

func (s *Login_fContext) AllLoginparam() []ILoginparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILoginparamContext)(nil)).Elem())
	var tst = make([]ILoginparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILoginparamContext)
		}
	}

	return tst
}

func (s *Login_fContext) Loginparam(i int) ILoginparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoginparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILoginparamContext)
}

func (s *Login_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Login_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Login_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterLogin_f(s)
	}
}

func (s *Login_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitLogin_f(s)
	}
}

func (p *GrammarParser) Login_f() (localctx ILogin_fContext) {
	localctx = NewLogin_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GrammarParserRULE_login_f)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(GrammarParserLOGIN)
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserID)|(1<<GrammarParserUSR)|(1<<GrammarParserPASSW))) != 0) {
		{
			p.SetState(160)
			p.Loginparam()
		}

		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.LoginS(info_LOGIN)
	initializeLOGIN(&info_LOGIN)

	return localctx
}

// ILoginparamContext is an interface to support dynamic dispatch.
type ILoginparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE_user returns the e_user token.
	GetE_user() antlr.Token

	// GetE_pass returns the e_pass token.
	GetE_pass() antlr.Token

	// Get_E_ID returns the _E_ID token.
	Get_E_ID() antlr.Token

	// SetE_user sets the e_user token.
	SetE_user(antlr.Token)

	// SetE_pass sets the e_pass token.
	SetE_pass(antlr.Token)

	// Set_E_ID sets the _E_ID token.
	Set_E_ID(antlr.Token)

	// IsLoginparamContext differentiates from other interfaces.
	IsLoginparamContext()
}

type LoginparamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	e_user antlr.Token
	e_pass antlr.Token
	_E_ID  antlr.Token
}

func NewEmptyLoginparamContext() *LoginparamContext {
	var p = new(LoginparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_loginparam
	return p
}

func (*LoginparamContext) IsLoginparamContext() {}

func NewLoginparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoginparamContext {
	var p = new(LoginparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_loginparam

	return p
}

func (s *LoginparamContext) GetParser() antlr.Parser { return s.parser }

func (s *LoginparamContext) GetE_user() antlr.Token { return s.e_user }

func (s *LoginparamContext) GetE_pass() antlr.Token { return s.e_pass }

func (s *LoginparamContext) Get_E_ID() antlr.Token { return s._E_ID }

func (s *LoginparamContext) SetE_user(v antlr.Token) { s.e_user = v }

func (s *LoginparamContext) SetE_pass(v antlr.Token) { s.e_pass = v }

func (s *LoginparamContext) Set_E_ID(v antlr.Token) { s._E_ID = v }

func (s *LoginparamContext) USR() antlr.TerminalNode {
	return s.GetToken(GrammarParserUSR, 0)
}

func (s *LoginparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *LoginparamContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFICADOR, 0)
}

func (s *LoginparamContext) COMPLEMENTO() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMPLEMENTO, 0)
}

func (s *LoginparamContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *LoginparamContext) E_USRS() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_USRS, 0)
}

func (s *LoginparamContext) PASSW() antlr.TerminalNode {
	return s.GetToken(GrammarParserPASSW, 0)
}

func (s *LoginparamContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *LoginparamContext) E_ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_ID, 0)
}

func (s *LoginparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoginparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoginparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterLoginparam(s)
	}
}

func (s *LoginparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitLoginparam(s)
	}
}

func (p *GrammarParser) Loginparam() (localctx ILoginparamContext) {
	localctx = NewLoginparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GrammarParserRULE_loginparam)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(179)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserUSR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(GrammarParserUSR)
		}
		{
			p.SetState(168)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(169)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LoginparamContext).e_user = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GrammarParserENTERO-33))|(1<<(GrammarParserIDENTIFICADOR-33))|(1<<(GrammarParserCOMPLEMENTO-33))|(1<<(GrammarParserE_USRS-33)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LoginparamContext).e_user = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_LOGIN.User = (func() string {
			if localctx.(*LoginparamContext).GetE_user() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetE_user().GetText()
			}
		}())

	case GrammarParserPASSW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
			p.Match(GrammarParserPASSW)
		}
		{
			p.SetState(172)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(173)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LoginparamContext).e_pass = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GrammarParserENTERO-33))|(1<<(GrammarParserIDENTIFICADOR-33))|(1<<(GrammarParserCOMPLEMENTO-33))|(1<<(GrammarParserE_USRS-33)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LoginparamContext).e_pass = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_LOGIN.Pass = (func() string {
			if localctx.(*LoginparamContext).GetE_pass() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetE_pass().GetText()
			}
		}())

	case GrammarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(175)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(176)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(177)

			var _m = p.Match(GrammarParserE_ID)

			localctx.(*LoginparamContext)._E_ID = _m
		}
		info_LOGIN.Id = (func() string {
			if localctx.(*LoginparamContext).Get_E_ID() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).Get_E_ID().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILogout_fContext is an interface to support dynamic dispatch.
type ILogout_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogout_fContext differentiates from other interfaces.
	IsLogout_fContext()
}

type Logout_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogout_fContext() *Logout_fContext {
	var p = new(Logout_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_logout_f
	return p
}

func (*Logout_fContext) IsLogout_fContext() {}

func NewLogout_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logout_fContext {
	var p = new(Logout_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_logout_f

	return p
}

func (s *Logout_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Logout_fContext) LOGOUT() antlr.TerminalNode {
	return s.GetToken(GrammarParserLOGOUT, 0)
}

func (s *Logout_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logout_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logout_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterLogout_f(s)
	}
}

func (s *Logout_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitLogout_f(s)
	}
}

func (p *GrammarParser) Logout_f() (localctx ILogout_fContext) {
	localctx = NewLogout_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GrammarParserRULE_logout_f)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		p.Match(GrammarParserLOGOUT)
	}
	Program.LogoutS()

	return localctx
}
