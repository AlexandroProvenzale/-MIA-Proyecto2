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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 227,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 6, 2, 38, 10, 2, 13, 2, 14, 2, 39, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 70, 10, 3, 3, 4, 3, 4, 6, 4, 74, 10, 4, 13, 4, 14, 4, 75, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 96, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 6, 7, 106, 10, 7, 13, 7, 14, 7, 107, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	136, 10, 8, 3, 9, 3, 9, 6, 9, 140, 10, 9, 13, 9, 14, 9, 141, 3, 9, 3, 9,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 154, 10,
	10, 3, 11, 3, 11, 6, 11, 158, 10, 11, 13, 11, 14, 11, 159, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 172, 10,
	12, 3, 13, 3, 13, 6, 13, 176, 10, 13, 13, 13, 14, 13, 177, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 194, 10, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 6, 17, 207, 10, 17, 13, 17, 14, 17,
	208, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 225, 10, 18, 3, 18, 2, 2, 19, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 3, 4, 2,
	38, 38, 40, 42, 2, 239, 2, 37, 3, 2, 2, 2, 4, 69, 3, 2, 2, 2, 6, 71, 3,
	2, 2, 2, 8, 95, 3, 2, 2, 2, 10, 97, 3, 2, 2, 2, 12, 103, 3, 2, 2, 2, 14,
	135, 3, 2, 2, 2, 16, 137, 3, 2, 2, 2, 18, 153, 3, 2, 2, 2, 20, 155, 3,
	2, 2, 2, 22, 171, 3, 2, 2, 2, 24, 173, 3, 2, 2, 2, 26, 193, 3, 2, 2, 2,
	28, 195, 3, 2, 2, 2, 30, 198, 3, 2, 2, 2, 32, 204, 3, 2, 2, 2, 34, 224,
	3, 2, 2, 2, 36, 38, 5, 4, 3, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2,
	39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 3, 3, 2, 2, 2, 41, 42, 5, 6,
	4, 2, 42, 43, 7, 43, 2, 2, 43, 70, 3, 2, 2, 2, 44, 45, 5, 10, 6, 2, 45,
	46, 7, 43, 2, 2, 46, 70, 3, 2, 2, 2, 47, 48, 5, 12, 7, 2, 48, 49, 7, 43,
	2, 2, 49, 70, 3, 2, 2, 2, 50, 51, 5, 16, 9, 2, 51, 52, 7, 43, 2, 2, 52,
	70, 3, 2, 2, 2, 53, 54, 5, 20, 11, 2, 54, 55, 7, 43, 2, 2, 55, 70, 3, 2,
	2, 2, 56, 57, 5, 24, 13, 2, 57, 58, 7, 43, 2, 2, 58, 70, 3, 2, 2, 2, 59,
	60, 5, 28, 15, 2, 60, 61, 7, 43, 2, 2, 61, 70, 3, 2, 2, 2, 62, 63, 5, 30,
	16, 2, 63, 64, 7, 43, 2, 2, 64, 70, 3, 2, 2, 2, 65, 66, 5, 32, 17, 2, 66,
	67, 7, 43, 2, 2, 67, 70, 3, 2, 2, 2, 68, 70, 7, 43, 2, 2, 69, 41, 3, 2,
	2, 2, 69, 44, 3, 2, 2, 2, 69, 47, 3, 2, 2, 2, 69, 50, 3, 2, 2, 2, 69, 53,
	3, 2, 2, 2, 69, 56, 3, 2, 2, 2, 69, 59, 3, 2, 2, 2, 69, 62, 3, 2, 2, 2,
	69, 65, 3, 2, 2, 2, 69, 68, 3, 2, 2, 2, 70, 5, 3, 2, 2, 2, 71, 73, 7, 6,
	2, 2, 72, 74, 5, 8, 5, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 8, 4, 1, 2,
	78, 7, 3, 2, 2, 2, 79, 80, 7, 16, 2, 2, 80, 81, 7, 37, 2, 2, 81, 82, 7,
	38, 2, 2, 82, 96, 8, 5, 1, 2, 83, 84, 7, 19, 2, 2, 84, 85, 7, 37, 2, 2,
	85, 86, 7, 33, 2, 2, 86, 96, 8, 5, 1, 2, 87, 88, 7, 17, 2, 2, 88, 89, 7,
	37, 2, 2, 89, 90, 7, 30, 2, 2, 90, 96, 8, 5, 1, 2, 91, 92, 7, 18, 2, 2,
	92, 93, 7, 37, 2, 2, 93, 94, 7, 31, 2, 2, 94, 96, 8, 5, 1, 2, 95, 79, 3,
	2, 2, 2, 95, 83, 3, 2, 2, 2, 95, 87, 3, 2, 2, 2, 95, 91, 3, 2, 2, 2, 96,
	9, 3, 2, 2, 2, 97, 98, 7, 7, 2, 2, 98, 99, 7, 19, 2, 2, 99, 100, 7, 37,
	2, 2, 100, 101, 7, 33, 2, 2, 101, 102, 8, 6, 1, 2, 102, 11, 3, 2, 2, 2,
	103, 105, 7, 8, 2, 2, 104, 106, 5, 14, 8, 2, 105, 104, 3, 2, 2, 2, 106,
	107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109,
	3, 2, 2, 2, 109, 110, 8, 7, 1, 2, 110, 13, 3, 2, 2, 2, 111, 112, 7, 16,
	2, 2, 112, 113, 7, 37, 2, 2, 113, 114, 7, 38, 2, 2, 114, 136, 8, 8, 1,
	2, 115, 116, 7, 19, 2, 2, 116, 117, 7, 37, 2, 2, 117, 118, 7, 33, 2, 2,
	118, 136, 8, 8, 1, 2, 119, 120, 7, 17, 2, 2, 120, 121, 7, 37, 2, 2, 121,
	122, 7, 30, 2, 2, 122, 136, 8, 8, 1, 2, 123, 124, 7, 18, 2, 2, 124, 125,
	7, 37, 2, 2, 125, 126, 7, 31, 2, 2, 126, 136, 8, 8, 1, 2, 127, 128, 7,
	20, 2, 2, 128, 129, 7, 37, 2, 2, 129, 130, 7, 32, 2, 2, 130, 136, 8, 8,
	1, 2, 131, 132, 7, 22, 2, 2, 132, 133, 7, 37, 2, 2, 133, 134, 7, 40, 2,
	2, 134, 136, 8, 8, 1, 2, 135, 111, 3, 2, 2, 2, 135, 115, 3, 2, 2, 2, 135,
	119, 3, 2, 2, 2, 135, 123, 3, 2, 2, 2, 135, 127, 3, 2, 2, 2, 135, 131,
	3, 2, 2, 2, 136, 15, 3, 2, 2, 2, 137, 139, 7, 9, 2, 2, 138, 140, 5, 18,
	10, 2, 139, 138, 3, 2, 2, 2, 140, 141, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2,
	141, 142, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 8, 9, 1, 2, 144,
	17, 3, 2, 2, 2, 145, 146, 7, 19, 2, 2, 146, 147, 7, 37, 2, 2, 147, 148,
	7, 33, 2, 2, 148, 154, 8, 10, 1, 2, 149, 150, 7, 22, 2, 2, 150, 151, 7,
	37, 2, 2, 151, 152, 7, 40, 2, 2, 152, 154, 8, 10, 1, 2, 153, 145, 3, 2,
	2, 2, 153, 149, 3, 2, 2, 2, 154, 19, 3, 2, 2, 2, 155, 157, 7, 11, 2, 2,
	156, 158, 5, 22, 12, 2, 157, 156, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159,
	157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162,
	8, 11, 1, 2, 162, 21, 3, 2, 2, 2, 163, 164, 7, 24, 2, 2, 164, 165, 7, 37,
	2, 2, 165, 166, 7, 34, 2, 2, 166, 172, 8, 12, 1, 2, 167, 168, 7, 20, 2,
	2, 168, 169, 7, 37, 2, 2, 169, 170, 7, 32, 2, 2, 170, 172, 8, 12, 1, 2,
	171, 163, 3, 2, 2, 2, 171, 167, 3, 2, 2, 2, 172, 23, 3, 2, 2, 2, 173, 175,
	7, 12, 2, 2, 174, 176, 5, 26, 14, 2, 175, 174, 3, 2, 2, 2, 176, 177, 3,
	2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 3, 2, 2,
	2, 179, 180, 8, 13, 1, 2, 180, 25, 3, 2, 2, 2, 181, 182, 7, 26, 2, 2, 182,
	183, 7, 37, 2, 2, 183, 184, 9, 2, 2, 2, 184, 194, 8, 14, 1, 2, 185, 186,
	7, 27, 2, 2, 186, 187, 7, 37, 2, 2, 187, 188, 9, 2, 2, 2, 188, 194, 8,
	14, 1, 2, 189, 190, 7, 24, 2, 2, 190, 191, 7, 37, 2, 2, 191, 192, 7, 34,
	2, 2, 192, 194, 8, 14, 1, 2, 193, 181, 3, 2, 2, 2, 193, 185, 3, 2, 2, 2,
	193, 189, 3, 2, 2, 2, 194, 27, 3, 2, 2, 2, 195, 196, 7, 13, 2, 2, 196,
	197, 8, 15, 1, 2, 197, 29, 3, 2, 2, 2, 198, 199, 7, 14, 2, 2, 199, 200,
	7, 22, 2, 2, 200, 201, 7, 37, 2, 2, 201, 202, 9, 2, 2, 2, 202, 203, 8,
	16, 1, 2, 203, 31, 3, 2, 2, 2, 204, 206, 7, 15, 2, 2, 205, 207, 5, 34,
	18, 2, 206, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2,
	208, 209, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 8, 17, 1, 2, 211,
	33, 3, 2, 2, 2, 212, 213, 7, 26, 2, 2, 213, 214, 7, 37, 2, 2, 214, 215,
	9, 2, 2, 2, 215, 225, 8, 18, 1, 2, 216, 217, 7, 28, 2, 2, 217, 218, 7,
	37, 2, 2, 218, 219, 9, 2, 2, 2, 219, 225, 8, 18, 1, 2, 220, 221, 7, 29,
	2, 2, 221, 222, 7, 37, 2, 2, 222, 223, 9, 2, 2, 2, 223, 225, 8, 18, 1,
	2, 224, 212, 3, 2, 2, 2, 224, 216, 3, 2, 2, 2, 224, 220, 3, 2, 2, 2, 225,
	35, 3, 2, 2, 2, 16, 39, 69, 75, 95, 107, 135, 141, 153, 159, 171, 177,
	193, 208, 224,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'='",
}
var symbolicNames = []string{
	"", "PAUSES", "EXEC", "REP", "MKDISK", "RMDISK", "FDISK", "MOUNT", "UNMOUNT",
	"MKFS", "LOGIN", "LOGOUT", "MKGRP", "MKUSR", "SIZE", "FIT", "UNIT", "PATH",
	"TYPE", "DELETEP", "NAME", "ADD", "ID", "FS", "USR", "PASSW", "PWD", "GRP",
	"E_FIT", "E_UNIT", "E_TYPE", "E_PATH", "E_ID", "PATH1", "PATH2", "IGUAL",
	"ENTERO", "NEGATIVO", "IDENTIFICADOR", "COMPLEMENTO", "E_USRS", "NEWLINE",
	"WHITESPACE", "COMENTARIO",
}

