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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 61, 338,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 6, 2, 54, 10, 2, 13, 2,
	14, 2, 55, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 101, 10, 3,
	3, 4, 3, 4, 6, 4, 105, 10, 4, 13, 4, 14, 4, 106, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 127, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 6, 7, 137, 10, 7, 13, 7, 14, 7, 138, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 167, 10, 8,
	3, 9, 3, 9, 6, 9, 171, 10, 9, 13, 9, 14, 9, 172, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 185, 10, 10, 3, 11,
	3, 11, 6, 11, 189, 10, 11, 13, 11, 14, 11, 190, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 203, 10, 12, 3, 13,
	3, 13, 6, 13, 207, 10, 13, 13, 13, 14, 13, 208, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	5, 14, 225, 10, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 6, 18,
	244, 10, 18, 13, 18, 14, 18, 245, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 268, 10, 20, 3, 21, 3, 21, 6, 21, 272,
	10, 21, 13, 21, 14, 21, 273, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 5, 22, 293, 10, 22, 3, 23, 3, 23, 6, 23, 297, 10, 23, 13, 23, 14, 23,
	298, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5,
	24, 310, 10, 24, 3, 25, 3, 25, 6, 25, 314, 10, 25, 13, 25, 14, 25, 315,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 336, 10, 26,
	3, 26, 2, 2, 27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 2, 5, 4, 2, 54, 54, 56, 58, 3,
	2, 4, 5, 3, 2, 6, 7, 2, 357, 2, 53, 3, 2, 2, 2, 4, 100, 3, 2, 2, 2, 6,
	102, 3, 2, 2, 2, 8, 126, 3, 2, 2, 2, 10, 128, 3, 2, 2, 2, 12, 134, 3, 2,
	2, 2, 14, 166, 3, 2, 2, 2, 16, 168, 3, 2, 2, 2, 18, 184, 3, 2, 2, 2, 20,
	186, 3, 2, 2, 2, 22, 202, 3, 2, 2, 2, 24, 204, 3, 2, 2, 2, 26, 224, 3,
	2, 2, 2, 28, 226, 3, 2, 2, 2, 30, 229, 3, 2, 2, 2, 32, 235, 3, 2, 2, 2,
	34, 241, 3, 2, 2, 2, 36, 249, 3, 2, 2, 2, 38, 267, 3, 2, 2, 2, 40, 269,
	3, 2, 2, 2, 42, 292, 3, 2, 2, 2, 44, 294, 3, 2, 2, 2, 46, 309, 3, 2, 2,
	2, 48, 311, 3, 2, 2, 2, 50, 335, 3, 2, 2, 2, 52, 54, 5, 4, 3, 2, 53, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2,
	56, 3, 3, 2, 2, 2, 57, 58, 5, 6, 4, 2, 58, 59, 7, 59, 2, 2, 59, 101, 3,
	2, 2, 2, 60, 61, 5, 10, 6, 2, 61, 62, 7, 59, 2, 2, 62, 101, 3, 2, 2, 2,
	63, 64, 5, 12, 7, 2, 64, 65, 7, 59, 2, 2, 65, 101, 3, 2, 2, 2, 66, 67,
	5, 16, 9, 2, 67, 68, 7, 59, 2, 2, 68, 101, 3, 2, 2, 2, 69, 70, 5, 20, 11,
	2, 70, 71, 7, 59, 2, 2, 71, 101, 3, 2, 2, 2, 72, 73, 5, 24, 13, 2, 73,
	74, 7, 59, 2, 2, 74, 101, 3, 2, 2, 2, 75, 76, 5, 28, 15, 2, 76, 77, 7,
	59, 2, 2, 77, 101, 3, 2, 2, 2, 78, 79, 5, 30, 16, 2, 79, 80, 7, 59, 2,
	2, 80, 101, 3, 2, 2, 2, 81, 82, 5, 32, 17, 2, 82, 83, 7, 59, 2, 2, 83,
	101, 3, 2, 2, 2, 84, 85, 5, 34, 18, 2, 85, 86, 7, 59, 2, 2, 86, 101, 3,
	2, 2, 2, 87, 88, 5, 36, 19, 2, 88, 89, 7, 59, 2, 2, 89, 101, 3, 2, 2, 2,
	90, 91, 5, 40, 21, 2, 91, 92, 7, 59, 2, 2, 92, 101, 3, 2, 2, 2, 93, 94,
	5, 44, 23, 2, 94, 95, 7, 59, 2, 2, 95, 101, 3, 2, 2, 2, 96, 97, 5, 48,
	25, 2, 97, 98, 7, 59, 2, 2, 98, 101, 3, 2, 2, 2, 99, 101, 7, 59, 2, 2,
	100, 57, 3, 2, 2, 2, 100, 60, 3, 2, 2, 2, 100, 63, 3, 2, 2, 2, 100, 66,
	3, 2, 2, 2, 100, 69, 3, 2, 2, 2, 100, 72, 3, 2, 2, 2, 100, 75, 3, 2, 2,
	2, 100, 78, 3, 2, 2, 2, 100, 81, 3, 2, 2, 2, 100, 84, 3, 2, 2, 2, 100,
	87, 3, 2, 2, 2, 100, 90, 3, 2, 2, 2, 100, 93, 3, 2, 2, 2, 100, 96, 3, 2,
	2, 2, 100, 99, 3, 2, 2, 2, 101, 5, 3, 2, 2, 2, 102, 104, 7, 11, 2, 2, 103,
	105, 5, 8, 5, 2, 104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 104,
	3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 8, 4,
	1, 2, 109, 7, 3, 2, 2, 2, 110, 111, 7, 25, 2, 2, 111, 112, 7, 53, 2, 2,
	112, 113, 7, 54, 2, 2, 113, 127, 8, 5, 1, 2, 114, 115, 7, 28, 2, 2, 115,
	116, 7, 53, 2, 2, 116, 117, 7, 44, 2, 2, 117, 127, 8, 5, 1, 2, 118, 119,
	7, 26, 2, 2, 119, 120, 7, 53, 2, 2, 120, 121, 7, 41, 2, 2, 121, 127, 8,
	5, 1, 2, 122, 123, 7, 27, 2, 2, 123, 124, 7, 53, 2, 2, 124, 125, 7, 42,
	2, 2, 125, 127, 8, 5, 1, 2, 126, 110, 3, 2, 2, 2, 126, 114, 3, 2, 2, 2,
	126, 118, 3, 2, 2, 2, 126, 122, 3, 2, 2, 2, 127, 9, 3, 2, 2, 2, 128, 129,
	7, 12, 2, 2, 129, 130, 7, 28, 2, 2, 130, 131, 7, 53, 2, 2, 131, 132, 7,
	44, 2, 2, 132, 133, 8, 6, 1, 2, 133, 11, 3, 2, 2, 2, 134, 136, 7, 13, 2,
	2, 135, 137, 5, 14, 8, 2, 136, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138,
	136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 141,
	8, 7, 1, 2, 141, 13, 3, 2, 2, 2, 142, 143, 7, 25, 2, 2, 143, 144, 7, 53,
	2, 2, 144, 145, 7, 54, 2, 2, 145, 167, 8, 8, 1, 2, 146, 147, 7, 28, 2,
	2, 147, 148, 7, 53, 2, 2, 148, 149, 7, 44, 2, 2, 149, 167, 8, 8, 1, 2,
	150, 151, 7, 26, 2, 2, 151, 152, 7, 53, 2, 2, 152, 153, 7, 41, 2, 2, 153,
	167, 8, 8, 1, 2, 154, 155, 7, 27, 2, 2, 155, 156, 7, 53, 2, 2, 156, 157,
	7, 42, 2, 2, 157, 167, 8, 8, 1, 2, 158, 159, 7, 29, 2, 2, 159, 160, 7,
	53, 2, 2, 160, 161, 7, 43, 2, 2, 161, 167, 8, 8, 1, 2, 162, 163, 7, 31,
	2, 2, 163, 164, 7, 53, 2, 2, 164, 165, 7, 56, 2, 2, 165, 167, 8, 8, 1,
	2, 166, 142, 3, 2, 2, 2, 166, 146, 3, 2, 2, 2, 166, 150, 3, 2, 2, 2, 166,
	154, 3, 2, 2, 2, 166, 158, 3, 2, 2, 2, 166, 162, 3, 2, 2, 2, 167, 15, 3,
	2, 2, 2, 168, 170, 7, 14, 2, 2, 169, 171, 5, 18, 10, 2, 170, 169, 3, 2,
	2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2,
	173, 174, 3, 2, 2, 2, 174, 175, 8, 9, 1, 2, 175, 17, 3, 2, 2, 2, 176, 177,
	7, 28, 2, 2, 177, 178, 7, 53, 2, 2, 178, 179, 7, 44, 2, 2, 179, 185, 8,
	10, 1, 2, 180, 181, 7, 31, 2, 2, 181, 182, 7, 53, 2, 2, 182, 183, 7, 56,
	2, 2, 183, 185, 8, 10, 1, 2, 184, 176, 3, 2, 2, 2, 184, 180, 3, 2, 2, 2,
	185, 19, 3, 2, 2, 2, 186, 188, 7, 16, 2, 2, 187, 189, 5, 22, 12, 2, 188,
	187, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191,
	3, 2, 2, 2, 191, 192, 3, 2, 2, 2, 192, 193, 8, 11, 1, 2, 193, 21, 3, 2,
	2, 2, 194, 195, 7, 33, 2, 2, 195, 196, 7, 53, 2, 2, 196, 197, 7, 46, 2,
	2, 197, 203, 8, 12, 1, 2, 198, 199, 7, 29, 2, 2, 199, 200, 7, 53, 2, 2,
	200, 201, 7, 43, 2, 2, 201, 203, 8, 12, 1, 2, 202, 194, 3, 2, 2, 2, 202,
	198, 3, 2, 2, 2, 203, 23, 3, 2, 2, 2, 204, 206, 7, 17, 2, 2, 205, 207,
	5, 26, 14, 2, 206, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 206, 3,
	2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 8, 13, 1,
	2, 211, 25, 3, 2, 2, 2, 212, 213, 7, 35, 2, 2, 213, 214, 7, 53, 2, 2, 214,
	215, 9, 2, 2, 2, 215, 225, 8, 14, 1, 2, 216, 217, 7, 36, 2, 2, 217, 218,
	7, 53, 2, 2, 218, 219, 9, 2, 2, 2, 219, 225, 8, 14, 1, 2, 220, 221, 7,
	33, 2, 2, 221, 222, 7, 53, 2, 2, 222, 223, 7, 46, 2, 2, 223, 225, 8, 14,
	1, 2, 224, 212, 3, 2, 2, 2, 224, 216, 3, 2, 2, 2, 224, 220, 3, 2, 2, 2,
	225, 27, 3, 2, 2, 2, 226, 227, 7, 18, 2, 2, 227, 228, 8, 15, 1, 2, 228,
	29, 3, 2, 2, 2, 229, 230, 7, 19, 2, 2, 230, 231, 7, 31, 2, 2, 231, 232,
	7, 53, 2, 2, 232, 233, 9, 2, 2, 2, 233, 234, 8, 16, 1, 2, 234, 31, 3, 2,
	2, 2, 235, 236, 7, 22, 2, 2, 236, 237, 7, 31, 2, 2, 237, 238, 7, 53, 2,
	2, 238, 239, 9, 2, 2, 2, 239, 240, 8, 17, 1, 2, 240, 33, 3, 2, 2, 2, 241,
	243, 7, 20, 2, 2, 242, 244, 5, 38, 20, 2, 243, 242, 3, 2, 2, 2, 244, 245,
	3, 2, 2, 2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 247, 3, 2,
	2, 2, 247, 248, 8, 18, 1, 2, 248, 35, 3, 2, 2, 2, 249, 250, 7, 21, 2, 2,
	250, 251, 7, 35, 2, 2, 251, 252, 7, 53, 2, 2, 252, 253, 9, 2, 2, 2, 253,
	254, 8, 19, 1, 2, 254, 37, 3, 2, 2, 2, 255, 256, 7, 35, 2, 2, 256, 257,
	7, 53, 2, 2, 257, 258, 9, 2, 2, 2, 258, 268, 8, 20, 1, 2, 259, 260, 7,
	37, 2, 2, 260, 261, 7, 53, 2, 2, 261, 262, 9, 2, 2, 2, 262, 268, 8, 20,
	1, 2, 263, 264, 7, 38, 2, 2, 264, 265, 7, 53, 2, 2, 265, 266, 9, 2, 2,
	2, 266, 268, 8, 20, 1, 2, 267, 255, 3, 2, 2, 2, 267, 259, 3, 2, 2, 2, 267,
	263, 3, 2, 2, 2, 268, 39, 3, 2, 2, 2, 269, 271, 7, 24, 2, 2, 270, 272,
	5, 42, 22, 2, 271, 270, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 271, 3,
	2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 276, 8, 21, 1,
	2, 276, 41, 3, 2, 2, 2, 277, 278, 7, 28, 2, 2, 278, 279, 7, 53, 2, 2, 279,
	280, 7, 44, 2, 2, 280, 293, 8, 22, 1, 2, 281, 282, 7, 3, 2, 2, 282, 283,
	9, 3, 2, 2, 283, 293, 8, 22, 1, 2, 284, 285, 7, 25, 2, 2, 285, 286, 7,
	53, 2, 2, 286, 287, 7, 54, 2, 2, 287, 293, 8, 22, 1, 2, 288, 289, 7, 40,
	2, 2, 289, 290, 7, 53, 2, 2, 290, 291, 7, 44, 2, 2, 291, 293, 8, 22, 1,
	2, 292, 277, 3, 2, 2, 2, 292, 281, 3, 2, 2, 2, 292, 284, 3, 2, 2, 2, 292,
	288, 3, 2, 2, 2, 293, 43, 3, 2, 2, 2, 294, 296, 7, 23, 2, 2, 295, 297,
	5, 46, 24, 2, 296, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 296, 3,
	2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 301, 8, 23, 1,
	2, 301, 45, 3, 2, 2, 2, 302, 303, 7, 28, 2, 2, 303, 304, 7, 53, 2, 2, 304,
	305, 7, 45, 2, 2, 305, 310, 8, 24, 1, 2, 306, 307, 7, 3, 2, 2, 307, 308,
	9, 4, 2, 2, 308, 310, 8, 24, 1, 2, 309, 302, 3, 2, 2, 2, 309, 306, 3, 2,
	2, 2, 310, 47, 3, 2, 2, 2, 311, 313, 7, 10, 2, 2, 312, 314, 5, 50, 26,
	2, 313, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 313, 3, 2, 2, 2, 315,
	316, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 318, 8, 25, 1, 2, 318, 49,
	3, 2, 2, 2, 319, 320, 7, 31, 2, 2, 320, 321, 7, 53, 2, 2, 321, 322, 7,
	47, 2, 2, 322, 336, 8, 26, 1, 2, 323, 324, 7, 28, 2, 2, 324, 325, 7, 53,
	2, 2, 325, 326, 7, 44, 2, 2, 326, 336, 8, 26, 1, 2, 327, 328, 7, 33, 2,
	2, 328, 329, 7, 53, 2, 2, 329, 330, 7, 46, 2, 2, 330, 336, 8, 26, 1, 2,
	331, 332, 7, 39, 2, 2, 332, 333, 7, 53, 2, 2, 333, 334, 7, 44, 2, 2, 334,
	336, 8, 26, 1, 2, 335, 319, 3, 2, 2, 2, 335, 323, 3, 2, 2, 2, 335, 327,
	3, 2, 2, 2, 335, 331, 3, 2, 2, 2, 336, 51, 3, 2, 2, 2, 22, 55, 100, 106,
	126, 138, 166, 172, 184, 190, 202, 208, 224, 245, 267, 273, 292, 298, 309,
	315, 335,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'-'", "'r'", "'R'", "'p'", "'P'", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "'='",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "PAUSES", "EXEC", "REP", "MKDISK", "RMDISK", "FDISK",
	"MOUNT", "UNMOUNT", "MKFS", "LOGIN", "LOGOUT", "MKGRP", "MKUSR", "RMUSR",
	"RMGRP", "MKDIR", "MKFILE", "SIZE", "FIT", "UNIT", "PATH", "TYPE", "DELETEP",
	"NAME", "ADD", "ID", "FS", "USR", "PASSW", "PWD", "GRP", "RUTA", "CONT",
	"E_FIT", "E_UNIT", "E_TYPE", "E_PATH", "E_PATHH", "E_ID", "E_REP", "PATH1",
	"PATH2", "TERMINAL", "PATHH1", "PATHH2", "IGUAL", "ENTERO", "NEGATIVO",
	"IDENTIFICADOR", "COMPLEMENTO", "E_USRS", "NEWLINE", "WHITESPACE", "COMENTARIO",
}

