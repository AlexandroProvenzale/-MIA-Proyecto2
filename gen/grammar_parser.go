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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 56, 295,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 6, 2, 48, 10, 2, 13, 2, 14, 2, 49, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 89, 10, 3, 3, 4,
	3, 4, 6, 4, 93, 10, 4, 13, 4, 14, 4, 94, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 5, 5, 115, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	6, 7, 125, 10, 7, 13, 7, 14, 7, 126, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 155, 10, 8, 3, 9,
	3, 9, 6, 9, 159, 10, 9, 13, 9, 14, 9, 160, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 173, 10, 10, 3, 11, 3, 11,
	6, 11, 177, 10, 11, 13, 11, 14, 11, 178, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 191, 10, 12, 3, 13, 3, 13,
	6, 13, 195, 10, 13, 13, 13, 14, 13, 196, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	213, 10, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 6, 17, 226, 10, 17, 13, 17, 14, 17, 227, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 250, 10, 19,
	3, 20, 3, 20, 6, 20, 254, 10, 20, 13, 20, 14, 20, 255, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 267, 10, 21, 3, 22,
	3, 22, 6, 22, 271, 10, 22, 13, 22, 14, 22, 272, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 293, 10, 23, 3, 23, 2, 2, 24, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 2, 4, 4, 2, 49, 49, 51, 53, 3, 2, 4, 5, 2, 311, 2, 47, 3, 2, 2, 2,
	4, 88, 3, 2, 2, 2, 6, 90, 3, 2, 2, 2, 8, 114, 3, 2, 2, 2, 10, 116, 3, 2,
	2, 2, 12, 122, 3, 2, 2, 2, 14, 154, 3, 2, 2, 2, 16, 156, 3, 2, 2, 2, 18,
	172, 3, 2, 2, 2, 20, 174, 3, 2, 2, 2, 22, 190, 3, 2, 2, 2, 24, 192, 3,
	2, 2, 2, 26, 212, 3, 2, 2, 2, 28, 214, 3, 2, 2, 2, 30, 217, 3, 2, 2, 2,
	32, 223, 3, 2, 2, 2, 34, 231, 3, 2, 2, 2, 36, 249, 3, 2, 2, 2, 38, 251,
	3, 2, 2, 2, 40, 266, 3, 2, 2, 2, 42, 268, 3, 2, 2, 2, 44, 292, 3, 2, 2,
	2, 46, 48, 5, 4, 3, 2, 47, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 47,
	3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 3, 3, 2, 2, 2, 51, 52, 5, 6, 4, 2,
	52, 53, 7, 54, 2, 2, 53, 89, 3, 2, 2, 2, 54, 55, 5, 10, 6, 2, 55, 56, 7,
	54, 2, 2, 56, 89, 3, 2, 2, 2, 57, 58, 5, 12, 7, 2, 58, 59, 7, 54, 2, 2,
	59, 89, 3, 2, 2, 2, 60, 61, 5, 16, 9, 2, 61, 62, 7, 54, 2, 2, 62, 89, 3,
	2, 2, 2, 63, 64, 5, 20, 11, 2, 64, 65, 7, 54, 2, 2, 65, 89, 3, 2, 2, 2,
	66, 67, 5, 24, 13, 2, 67, 68, 7, 54, 2, 2, 68, 89, 3, 2, 2, 2, 69, 70,
	5, 28, 15, 2, 70, 71, 7, 54, 2, 2, 71, 89, 3, 2, 2, 2, 72, 73, 5, 30, 16,
	2, 73, 74, 7, 54, 2, 2, 74, 89, 3, 2, 2, 2, 75, 76, 5, 32, 17, 2, 76, 77,
	7, 54, 2, 2, 77, 89, 3, 2, 2, 2, 78, 79, 5, 34, 18, 2, 79, 80, 7, 54, 2,
	2, 80, 89, 3, 2, 2, 2, 81, 82, 5, 38, 20, 2, 82, 83, 7, 54, 2, 2, 83, 89,
	3, 2, 2, 2, 84, 85, 5, 42, 22, 2, 85, 86, 7, 54, 2, 2, 86, 89, 3, 2, 2,
	2, 87, 89, 7, 54, 2, 2, 88, 51, 3, 2, 2, 2, 88, 54, 3, 2, 2, 2, 88, 57,
	3, 2, 2, 2, 88, 60, 3, 2, 2, 2, 88, 63, 3, 2, 2, 2, 88, 66, 3, 2, 2, 2,
	88, 69, 3, 2, 2, 2, 88, 72, 3, 2, 2, 2, 88, 75, 3, 2, 2, 2, 88, 78, 3,
	2, 2, 2, 88, 81, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 88, 87, 3, 2, 2, 2, 89,
	5, 3, 2, 2, 2, 90, 92, 7, 9, 2, 2, 91, 93, 5, 8, 5, 2, 92, 91, 3, 2, 2,
	2, 93, 94, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96,
	3, 2, 2, 2, 96, 97, 8, 4, 1, 2, 97, 7, 3, 2, 2, 2, 98, 99, 7, 21, 2, 2,
	99, 100, 7, 48, 2, 2, 100, 101, 7, 49, 2, 2, 101, 115, 8, 5, 1, 2, 102,
	103, 7, 24, 2, 2, 103, 104, 7, 48, 2, 2, 104, 105, 7, 39, 2, 2, 105, 115,
	8, 5, 1, 2, 106, 107, 7, 22, 2, 2, 107, 108, 7, 48, 2, 2, 108, 109, 7,
	36, 2, 2, 109, 115, 8, 5, 1, 2, 110, 111, 7, 23, 2, 2, 111, 112, 7, 48,
	2, 2, 112, 113, 7, 37, 2, 2, 113, 115, 8, 5, 1, 2, 114, 98, 3, 2, 2, 2,
	114, 102, 3, 2, 2, 2, 114, 106, 3, 2, 2, 2, 114, 110, 3, 2, 2, 2, 115,
	9, 3, 2, 2, 2, 116, 117, 7, 10, 2, 2, 117, 118, 7, 24, 2, 2, 118, 119,
	7, 48, 2, 2, 119, 120, 7, 39, 2, 2, 120, 121, 8, 6, 1, 2, 121, 11, 3, 2,
	2, 2, 122, 124, 7, 11, 2, 2, 123, 125, 5, 14, 8, 2, 124, 123, 3, 2, 2,
	2, 125, 126, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 129, 8, 7, 1, 2, 129, 13, 3, 2, 2, 2, 130, 131, 7,
	21, 2, 2, 131, 132, 7, 48, 2, 2, 132, 133, 7, 49, 2, 2, 133, 155, 8, 8,
	1, 2, 134, 135, 7, 24, 2, 2, 135, 136, 7, 48, 2, 2, 136, 137, 7, 39, 2,
	2, 137, 155, 8, 8, 1, 2, 138, 139, 7, 22, 2, 2, 139, 140, 7, 48, 2, 2,
	140, 141, 7, 36, 2, 2, 141, 155, 8, 8, 1, 2, 142, 143, 7, 23, 2, 2, 143,
	144, 7, 48, 2, 2, 144, 145, 7, 37, 2, 2, 145, 155, 8, 8, 1, 2, 146, 147,
	7, 25, 2, 2, 147, 148, 7, 48, 2, 2, 148, 149, 7, 38, 2, 2, 149, 155, 8,
	8, 1, 2, 150, 151, 7, 27, 2, 2, 151, 152, 7, 48, 2, 2, 152, 153, 7, 51,
	2, 2, 153, 155, 8, 8, 1, 2, 154, 130, 3, 2, 2, 2, 154, 134, 3, 2, 2, 2,
	154, 138, 3, 2, 2, 2, 154, 142, 3, 2, 2, 2, 154, 146, 3, 2, 2, 2, 154,
	150, 3, 2, 2, 2, 155, 15, 3, 2, 2, 2, 156, 158, 7, 12, 2, 2, 157, 159,
	5, 18, 10, 2, 158, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 158, 3,
	2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 163, 8, 9, 1,
	2, 163, 17, 3, 2, 2, 2, 164, 165, 7, 24, 2, 2, 165, 166, 7, 48, 2, 2, 166,
	167, 7, 39, 2, 2, 167, 173, 8, 10, 1, 2, 168, 169, 7, 27, 2, 2, 169, 170,
	7, 48, 2, 2, 170, 171, 7, 51, 2, 2, 171, 173, 8, 10, 1, 2, 172, 164, 3,
	2, 2, 2, 172, 168, 3, 2, 2, 2, 173, 19, 3, 2, 2, 2, 174, 176, 7, 14, 2,
	2, 175, 177, 5, 22, 12, 2, 176, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2,
	178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180,
	181, 8, 11, 1, 2, 181, 21, 3, 2, 2, 2, 182, 183, 7, 29, 2, 2, 183, 184,
	7, 48, 2, 2, 184, 185, 7, 41, 2, 2, 185, 191, 8, 12, 1, 2, 186, 187, 7,
	25, 2, 2, 187, 188, 7, 48, 2, 2, 188, 189, 7, 38, 2, 2, 189, 191, 8, 12,
	1, 2, 190, 182, 3, 2, 2, 2, 190, 186, 3, 2, 2, 2, 191, 23, 3, 2, 2, 2,
	192, 194, 7, 15, 2, 2, 193, 195, 5, 26, 14, 2, 194, 193, 3, 2, 2, 2, 195,
	196, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198,
	3, 2, 2, 2, 198, 199, 8, 13, 1, 2, 199, 25, 3, 2, 2, 2, 200, 201, 7, 31,
	2, 2, 201, 202, 7, 48, 2, 2, 202, 203, 9, 2, 2, 2, 203, 213, 8, 14, 1,
	2, 204, 205, 7, 32, 2, 2, 205, 206, 7, 48, 2, 2, 206, 207, 9, 2, 2, 2,
	207, 213, 8, 14, 1, 2, 208, 209, 7, 29, 2, 2, 209, 210, 7, 48, 2, 2, 210,
	211, 7, 41, 2, 2, 211, 213, 8, 14, 1, 2, 212, 200, 3, 2, 2, 2, 212, 204,
	3, 2, 2, 2, 212, 208, 3, 2, 2, 2, 213, 27, 3, 2, 2, 2, 214, 215, 7, 16,
	2, 2, 215, 216, 8, 15, 1, 2, 216, 29, 3, 2, 2, 2, 217, 218, 7, 17, 2, 2,
	218, 219, 7, 27, 2, 2, 219, 220, 7, 48, 2, 2, 220, 221, 9, 2, 2, 2, 221,
	222, 8, 16, 1, 2, 222, 31, 3, 2, 2, 2, 223, 225, 7, 18, 2, 2, 224, 226,
	5, 36, 19, 2, 225, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 225, 3,
	2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 8, 17, 1,
	2, 230, 33, 3, 2, 2, 2, 231, 232, 7, 19, 2, 2, 232, 233, 7, 31, 2, 2, 233,
	234, 7, 48, 2, 2, 234, 235, 9, 2, 2, 2, 235, 236, 8, 18, 1, 2, 236, 35,
	3, 2, 2, 2, 237, 238, 7, 31, 2, 2, 238, 239, 7, 48, 2, 2, 239, 240, 9,
	2, 2, 2, 240, 250, 8, 19, 1, 2, 241, 242, 7, 33, 2, 2, 242, 243, 7, 48,
	2, 2, 243, 244, 9, 2, 2, 2, 244, 250, 8, 19, 1, 2, 245, 246, 7, 34, 2,
	2, 246, 247, 7, 48, 2, 2, 247, 248, 9, 2, 2, 2, 248, 250, 8, 19, 1, 2,
	249, 237, 3, 2, 2, 2, 249, 241, 3, 2, 2, 2, 249, 245, 3, 2, 2, 2, 250,
	37, 3, 2, 2, 2, 251, 253, 7, 20, 2, 2, 252, 254, 5, 40, 21, 2, 253, 252,
	3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 255, 256, 3, 2,
	2, 2, 256, 257, 3, 2, 2, 2, 257, 258, 8, 20, 1, 2, 258, 39, 3, 2, 2, 2,
	259, 260, 7, 24, 2, 2, 260, 261, 7, 48, 2, 2, 261, 262, 7, 40, 2, 2, 262,
	267, 8, 21, 1, 2, 263, 264, 7, 3, 2, 2, 264, 265, 9, 3, 2, 2, 265, 267,
	8, 21, 1, 2, 266, 259, 3, 2, 2, 2, 266, 263, 3, 2, 2, 2, 267, 41, 3, 2,
	2, 2, 268, 270, 7, 8, 2, 2, 269, 271, 5, 44, 23, 2, 270, 269, 3, 2, 2,
	2, 271, 272, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273,
	274, 3, 2, 2, 2, 274, 275, 8, 22, 1, 2, 275, 43, 3, 2, 2, 2, 276, 277,
	7, 27, 2, 2, 277, 278, 7, 48, 2, 2, 278, 279, 7, 42, 2, 2, 279, 293, 8,
	23, 1, 2, 280, 281, 7, 24, 2, 2, 281, 282, 7, 48, 2, 2, 282, 283, 7, 39,
	2, 2, 283, 293, 8, 23, 1, 2, 284, 285, 7, 29, 2, 2, 285, 286, 7, 48, 2,
	2, 286, 287, 7, 41, 2, 2, 287, 293, 8, 23, 1, 2, 288, 289, 7, 35, 2, 2,
	289, 290, 7, 48, 2, 2, 290, 291, 7, 39, 2, 2, 291, 293, 8, 23, 1, 2, 292,
	276, 3, 2, 2, 2, 292, 280, 3, 2, 2, 2, 292, 284, 3, 2, 2, 2, 292, 288,
	3, 2, 2, 2, 293, 45, 3, 2, 2, 2, 20, 49, 88, 94, 114, 126, 154, 160, 172,
	178, 190, 196, 212, 227, 249, 255, 266, 272, 292,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'-'", "'p'", "'P'", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "PAUSES", "EXEC", "REP", "MKDISK", "RMDISK", "FDISK", "MOUNT",
	"UNMOUNT", "MKFS", "LOGIN", "LOGOUT", "MKGRP", "MKUSR", "RMUSR", "MKDIR",
	"SIZE", "FIT", "UNIT", "PATH", "TYPE", "DELETEP", "NAME", "ADD", "ID",
	"FS", "USR", "PASSW", "PWD", "GRP", "RUTA", "E_FIT", "E_UNIT", "E_TYPE",
	"E_PATH", "E_PATHH", "E_ID", "E_REP", "PATH1", "PATH2", "TERMINAL", "PATHH1",
	"PATHH2", "IGUAL", "ENTERO", "NEGATIVO", "IDENTIFICADOR", "COMPLEMENTO",
	"E_USRS", "NEWLINE", "WHITESPACE", "COMENTARIO",
}