var ruleNames = []string{
	"start", "comando", "mkdisk_f", "mkparam", "rmdisk_f", "fdisk_f", "fdiskparam",
	"mount_f", "mountparam", "mkfs_f", "mkfsparam", "login_f", "loginparam",
	"logout_f", "mkgroup_f", "mkuser_f", "mkuserparam",
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
var info_MKUSER Program.InfoMkuser

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

func initializeMKUSER(MKUSER *Program.InfoMkuser) {
	MKUSER.User = ""
	MKUSER.Pass = ""
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
	GrammarParserMKUSR         = 13
	GrammarParserSIZE          = 14
	GrammarParserFIT           = 15
	GrammarParserUNIT          = 16
	GrammarParserPATH          = 17
	GrammarParserTYPE          = 18
	GrammarParserDELETEP       = 19
	GrammarParserNAME          = 20
	GrammarParserADD           = 21
	GrammarParserID            = 22
	GrammarParserFS            = 23
	GrammarParserUSR           = 24
	GrammarParserPASSW         = 25
	GrammarParserPWD           = 26
	GrammarParserGRP           = 27
	GrammarParserE_FIT         = 28
	GrammarParserE_UNIT        = 29
	GrammarParserE_TYPE        = 30
	GrammarParserE_PATH        = 31
	GrammarParserE_ID          = 32
	GrammarParserPATH1         = 33
	GrammarParserPATH2         = 34
	GrammarParserIGUAL         = 35
	GrammarParserENTERO        = 36
	GrammarParserNEGATIVO      = 37
	GrammarParserIDENTIFICADOR = 38
	GrammarParserCOMPLEMENTO   = 39
	GrammarParserE_USRS        = 40
	GrammarParserNEWLINE       = 41
	GrammarParserWHITESPACE    = 42
	GrammarParserCOMENTARIO    = 43
)

// GrammarParser rules.
const (
	GrammarParserRULE_start       = 0
	GrammarParserRULE_comando     = 1
	GrammarParserRULE_mkdisk_f    = 2
	GrammarParserRULE_mkparam     = 3
	GrammarParserRULE_rmdisk_f    = 4
	GrammarParserRULE_fdisk_f     = 5
	GrammarParserRULE_fdiskparam  = 6
	GrammarParserRULE_mount_f     = 7
	GrammarParserRULE_mountparam  = 8
	GrammarParserRULE_mkfs_f      = 9
	GrammarParserRULE_mkfsparam   = 10
	GrammarParserRULE_login_f     = 11
	GrammarParserRULE_loginparam  = 12
	GrammarParserRULE_logout_f    = 13
	GrammarParserRULE_mkgroup_f   = 14
	GrammarParserRULE_mkuser_f    = 15
	GrammarParserRULE_mkuserparam = 16
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserMKDISK)|(1<<GrammarParserRMDISK)|(1<<GrammarParserFDISK)|(1<<GrammarParserMOUNT)|(1<<GrammarParserMKFS)|(1<<GrammarParserLOGIN)|(1<<GrammarParserLOGOUT)|(1<<GrammarParserMKGRP)|(1<<GrammarParserMKUSR))) != 0) || _la == GrammarParserNEWLINE {
		{
			p.SetState(34)
			p.Comando()
		}

		p.SetState(37)
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

func (s *ComandoContext) Mkgroup_f() IMkgroup_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkgroup_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkgroup_fContext)
}