var ruleNames = []string{
	"start", "comando", "mkdisk_f", "mkparam", "rmdisk_f", "fdisk_f", "fdiskparam",
	"mount_f", "mountparam", "mkfs_f", "mkfsparam", "login_f", "loginparam",
	"logout_f", "mkgroup_f", "rmgroup_f", "mkuser_f", "rmuser_f", "mkuserparam",
	"mkfile_f", "mkfileparam", "mkdir_f", "mkdirparam", "rep_f", "repparam",
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
var info_MKFILE Program.InfoMkfile
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

func initializeMKFILE(MKFILE *Program.InfoMkfile) {
	MKFILE.Path = ""
	MKFILE.R = false
	MKFILE.Size = 0
	MKFILE.Cont = ""
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
	GrammarParserT__3          = 4
	GrammarParserT__4          = 5
	GrammarParserPAUSES        = 6
	GrammarParserEXEC          = 7
	GrammarParserREP           = 8
	GrammarParserMKDISK        = 9
	GrammarParserRMDISK        = 10
	GrammarParserFDISK         = 11
	GrammarParserMOUNT         = 12
	GrammarParserUNMOUNT       = 13
	GrammarParserMKFS          = 14
	GrammarParserLOGIN         = 15
	GrammarParserLOGOUT        = 16
	GrammarParserMKGRP         = 17
	GrammarParserMKUSR         = 18
	GrammarParserRMUSR         = 19
	GrammarParserRMGRP         = 20
	GrammarParserMKDIR         = 21
	GrammarParserMKFILE        = 22
	GrammarParserSIZE          = 23
	GrammarParserFIT           = 24
	GrammarParserUNIT          = 25
	GrammarParserPATH          = 26
	GrammarParserTYPE          = 27
	GrammarParserDELETEP       = 28
	GrammarParserNAME          = 29
	GrammarParserADD           = 30
	GrammarParserID            = 31
	GrammarParserFS            = 32
	GrammarParserUSR           = 33
	GrammarParserPASSW         = 34
	GrammarParserPWD           = 35
	GrammarParserGRP           = 36
	GrammarParserRUTA          = 37
	GrammarParserCONT          = 38
	GrammarParserE_FIT         = 39
	GrammarParserE_UNIT        = 40
	GrammarParserE_TYPE        = 41
	GrammarParserE_PATH        = 42
	GrammarParserE_PATHH       = 43
	GrammarParserE_ID          = 44
	GrammarParserE_REP         = 45
	GrammarParserPATH1         = 46
	GrammarParserPATH2         = 47
	GrammarParserTERMINAL      = 48
	GrammarParserPATHH1        = 49
	GrammarParserPATHH2        = 50
	GrammarParserIGUAL         = 51
	GrammarParserENTERO        = 52
	GrammarParserNEGATIVO      = 53
	GrammarParserIDENTIFICADOR = 54
	GrammarParserCOMPLEMENTO   = 55
	GrammarParserE_USRS        = 56
	GrammarParserNEWLINE       = 57
	GrammarParserWHITESPACE    = 58
	GrammarParserCOMENTARIO    = 59
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
	GrammarParserRULE_rmgroup_f   = 15
	GrammarParserRULE_mkuser_f    = 16
	GrammarParserRULE_rmuser_f    = 17
	GrammarParserRULE_mkuserparam = 18
	GrammarParserRULE_mkfile_f    = 19
	GrammarParserRULE_mkfileparam = 20
	GrammarParserRULE_mkdir_f     = 21
	GrammarParserRULE_mkdirparam  = 22
	GrammarParserRULE_rep_f       = 23
	GrammarParserRULE_repparam    = 24
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
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserREP)|(1<<GrammarParserMKDISK)|(1<<GrammarParserRMDISK)|(1<<GrammarParserFDISK)|(1<<GrammarParserMOUNT)|(1<<GrammarParserMKFS)|(1<<GrammarParserLOGIN)|(1<<GrammarParserLOGOUT)|(1<<GrammarParserMKGRP)|(1<<GrammarParserMKUSR)|(1<<GrammarParserRMUSR)|(1<<GrammarParserRMGRP)|(1<<GrammarParserMKDIR)|(1<<GrammarParserMKFILE))) != 0) || _la == GrammarParserNEWLINE {
		{
			p.SetState(50)
			p.Comando()
		}

		p.SetState(53)
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

func (s *ComandoContext) Rmgroup_f() IRmgroup_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRmgroup_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRmgroup_fContext)
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