var ruleNames = []string{
	"start", "comando", "mkdisk_f", "mkparam", "rmdisk_f", "fdisk_f", "fdiskparam",
	"mount_f", "mountparam", "mkfs_f", "mkfsparam", "login_f", "loginparam",
	"logout_f", "mkgroup_f", "mkuser_f", "rmuser_f", "mkuserparam", "mkdir_f",
	"mkdirparam", "rep_f", "repparam",
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
var info_MKDIR Program.InfoMkdir
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

func initializeMKDIR(MKDIR *Program.InfoMkdir) {
	MKDIR.Path = ""
	MKDIR.P = false
}

func initializeREP(REP *Program.InfoRep) {
	REP.Name = ""
	REP.Path = ""
	REP.Id = ""
	REP.Ruta = ""
}

// GrammarParser tokens.
const (
	GrammarParserEOF           = antlr.TokenEOF
	GrammarParserT__0          = 1
	GrammarParserT__1          = 2
	GrammarParserT__2          = 3
	GrammarParserPAUSES        = 4
	GrammarParserEXEC          = 5
	GrammarParserREP           = 6
	GrammarParserMKDISK        = 7
	GrammarParserRMDISK        = 8
	GrammarParserFDISK         = 9
	GrammarParserMOUNT         = 10
	GrammarParserUNMOUNT       = 11
	GrammarParserMKFS          = 12
	GrammarParserLOGIN         = 13
	GrammarParserLOGOUT        = 14
	GrammarParserMKGRP         = 15
	GrammarParserMKUSR         = 16
	GrammarParserRMUSR         = 17
	GrammarParserMKDIR         = 18
	GrammarParserSIZE          = 19
	GrammarParserFIT           = 20
	GrammarParserUNIT          = 21
	GrammarParserPATH          = 22
	GrammarParserTYPE          = 23
	GrammarParserDELETEP       = 24
	GrammarParserNAME          = 25
	GrammarParserADD           = 26
	GrammarParserID            = 27
	GrammarParserFS            = 28
	GrammarParserUSR           = 29
	GrammarParserPASSW         = 30
	GrammarParserPWD           = 31
	GrammarParserGRP           = 32
	GrammarParserRUTA          = 33
	GrammarParserE_FIT         = 34
	GrammarParserE_UNIT        = 35
	GrammarParserE_TYPE        = 36
	GrammarParserE_PATH        = 37
	GrammarParserE_PATHH       = 38
	GrammarParserE_ID          = 39
	GrammarParserE_REP         = 40
	GrammarParserPATH1         = 41
	GrammarParserPATH2         = 42
	GrammarParserTERMINAL      = 43
	GrammarParserPATHH1        = 44
	GrammarParserPATHH2        = 45
	GrammarParserIGUAL         = 46
	GrammarParserENTERO        = 47
	GrammarParserNEGATIVO      = 48
	GrammarParserIDENTIFICADOR = 49
	GrammarParserCOMPLEMENTO   = 50
	GrammarParserE_USRS        = 51
	GrammarParserNEWLINE       = 52
	GrammarParserWHITESPACE    = 53
	GrammarParserCOMENTARIO    = 54
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
	GrammarParserRULE_rmuser_f    = 16
	GrammarParserRULE_mkuserparam = 17
	GrammarParserRULE_mkdir_f     = 18
	GrammarParserRULE_mkdirparam  = 19
	GrammarParserRULE_rep_f       = 20
	GrammarParserRULE_repparam    = 21
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserREP)|(1<<GrammarParserMKDISK)|(1<<GrammarParserRMDISK)|(1<<GrammarParserFDISK)|(1<<GrammarParserMOUNT)|(1<<GrammarParserMKFS)|(1<<GrammarParserLOGIN)|(1<<GrammarParserLOGOUT)|(1<<GrammarParserMKGRP)|(1<<GrammarParserMKUSR)|(1<<GrammarParserRMUSR)|(1<<GrammarParserMKDIR))) != 0) || _la == GrammarParserNEWLINE {
		{
			p.SetState(44)
			p.Comando()
		}

		p.SetState(47)
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

func (s *ComandoContext) Rmuser_f() IRmuser_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRmuser_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRmuser_fContext)
}