func (s *ComandoContext) Mkuser_f() IMkuser_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkuser_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkuser_fContext)
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

	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserMKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(39)
			p.Mkdisk_f()
		}
		{
			p.SetState(40)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserRMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Rmdisk_f()
		}
		{
			p.SetState(43)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserFDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.Fdisk_f()
		}
		{
			p.SetState(46)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(48)
			p.Mount_f()
		}
		{
			p.SetState(49)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKFS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(51)
			p.Mkfs_f()
		}
		{
			p.SetState(52)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(54)
			p.Login_f()
		}
		{
			p.SetState(55)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(57)
			p.Logout_f()
		}
		{
			p.SetState(58)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKGRP:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(60)
			p.Mkgroup_f()
		}
		{
			p.SetState(61)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKUSR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(63)
			p.Mkuser_f()
		}
		{
			p.SetState(64)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserNEWLINE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(66)
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
		p.SetState(69)
		p.Match(GrammarParserMKDISK)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH))) != 0) {
		{
			p.SetState(70)
			p.Mkparam()
		}

		p.SetState(73)
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

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(78)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(79)

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
			p.SetState(81)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(82)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(83)

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
			p.SetState(85)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(86)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(87)

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
			p.SetState(89)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(90)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(91)

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
		p.SetState(95)
		p.Match(GrammarParserRMDISK)
	}
	{
		p.SetState(96)
		p.Match(GrammarParserPATH)
	}
	{
		p.SetState(97)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(98)

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
		p.SetState(101)
		p.Match(GrammarParserFDISK)
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH)|(1<<GrammarParserTYPE)|(1<<GrammarParserNAME))) != 0) {
		{
			p.SetState(102)
			p.Fdiskparam()
		}

		p.SetState(105)
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

	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(110)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(111)

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
			p.SetState(113)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(114)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(115)

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
			p.SetState(117)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(118)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(119)

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
			p.SetState(121)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(122)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(123)

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
			p.SetState(125)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(126)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(127)

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
			p.SetState(129)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(130)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(131)

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
		p.SetState(135)
		p.Match(GrammarParserMOUNT)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserPATH || _la == GrammarParserNAME {
		{
			p.SetState(136)
			p.Mountparam()
		}

		p.SetState(139)
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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(144)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(145)

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
			p.SetState(147)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(148)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(149)

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
		p.SetState(153)
		p.Match(GrammarParserMKFS)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserTYPE || _la == GrammarParserID {
		{
			p.SetState(154)
			p.Mkfsparam()
		}

		p.SetState(157)
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

	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(162)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(163)

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
			p.SetState(165)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(166)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(167)

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
		p.SetState(171)
		p.Match(GrammarParserLOGIN)
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserID)|(1<<GrammarParserUSR)|(1<<GrammarParserPASSW))) != 0) {
		{
			p.SetState(172)
			p.Loginparam()
		}

		p.SetState(175)
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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserUSR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Match(GrammarParserUSR)
		}
		{
			p.SetState(180)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(181)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LoginparamContext).e_user = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GrammarParserENTERO-36))|(1<<(GrammarParserIDENTIFICADOR-36))|(1<<(GrammarParserCOMPLEMENTO-36))|(1<<(GrammarParserE_USRS-36)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LoginparamContext).e_user = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_LOGIN.User = strings.ReplaceAll((func() string {
			if localctx.(*LoginparamContext).GetE_user() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetE_user().GetText()
			}
		}()), "\"", "")

	case GrammarParserPASSW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(GrammarParserPASSW)
		}
		{
			p.SetState(184)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(185)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LoginparamContext).e_pass = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GrammarParserENTERO-36))|(1<<(GrammarParserIDENTIFICADOR-36))|(1<<(GrammarParserCOMPLEMENTO-36))|(1<<(GrammarParserE_USRS-36)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LoginparamContext).e_pass = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_LOGIN.Pass = strings.ReplaceAll((func() string {
			if localctx.(*LoginparamContext).GetE_pass() == nil {
				return ""
			} else {
				return localctx.(*LoginparamContext).GetE_pass().GetText()
			}
		}()), "\"", "")

	case GrammarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(188)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(189)

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
		p.SetState(193)
		p.Match(GrammarParserLOGOUT)
	}
	Program.LogoutS()

	return localctx
}

