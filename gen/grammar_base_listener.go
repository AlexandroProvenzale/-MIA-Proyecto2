// Code generated from Grammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseGrammarListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseGrammarListener) ExitStart(ctx *StartContext) {}

// EnterComando is called when production comando is entered.
func (s *BaseGrammarListener) EnterComando(ctx *ComandoContext) {}

// ExitComando is called when production comando is exited.
func (s *BaseGrammarListener) ExitComando(ctx *ComandoContext) {}

// EnterMkdisk_f is called when production mkdisk_f is entered.
func (s *BaseGrammarListener) EnterMkdisk_f(ctx *Mkdisk_fContext) {}

// ExitMkdisk_f is called when production mkdisk_f is exited.
func (s *BaseGrammarListener) ExitMkdisk_f(ctx *Mkdisk_fContext) {}

// EnterMkparam is called when production mkparam is entered.
func (s *BaseGrammarListener) EnterMkparam(ctx *MkparamContext) {}

// ExitMkparam is called when production mkparam is exited.
func (s *BaseGrammarListener) ExitMkparam(ctx *MkparamContext) {}

// EnterRmdisk_f is called when production rmdisk_f is entered.
func (s *BaseGrammarListener) EnterRmdisk_f(ctx *Rmdisk_fContext) {}

// ExitRmdisk_f is called when production rmdisk_f is exited.
func (s *BaseGrammarListener) ExitRmdisk_f(ctx *Rmdisk_fContext) {}

// EnterFdisk_f is called when production fdisk_f is entered.
func (s *BaseGrammarListener) EnterFdisk_f(ctx *Fdisk_fContext) {}

// ExitFdisk_f is called when production fdisk_f is exited.
func (s *BaseGrammarListener) ExitFdisk_f(ctx *Fdisk_fContext) {}

// EnterFdiskparam is called when production fdiskparam is entered.
func (s *BaseGrammarListener) EnterFdiskparam(ctx *FdiskparamContext) {}

// ExitFdiskparam is called when production fdiskparam is exited.
func (s *BaseGrammarListener) ExitFdiskparam(ctx *FdiskparamContext) {}