func (s *ComandoContext) Mkdir_f() IMkdir_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkdir_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkdir_fContext)
}

func (s *ComandoContext) Rep_f() IRep_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRep_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRep_fContext)
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

	p.SetState(86)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserMKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(49)
			p.Mkdisk_f()
		}
		{
			p.SetState(50)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserRMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Rmdisk_f()
		}
		{
			p.SetState(53)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserFDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Fdisk_f()
		}
		{
			p.SetState(56)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(58)
			p.Mount_f()
		}
		{
			p.SetState(59)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKFS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(61)
			p.Mkfs_f()
		}
		{
			p.SetState(62)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(64)
			p.Login_f()
		}
		{
			p.SetState(65)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(67)
			p.Logout_f()
		}
		{
			p.SetState(68)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKGRP:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(70)
			p.Mkgroup_f()
		}
		{
			p.SetState(71)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKUSR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(73)
			p.Mkuser_f()
		}
		{
			p.SetState(74)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserRMUSR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(76)
			p.Rmuser_f()
		}
		{
			p.SetState(77)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKDIR:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(79)
			p.Mkdir_f()
		}
		{
			p.SetState(80)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserREP:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(82)
			p.Rep_f()
		}
		{
			p.SetState(83)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserNEWLINE:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(85)
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
		p.SetState(88)
		p.Match(GrammarParserMKDISK)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH))) != 0) {
		{
			p.SetState(89)
			p.Mkparam()
		}

		p.SetState(92)
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

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(97)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(98)

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
			p.SetState(100)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(101)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(102)

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
			p.SetState(104)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(105)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(106)

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
			p.SetState(108)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(109)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(110)

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
		p.SetState(114)
		p.Match(GrammarParserRMDISK)
	}
	{
		p.SetState(115)
		p.Match(GrammarParserPATH)
	}
	{
		p.SetState(116)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(117)

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
		p.SetState(120)
		p.Match(GrammarParserFDISK)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH)|(1<<GrammarParserTYPE)|(1<<GrammarParserNAME))) != 0) {
		{
			p.SetState(121)
			p.Fdiskparam()
		}

		p.SetState(124)
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

	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(129)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(130)

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
			p.SetState(132)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(133)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(134)

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
			p.SetState(136)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(137)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(138)

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
			p.SetState(140)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(141)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(142)

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
			p.SetState(144)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(145)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(146)

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
			p.SetState(148)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(149)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(150)

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
		p.SetState(154)
		p.Match(GrammarParserMOUNT)
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserPATH || _la == GrammarParserNAME {
		{
			p.SetState(155)
			p.Mountparam()
		}

		p.SetState(158)
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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(163)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(164)

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
			p.SetState(166)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(167)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(168)

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
		p.SetState(172)
		p.Match(GrammarParserMKFS)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserTYPE || _la == GrammarParserID {
		{
			p.SetState(173)
			p.Mkfsparam()
		}

		p.SetState(176)
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

	p.SetState(188)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(181)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(182)

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
			p.SetState(184)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(185)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(186)

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
		p.SetState(190)
		p.Match(GrammarParserLOGIN)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserID)|(1<<GrammarParserUSR)|(1<<GrammarParserPASSW))) != 0) {
		{
			p.SetState(191)
			p.Loginparam()
		}

		p.SetState(194)
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

	p.SetState(210)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserUSR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Match(GrammarParserUSR)
		}
		{
			p.SetState(199)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(200)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LoginparamContext).e_user = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(GrammarParserENTERO-47))|(1<<(GrammarParserIDENTIFICADOR-47))|(1<<(GrammarParserCOMPLEMENTO-47))|(1<<(GrammarParserE_USRS-47)))) != 0) {
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
			p.SetState(202)
			p.Match(GrammarParserPASSW)
		}
		{
			p.SetState(203)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(204)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LoginparamContext).e_pass = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(GrammarParserENTERO-47))|(1<<(GrammarParserIDENTIFICADOR-47))|(1<<(GrammarParserCOMPLEMENTO-47))|(1<<(GrammarParserE_USRS-47)))) != 0) {
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
			p.SetState(206)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(207)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(208)

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
		p.SetState(212)
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
		p.SetState(215)
		p.Match(GrammarParserMKGRP)
	}
	{
		p.SetState(216)
		p.Match(GrammarParserNAME)
	}
	{
		p.SetState(217)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(218)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Mkgroup_fContext).e_name = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(GrammarParserENTERO-47))|(1<<(GrammarParserIDENTIFICADOR-47))|(1<<(GrammarParserCOMPLEMENTO-47))|(1<<(GrammarParserE_USRS-47)))) != 0) {
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
		p.SetState(221)
		p.Match(GrammarParserMKUSR)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-29)&-(0x1f+1)) == 0 && ((1<<uint((_la-29)))&((1<<(GrammarParserUSR-29))|(1<<(GrammarParserPWD-29))|(1<<(GrammarParserGRP-29)))) != 0) {
		{
			p.SetState(222)
			p.Mkuserparam()
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.Mkuser(info_MKUSER)
	initializeMKUSER(&info_MKUSER)

	return localctx
}

// IRmuser_fContext is an interface to support dynamic dispatch.
type IRmuser_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE_name returns the e_name token.
	GetE_name() antlr.Token

	// SetE_name sets the e_name token.
	SetE_name(antlr.Token)

	// IsRmuser_fContext differentiates from other interfaces.
	IsRmuser_fContext()
}

