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

// EnterMount_f is called when production mount_f is entered.
func (s *BaseGrammarListener) EnterMount_f(ctx *Mount_fContext) {}

// ExitMount_f is called when production mount_f is exited.
func (s *BaseGrammarListener) ExitMount_f(ctx *Mount_fContext) {}

// EnterMountparam is called when production mountparam is entered.
func (s *BaseGrammarListener) EnterMountparam(ctx *MountparamContext) {}

// ExitMountparam is called when production mountparam is exited.
func (s *BaseGrammarListener) ExitMountparam(ctx *MountparamContext) {}

// EnterMkfs_f is called when production mkfs_f is entered.
func (s *BaseGrammarListener) EnterMkfs_f(ctx *Mkfs_fContext) {}

// ExitMkfs_f is called when production mkfs_f is exited.
func (s *BaseGrammarListener) ExitMkfs_f(ctx *Mkfs_fContext) {}

// EnterMkfsparam is called when production mkfsparam is entered.
func (s *BaseGrammarListener) EnterMkfsparam(ctx *MkfsparamContext) {}

// ExitMkfsparam is called when production mkfsparam is exited.
func (s *BaseGrammarListener) ExitMkfsparam(ctx *MkfsparamContext) {}

// EnterLogin_f is called when production login_f is entered.
func (s *BaseGrammarListener) EnterLogin_f(ctx *Login_fContext) {}

// ExitLogin_f is called when production login_f is exited.
func (s *BaseGrammarListener) ExitLogin_f(ctx *Login_fContext) {}

// EnterLoginparam is called when production loginparam is entered.
func (s *BaseGrammarListener) EnterLoginparam(ctx *LoginparamContext) {}

// ExitLoginparam is called when production loginparam is exited.
func (s *BaseGrammarListener) ExitLoginparam(ctx *LoginparamContext) {}

// EnterLogout_f is called when production logout_f is entered.
func (s *BaseGrammarListener) EnterLogout_f(ctx *Logout_fContext) {}

// ExitLogout_f is called when production logout_f is exited.
func (s *BaseGrammarListener) ExitLogout_f(ctx *Logout_fContext) {}

// EnterMkgroup_f is called when production mkgroup_f is entered.
func (s *BaseGrammarListener) EnterMkgroup_f(ctx *Mkgroup_fContext) {}

// ExitMkgroup_f is called when production mkgroup_f is exited.
func (s *BaseGrammarListener) ExitMkgroup_f(ctx *Mkgroup_fContext) {}

// EnterMkuser_f is called when production mkuser_f is entered.
func (s *BaseGrammarListener) EnterMkuser_f(ctx *Mkuser_fContext) {}

// ExitMkuser_f is called when production mkuser_f is exited.
func (s *BaseGrammarListener) ExitMkuser_f(ctx *Mkuser_fContext) {}

// EnterMkuserparam is called when production mkuserparam is entered.
func (s *BaseGrammarListener) EnterMkuserparam(ctx *MkuserparamContext) {}

// ExitMkuserparam is called when production mkuserparam is exited.
func (s *BaseGrammarListener) ExitMkuserparam(ctx *MkuserparamContext) {}