func (s *ComandoContext) Mkfile_f() IMkfile_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkfile_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMkfile_fContext)
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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserMKDISK:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Mkdisk_f()
		}
		{
			p.SetState(56)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserRMDISK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(58)
			p.Rmdisk_f()
		}
		{
			p.SetState(59)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserFDISK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Fdisk_f()
		}
		{
			p.SetState(62)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMOUNT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.Mount_f()
		}
		{
			p.SetState(65)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKFS:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(67)
			p.Mkfs_f()
		}
		{
			p.SetState(68)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGIN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.Login_f()
		}
		{
			p.SetState(71)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserLOGOUT:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(73)
			p.Logout_f()
		}
		{
			p.SetState(74)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKGRP:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(76)
			p.Mkgroup_f()
		}
		{
			p.SetState(77)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserRMGRP:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(79)
			p.Rmgroup_f()
		}
		{
			p.SetState(80)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKUSR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(82)
			p.Mkuser_f()
		}
		{
			p.SetState(83)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserRMUSR:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(85)
			p.Rmuser_f()
		}
		{
			p.SetState(86)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKFILE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(88)
			p.Mkfile_f()
		}
		{
			p.SetState(89)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserMKDIR:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(91)
			p.Mkdir_f()
		}
		{
			p.SetState(92)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserREP:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(94)
			p.Rep_f()
		}
		{
			p.SetState(95)
			p.Match(GrammarParserNEWLINE)
		}

	case GrammarParserNEWLINE:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(97)
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
		p.SetState(100)
		p.Match(GrammarParserMKDISK)
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH))) != 0) {
		{
			p.SetState(101)
			p.Mkparam()
		}

		p.SetState(104)
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

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(109)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(110)

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
			p.SetState(112)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(113)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(114)

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
			p.SetState(116)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(117)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(118)

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
			p.SetState(120)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(121)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(122)

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
		p.SetState(126)
		p.Match(GrammarParserRMDISK)
	}
	{
		p.SetState(127)
		p.Match(GrammarParserPATH)
	}
	{
		p.SetState(128)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(129)

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
		p.SetState(132)
		p.Match(GrammarParserFDISK)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserSIZE)|(1<<GrammarParserFIT)|(1<<GrammarParserUNIT)|(1<<GrammarParserPATH)|(1<<GrammarParserTYPE)|(1<<GrammarParserNAME))) != 0) {
		{
			p.SetState(133)
			p.Fdiskparam()
		}

		p.SetState(136)
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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(141)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(142)

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
			p.SetState(144)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(145)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(146)

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
			p.SetState(148)
			p.Match(GrammarParserFIT)
		}
		{
			p.SetState(149)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(150)

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
			p.SetState(152)
			p.Match(GrammarParserUNIT)
		}
		{
			p.SetState(153)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(154)

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
			p.SetState(156)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(157)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(158)

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
			p.SetState(160)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(161)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(162)

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
		p.SetState(166)
		p.Match(GrammarParserMOUNT)
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserPATH || _la == GrammarParserNAME {
		{
			p.SetState(167)
			p.Mountparam()
		}

		p.SetState(170)
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

	p.SetState(182)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(174)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(175)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(176)

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
			p.SetState(178)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(179)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(180)

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
		p.SetState(184)
		p.Match(GrammarParserMKFS)
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserTYPE || _la == GrammarParserID {
		{
			p.SetState(185)
			p.Mkfsparam()
		}

		p.SetState(188)
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

	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(193)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(194)

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
			p.SetState(196)
			p.Match(GrammarParserTYPE)
		}
		{
			p.SetState(197)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(198)

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
		p.SetState(202)
		p.Match(GrammarParserLOGIN)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(GrammarParserID-31))|(1<<(GrammarParserUSR-31))|(1<<(GrammarParserPASSW-31)))) != 0) {
		{
			p.SetState(203)
			p.Loginparam()
		}

		p.SetState(206)
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

			localctx.(*LoginparamContext).e_user = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
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
			p.SetState(214)
			p.Match(GrammarParserPASSW)
		}
		{
			p.SetState(215)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(216)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LoginparamContext).e_pass = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
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
			p.SetState(218)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(219)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(220)

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
		p.SetState(224)
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
		p.SetState(227)
		p.Match(GrammarParserMKGRP)
	}
	{
		p.SetState(228)
		p.Match(GrammarParserNAME)
	}
	{
		p.SetState(229)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(230)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Mkgroup_fContext).e_name = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
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

