grammar Grammar;

options {
    language='Go';
}

@parser::header {
    import "Proyecto2/Program"
    import "strings"
}

@parser::members{
    var info_MKDISK Program.InfoMKDisk
    var info_FDISK Program.InfoFDisk
    var info_MOUNT Program.InfoMount
    var info_MKFS Program.InfoMkfs
    var info_LOGIN Program.InfoLogin
    var info_MKUSER Program.InfoMkuser
    var info_REP Program.InfoRep

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

    func initializeREP(REP *Program.InfoRep) {
        REP.Name = ""
        REP.Path = ""
        REP.Id = ""
        REP.Ruta = ""
    }
}

// Rules

start:  (comando)+
;

comando: mkdisk_f NEWLINE
       | rmdisk_f NEWLINE
       | fdisk_f NEWLINE
       | mount_f NEWLINE
       | mkfs_f NEWLINE
       | login_f NEWLINE
       | logout_f NEWLINE
       | mkgroup_f NEWLINE
       | mkuser_f NEWLINE
       | rep_f NEWLINE
       | NEWLINE
;

mkdisk_f:
    MKDISK mkparam+     {
                         Program.CreateDisk(info_MKDISK)
                         initializeMKDISK(&info_MKDISK)
                        }
;

mkparam:
    SIZE IGUAL ENTERO   {info_MKDISK.Size = $ENTERO.int}
|   PATH IGUAL E_PATH   {info_MKDISK.Path = strings.ReplaceAll($E_PATH.text, "\"", "")}
|   FIT IGUAL E_FIT     {info_MKDISK.Fit = $E_FIT.text}
|   UNIT IGUAL E_UNIT   {info_MKDISK.Unit = $E_UNIT.text}
;

rmdisk_f:
RMDISK PATH IGUAL E_PATH     {Program.RemoveDisk(strings.ReplaceAll($E_PATH.text, "\"", ""))}
;

fdisk_f:
    FDISK fdiskparam+   {
                         Program.CrearParticion(info_FDISK)
                         initializeFDISK(&info_FDISK)
                        }
;

fdiskparam:
    SIZE IGUAL ENTERO                   {info_FDISK.Size = $ENTERO.int}
|   PATH IGUAL E_PATH                   {info_FDISK.Path = strings.ReplaceAll($E_PATH.text, "\"", "")}
|   FIT IGUAL E_FIT                     {info_FDISK.Fit = $E_FIT.text}
|   UNIT IGUAL E_UNIT                   {info_FDISK.Unit = $E_UNIT.text}
|   TYPE IGUAL E_TYPE                   {info_FDISK.Type = $E_TYPE.text}
|   NAME IGUAL e_name=IDENTIFICADOR     {info_FDISK.Name = $e_name.text}
;

mount_f:
    MOUNT mountparam+   {
                         Program.MontarDisco(info_MOUNT)
                         initializeMOUNT(&info_MOUNT)
                        }
;

mountparam:
    PATH IGUAL E_PATH                       {info_MOUNT.Path = strings.ReplaceAll($E_PATH.text, "\"", "")}
    |   NAME IGUAL e_name=IDENTIFICADOR     {info_MOUNT.Name = $e_name.text}
;

mkfs_f:
    MKFS mkfsparam+     {
                         Program.Formatear(info_MKFS)
                         initializeMKFS(&info_MKFS)
                        }
;

mkfsparam:
    ID IGUAL E_ID       {info_MKFS.Id = $E_ID.text}
|   TYPE IGUAL E_TYPE   {info_MKFS.Type = $E_TYPE.text}
;

login_f:
    LOGIN loginparam+   {
                         Program.LoginS(info_LOGIN)
                         initializeLOGIN(&info_LOGIN)
                        }
;

loginparam:
    USR IGUAL e_user=(IDENTIFICADOR|COMPLEMENTO|ENTERO|E_USRS)      {info_LOGIN.User = strings.ReplaceAll($e_user.text, "\"", "")}
|   PASSW IGUAL e_pass=(IDENTIFICADOR|COMPLEMENTO|ENTERO|E_USRS)    {info_LOGIN.Pass = strings.ReplaceAll($e_pass.text, "\"", "")}
|   ID IGUAL E_ID                                                   {info_LOGIN.Id = $E_ID.text}
;

