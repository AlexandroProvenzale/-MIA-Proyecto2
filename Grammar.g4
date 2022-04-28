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
}

// Rules

start:  (comando)+
;

comando: mkdisk_f NEWLINE
       | rmdisk_f NEWLINE
       | fdisk_f NEWLINE
       | mount_f NEWLINE
       | mkfs_f NEWLINE
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

// Parametros
SIZE:       '-' S I Z E;
FIT:        '-' F I T;
UNIT:      '-' U N I T;
PATH:       '-' P A T H;
TYPE:       '-' T Y P E;
DELETEP:    '-' D E L E T E;
NAME:       '-' N A M E;
ADD:        '-' A D D;
ID:         '-' I D;
FS:         '-' F S;
USR:        '-' U S U A R I O;
PASS:       '-' P A S S W O R D;

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

PATH1:          '/' IDENTIFICADOR ('/' IDENTIFICADOR)* ('.dk');
PATH2:          '"' '/' IDENTIFICADOR ((' ')* COMPLEMENTO)* ('/' IDENTIFICADOR ((' ')* COMPLEMENTO)*)* ('.dk') '"';

IGUAL: '=';

ENTERO:         [0-9]+;
NEGATIVO:       '-' ENTERO;
IDENTIFICADOR:  [a-zA-Z][a-zA-Z0-9_]*;
COMPLEMENTO: [a-zA-Z0-9][a-zA-Z0-9_]*;

NEWLINE:'\r'? '\n' | COMENTARIO | EOF;
WHITESPACE: [ \t]+ -> skip;
COMENTARIO: '#' (.*? '\r'? '\n') -> skip ;