// IRmgroup_fContext is an interface to support dynamic dispatch.
type IRmgroup_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE_name returns the e_name token.
	GetE_name() antlr.Token

	// SetE_name sets the e_name token.
	SetE_name(antlr.Token)

	// IsRmgroup_fContext differentiates from other interfaces.
	IsRmgroup_fContext()
}

type Rmgroup_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	e_name antlr.Token
}

func NewEmptyRmgroup_fContext() *Rmgroup_fContext {
	var p = new(Rmgroup_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_rmgroup_f
	return p
}

func (*Rmgroup_fContext) IsRmgroup_fContext() {}

func NewRmgroup_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rmgroup_fContext {
	var p = new(Rmgroup_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_rmgroup_f

	return p
}

func (s *Rmgroup_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Rmgroup_fContext) GetE_name() antlr.Token { return s.e_name }

func (s *Rmgroup_fContext) SetE_name(v antlr.Token) { s.e_name = v }

func (s *Rmgroup_fContext) RMGRP() antlr.TerminalNode {
	return s.GetToken(GrammarParserRMGRP, 0)
}

func (s *Rmgroup_fContext) NAME() antlr.TerminalNode {
	return s.GetToken(GrammarParserNAME, 0)
}

func (s *Rmgroup_fContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *Rmgroup_fContext) IDENTIFICADOR() antlr.TerminalNode {
	return s.GetToken(GrammarParserIDENTIFICADOR, 0)
}

func (s *Rmgroup_fContext) COMPLEMENTO() antlr.TerminalNode {
	return s.GetToken(GrammarParserCOMPLEMENTO, 0)
}

func (s *Rmgroup_fContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *Rmgroup_fContext) E_USRS() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_USRS, 0)
}