logout_f: LOGOUT    {Program.LogoutS()}
;

mkgroup_f: MKGRP NAME IGUAL e_name=(IDENTIFICADOR|COMPLEMENTO|ENTERO|E_USRS)    {Program.Mkgroup(strings.ReplaceAll($e_name.text, "\"", ""))}
;

mkuser_f: MKUSR mkuserparam+    {
                                 Program.Mkuser(info_MKUSER)
                                 initializeMKUSER(&info_MKUSER)
                                }
;

mkuserparam:
    USR IGUAL e_user=(IDENTIFICADOR|COMPLEMENTO|ENTERO|E_USRS)      {info_MKUSER.User = strings.ReplaceAll($e_user.text, "\"", "")}
|   PWD IGUAL e_pass=(IDENTIFICADOR|COMPLEMENTO|ENTERO|E_USRS)      {info_MKUSER.Pass = strings.ReplaceAll($e_pass.text, "\"", "")}
|   GRP IGUAL e_group=(IDENTIFICADOR|COMPLEMENTO|ENTERO|E_USRS)     {info_MKUSER.Grp = strings.ReplaceAll($e_group.text, "\"", "")}
;

rep_f: REP repparam+    {
                         Program.Reportar(info_REP)
                         initializeREP(&info_REP)
                        }
;

repparam:
    NAME IGUAL E_REP    {info_REP.Name = $E_REP.text}
|   PATH IGUAL E_PATH   {info_REP.Path = strings.ReplaceAll($E_PATH.text, "\"", "")}
|   ID IGUAL E_ID       {info_REP.Id = $E_ID.text}
|   RUTA IGUAL E_PATH   {info_REP.Ruta = strings.ReplaceAll($E_PATH.text, "\"", "")}
;

// Tokens

// Comandos
fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];

PAUSES:     P A U S E;
EXEC:       E X E C;
REP:        R E P;
MKDISK:     M K D I S K;
RMDISK:     R M D I S K;
FDISK:      F D I S K;
MOUNT:      M O U N T;
UNMOUNT:    U N M O U N T;
MKFS:       M K F S;

LOGIN:      L O G I N;
LOGOUT:     L O G O U T;
MKGRP:      M K G R P;
MKUSR:      M K U S R;

// Parametros
SIZE:       '-' S I Z E;
FIT:        '-' F I T;
UNIT:       '-' U N I T;
PATH:       '-' P A T H;
TYPE:       '-' T Y P E;
DELETEP:    '-' D E L E T E;
NAME:       '-' N A M E;
ADD:        '-' A D D;
ID:         '-' I D;
FS:         '-' F S;
USR:        '-' U S U A R I O;
PASSW:      '-' P A S S W O R D;
PWD:        '-' P W D;
GRP:        '-' G R P;
RUTA:       '-' R U T A;

// Entradas
E_FIT:  B F
     |  F F
     |  W F
;
E_UNIT: K
     |  M
     |  B
;
E_TYPE: P
     |  E
     |  L
     |  F A S T
     |  F U L L
;
E_PATH: PATH1
     |  PATH2
;
E_ID:     [0-9]+[a-zA-Z]
;
E_REP: D I S K
    |  T R E E
    |  F I L E
;

PATH1:          '/' IDENTIFICADOR ('/' IDENTIFICADOR)* ('.dk');
PATH2:          '"' '/' IDENTIFICADOR ((' ')* COMPLEMENTO)* ('/' IDENTIFICADOR ((' ')* COMPLEMENTO)*)* TERMINAL '"';
TERMINAL:       '.' (D K | T X T | D S K | J P G | P N G);

IGUAL: '=';

ENTERO:         [0-9]+;
NEGATIVO:       '-' ENTERO;
IDENTIFICADOR:  [a-zA-Z][a-zA-Z0-9_]*;
COMPLEMENTO: [a-zA-Z0-9][a-zA-Z0-9_]*;
E_USRS: '"' COMPLEMENTO ((' ')* COMPLEMENTO)* '"';

NEWLINE:'\r'? '\n' | COMENTARIO | EOF;
WHITESPACE: [ \t]+ -> skip;
COMENTARIO: '#' (.*? '\r'? '\n') -> skip ;