type Rmuser_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	e_name antlr.Token
}

func NewEmptyRmuser_fContext() *Rmuser_fContext {
	var p = new(Rmuser_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_rmuser_f
	return p
}

func (*Rmuser_fContext) IsRmuser_fContext() {}

func NewRmuser_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rmuser_fContext {
	var p = new(Rmuser_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_rmuser_f

	return p
}

func (s *Rmuser_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Rmuser_fContext) GetE_name() antlr.Token { return s.e_name }

func (s *Rmuser_fContext) SetE_name(v antlr.Token) { s.e_name = v }

func (s *Rmuser_fContext) RMUSR() antlr.TerminalNode {
	return s.GetToken(GrammarParserRMUSR, 0)
}

func (s *Rmuser_fContext) USR() antlr.TerminalNode {
	return s.GetToken(GrammarParserUSR, 0)
}

func (s *Rmuser_fContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *Rmuser_fContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFICADOR, 0)
}

func (s *Rmuser_fContext) COMPLEMENTO() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMPLEMENTO, 0)
}

func (s *Rmuser_fContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *Rmuser_fContext) E_USRS() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_USRS, 0)
}

func (s *Rmuser_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rmuser_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rmuser_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRmuser_f(s)
	}
}

func (s *Rmuser_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRmuser_f(s)
	}
}