func (s *Rmgroup_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rmgroup_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rmgroup_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterRmgroup_f(s)
	}
}

func (s *Rmgroup_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitRmgroup_f(s)
	}
}

func (p *GrammarParser) Rmgroup_f() (localctx IRmgroup_fContext) {
	localctx = NewRmgroup_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GrammarParserRULE_rmgroup_f)
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
		p.SetState(233)
		p.Match(GrammarParserRMGRP)
	}
	{
		p.SetState(234)
		p.Match(GrammarParserNAME)
	}
	{
		p.SetState(235)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(236)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Rmgroup_fContext).e_name = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Rmgroup_fContext).e_name = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	Program.Rmgroup(strings.ReplaceAll((func() string {
		if localctx.(*Rmgroup_fContext).GetE_name() == nil {
			return ""
		} else {
			return localctx.(*Rmgroup_fContext).GetE_name().GetText()
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
	p.EnterRule(localctx, 32, GrammarParserRULE_mkuser_f)
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
		p.SetState(239)
		p.Match(GrammarParserMKUSR)
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(GrammarParserUSR-33))|(1<<(GrammarParserPWD-33))|(1<<(GrammarParserGRP-33)))) != 0) {
		{
			p.SetState(240)
			p.Mkuserparam()
		}

		p.SetState(243)
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
	p.EnterRule(localctx, 34, GrammarParserRULE_rmuser_f)
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
		p.SetState(247)
		p.Match(GrammarParserRMUSR)
	}
	{
		p.SetState(248)
		p.Match(GrammarParserUSR)
	}
	{
		p.SetState(249)
		p.Match(GrammarParserIGUAL)
	}
	{
		p.SetState(250)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Rmuser_fContext).e_name = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
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
	p.EnterRule(localctx, 36, GrammarParserRULE_mkuserparam)
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

	p.SetState(265)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserUSR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Match(GrammarParserUSR)
		}
		{
			p.SetState(254)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(255)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_user = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
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
			p.SetState(257)
			p.Match(GrammarParserPWD)
		}
		{
			p.SetState(258)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(259)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_pass = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
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
			p.SetState(261)
			p.Match(GrammarParserGRP)
		}
		{
			p.SetState(262)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(263)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*MkuserparamContext).e_group = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-52)&-(0x1f+1)) == 0 && ((1<<uint((_la-52)))&((1<<(GrammarParserENTERO-52))|(1<<(GrammarParserIDENTIFICADOR-52))|(1<<(GrammarParserCOMPLEMENTO-52))|(1<<(GrammarParserE_USRS-52)))) != 0) {
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

// IMkfile_fContext is an interface to support dynamic dispatch.
type IMkfile_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMkfile_fContext differentiates from other interfaces.
	IsMkfile_fContext()
}