// IMkgroup_fContext is an interface to support dynamic dispatch.
type IMkgroup_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE_name returns the e_name token.
	GetE_name() antlr.Token

	// SetE_name sets the e_name token.
	SetE_name(antlr.Token)

	// IsMkgroup_fContext differentiates from other interfaces.
	IsMkgroup_fContext()
}

type Mkgroup_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	e_name antlr.Token
}

func NewEmptyMkgroup_fContext() *Mkgroup_fContext {
	var p = new(Mkgroup_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkgroup_f
	return p
}

func (*Mkgroup_fContext) IsMkgroup_fContext() {}

func NewMkgroup_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkgroup_fContext {
	var p = new(Mkgroup_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkgroup_f

	return p
}

func (s *Mkgroup_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkgroup_fContext) GetE_name() antlr.Token { return s.e_name }

func (s *Mkgroup_fContext) SetE_name(v antlr.Token) { s.e_name = v }

func (s *Mkgroup_fContext) MKGRP() antlr.TerminalNode {
	return s.GetToken(GrammarParserMKGRP, 0)
}

func (s *Mkgroup_fContext) NAME() antlr.TerminalNode {
	return s.GetToken(GrammarParserNAME, 0)
}

func (s *Mkgroup_fContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *Mkgroup_fContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFICADOR, 0)
}

func (s *Mkgroup_fContext) COMPLEMENTO() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMPLEMENTO, 0)
}