func (p *GrammarParser) Rmuser_f() (localctx IRmuser_fContext) {
	localctx = NewRmuser_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GrammarParserRULE_rmuser_f)
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
		p.SetState(229)
		p.Match(GrammarParserRMUSR)
	}
	{
		p.SetState(230)
		p.Match(GrammarParserUSR)
	}
	{
		p.SetState(231)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(232)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Rmuser_fContext).e_name = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(GrammarParserENTERO-47))|(1<<(GrammarParserIDENTIFICADOR-47))|(1<<(GrammarParserCOMPLEMENTO-47))|(1<<(GrammarParserE_USRS-47)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Rmuser_fContext).e_name = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	Program.Rmuser(strings.ReplaceAll((func() string {
		if localctx.(*Rmuser_fContext).GetE_name() == nil {
			return ""
		} else {
			return localctx.(*Rmuser_fContext).GetE_name().GetText()
		}
	}()), "\"", ""))

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
	p.EnterRule(localctx, 34, GrammarParserRULE_mkuserparam)
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

	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserUSR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Match(GrammarParserUSR)
		}
		{
			p.SetState(236)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(237)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_user = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(GrammarParserENTERO-47))|(1<<(GrammarParserIDENTIFICADOR-47))|(1<<(GrammarParserCOMPLEMENTO-47))|(1<<(GrammarParserE_USRS-47)))) != 0) {
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
			p.SetState(239)
			p.Match(GrammarParserPWD)
		}
		{
			p.SetState(240)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(241)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_pass = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(GrammarParserENTERO-47))|(1<<(GrammarParserIDENTIFICADOR-47))|(1<<(GrammarParserCOMPLEMENTO-47))|(1<<(GrammarParserE_USRS-47)))) != 0) {
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
			p.SetState(243)
			p.Match(GrammarParserGRP)
		}
		{
			p.SetState(244)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(245)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_group = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-47)&-(0x1f+1)) == 0 && ((1<<uint((_la-47)))&((1<<(GrammarParserENTERO-47))|(1<<(GrammarParserIDENTIFICADOR-47))|(1<<(GrammarParserCOMPLEMENTO-47))|(1<<(GrammarParserE_USRS-47)))) != 0) {
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

// IMkdir_fContext is an interface to support dynamic dispatch.
type IMkdir_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMkdir_fContext differentiates from other interfaces.
	IsMkdir_fContext()
}