type Mkfile_fContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfile_fContext() *Mkfile_fContext {
	var p = new(Mkfile_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkfile_f
	return p
}

func (*Mkfile_fContext) IsMkfile_fContext() {}

func NewMkfile_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkfile_fContext {
	var p = new(Mkfile_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkfile_f

	return p
}

func (s *Mkfile_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkfile_fContext) MKFILE() antlr.TerminalNode {
	return s.GetToken(GrammarParserMKFILE, 0)
}

func (s *Mkfile_fContext) AllMkfileparam() []IMkfileparamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMkfileparamContext)(nil)).Elem())
	var tst = make([]IMkfileparamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMkfileparamContext)
		}
	}

	return tst
}

func (s *Mkfile_fContext) Mkfileparam(i int) IMkfileparamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMkfileparamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMkfileparamContext)
}

func (s *Mkfile_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkfile_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkfile_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkfile_f(s)
	}
}

func (s *Mkfile_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkfile_f(s)
	}
}

func (p *GrammarParser) Mkfile_f() (localctx IMkfile_fContext) {
	localctx = NewMkfile_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GrammarParserRULE_mkfile_f)
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
		p.SetState(267)
		p.Match(GrammarParserMKFILE)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GrammarParserT__0)|(1<<GrammarParserSIZE)|(1<<GrammarParserPATH))) != 0) || _la == GrammarParserCONT {
		{
			p.SetState(268)
			p.Mkfileparam()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	Program.EscribirArchivo(info_MKFILE)
	initializeMKFILE(&info_MKFILE)

	return localctx
}

// IMkfileparamContext is an interface to support dynamic dispatch.
type IMkfileparamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_E_PATH returns the _E_PATH token.
	Get_E_PATH() antlr.Token

	// Get_ENTERO returns the _ENTERO token.
	Get_ENTERO() antlr.Token

	// Set_E_PATH sets the _E_PATH token.
	Set_E_PATH(antlr.Token)

	// Set_ENTERO sets the _ENTERO token.
	Set_ENTERO(antlr.Token)

	// IsMkfileparamContext differentiates from other interfaces.
	IsMkfileparamContext()
}

type MkfileparamContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_E_PATH antlr.Token
	_ENTERO antlr.Token
}

func NewEmptyMkfileparamContext() *MkfileparamContext {
	var p = new(MkfileparamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GrammarParserRULE_mkfileparam
	return p
}

func (*MkfileparamContext) IsMkfileparamContext() {}

func NewMkfileparamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MkfileparamContext {
	var p = new(MkfileparamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_mkfileparam

	return p
}

func (s *MkfileparamContext) GetParser() antlr.Parser { return s.parser }

func (s *MkfileparamContext) Get_E_PATH() antlr.Token { return s._E_PATH }

func (s *MkfileparamContext) Get_ENTERO() antlr.Token { return s._ENTERO }

func (s *MkfileparamContext) Set_E_PATH(v antlr.Token) { s._E_PATH = v }

func (s *MkfileparamContext) Set_ENTERO(v antlr.Token) { s._ENTERO = v }

func (s *MkfileparamContext) PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserPATH, 0)
}

func (s *MkfileparamContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(GrammarParserIGUAL, 0)
}

func (s *MkfileparamContext) E_PATH() antlr.TerminalNode {
	return s.GetToken(GrammarParserE_PATH, 0)
}

func (s *MkfileparamContext) SIZE() antlr.TerminalNode {
	return s.GetToken(GrammarParserSIZE, 0)
}