func (s *Mkgroup_fContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *Mkgroup_fContext) E_USRS() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_USRS, 0)
}

func (s *Mkgroup_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkgroup_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkgroup_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkgroup_f(s)
	}
}

func (s *Mkgroup_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkgroup_f(s)
	}
}

func (p *GrammarParser) Mkgroup_f() (localctx IMkgroup_fContext) {
	localctx = NewMkgroup_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GrammarParserRULE_mkgroup_f)
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
		p.SetState(196)
		p.Match(GrammarParserMKGRP)
	}
	{
		p.SetState(197)
		p.Match(GrammarParserNAME)
	}
	{
		p.SetState(198)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(199)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Mkgroup_fContext).e_name = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GrammarParserENTERO-36))|(1<<(GrammarParserIDENTIFICADOR-36))|(1<<(GrammarParserCOMPLEMENTO-36))|(1<<(GrammarParserE_USRS-36)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Mkgroup_fContext).e_name = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	Program.Mkgroup(strings.ReplaceAll((func() string {
		if localctx.(*Mkgroup_fContext).GetE_name() == nil {
			return ""
		} else {
			return localctx.(*Mkgroup_fContext).GetE_name().GetText()
		}
	}()), "\"", ""))

	return localctx
}

// IMkuser_fContext is an interface to support dynamic dispatch.
type IMkuser_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMkuser_fContext differentiates from other interfaces.
	IsMkuser_fContext()
}