type Mkdir_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdir_fContext() *Mkdir_fContext {
	var p = new(Mkdir_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkdir_f
	return p
}

func (*Mkdir_fContext) IsMkdir_fContext() {}

func NewMkdir_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkdir_fContext {
	var p = new(Mkdir_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkdir_f

	return p
}

func (s *Mkdir_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkdir_fContext) MKDIR() antlr.TerminalNode {
	return s.GetToken(GrammarParserMKDIR, 0)
}

func (s *Mkdir_fContext) AllMkdirparam() []IMkdirparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMkdirparamContext)(nil)).Elem())
	var tst = make([]IMkdirparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMkdirparamContext)
		}
	}

	return tst
}

func (s *Mkdir_fContext) Mkdirparam(i int) IMkdirparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkdirparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMkdirparamContext)
}

func (s *Mkdir_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkdir_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkdir_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkdir_f(s)
	}
}

func (s *Mkdir_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkdir_f(s)
	}
}

func (p *GrammarParser) Mkdir_f() (localctx IMkdir_fContext) {
	localctx = NewMkdir_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GrammarParserRULE_mkdir_f)
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
		p.SetState(249)
		p.Match(GrammarParserMKDIR)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserT__0 || _la == GrammarParserPATH {
		{
			p.SetState(250)
			p.Mkdirparam()
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.EscribirDirectorio(info_MKDIR)
	initializeMKDIR(&info_MKDIR)

	return localctx
}

// IMkdirparamContext is an interface to support dynamic dispatch.
type IMkdirparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_E_PATHH returns the _E_PATHH token.
	Get_E_PATHH() antlr.Token

	// Set_E_PATHH sets the _E_PATHH token.
	Set_E_PATHH(antlr.Token)

	// IsMkdirparamContext differentiates from other interfaces.
	IsMkdirparamContext()
}

type MkdirparamContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_E_PATHH antlr.Token
}

func NewEmptyMkdirparamContext() *MkdirparamContext {
	var p = new(MkdirparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkdirparam
	return p
}

func (*MkdirparamContext) IsMkdirparamContext() {}

func NewMkdirparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkdirparamContext {
	var p = new(MkdirparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkdirparam

	return p
}

func (s *MkdirparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkdirparamContext) Get_E_PATHH() antlr.Token { return s._E_PATHH }

func (s *MkdirparamContext) Set_E_PATHH(v antlr.Token) { s._E_PATHH = v }

func (s *MkdirparamContext) PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserPATH, 0)
}

func (s *MkdirparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *MkdirparamContext) E_PATHH() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_PATHH, 0)
}

func (s *MkdirparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkdirparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkdirparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkdirparam(s)
	}
}

func (s *MkdirparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkdirparam(s)
	}
}

func (p *GrammarParser) Mkdirparam() (localctx IMkdirparamContext) {
	localctx = NewMkdirparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrammarParserRULE_mkdirparam)
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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(258)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(259)

			var _m = p.Match(GrammarParserE_PATHH)

			localctx.(*MkdirparamContext)._E_PATHH = _m
		}
		info_MKDIR.Path = strings.ReplaceAll((func() string {
			if localctx.(*MkdirparamContext).Get_E_PATHH() == nil {
				return ""
			} else {
				return localctx.(*MkdirparamContext).Get_E_PATHH().GetText()
			}
		}()), "\"", "")

	case GrammarParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Match(GrammarParserT__0)
		}
		{
			p.SetState(262)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__1 || _la == GrammarParserT__2) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_MKDIR.P = true

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRep_fContext is an interface to support dynamic dispatch.
type IRep_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRep_fContext differentiates from other interfaces.
	IsRep_fContext()
}

type Rep_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRep_fContext() *Rep_fContext {
	var p = new(Rep_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_rep_f
	return p
}

func (*Rep_fContext) IsRep_fContext() {}

func NewRep_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rep_fContext {
	var p = new(Rep_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_rep_f

	return p
}

func (s *Rep_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Rep_fContext) REP() antlr.TerminalNode {
	return s.GetToken(GrammarParserREP, 0)
}

func (s *Rep_fContext) AllRepparam() []IRepparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRepparamContext)(nil)).Elem())
	var tst = make([]IRepparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRepparamContext)
		}
	}

	return tst
}

func (s *Rep_fContext) Repparam(i int) IRepparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRepparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRepparamContext)
}

func (s *Rep_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rep_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rep_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRep_f(s)
	}
}

