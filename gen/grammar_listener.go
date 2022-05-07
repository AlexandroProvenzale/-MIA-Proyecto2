// Code generated from Grammar.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterComando is called when entering the comando production.
	EnterComando(c *ComandoContext)

	// EnterMkdisk_f is called when entering the mkdisk_f production.
	EnterMkdisk_f(c *Mkdisk_fContext)

	// EnterMkparam is called when entering the mkparam production.
	EnterMkparam(c *MkparamContext)

	// EnterRmdisk_f is called when entering the rmdisk_f production.
	EnterRmdisk_f(c *Rmdisk_fContext)

	// EnterFdisk_f is called when entering the fdisk_f production.
	EnterFdisk_f(c *Fdisk_fContext)

	// EnterFdiskparam is called when entering the fdiskparam production.
	EnterFdiskparam(c *FdiskparamContext)

	// EnterMount_f is called when entering the mount_f production.
	EnterMount_f(c *Mount_fContext)

	// EnterMountparam is called when entering the mountparam production.
	EnterMountparam(c *MountparamContext)

	// EnterMkfs_f is called when entering the mkfs_f production.
	EnterMkfs_f(c *Mkfs_fContext)

	// EnterMkfsparam is called when entering the mkfsparam production.
	EnterMkfsparam(c *MkfsparamContext)

	// EnterLogin_f is called when entering the login_f production.
	EnterLogin_f(c *Login_fContext)

	// EnterLoginparam is called when entering the loginparam production.
	EnterLoginparam(c *LoginparamContext)

	// EnterLogout_f is called when entering the logout_f production.
	EnterLogout_f(c *Logout_fContext)

	// EnterMkgroup_f is called when entering the mkgroup_f production.
	EnterMkgroup_f(c *Mkgroup_fContext)

	// EnterRmgroup_f is called when entering the rmgroup_f production.
	EnterRmgroup_f(c *Rmgroup_fContext)

	// EnterMkuser_f is called when entering the mkuser_f production.
	EnterMkuser_f(c *Mkuser_fContext)

	// EnterRmuser_f is called when entering the rmuser_f production.
	EnterRmuser_f(c *Rmuser_fContext)

	// EnterMkuserparam is called when entering the mkuserparam production.
	EnterMkuserparam(c *MkuserparamContext)

	// EnterMkfile_f is called when entering the mkfile_f production.
	EnterMkfile_f(c *Mkfile_fContext)

	// EnterMkfileparam is called when entering the mkfileparam production.
	EnterMkfileparam(c *MkfileparamContext)

	// EnterMkdir_f is called when entering the mkdir_f production.
	EnterMkdir_f(c *Mkdir_fContext)

	// EnterMkdirparam is called when entering the mkdirparam production.
	EnterMkdirparam(c *MkdirparamContext)

	// EnterRep_f is called when entering the rep_f production.
	EnterRep_f(c *Rep_fContext)

	// EnterRepparam is called when entering the repparam production.
	EnterRepparam(c *RepparamContext)

	// EnterPause_f is called when entering the pause_f production.
	EnterPause_f(c *Pause_fContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitComando is called when exiting the comando production.
	ExitComando(c *ComandoContext)

	// ExitMkdisk_f is called when exiting the mkdisk_f production.
	ExitMkdisk_f(c *Mkdisk_fContext)

	// ExitMkparam is called when exiting the mkparam production.
	ExitMkparam(c *MkparamContext)

	// ExitRmdisk_f is called when exiting the rmdisk_f production.
	ExitRmdisk_f(c *Rmdisk_fContext)

	// ExitFdisk_f is called when exiting the fdisk_f production.
	ExitFdisk_f(c *Fdisk_fContext)

	// ExitFdiskparam is called when exiting the fdiskparam production.
	ExitFdiskparam(c *FdiskparamContext)

	// ExitMount_f is called when exiting the mount_f production.
	ExitMount_f(c *Mount_fContext)

	// ExitMountparam is called when exiting the mountparam production.
	ExitMountparam(c *MountparamContext)

	// ExitMkfs_f is called when exiting the mkfs_f production.
	ExitMkfs_f(c *Mkfs_fContext)

	// ExitMkfsparam is called when exiting the mkfsparam production.
	ExitMkfsparam(c *MkfsparamContext)

	// ExitLogin_f is called when exiting the login_f production.
	ExitLogin_f(c *Login_fContext)

	// ExitLoginparam is called when exiting the loginparam production.
	ExitLoginparam(c *LoginparamContext)

	// ExitLogout_f is called when exiting the logout_f production.
	ExitLogout_f(c *Logout_fContext)

	// ExitMkgroup_f is called when exiting the mkgroup_f production.
	ExitMkgroup_f(c *Mkgroup_fContext)

	// ExitRmgroup_f is called when exiting the rmgroup_f production.
	ExitRmgroup_f(c *Rmgroup_fContext)

	// ExitMkuser_f is called when exiting the mkuser_f production.
	ExitMkuser_f(c *Mkuser_fContext)

	// ExitRmuser_f is called when exiting the rmuser_f production.
	ExitRmuser_f(c *Rmuser_fContext)

	// ExitMkuserparam is called when exiting the mkuserparam production.
	ExitMkuserparam(c *MkuserparamContext)

	// ExitMkfile_f is called when exiting the mkfile_f production.
	ExitMkfile_f(c *Mkfile_fContext)

	// ExitMkfileparam is called when exiting the mkfileparam production.
	ExitMkfileparam(c *MkfileparamContext)

	// ExitMkdir_f is called when exiting the mkdir_f production.
	ExitMkdir_f(c *Mkdir_fContext)

	// ExitMkdirparam is called when exiting the mkdirparam production.
	ExitMkdirparam(c *MkdirparamContext)

	// ExitRep_f is called when exiting the rep_f production.
	ExitRep_f(c *Rep_fContext)

	// ExitRepparam is called when exiting the repparam production.
	ExitRepparam(c *RepparamContext)

	// ExitPause_f is called when exiting the pause_f production.
	ExitPause_f(c *Pause_fContext)
}