type Mkuser_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkuser_fContext() *Mkuser_fContext {
	var p = new(Mkuser_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkuser_f
	return p
}

func (*Mkuser_fContext) IsMkuser_fContext() {}

func NewMkuser_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkuser_fContext {
	var p = new(Mkuser_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkuser_f

	return p
}

func (s *Mkuser_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkuser_fContext) MKUSR() antlr.TerminalNode {
	return s.GetToken(GrammarParserMKUSR, 0)
}

func (s *Mkuser_fContext) AllMkuserparam() []IMkuserparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMkuserparamContext)(nil)).Elem())
	var tst = make([]IMkuserparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMkuserparamContext)
		}
	}

	return tst
}

func (s *Mkuser_fContext) Mkuserparam(i int) IMkuserparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkuserparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMkuserparamContext)
}

func (s *Mkuser_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkuser_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkuser_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkuser_f(s)
	}
}

func (s *Mkuser_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkuser_f(s)
	}
}

func (p *GrammarParser) Mkuser_f() (localctx IMkuser_fContext) {
	localctx = NewMkuser_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_mkuser_f)
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
		p.SetState(202)
		p.Match(GrammarParserMKUSR)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserUSR)|(1<<GrammarParserPWD)|(1<<GrammarParserGRP))) != 0) {
		{
			p.SetState(203)
			p.Mkuserparam()
		}

		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.Mkuser(info_MKUSER)
	initializeMKUSER(&info_MKUSER)

	return localctx
}

// IMkuserparamContext is an interface to support dynamic dispatch.
type IMkuserparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE_user returns the e_user token.
	GetE_user() antlr.Token

	// GetE_pass returns the e_pass token.
	GetE_pass() antlr.Token

	// GetE_group returns the e_group token.
	GetE_group() antlr.Token

	// SetE_user sets the e_user token.
	SetE_user(antlr.Token)

	// SetE_pass sets the e_pass token.
	SetE_pass(antlr.Token)

	// SetE_group sets the e_group token.
	SetE_group(antlr.Token)

	// IsMkuserparamContext differentiates from other interfaces.
	IsMkuserparamContext()
}

type MkuserparamContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	e_user  antlr.Token
	e_pass  antlr.Token
	e_group antlr.Token
}

func NewEmptyMkuserparamContext() *MkuserparamContext {
	var p = new(MkuserparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkuserparam
	return p
}

func (*MkuserparamContext) IsMkuserparamContext() {}

func NewMkuserparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkuserparamContext {
	var p = new(MkuserparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkuserparam

	return p
}

func (s *MkuserparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkuserparamContext) GetE_user() antlr.Token { return s.e_user }

func (s *MkuserparamContext) GetE_pass() antlr.Token { return s.e_pass }

func (s *MkuserparamContext) GetE_group() antlr.Token { return s.e_group }

func (s *MkuserparamContext) SetE_user(v antlr.Token) { s.e_user = v }

func (s *MkuserparamContext) SetE_pass(v antlr.Token) { s.e_pass = v }

func (s *MkuserparamContext) SetE_group(v antlr.Token) { s.e_group = v }

func (s *MkuserparamContext) USR() antlr.TerminalNode {
	return s.GetToken(GrammarParserUSR, 0)
}

func (s *MkuserparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *MkuserparamContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFICADOR, 0)
}

func (s *MkuserparamContext) COMPLEMENTO() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMPLEMENTO, 0)
}

func (s *MkuserparamContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *MkuserparamContext) E_USRS() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_USRS, 0)
}

func (s *MkuserparamContext) PWD() antlr.TerminalNode {
	return s.GetToken(GrammarParserPWD, 0)
}

func (s *MkuserparamContext) GRP() antlr.TerminalNode {
	return s.GetToken(GrammarParserGRP, 0)
}

func (s *MkuserparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkuserparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkuserparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkuserparam(s)
	}
}

func (s *MkuserparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkuserparam(s)
	}
}

func (p *GrammarParser) Mkuserparam() (localctx IMkuserparamContext) {
	localctx = NewMkuserparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_mkuserparam)
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

	p.SetState(222)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserUSR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(GrammarParserUSR)
		}
		{
			p.SetState(211)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(212)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_user = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GrammarParserENTERO-36))|(1<<(GrammarParserIDENTIFICADOR-36))|(1<<(GrammarParserCOMPLEMENTO-36))|(1<<(GrammarParserE_USRS-36)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MkuserparamContext).e_user = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_MKUSER.User = strings.ReplaceAll((func() string {
			if localctx.(*MkuserparamContext).GetE_user() == nil {
				return ""
			} else {
				return localctx.(*MkuserparamContext).GetE_user().GetText()
			}
		}()), "\"", "")

	case GrammarParserPWD:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(GrammarParserPWD)
		}
		{
			p.SetState(215)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(216)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_pass = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GrammarParserENTERO-36))|(1<<(GrammarParserIDENTIFICADOR-36))|(1<<(GrammarParserCOMPLEMENTO-36))|(1<<(GrammarParserE_USRS-36)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MkuserparamContext).e_pass = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_MKUSER.Pass = strings.ReplaceAll((func() string {
			if localctx.(*MkuserparamContext).GetE_pass() == nil {
				return ""
			} else {
				return localctx.(*MkuserparamContext).GetE_pass().GetText()
			}
		}()), "\"", "")

	case GrammarParserGRP:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.Match(GrammarParserGRP)
		}
		{
			p.SetState(219)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(220)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_group = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(GrammarParserENTERO-36))|(1<<(GrammarParserIDENTIFICADOR-36))|(1<<(GrammarParserCOMPLEMENTO-36))|(1<<(GrammarParserE_USRS-36)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*MkuserparamContext).e_group = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_MKUSER.Grp = strings.ReplaceAll((func() string {
			if localctx.(*MkuserparamContext).GetE_group() == nil {
				return ""
			} else {
				return localctx.(*MkuserparamContext).GetE_group().GetText()
			}
		}()), "\"", "")

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