func (s *Rep_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRep_f(s)
	}
}

func (p *GrammarParser) Rep_f() (localctx IRep_fContext) {
	localctx = NewRep_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrammarParserRULE_rep_f)
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
		p.SetState(266)
		p.Match(GrammarParserREP)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(GrammarParserPATH-22))|(1<<(GrammarParserNAME-22))|(1<<(GrammarParserID-22))|(1<<(GrammarParserRUTA-22)))) != 0) {
		{
			p.SetState(267)
			p.Repparam()
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.Reportar(info_REP)
	initializeREP(&info_REP)

	return localctx
}

// IRepparamContext is an interface to support dynamic dispatch.
type IRepparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_E_REP returns the _E_REP token.
	Get_E_REP() antlr.Token

	// Get_E_PATH returns the _E_PATH token.
	Get_E_PATH() antlr.Token

	// Get_E_ID returns the _E_ID token.
	Get_E_ID() antlr.Token

	// Set_E_REP sets the _E_REP token.
	Set_E_REP(antlr.Token)

	// Set_E_PATH sets the _E_PATH token.
	Set_E_PATH(antlr.Token)

	// Set_E_ID sets the _E_ID token.
	Set_E_ID(antlr.Token)

	// IsRepparamContext differentiates from other interfaces.
	IsRepparamContext()
}

type RepparamContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_E_REP  antlr.Token
	_E_PATH antlr.Token
	_E_ID   antlr.Token
}

func NewEmptyRepparamContext() *RepparamContext {
	var p = new(RepparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_repparam
	return p
}

func (*RepparamContext) IsRepparamContext() {}

func NewRepparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RepparamContext {
	var p = new(RepparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_repparam

	return p
}

func (s *RepparamContext) GetParser() antlr.Parser { return s.parser }

func (s *RepparamContext) Get_E_REP() antlr.Token { return s._E_REP }

func (s *RepparamContext) Get_E_PATH() antlr.Token { return s._E_PATH }

func (s *RepparamContext) Get_E_ID() antlr.Token { return s._E_ID }

func (s *RepparamContext) Set_E_REP(v antlr.Token) { s._E_REP = v }

func (s *RepparamContext) Set_E_PATH(v antlr.Token) { s._E_PATH = v }

func (s *RepparamContext) Set_E_ID(v antlr.Token) { s._E_ID = v }

func (s *RepparamContext) NAME() antlr.TerminalNode {
	return s.GetToken(GrammarParserNAME, 0)
}

func (s *RepparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *RepparamContext) E_REP() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_REP, 0)
}

func (s *RepparamContext) PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserPATH, 0)
}

func (s *RepparamContext) E_PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_PATH, 0)
}

func (s *RepparamContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *RepparamContext) E_ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_ID, 0)
}

func (s *RepparamContext) RUTA() antlr.TerminalNode {
	return s.GetToken(GrammarParserRUTA, 0)
}

func (s *RepparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RepparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRepparam(s)
	}
}

func (s *RepparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRepparam(s)
	}
}

func (p *GrammarParser) Repparam() (localctx IRepparamContext) {
	localctx = NewRepparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GrammarParserRULE_repparam)

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

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(275)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(276)

			var _m = p.Match(GrammarParserE_REP)

			localctx.(*RepparamContext)._E_REP = _m
		}
		info_REP.Name = (func() string {
			if localctx.(*RepparamContext).Get_E_REP() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).Get_E_REP().GetText()
			}
		}())

	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(278)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(279)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(280)

			var _m = p.Match(GrammarParserE_PATH)

			localctx.(*RepparamContext)._E_PATH = _m
		}
		info_REP.Path = strings.ReplaceAll((func() string {
			if localctx.(*RepparamContext).Get_E_PATH() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).Get_E_PATH().GetText()
			}
		}()), "\"", "")

	case GrammarParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(283)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(284)

			var _m = p.Match(GrammarParserE_ID)

			localctx.(*RepparamContext)._E_ID = _m
		}
		info_REP.Id = (func() string {
			if localctx.(*RepparamContext).Get_E_ID() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).Get_E_ID().GetText()
			}
		}())

	case GrammarParserRUTA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(286)
			p.Match(GrammarParserRUTA)
		}
		{
			p.SetState(287)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(288)

			var _m = p.Match(GrammarParserE_PATH)

			localctx.(*RepparamContext)._E_PATH = _m
		}
		info_REP.Ruta = strings.ReplaceAll((func() string {
			if localctx.(*RepparamContext).Get_E_PATH() == nil {
				return ""
			} else {
				return localctx.(*RepparamContext).Get_E_PATH().GetText()
			}
		}()), "\"", "")

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