func (s *MkfileparamContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(GrammarParserENTERO, 0)
}

func (s *MkfileparamContext) CONT() antlr.TerminalNode {
	return s.GetToken(GrammarParserCONT, 0)
}

func (s *MkfileparamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MkfileparamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MkfileparamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMkfileparam(s)
	}
}

func (s *MkfileparamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMkfileparam(s)
	}
}

func (p *GrammarParser) Mkfileparam() (localctx IMkfileparamContext) {
	localctx = NewMkfileparamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GrammarParserRULE_mkfileparam)
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

	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(276)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(277)

			var _m = p.Match(GrammarParserE_PATH)

			localctx.(*MkfileparamContext)._E_PATH = _m
		}
		info_MKFILE.Path = strings.ReplaceAll((func() string {
			if localctx.(*MkfileparamContext).Get_E_PATH() == nil {
				return ""
			} else {
				return localctx.(*MkfileparamContext).Get_E_PATH().GetText()
			}
		}()), "\"", "")

	case GrammarParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
			p.Match(GrammarParserT__0)
		}
		{
			p.SetState(280)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__1 || _la == GrammarParserT__2) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		info_MKFILE.R = true

	case GrammarParserSIZE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
			p.Match(GrammarParserSIZE)
		}
		{
			p.SetState(283)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(284)

			var _m = p.Match(GrammarParserENTERO)

			localctx.(*MkfileparamContext)._ENTERO = _m
		}
		info_MKFILE.Size = (func() int {
			if localctx.(*MkfileparamContext).Get_ENTERO() == nil {
				return 0
			} else {
				i, _ := strconv.Atoi(localctx.(*MkfileparamContext).Get_ENTERO().GetText())
				return i
			}
		}())

	case GrammarParserCONT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(286)
			p.Match(GrammarParserCONT)
		}
		{
			p.SetState(287)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(288)

			var _m = p.Match(GrammarParserE_PATH)

			localctx.(*MkfileparamContext)._E_PATH = _m
		}
		info_MKFILE.Cont = strings.ReplaceAll((func() string {
			if localctx.(*MkfileparamContext).Get_E_PATH() == nil {
				return ""
			} else {
				return localctx.(*MkfileparamContext).Get_E_PATH().GetText()
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
	p.EnterRule(localctx, 42, GrammarParserRULE_mkdir_f)
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
		p.SetState(292)
		p.Match(GrammarParserMKDIR)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserT__0 || _la == GrammarParserPATH {
		{
			p.SetState(293)
			p.Mkdirparam()
		}

		p.SetState(296)
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
	p.EnterRule(localctx, 44, GrammarParserRULE_mkdirparam)
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

	p.SetState(307)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserPATH:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(301)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(302)

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
			p.SetState(304)
			p.Match(GrammarParserT__0)
		}
		{
			p.SetState(305)
			_la = p.GetTokenStream().LA(1)

			if !(_la == GrammarParserT__3 || _la == GrammarParserT__4) {
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
	p.EnterRule(localctx, 46, GrammarParserRULE_rep_f)
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
		p.SetState(309)
		p.Match(GrammarParserREP)
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(GrammarParserPATH-26))|(1<<(GrammarParserNAME-26))|(1<<(GrammarParserID-26))|(1<<(GrammarParserRUTA-26)))) != 0) {
		{
			p.SetState(310)
			p.Repparam()
		}

		p.SetState(313)
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
	p.EnterRule(localctx, 48, GrammarParserRULE_repparam)

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

	p.SetState(333)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GrammarParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.Match(GrammarParserNAME)
		}
		{
			p.SetState(318)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(319)

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
			p.SetState(321)
			p.Match(GrammarParserPATH)
		}
		{
			p.SetState(322)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(323)

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
			p.SetState(325)
			p.Match(GrammarParserID)
		}
		{
			p.SetState(326)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(327)

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
			p.SetState(329)
			p.Match(GrammarParserRUTA)
		}
		{
			p.SetState(330)
			p.Match(GrammarParserIGUAL)
		}
		{
			p.SetState(331)

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
