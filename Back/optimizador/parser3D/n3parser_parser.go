// Code generated from N3parser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // N3parser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "github.com/colegno/arraylist"
import "Back/optimizador/elementos/bloque3d"
import "Back/optimizador/elementos/control"
import "Back/optimizador/elementos/expresiones3d"
import "Back/optimizador/elementos/funciones3d"
import "Back/optimizador/elementos/headers3d"
import "Back/optimizador/elementos/instrucciones3d"
import "Back/optimizador/elementos/interfaces3d"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 50, 385,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 6, 3, 55, 10, 3, 13,
	3, 14, 3, 56, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 77, 10, 4, 3,
	5, 6, 5, 80, 10, 5, 13, 5, 14, 5, 81, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 98, 10, 6, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 118, 10, 7, 3, 8, 6, 8, 121, 10,
	8, 13, 8, 14, 8, 122, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 139, 10, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 153,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 192, 10, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5,
	13, 215, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 5, 14, 249, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 5, 15, 288, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 303, 10, 18, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 20, 6, 20, 310, 10, 20, 13, 20, 14, 20, 311, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 343, 10, 21,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 7, 23, 356, 10, 23, 12, 23, 14, 23, 359, 11, 23, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 383,
	10, 24, 3, 24, 2, 3, 44, 25, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 2, 2, 2, 408, 2, 48, 3, 2,
	2, 2, 4, 54, 3, 2, 2, 2, 6, 76, 3, 2, 2, 2, 8, 79, 3, 2, 2, 2, 10, 97,
	3, 2, 2, 2, 12, 117, 3, 2, 2, 2, 14, 120, 3, 2, 2, 2, 16, 138, 3, 2, 2,
	2, 18, 152, 3, 2, 2, 2, 20, 191, 3, 2, 2, 2, 22, 193, 3, 2, 2, 2, 24, 214,
	3, 2, 2, 2, 26, 248, 3, 2, 2, 2, 28, 287, 3, 2, 2, 2, 30, 289, 3, 2, 2,
	2, 32, 293, 3, 2, 2, 2, 34, 302, 3, 2, 2, 2, 36, 304, 3, 2, 2, 2, 38, 309,
	3, 2, 2, 2, 40, 342, 3, 2, 2, 2, 42, 344, 3, 2, 2, 2, 44, 347, 3, 2, 2,
	2, 46, 382, 3, 2, 2, 2, 48, 49, 5, 36, 19, 2, 49, 50, 5, 38, 20, 2, 50,
	51, 5, 4, 3, 2, 51, 52, 8, 2, 1, 2, 52, 3, 3, 2, 2, 2, 53, 55, 5, 6, 4,
	2, 54, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57,
	3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 59, 8, 3, 1, 2, 59, 5, 3, 2, 2, 2,
	60, 61, 7, 3, 2, 2, 61, 62, 7, 5, 2, 2, 62, 63, 7, 41, 2, 2, 63, 64, 7,
	42, 2, 2, 64, 65, 7, 43, 2, 2, 65, 66, 5, 8, 5, 2, 66, 67, 8, 4, 1, 2,
	67, 77, 3, 2, 2, 2, 68, 69, 7, 8, 2, 2, 69, 70, 7, 22, 2, 2, 70, 71, 7,
	41, 2, 2, 71, 72, 7, 42, 2, 2, 72, 73, 7, 43, 2, 2, 73, 74, 5, 8, 5, 2,
	74, 75, 8, 4, 1, 2, 75, 77, 3, 2, 2, 2, 76, 60, 3, 2, 2, 2, 76, 68, 3,
	2, 2, 2, 77, 7, 3, 2, 2, 2, 78, 80, 5, 10, 6, 2, 79, 78, 3, 2, 2, 2, 80,
	81, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 3, 2, 2,
	2, 83, 84, 8, 5, 1, 2, 84, 9, 3, 2, 2, 2, 85, 86, 5, 12, 7, 2, 86, 87,
	5, 14, 8, 2, 87, 88, 5, 18, 10, 2, 88, 89, 8, 6, 1, 2, 89, 98, 3, 2, 2,
	2, 90, 91, 5, 14, 8, 2, 91, 92, 8, 6, 1, 2, 92, 98, 3, 2, 2, 2, 93, 94,
	5, 12, 7, 2, 94, 95, 5, 18, 10, 2, 95, 96, 8, 6, 1, 2, 96, 98, 3, 2, 2,
	2, 97, 85, 3, 2, 2, 2, 97, 90, 3, 2, 2, 2, 97, 93, 3, 2, 2, 2, 98, 11,
	3, 2, 2, 2, 99, 100, 5, 24, 13, 2, 100, 101, 7, 26, 2, 2, 101, 102, 8,
	7, 1, 2, 102, 118, 3, 2, 2, 2, 103, 104, 5, 30, 16, 2, 104, 105, 8, 7,
	1, 2, 105, 118, 3, 2, 2, 2, 106, 107, 5, 20, 11, 2, 107, 108, 7, 26, 2,
	2, 108, 109, 8, 7, 1, 2, 109, 118, 3, 2, 2, 2, 110, 111, 5, 22, 12, 2,
	111, 112, 8, 7, 1, 2, 112, 118, 3, 2, 2, 2, 113, 114, 5, 32, 17, 2, 114,
	115, 7, 26, 2, 2, 115, 116, 8, 7, 1, 2, 116, 118, 3, 2, 2, 2, 117, 99,
	3, 2, 2, 2, 117, 103, 3, 2, 2, 2, 117, 106, 3, 2, 2, 2, 117, 110, 3, 2,
	2, 2, 117, 113, 3, 2, 2, 2, 118, 13, 3, 2, 2, 2, 119, 121, 5, 16, 9, 2,
	120, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122,
	123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 125, 8, 8, 1, 2, 125, 15, 3,
	2, 2, 2, 126, 127, 5, 24, 13, 2, 127, 128, 7, 26, 2, 2, 128, 129, 8, 9,
	1, 2, 129, 139, 3, 2, 2, 2, 130, 131, 5, 34, 18, 2, 131, 132, 7, 26, 2,
	2, 132, 133, 8, 9, 1, 2, 133, 139, 3, 2, 2, 2, 134, 135, 5, 20, 11, 2,
	135, 136, 7, 26, 2, 2, 136, 137, 8, 9, 1, 2, 137, 139, 3, 2, 2, 2, 138,
	126, 3, 2, 2, 2, 138, 130, 3, 2, 2, 2, 138, 134, 3, 2, 2, 2, 139, 17, 3,
	2, 2, 2, 140, 141, 5, 30, 16, 2, 141, 142, 8, 10, 1, 2, 142, 153, 3, 2,
	2, 2, 143, 144, 5, 32, 17, 2, 144, 145, 7, 26, 2, 2, 145, 146, 8, 10, 1,
	2, 146, 153, 3, 2, 2, 2, 147, 148, 5, 22, 12, 2, 148, 149, 8, 10, 1, 2,
	149, 153, 3, 2, 2, 2, 150, 151, 7, 44, 2, 2, 151, 153, 8, 10, 1, 2, 152,
	140, 3, 2, 2, 2, 152, 143, 3, 2, 2, 2, 152, 147, 3, 2, 2, 2, 152, 150,
	3, 2, 2, 2, 153, 19, 3, 2, 2, 2, 154, 155, 7, 9, 2, 2, 155, 156, 7, 41,
	2, 2, 156, 157, 7, 47, 2, 2, 157, 158, 7, 25, 2, 2, 158, 159, 7, 41, 2,
	2, 159, 160, 7, 3, 2, 2, 160, 161, 7, 42, 2, 2, 161, 162, 7, 19, 2, 2,
	162, 163, 7, 42, 2, 2, 163, 192, 8, 11, 1, 2, 164, 165, 7, 9, 2, 2, 165,
	166, 7, 41, 2, 2, 166, 167, 7, 47, 2, 2, 167, 168, 7, 25, 2, 2, 168, 169,
	7, 17, 2, 2, 169, 170, 7, 42, 2, 2, 170, 192, 8, 11, 1, 2, 171, 172, 7,
	9, 2, 2, 172, 173, 7, 41, 2, 2, 173, 174, 7, 47, 2, 2, 174, 175, 7, 25,
	2, 2, 175, 176, 7, 41, 2, 2, 176, 177, 7, 3, 2, 2, 177, 178, 7, 42, 2,
	2, 178, 179, 7, 17, 2, 2, 179, 180, 7, 42, 2, 2, 180, 192, 8, 11, 1, 2,
	181, 182, 7, 9, 2, 2, 182, 183, 7, 41, 2, 2, 183, 184, 7, 47, 2, 2, 184,
	185, 7, 25, 2, 2, 185, 186, 7, 41, 2, 2, 186, 187, 7, 3, 2, 2, 187, 188,
	7, 42, 2, 2, 188, 189, 7, 18, 2, 2, 189, 190, 7, 42, 2, 2, 190, 192, 8,
	11, 1, 2, 191, 154, 3, 2, 2, 2, 191, 164, 3, 2, 2, 2, 191, 171, 3, 2, 2,
	2, 191, 181, 3, 2, 2, 2, 192, 21, 3, 2, 2, 2, 193, 194, 7, 6, 2, 2, 194,
	195, 7, 41, 2, 2, 195, 196, 5, 26, 14, 2, 196, 197, 7, 42, 2, 2, 197, 198,
	5, 32, 17, 2, 198, 199, 7, 26, 2, 2, 199, 200, 5, 32, 17, 2, 200, 201,
	7, 26, 2, 2, 201, 202, 5, 30, 16, 2, 202, 203, 8, 12, 1, 2, 203, 23, 3,
	2, 2, 2, 204, 205, 5, 28, 15, 2, 205, 206, 7, 34, 2, 2, 206, 207, 5, 26,
	14, 2, 207, 208, 8, 13, 1, 2, 208, 215, 3, 2, 2, 2, 209, 210, 5, 28, 15,
	2, 210, 211, 7, 34, 2, 2, 211, 212, 5, 28, 15, 2, 212, 213, 8, 13, 1, 2,
	213, 215, 3, 2, 2, 2, 214, 204, 3, 2, 2, 2, 214, 209, 3, 2, 2, 2, 215,
	25, 3, 2, 2, 2, 216, 217, 5, 28, 15, 2, 217, 218, 5, 46, 24, 2, 218, 219,
	5, 28, 15, 2, 219, 220, 8, 14, 1, 2, 220, 249, 3, 2, 2, 2, 221, 222, 7,
	41, 2, 2, 222, 223, 7, 3, 2, 2, 223, 224, 7, 42, 2, 2, 224, 225, 5, 28,
	15, 2, 225, 226, 5, 46, 24, 2, 226, 227, 7, 41, 2, 2, 227, 228, 7, 3, 2,
	2, 228, 229, 7, 42, 2, 2, 229, 230, 5, 28, 15, 2, 230, 231, 8, 14, 1, 2,
	231, 249, 3, 2, 2, 2, 232, 233, 7, 41, 2, 2, 233, 234, 7, 3, 2, 2, 234,
	235, 7, 42, 2, 2, 235, 236, 5, 28, 15, 2, 236, 237, 5, 46, 24, 2, 237,
	238, 5, 28, 15, 2, 238, 239, 8, 14, 1, 2, 239, 249, 3, 2, 2, 2, 240, 241,
	5, 28, 15, 2, 241, 242, 5, 46, 24, 2, 242, 243, 7, 41, 2, 2, 243, 244,
	7, 3, 2, 2, 244, 245, 7, 42, 2, 2, 245, 246, 5, 28, 15, 2, 246, 247, 8,
	14, 1, 2, 247, 249, 3, 2, 2, 2, 248, 216, 3, 2, 2, 2, 248, 221, 3, 2, 2,
	2, 248, 232, 3, 2, 2, 2, 248, 240, 3, 2, 2, 2, 249, 27, 3, 2, 2, 2, 250,
	251, 7, 13, 2, 2, 251, 288, 8, 15, 1, 2, 252, 253, 7, 19, 2, 2, 253, 288,
	8, 15, 1, 2, 254, 255, 7, 14, 2, 2, 255, 288, 8, 15, 1, 2, 256, 257, 7,
	38, 2, 2, 257, 258, 7, 17, 2, 2, 258, 288, 8, 15, 1, 2, 259, 260, 7, 17,
	2, 2, 260, 288, 8, 15, 1, 2, 261, 262, 7, 18, 2, 2, 262, 288, 8, 15, 1,
	2, 263, 264, 7, 11, 2, 2, 264, 265, 7, 45, 2, 2, 265, 266, 7, 41, 2, 2,
	266, 267, 7, 3, 2, 2, 267, 268, 7, 42, 2, 2, 268, 269, 7, 19, 2, 2, 269,
	270, 7, 46, 2, 2, 270, 288, 8, 15, 1, 2, 271, 272, 7, 11, 2, 2, 272, 273,
	7, 45, 2, 2, 273, 274, 7, 41, 2, 2, 274, 275, 7, 3, 2, 2, 275, 276, 7,
	42, 2, 2, 276, 277, 7, 14, 2, 2, 277, 278, 7, 46, 2, 2, 278, 288, 8, 15,
	1, 2, 279, 280, 7, 12, 2, 2, 280, 281, 7, 45, 2, 2, 281, 282, 7, 41, 2,
	2, 282, 283, 7, 3, 2, 2, 283, 284, 7, 42, 2, 2, 284, 285, 7, 19, 2, 2,
	285, 286, 7, 46, 2, 2, 286, 288, 8, 15, 1, 2, 287, 250, 3, 2, 2, 2, 287,
	252, 3, 2, 2, 2, 287, 254, 3, 2, 2, 2, 287, 256, 3, 2, 2, 2, 287, 259,
	3, 2, 2, 2, 287, 261, 3, 2, 2, 2, 287, 263, 3, 2, 2, 2, 287, 271, 3, 2,
	2, 2, 287, 279, 3, 2, 2, 2, 288, 29, 3, 2, 2, 2, 289, 290, 7, 20, 2, 2,
	290, 291, 7, 23, 2, 2, 291, 292, 8, 16, 1, 2, 292, 31, 3, 2, 2, 2, 293,
	294, 7, 10, 2, 2, 294, 295, 7, 20, 2, 2, 295, 296, 8, 17, 1, 2, 296, 33,
	3, 2, 2, 2, 297, 298, 7, 7, 2, 2, 298, 303, 8, 18, 1, 2, 299, 300, 7, 7,
	2, 2, 300, 301, 7, 17, 2, 2, 301, 303, 8, 18, 1, 2, 302, 297, 3, 2, 2,
	2, 302, 299, 3, 2, 2, 2, 303, 35, 3, 2, 2, 2, 304, 305, 7, 15, 2, 2, 305,
	306, 7, 16, 2, 2, 306, 307, 8, 19, 1, 2, 307, 37, 3, 2, 2, 2, 308, 310,
	5, 40, 21, 2, 309, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 309, 3,
	2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 314, 8, 20, 1,
	2, 314, 39, 3, 2, 2, 2, 315, 316, 7, 4, 2, 2, 316, 317, 7, 12, 2, 2, 317,
	318, 7, 45, 2, 2, 318, 319, 7, 17, 2, 2, 319, 320, 7, 46, 2, 2, 320, 321,
	7, 26, 2, 2, 321, 343, 8, 21, 1, 2, 322, 323, 7, 4, 2, 2, 323, 324, 7,
	11, 2, 2, 324, 325, 7, 45, 2, 2, 325, 326, 7, 17, 2, 2, 326, 327, 7, 46,
	2, 2, 327, 328, 7, 26, 2, 2, 328, 343, 8, 21, 1, 2, 329, 330, 7, 4, 2,
	2, 330, 331, 7, 13, 2, 2, 331, 332, 7, 26, 2, 2, 332, 343, 8, 21, 1, 2,
	333, 334, 7, 4, 2, 2, 334, 335, 7, 14, 2, 2, 335, 336, 7, 26, 2, 2, 336,
	343, 8, 21, 1, 2, 337, 338, 7, 4, 2, 2, 338, 339, 5, 42, 22, 2, 339, 340,
	7, 26, 2, 2, 340, 341, 8, 21, 1, 2, 341, 343, 3, 2, 2, 2, 342, 315, 3,
	2, 2, 2, 342, 322, 3, 2, 2, 2, 342, 329, 3, 2, 2, 2, 342, 333, 3, 2, 2,
	2, 342, 337, 3, 2, 2, 2, 343, 41, 3, 2, 2, 2, 344, 345, 5, 44, 23, 2, 345,
	346, 8, 22, 1, 2, 346, 43, 3, 2, 2, 2, 347, 348, 8, 23, 1, 2, 348, 349,
	7, 19, 2, 2, 349, 350, 8, 23, 1, 2, 350, 357, 3, 2, 2, 2, 351, 352, 12,
	4, 2, 2, 352, 353, 7, 25, 2, 2, 353, 354, 7, 19, 2, 2, 354, 356, 8, 23,
	1, 2, 355, 351, 3, 2, 2, 2, 356, 359, 3, 2, 2, 2, 357, 355, 3, 2, 2, 2,
	357, 358, 3, 2, 2, 2, 358, 45, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 360, 361,
	7, 27, 2, 2, 361, 383, 8, 24, 1, 2, 362, 363, 7, 28, 2, 2, 363, 383, 8,
	24, 1, 2, 364, 365, 7, 29, 2, 2, 365, 383, 8, 24, 1, 2, 366, 367, 7, 30,
	2, 2, 367, 383, 8, 24, 1, 2, 368, 369, 7, 31, 2, 2, 369, 383, 8, 24, 1,
	2, 370, 371, 7, 33, 2, 2, 371, 383, 8, 24, 1, 2, 372, 373, 7, 35, 2, 2,
	373, 383, 8, 24, 1, 2, 374, 375, 7, 36, 2, 2, 375, 383, 8, 24, 1, 2, 376,
	377, 7, 37, 2, 2, 377, 383, 8, 24, 1, 2, 378, 379, 7, 38, 2, 2, 379, 383,
	8, 24, 1, 2, 380, 381, 7, 39, 2, 2, 381, 383, 8, 24, 1, 2, 382, 360, 3,
	2, 2, 2, 382, 362, 3, 2, 2, 2, 382, 364, 3, 2, 2, 2, 382, 366, 3, 2, 2,
	2, 382, 368, 3, 2, 2, 2, 382, 370, 3, 2, 2, 2, 382, 372, 3, 2, 2, 2, 382,
	374, 3, 2, 2, 2, 382, 376, 3, 2, 2, 2, 382, 378, 3, 2, 2, 2, 382, 380,
	3, 2, 2, 2, 383, 47, 3, 2, 2, 2, 19, 56, 76, 81, 97, 117, 122, 138, 152,
	191, 214, 248, 287, 302, 311, 342, 357, 382,
}
var literalNames = []string{
	"", "'int'", "'float'", "'main'", "'if'", "'return'", "'void'", "'printf'",
	"'goto'", "'heap'", "'stack'", "'P'", "'H'", "'#include'", "'<stdio.h>'",
	"", "", "", "", "", "", "':'", "'.'", "','", "';'", "'>='", "'>'", "'<='",
	"'<'", "'=='", "'=>'", "'!='", "'='", "'%'", "'*'", "'/'", "'-'", "'+'",
	"'!'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "INT", "FLOAT", "MAIN", "IF", "RETURN", "VOID", "PRINTF", "GOTO", "HEAP",
	"STACK", "P", "H", "INCLUDE", "STDIO", "NUMERO", "DECIMAL", "TEMPORAL",
	"LABEL", "ID_CAMEL", "ID", "DOSPUNTOS", "PUNTO", "COMA", "PUNTOCOMA", "MAYOR_I",
	"MAYOR", "MENOR_I", "MENOR", "IGUALDAD", "CASE", "DISTINTO", "IGUAL", "MODULO",
	"MULTIPLICACION", "DIVISION", "RESTA", "SUMA", "NOT", "PAR_IZQ", "PAR_DER",
	"LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "WHITESPACE",
	"COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"inicio", "funciones", "funcion", "bloques", "bloque", "bloque_i", "instrucciones",
	"instruccion", "bloque_f", "print3d", "if3d", "asignacion", "operacion",
	"expresion", "etiqueta", "salto", "retorno", "include", "declaraciones",
	"declaracion", "temporalesTexto", "temporalesLista", "operador",
}

type N3parser struct {
	*antlr.BaseParser
}

// NewN3parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *N3parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewN3parser(input antlr.TokenStream) *N3parser {
	this := new(N3parser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "N3parser.g4"

	return this
}

// N3parser tokens.
const (
	N3parserEOF            = antlr.TokenEOF
	N3parserINT            = 1
	N3parserFLOAT          = 2
	N3parserMAIN           = 3
	N3parserIF             = 4
	N3parserRETURN         = 5
	N3parserVOID           = 6
	N3parserPRINTF         = 7
	N3parserGOTO           = 8
	N3parserHEAP           = 9
	N3parserSTACK          = 10
	N3parserP              = 11
	N3parserH              = 12
	N3parserINCLUDE        = 13
	N3parserSTDIO          = 14
	N3parserNUMERO         = 15
	N3parserDECIMAL        = 16
	N3parserTEMPORAL       = 17
	N3parserLABEL          = 18
	N3parserID_CAMEL       = 19
	N3parserID             = 20
	N3parserDOSPUNTOS      = 21
	N3parserPUNTO          = 22
	N3parserCOMA           = 23
	N3parserPUNTOCOMA      = 24
	N3parserMAYOR_I        = 25
	N3parserMAYOR          = 26
	N3parserMENOR_I        = 27
	N3parserMENOR          = 28
	N3parserIGUALDAD       = 29
	N3parserCASE           = 30
	N3parserDISTINTO       = 31
	N3parserIGUAL          = 32
	N3parserMODULO         = 33
	N3parserMULTIPLICACION = 34
	N3parserDIVISION       = 35
	N3parserRESTA          = 36
	N3parserSUMA           = 37
	N3parserNOT            = 38
	N3parserPAR_IZQ        = 39
	N3parserPAR_DER        = 40
	N3parserLLAVE_IZQ      = 41
	N3parserLLAVE_DER      = 42
	N3parserCORCHETE_IZQ   = 43
	N3parserCORCHETE_DER   = 44
	N3parserCADENA         = 45
	N3parserWHITESPACE     = 46
	N3parserCOMMENT        = 47
	N3parserLINE_COMMENT   = 48
)

// N3parser rules.
const (
	N3parserRULE_inicio          = 0
	N3parserRULE_funciones       = 1
	N3parserRULE_funcion         = 2
	N3parserRULE_bloques         = 3
	N3parserRULE_bloque          = 4
	N3parserRULE_bloque_i        = 5
	N3parserRULE_instrucciones   = 6
	N3parserRULE_instruccion     = 7
	N3parserRULE_bloque_f        = 8
	N3parserRULE_print3d         = 9
	N3parserRULE_if3d            = 10
	N3parserRULE_asignacion      = 11
	N3parserRULE_operacion       = 12
	N3parserRULE_expresion       = 13
	N3parserRULE_etiqueta        = 14
	N3parserRULE_salto           = 15
	N3parserRULE_retorno         = 16
	N3parserRULE_include         = 17
	N3parserRULE_declaraciones   = 18
	N3parserRULE_declaracion     = 19
	N3parserRULE_temporalesTexto = 20
	N3parserRULE_temporalesLista = 21
	N3parserRULE_operador        = 22
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_include returns the _include rule contexts.
	Get_include() IIncludeContext

	// Get_declaraciones returns the _declaraciones rule contexts.
	Get_declaraciones() IDeclaracionesContext

	// Get_funciones returns the _funciones rule contexts.
	Get_funciones() IFuncionesContext

	// Set_include sets the _include rule contexts.
	Set_include(IIncludeContext)

	// Set_declaraciones sets the _declaraciones rule contexts.
	Set_declaraciones(IDeclaracionesContext)

	// Set_funciones sets the _funciones rule contexts.
	Set_funciones(IFuncionesContext)

	// GetEx returns the ex attribute.
	GetEx() bloque3d.Bloque3dGlobal

	// SetEx sets the ex attribute.
	SetEx(bloque3d.Bloque3dGlobal)

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	ex             bloque3d.Bloque3dGlobal
	_include       IIncludeContext
	_declaraciones IDeclaracionesContext
	_funciones     IFuncionesContext
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_inicio
	return p
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) Get_include() IIncludeContext { return s._include }

func (s *InicioContext) Get_declaraciones() IDeclaracionesContext { return s._declaraciones }

func (s *InicioContext) Get_funciones() IFuncionesContext { return s._funciones }

func (s *InicioContext) Set_include(v IIncludeContext) { s._include = v }

func (s *InicioContext) Set_declaraciones(v IDeclaracionesContext) { s._declaraciones = v }

func (s *InicioContext) Set_funciones(v IFuncionesContext) { s._funciones = v }

func (s *InicioContext) GetEx() bloque3d.Bloque3dGlobal { return s.ex }

func (s *InicioContext) SetEx(v bloque3d.Bloque3dGlobal) { s.ex = v }

func (s *InicioContext) Include() IIncludeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncludeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncludeContext)
}

func (s *InicioContext) Declaraciones() IDeclaracionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionesContext)
}

func (s *InicioContext) Funciones() IFuncionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncionesContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterInicio(s)
	}
}

func (s *InicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitInicio(s)
	}
}

func (p *N3parser) Inicio() (localctx IInicioContext) {
	this := p
	_ = this

	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, N3parserRULE_inicio)

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
		p.SetState(46)

		var _x = p.Include()

		localctx.(*InicioContext)._include = _x
	}
	{
		p.SetState(47)

		var _x = p.Declaraciones()

		localctx.(*InicioContext)._declaraciones = _x
	}
	{
		p.SetState(48)

		var _x = p.Funciones()

		localctx.(*InicioContext)._funciones = _x
	}

	localctx.(*InicioContext).ex = bloque3d.NewBloque3dGlobal(localctx.(*InicioContext).Get_declaraciones().GetList(),
		localctx.(*InicioContext).Get_funciones().GetList(), localctx.(*InicioContext).Get_include().GetEx())

	return localctx
}

// IFuncionesContext is an interface to support dynamic dispatch.
type IFuncionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_funcion returns the _funcion rule contexts.
	Get_funcion() IFuncionContext

	// Set_funcion sets the _funcion rule contexts.
	Set_funcion(IFuncionContext)

	// GetLista returns the lista rule context list.
	GetLista() []IFuncionContext

	// SetLista sets the lista rule context list.
	SetLista([]IFuncionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsFuncionesContext differentiates from other interfaces.
	IsFuncionesContext()
}

type FuncionesContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	list     *arraylist.List
	_funcion IFuncionContext
	lista    []IFuncionContext
}

func NewEmptyFuncionesContext() *FuncionesContext {
	var p = new(FuncionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_funciones
	return p
}

func (*FuncionesContext) IsFuncionesContext() {}

func NewFuncionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionesContext {
	var p = new(FuncionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_funciones

	return p
}

func (s *FuncionesContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionesContext) Get_funcion() IFuncionContext { return s._funcion }

func (s *FuncionesContext) Set_funcion(v IFuncionContext) { s._funcion = v }

func (s *FuncionesContext) GetLista() []IFuncionContext { return s.lista }

func (s *FuncionesContext) SetLista(v []IFuncionContext) { s.lista = v }

func (s *FuncionesContext) GetList() *arraylist.List { return s.list }

func (s *FuncionesContext) SetList(v *arraylist.List) { s.list = v }

func (s *FuncionesContext) AllFuncion() []IFuncionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncionContext)(nil)).Elem())
	var tst = make([]IFuncionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncionContext)
		}
	}

	return tst
}

func (s *FuncionesContext) Funcion(i int) IFuncionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncionContext)
}

func (s *FuncionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterFunciones(s)
	}
}

func (s *FuncionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitFunciones(s)
	}
}

func (p *N3parser) Funciones() (localctx IFuncionesContext) {
	this := p
	_ = this

	localctx = NewFuncionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, N3parserRULE_funciones)
	localctx.(*FuncionesContext).list = arraylist.New()
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
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == N3parserINT || _la == N3parserVOID {
		{
			p.SetState(51)

			var _x = p.Funcion()

			localctx.(*FuncionesContext)._funcion = _x
		}
		localctx.(*FuncionesContext).lista = append(localctx.(*FuncionesContext).lista, localctx.(*FuncionesContext)._funcion)

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listas := localctx.(*FuncionesContext).GetLista()
	for _, e := range listas {
		localctx.(*FuncionesContext).list.Add(e.GetEx())
	}

	return localctx
}

// IFuncionContext is an interface to support dynamic dispatch.
type IFuncionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MAIN returns the _MAIN token.
	Get_MAIN() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_MAIN sets the _MAIN token.
	Set_MAIN(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_bloques returns the _bloques rule contexts.
	Get_bloques() IBloquesContext

	// Set_bloques sets the _bloques rule contexts.
	Set_bloques(IBloquesContext)

	// GetEx returns the ex attribute.
	GetEx() funciones3d.Funcion3D

	// SetEx sets the ex attribute.
	SetEx(funciones3d.Funcion3D)

	// IsFuncionContext differentiates from other interfaces.
	IsFuncionContext()
}

type FuncionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	ex       funciones3d.Funcion3D
	_MAIN    antlr.Token
	_bloques IBloquesContext
	_ID      antlr.Token
}

func NewEmptyFuncionContext() *FuncionContext {
	var p = new(FuncionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_funcion
	return p
}

func (*FuncionContext) IsFuncionContext() {}

func NewFuncionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionContext {
	var p = new(FuncionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_funcion

	return p
}

func (s *FuncionContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionContext) Get_MAIN() antlr.Token { return s._MAIN }

func (s *FuncionContext) Get_ID() antlr.Token { return s._ID }

func (s *FuncionContext) Set_MAIN(v antlr.Token) { s._MAIN = v }

func (s *FuncionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FuncionContext) Get_bloques() IBloquesContext { return s._bloques }

func (s *FuncionContext) Set_bloques(v IBloquesContext) { s._bloques = v }

func (s *FuncionContext) GetEx() funciones3d.Funcion3D { return s.ex }

func (s *FuncionContext) SetEx(v funciones3d.Funcion3D) { s.ex = v }

func (s *FuncionContext) INT() antlr.TerminalNode {
	return s.GetToken(N3parserINT, 0)
}

func (s *FuncionContext) MAIN() antlr.TerminalNode {
	return s.GetToken(N3parserMAIN, 0)
}

func (s *FuncionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(N3parserPAR_IZQ, 0)
}

func (s *FuncionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(N3parserPAR_DER, 0)
}

func (s *FuncionContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(N3parserLLAVE_IZQ, 0)
}

func (s *FuncionContext) Bloques() IBloquesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloquesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloquesContext)
}

func (s *FuncionContext) VOID() antlr.TerminalNode {
	return s.GetToken(N3parserVOID, 0)
}

func (s *FuncionContext) ID() antlr.TerminalNode {
	return s.GetToken(N3parserID, 0)
}

func (s *FuncionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterFuncion(s)
	}
}

func (s *FuncionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitFuncion(s)
	}
}

func (p *N3parser) Funcion() (localctx IFuncionContext) {
	this := p
	_ = this

	localctx = NewFuncionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, N3parserRULE_funcion)

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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case N3parserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(N3parserINT)
		}
		{
			p.SetState(59)

			var _m = p.Match(N3parserMAIN)

			localctx.(*FuncionContext)._MAIN = _m
		}
		{
			p.SetState(60)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(61)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(62)
			p.Match(N3parserLLAVE_IZQ)
		}
		{
			p.SetState(63)

			var _x = p.Bloques()

			localctx.(*FuncionContext)._bloques = _x
		}

		localctx.(*FuncionContext).ex = funciones3d.NewFuncion3D((func() string {
			if localctx.(*FuncionContext).Get_MAIN() == nil {
				return ""
			} else {
				return localctx.(*FuncionContext).Get_MAIN().GetText()
			}
		}()), localctx.(*FuncionContext).Get_bloques().GetList(), interfaces3d.MAIN3D)

	case N3parserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(N3parserVOID)
		}
		{
			p.SetState(67)

			var _m = p.Match(N3parserID)

			localctx.(*FuncionContext)._ID = _m
		}
		{
			p.SetState(68)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(69)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(70)
			p.Match(N3parserLLAVE_IZQ)
		}
		{
			p.SetState(71)

			var _x = p.Bloques()

			localctx.(*FuncionContext)._bloques = _x
		}

		localctx.(*FuncionContext).ex = funciones3d.NewFuncion3D((func() string {
			if localctx.(*FuncionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FuncionContext).Get_ID().GetText()
			}
		}()), localctx.(*FuncionContext).Get_bloques().GetList(), interfaces3d.FUNCION3D)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBloquesContext is an interface to support dynamic dispatch.
type IBloquesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_bloque returns the _bloque rule contexts.
	Get_bloque() IBloqueContext

	// Set_bloque sets the _bloque rule contexts.
	Set_bloque(IBloqueContext)

	// GetLista returns the lista rule context list.
	GetLista() []IBloqueContext

	// SetLista sets the lista rule context list.
	SetLista([]IBloqueContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsBloquesContext differentiates from other interfaces.
	IsBloquesContext()
}

type BloquesContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	list    *arraylist.List
	_bloque IBloqueContext
	lista   []IBloqueContext
}

func NewEmptyBloquesContext() *BloquesContext {
	var p = new(BloquesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_bloques
	return p
}

func (*BloquesContext) IsBloquesContext() {}

func NewBloquesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloquesContext {
	var p = new(BloquesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_bloques

	return p
}

func (s *BloquesContext) GetParser() antlr.Parser { return s.parser }

func (s *BloquesContext) Get_bloque() IBloqueContext { return s._bloque }

func (s *BloquesContext) Set_bloque(v IBloqueContext) { s._bloque = v }

func (s *BloquesContext) GetLista() []IBloqueContext { return s.lista }

func (s *BloquesContext) SetLista(v []IBloqueContext) { s.lista = v }

func (s *BloquesContext) GetList() *arraylist.List { return s.list }

func (s *BloquesContext) SetList(v *arraylist.List) { s.list = v }

func (s *BloquesContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *BloquesContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *BloquesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloquesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloquesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterBloques(s)
	}
}

func (s *BloquesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitBloques(s)
	}
}

func (p *N3parser) Bloques() (localctx IBloquesContext) {
	this := p
	_ = this

	localctx = NewBloquesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, N3parserRULE_bloques)
	localctx.(*BloquesContext).list = arraylist.New()
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<N3parserIF)|(1<<N3parserRETURN)|(1<<N3parserPRINTF)|(1<<N3parserGOTO)|(1<<N3parserHEAP)|(1<<N3parserSTACK)|(1<<N3parserP)|(1<<N3parserH)|(1<<N3parserNUMERO)|(1<<N3parserDECIMAL)|(1<<N3parserTEMPORAL)|(1<<N3parserLABEL))) != 0) || _la == N3parserRESTA {
		{
			p.SetState(76)

			var _x = p.Bloque()

			localctx.(*BloquesContext)._bloque = _x
		}
		localctx.(*BloquesContext).lista = append(localctx.(*BloquesContext).lista, localctx.(*BloquesContext)._bloque)

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listas := localctx.(*BloquesContext).GetLista()
	for _, e := range listas {
		localctx.(*BloquesContext).list.Add(e.GetEx())
	}

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInit returns the init rule contexts.
	GetInit() IBloque_iContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// GetFin returns the fin rule contexts.
	GetFin() IBloque_fContext

	// SetInit sets the init rule contexts.
	SetInit(IBloque_iContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// SetFin sets the fin rule contexts.
	SetFin(IBloque_fContext)

	// GetEx returns the ex attribute.
	GetEx() bloque3d.Bloque3d

	// SetEx sets the ex attribute.
	SetEx(bloque3d.Bloque3d)

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	ex             bloque3d.Bloque3d
	init           IBloque_iContext
	_instrucciones IInstruccionesContext
	fin            IBloque_fContext
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) GetInit() IBloque_iContext { return s.init }

func (s *BloqueContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *BloqueContext) GetFin() IBloque_fContext { return s.fin }

func (s *BloqueContext) SetInit(v IBloque_iContext) { s.init = v }

func (s *BloqueContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *BloqueContext) SetFin(v IBloque_fContext) { s.fin = v }

func (s *BloqueContext) GetEx() bloque3d.Bloque3d { return s.ex }

func (s *BloqueContext) SetEx(v bloque3d.Bloque3d) { s.ex = v }

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) Bloque_i() IBloque_iContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloque_iContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloque_iContext)
}

func (s *BloqueContext) Bloque_f() IBloque_fContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloque_fContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloque_fContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (p *N3parser) Bloque() (localctx IBloqueContext) {
	this := p
	_ = this

	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, N3parserRULE_bloque)

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

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)

			var _x = p.Bloque_i()

			localctx.(*BloqueContext).init = _x
		}
		{
			p.SetState(84)

			var _x = p.Instrucciones()

			localctx.(*BloqueContext)._instrucciones = _x
		}
		{
			p.SetState(85)

			var _x = p.Bloque_f()

			localctx.(*BloqueContext).fin = _x
		}

		nLista := arraylist.New()
		if localctx.(*BloqueContext).GetInit().GetEx() != nil {
			nLista.Add(localctx.(*BloqueContext).GetInit().GetEx())
		}
		listas := localctx.(*BloqueContext).Get_instrucciones().GetList()
		for i := 0; i < listas.Len(); i++ {
			nLista.Add(listas.GetValue(i))
		}
		if localctx.(*BloqueContext).GetFin().GetEx() != nil {
			nLista.Add(localctx.(*BloqueContext).GetFin().GetEx())
		}
		localctx.(*BloqueContext).ex = bloque3d.NewBloque3d(nLista)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(88)

			var _x = p.Instrucciones()

			localctx.(*BloqueContext)._instrucciones = _x
		}

		nLista := arraylist.New()
		listas := localctx.(*BloqueContext).Get_instrucciones().GetList()
		for i := 0; i < listas.Len(); i++ {
			nLista.Add(listas.GetValue(i))
		}
		localctx.(*BloqueContext).ex = bloque3d.NewBloque3d(nLista)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)

			var _x = p.Bloque_i()

			localctx.(*BloqueContext).init = _x
		}
		{
			p.SetState(92)

			var _x = p.Bloque_f()

			localctx.(*BloqueContext).fin = _x
		}

		nLista := arraylist.New()
		if localctx.(*BloqueContext).GetInit().GetEx() != nil {
			nLista.Add(localctx.(*BloqueContext).GetInit().GetEx())
		}
		if localctx.(*BloqueContext).GetFin().GetEx() != nil {
			nLista.Add(localctx.(*BloqueContext).GetFin().GetEx())
		}
		localctx.(*BloqueContext).ex = bloque3d.NewBloque3d(nLista)

	}

	return localctx
}

// IBloque_iContext is an interface to support dynamic dispatch.
type IBloque_iContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_etiqueta returns the _etiqueta rule contexts.
	Get_etiqueta() IEtiquetaContext

	// Get_print3d returns the _print3d rule contexts.
	Get_print3d() IPrint3dContext

	// Get_if3d returns the _if3d rule contexts.
	Get_if3d() IIf3dContext

	// Get_salto returns the _salto rule contexts.
	Get_salto() ISaltoContext

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_etiqueta sets the _etiqueta rule contexts.
	Set_etiqueta(IEtiquetaContext)

	// Set_print3d sets the _print3d rule contexts.
	Set_print3d(IPrint3dContext)

	// Set_if3d sets the _if3d rule contexts.
	Set_if3d(IIf3dContext)

	// Set_salto sets the _salto rule contexts.
	Set_salto(ISaltoContext)

	// GetEx returns the ex attribute.
	GetEx() interfaces3d.Expresion3D

	// SetEx sets the ex attribute.
	SetEx(interfaces3d.Expresion3D)

	// IsBloque_iContext differentiates from other interfaces.
	IsBloque_iContext()
}

type Bloque_iContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ex          interfaces3d.Expresion3D
	_asignacion IAsignacionContext
	_etiqueta   IEtiquetaContext
	_print3d    IPrint3dContext
	_if3d       IIf3dContext
	_salto      ISaltoContext
}

func NewEmptyBloque_iContext() *Bloque_iContext {
	var p = new(Bloque_iContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_bloque_i
	return p
}

func (*Bloque_iContext) IsBloque_iContext() {}

func NewBloque_iContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bloque_iContext {
	var p = new(Bloque_iContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_bloque_i

	return p
}

func (s *Bloque_iContext) GetParser() antlr.Parser { return s.parser }

func (s *Bloque_iContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *Bloque_iContext) Get_etiqueta() IEtiquetaContext { return s._etiqueta }

func (s *Bloque_iContext) Get_print3d() IPrint3dContext { return s._print3d }

func (s *Bloque_iContext) Get_if3d() IIf3dContext { return s._if3d }

func (s *Bloque_iContext) Get_salto() ISaltoContext { return s._salto }

func (s *Bloque_iContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *Bloque_iContext) Set_etiqueta(v IEtiquetaContext) { s._etiqueta = v }

func (s *Bloque_iContext) Set_print3d(v IPrint3dContext) { s._print3d = v }

func (s *Bloque_iContext) Set_if3d(v IIf3dContext) { s._if3d = v }

func (s *Bloque_iContext) Set_salto(v ISaltoContext) { s._salto = v }

func (s *Bloque_iContext) GetEx() interfaces3d.Expresion3D { return s.ex }

func (s *Bloque_iContext) SetEx(v interfaces3d.Expresion3D) { s.ex = v }

func (s *Bloque_iContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *Bloque_iContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(N3parserPUNTOCOMA, 0)
}

func (s *Bloque_iContext) Etiqueta() IEtiquetaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEtiquetaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEtiquetaContext)
}

func (s *Bloque_iContext) Print3d() IPrint3dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrint3dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrint3dContext)
}

func (s *Bloque_iContext) If3d() IIf3dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf3dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf3dContext)
}

func (s *Bloque_iContext) Salto() ISaltoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISaltoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISaltoContext)
}

func (s *Bloque_iContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bloque_iContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bloque_iContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterBloque_i(s)
	}
}

func (s *Bloque_iContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitBloque_i(s)
	}
}

func (p *N3parser) Bloque_i() (localctx IBloque_iContext) {
	this := p
	_ = this

	localctx = NewBloque_iContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, N3parserRULE_bloque_i)

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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case N3parserHEAP, N3parserSTACK, N3parserP, N3parserH, N3parserNUMERO, N3parserDECIMAL, N3parserTEMPORAL, N3parserRESTA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)

			var _x = p.Asignacion()

			localctx.(*Bloque_iContext)._asignacion = _x
		}
		{
			p.SetState(98)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*Bloque_iContext).ex = localctx.(*Bloque_iContext).Get_asignacion().GetEx()

	case N3parserLABEL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)

			var _x = p.Etiqueta()

			localctx.(*Bloque_iContext)._etiqueta = _x
		}

		localctx.(*Bloque_iContext).ex = localctx.(*Bloque_iContext).Get_etiqueta().GetEx()

	case N3parserPRINTF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)

			var _x = p.Print3d()

			localctx.(*Bloque_iContext)._print3d = _x
		}
		{
			p.SetState(105)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*Bloque_iContext).ex = localctx.(*Bloque_iContext).Get_print3d().GetEx()

	case N3parserIF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)

			var _x = p.If3d()

			localctx.(*Bloque_iContext)._if3d = _x
		}

		localctx.(*Bloque_iContext).ex = localctx.(*Bloque_iContext).Get_if3d().GetEx()

	case N3parserGOTO:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)

			var _x = p.Salto()

			localctx.(*Bloque_iContext)._salto = _x
		}
		{
			p.SetState(112)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*Bloque_iContext).ex = localctx.(*Bloque_iContext).Get_salto().GetEx()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetLista returns the lista rule context list.
	GetLista() []IInstruccionContext

	// SetLista sets the lista rule context list.
	SetLista([]IInstruccionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	list         *arraylist.List
	_instruccion IInstruccionContext
	lista        []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetLista() []IInstruccionContext { return s.lista }

func (s *InstruccionesContext) SetLista(v []IInstruccionContext) { s.lista = v }

func (s *InstruccionesContext) GetList() *arraylist.List { return s.list }

func (s *InstruccionesContext) SetList(v *arraylist.List) { s.list = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *N3parser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, N3parserRULE_instrucciones)
	localctx.(*InstruccionesContext).list = arraylist.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(117)

				var _x = p.Instruccion()

				localctx.(*InstruccionesContext)._instruccion = _x
			}
			localctx.(*InstruccionesContext).lista = append(localctx.(*InstruccionesContext).lista, localctx.(*InstruccionesContext)._instruccion)

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	listas := localctx.(*InstruccionesContext).GetLista()
	for _, e := range listas {
		localctx.(*InstruccionesContext).list.Add(e.GetEx())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_retorno returns the _retorno rule contexts.
	Get_retorno() IRetornoContext

	// Get_print3d returns the _print3d rule contexts.
	Get_print3d() IPrint3dContext

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_retorno sets the _retorno rule contexts.
	Set_retorno(IRetornoContext)

	// Set_print3d sets the _print3d rule contexts.
	Set_print3d(IPrint3dContext)

	// GetEx returns the ex attribute.
	GetEx() interfaces3d.Expresion3D

	// SetEx sets the ex attribute.
	SetEx(interfaces3d.Expresion3D)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	ex          interfaces3d.Expresion3D
	_asignacion IAsignacionContext
	_retorno    IRetornoContext
	_print3d    IPrint3dContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Get_retorno() IRetornoContext { return s._retorno }

func (s *InstruccionContext) Get_print3d() IPrint3dContext { return s._print3d }

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_retorno(v IRetornoContext) { s._retorno = v }

func (s *InstruccionContext) Set_print3d(v IPrint3dContext) { s._print3d = v }

func (s *InstruccionContext) GetEx() interfaces3d.Expresion3D { return s.ex }

func (s *InstruccionContext) SetEx(v interfaces3d.Expresion3D) { s.ex = v }

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(N3parserPUNTOCOMA, 0)
}

func (s *InstruccionContext) Retorno() IRetornoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetornoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetornoContext)
}

func (s *InstruccionContext) Print3d() IPrint3dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrint3dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrint3dContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *N3parser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, N3parserRULE_instruccion)

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

	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case N3parserHEAP, N3parserSTACK, N3parserP, N3parserH, N3parserNUMERO, N3parserDECIMAL, N3parserTEMPORAL, N3parserRESTA:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		{
			p.SetState(125)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_asignacion().GetEx()

	case N3parserRETURN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)

			var _x = p.Retorno()

			localctx.(*InstruccionContext)._retorno = _x
		}
		{
			p.SetState(129)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_retorno().GetEx()

	case N3parserPRINTF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)

			var _x = p.Print3d()

			localctx.(*InstruccionContext)._print3d = _x
		}
		{
			p.SetState(133)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_print3d().GetEx()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBloque_fContext is an interface to support dynamic dispatch.
type IBloque_fContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_etiqueta returns the _etiqueta rule contexts.
	Get_etiqueta() IEtiquetaContext

	// Get_salto returns the _salto rule contexts.
	Get_salto() ISaltoContext

	// Get_if3d returns the _if3d rule contexts.
	Get_if3d() IIf3dContext

	// Set_etiqueta sets the _etiqueta rule contexts.
	Set_etiqueta(IEtiquetaContext)

	// Set_salto sets the _salto rule contexts.
	Set_salto(ISaltoContext)

	// Set_if3d sets the _if3d rule contexts.
	Set_if3d(IIf3dContext)

	// GetEx returns the ex attribute.
	GetEx() interfaces3d.Expresion3D

	// SetEx sets the ex attribute.
	SetEx(interfaces3d.Expresion3D)

	// IsBloque_fContext differentiates from other interfaces.
	IsBloque_fContext()
}

type Bloque_fContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	ex        interfaces3d.Expresion3D
	_etiqueta IEtiquetaContext
	_salto    ISaltoContext
	_if3d     IIf3dContext
}

func NewEmptyBloque_fContext() *Bloque_fContext {
	var p = new(Bloque_fContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_bloque_f
	return p
}

func (*Bloque_fContext) IsBloque_fContext() {}

func NewBloque_fContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bloque_fContext {
	var p = new(Bloque_fContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_bloque_f

	return p
}

func (s *Bloque_fContext) GetParser() antlr.Parser { return s.parser }

func (s *Bloque_fContext) Get_etiqueta() IEtiquetaContext { return s._etiqueta }

func (s *Bloque_fContext) Get_salto() ISaltoContext { return s._salto }

func (s *Bloque_fContext) Get_if3d() IIf3dContext { return s._if3d }

func (s *Bloque_fContext) Set_etiqueta(v IEtiquetaContext) { s._etiqueta = v }

func (s *Bloque_fContext) Set_salto(v ISaltoContext) { s._salto = v }

func (s *Bloque_fContext) Set_if3d(v IIf3dContext) { s._if3d = v }

func (s *Bloque_fContext) GetEx() interfaces3d.Expresion3D { return s.ex }

func (s *Bloque_fContext) SetEx(v interfaces3d.Expresion3D) { s.ex = v }

func (s *Bloque_fContext) Etiqueta() IEtiquetaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEtiquetaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEtiquetaContext)
}

func (s *Bloque_fContext) Salto() ISaltoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISaltoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISaltoContext)
}

func (s *Bloque_fContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(N3parserPUNTOCOMA, 0)
}

func (s *Bloque_fContext) If3d() IIf3dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf3dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf3dContext)
}

func (s *Bloque_fContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(N3parserLLAVE_DER, 0)
}

func (s *Bloque_fContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bloque_fContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bloque_fContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterBloque_f(s)
	}
}

func (s *Bloque_fContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitBloque_f(s)
	}
}

func (p *N3parser) Bloque_f() (localctx IBloque_fContext) {
	this := p
	_ = this

	localctx = NewBloque_fContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, N3parserRULE_bloque_f)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case N3parserLABEL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)

			var _x = p.Etiqueta()

			localctx.(*Bloque_fContext)._etiqueta = _x
		}

		localctx.(*Bloque_fContext).ex = localctx.(*Bloque_fContext).Get_etiqueta().GetEx()

	case N3parserGOTO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)

			var _x = p.Salto()

			localctx.(*Bloque_fContext)._salto = _x
		}
		{
			p.SetState(142)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*Bloque_fContext).ex = localctx.(*Bloque_fContext).Get_salto().GetEx()

	case N3parserIF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)

			var _x = p.If3d()

			localctx.(*Bloque_fContext)._if3d = _x
		}

		localctx.(*Bloque_fContext).ex = localctx.(*Bloque_fContext).Get_if3d().GetEx()

	case N3parserLLAVE_DER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(148)
			p.Match(N3parserLLAVE_DER)
		}

		localctx.(*Bloque_fContext).ex = expresiones3d.NewPrimitivo3D("", interfaces3d.PUNTERO_STACK)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrint3dContext is an interface to support dynamic dispatch.
type IPrint3dContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_TEMPORAL returns the _TEMPORAL token.
	Get_TEMPORAL() antlr.Token

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_TEMPORAL sets the _TEMPORAL token.
	Set_TEMPORAL(antlr.Token)

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() instrucciones3d.Print3D

	// SetEx sets the ex attribute.
	SetEx(instrucciones3d.Print3D)

	// IsPrint3dContext differentiates from other interfaces.
	IsPrint3dContext()
}

type Print3dContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	ex        instrucciones3d.Print3D
	_CADENA   antlr.Token
	_TEMPORAL antlr.Token
	_NUMERO   antlr.Token
	_DECIMAL  antlr.Token
}

func NewEmptyPrint3dContext() *Print3dContext {
	var p = new(Print3dContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_print3d
	return p
}

func (*Print3dContext) IsPrint3dContext() {}

func NewPrint3dContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print3dContext {
	var p = new(Print3dContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_print3d

	return p
}

func (s *Print3dContext) GetParser() antlr.Parser { return s.parser }

func (s *Print3dContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *Print3dContext) Get_TEMPORAL() antlr.Token { return s._TEMPORAL }

func (s *Print3dContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *Print3dContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *Print3dContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *Print3dContext) Set_TEMPORAL(v antlr.Token) { s._TEMPORAL = v }

func (s *Print3dContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *Print3dContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *Print3dContext) GetEx() instrucciones3d.Print3D { return s.ex }

func (s *Print3dContext) SetEx(v instrucciones3d.Print3D) { s.ex = v }

func (s *Print3dContext) PRINTF() antlr.TerminalNode {
	return s.GetToken(N3parserPRINTF, 0)
}

func (s *Print3dContext) AllPAR_IZQ() []antlr.TerminalNode {
	return s.GetTokens(N3parserPAR_IZQ)
}

func (s *Print3dContext) PAR_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(N3parserPAR_IZQ, i)
}

func (s *Print3dContext) CADENA() antlr.TerminalNode {
	return s.GetToken(N3parserCADENA, 0)
}

func (s *Print3dContext) COMA() antlr.TerminalNode {
	return s.GetToken(N3parserCOMA, 0)
}

func (s *Print3dContext) INT() antlr.TerminalNode {
	return s.GetToken(N3parserINT, 0)
}

func (s *Print3dContext) AllPAR_DER() []antlr.TerminalNode {
	return s.GetTokens(N3parserPAR_DER)
}

func (s *Print3dContext) PAR_DER(i int) antlr.TerminalNode {
	return s.GetToken(N3parserPAR_DER, i)
}

func (s *Print3dContext) TEMPORAL() antlr.TerminalNode {
	return s.GetToken(N3parserTEMPORAL, 0)
}

func (s *Print3dContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(N3parserNUMERO, 0)
}

func (s *Print3dContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(N3parserDECIMAL, 0)
}

func (s *Print3dContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print3dContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Print3dContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterPrint3d(s)
	}
}

func (s *Print3dContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitPrint3d(s)
	}
}

func (p *N3parser) Print3d() (localctx IPrint3dContext) {
	this := p
	_ = this

	localctx = NewPrint3dContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, N3parserRULE_print3d)

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

	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(N3parserPRINTF)
		}
		{
			p.SetState(153)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(154)

			var _m = p.Match(N3parserCADENA)

			localctx.(*Print3dContext)._CADENA = _m
		}
		{
			p.SetState(155)
			p.Match(N3parserCOMA)
		}
		{
			p.SetState(156)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(157)
			p.Match(N3parserINT)
		}
		{
			p.SetState(158)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(159)

			var _m = p.Match(N3parserTEMPORAL)

			localctx.(*Print3dContext)._TEMPORAL = _m
		}
		{
			p.SetState(160)
			p.Match(N3parserPAR_DER)
		}

		cad := (func() string {
			if localctx.(*Print3dContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_CADENA().GetText()
			}
		}())
		if strings.Contains(cad, "c") {
			cad = "%%c"
		} else {
			cad = "%%d"
		}
		valor := "printf(\"" + cad + "\",(int)" + (func() string {
			if localctx.(*Print3dContext).Get_TEMPORAL() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_TEMPORAL().GetText()
			}
		}()) + ")"
		localctx.(*Print3dContext).ex = instrucciones3d.NewPrint3D(valor)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(N3parserPRINTF)
		}
		{
			p.SetState(163)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(164)

			var _m = p.Match(N3parserCADENA)

			localctx.(*Print3dContext)._CADENA = _m
		}
		{
			p.SetState(165)
			p.Match(N3parserCOMA)
		}
		{
			p.SetState(166)

			var _m = p.Match(N3parserNUMERO)

			localctx.(*Print3dContext)._NUMERO = _m
		}
		{
			p.SetState(167)
			p.Match(N3parserPAR_DER)
		}

		cad := (func() string {
			if localctx.(*Print3dContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_CADENA().GetText()
			}
		}())
		if strings.Contains(cad, "c") {
			cad = "%%c"
		} else {
			cad = "%%d"
		}
		valor := "printf(\"" + cad + "\",(int)" + (func() string {
			if localctx.(*Print3dContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_NUMERO().GetText()
			}
		}()) + ")"
		localctx.(*Print3dContext).ex = instrucciones3d.NewPrint3D(valor)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(169)
			p.Match(N3parserPRINTF)
		}
		{
			p.SetState(170)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(171)

			var _m = p.Match(N3parserCADENA)

			localctx.(*Print3dContext)._CADENA = _m
		}
		{
			p.SetState(172)
			p.Match(N3parserCOMA)
		}
		{
			p.SetState(173)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(174)
			p.Match(N3parserINT)
		}
		{
			p.SetState(175)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(176)

			var _m = p.Match(N3parserNUMERO)

			localctx.(*Print3dContext)._NUMERO = _m
		}
		{
			p.SetState(177)
			p.Match(N3parserPAR_DER)
		}

		cad := (func() string {
			if localctx.(*Print3dContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_CADENA().GetText()
			}
		}())
		if strings.Contains(cad, "c") {
			cad = "%%c"
		} else {
			cad = "%%d"
		}
		valor := "printf(\"" + cad + "\",(int)" + (func() string {
			if localctx.(*Print3dContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_NUMERO().GetText()
			}
		}()) + ")"
		localctx.(*Print3dContext).ex = instrucciones3d.NewPrint3D(valor)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(179)
			p.Match(N3parserPRINTF)
		}
		{
			p.SetState(180)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(181)

			var _m = p.Match(N3parserCADENA)

			localctx.(*Print3dContext)._CADENA = _m
		}
		{
			p.SetState(182)
			p.Match(N3parserCOMA)
		}
		{
			p.SetState(183)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(184)
			p.Match(N3parserINT)
		}
		{
			p.SetState(185)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(186)

			var _m = p.Match(N3parserDECIMAL)

			localctx.(*Print3dContext)._DECIMAL = _m
		}
		{
			p.SetState(187)
			p.Match(N3parserPAR_DER)
		}

		cad := (func() string {
			if localctx.(*Print3dContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_CADENA().GetText()
			}
		}())
		if strings.Contains(cad, "c") {
			cad = "%%c"
		} else {
			cad = "%%d"
		}
		valor := "printf(\"" + cad + "\",(int)" + (func() string {
			if localctx.(*Print3dContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*Print3dContext).Get_DECIMAL().GetText()
			}
		}()) + ")"
		localctx.(*Print3dContext).ex = instrucciones3d.NewPrint3D(valor)

	}

	return localctx
}

// IIf3dContext is an interface to support dynamic dispatch.
type IIf3dContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_operacion returns the _operacion rule contexts.
	Get_operacion() IOperacionContext

	// GetGo1 returns the go1 rule contexts.
	GetGo1() ISaltoContext

	// GetGo2 returns the go2 rule contexts.
	GetGo2() ISaltoContext

	// Get_etiqueta returns the _etiqueta rule contexts.
	Get_etiqueta() IEtiquetaContext

	// Set_operacion sets the _operacion rule contexts.
	Set_operacion(IOperacionContext)

	// SetGo1 sets the go1 rule contexts.
	SetGo1(ISaltoContext)

	// SetGo2 sets the go2 rule contexts.
	SetGo2(ISaltoContext)

	// Set_etiqueta sets the _etiqueta rule contexts.
	Set_etiqueta(IEtiquetaContext)

	// GetEx returns the ex attribute.
	GetEx() control.IF3D

	// SetEx sets the ex attribute.
	SetEx(control.IF3D)

	// IsIf3dContext differentiates from other interfaces.
	IsIf3dContext()
}

type If3dContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         control.IF3D
	_operacion IOperacionContext
	go1        ISaltoContext
	go2        ISaltoContext
	_etiqueta  IEtiquetaContext
}

func NewEmptyIf3dContext() *If3dContext {
	var p = new(If3dContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_if3d
	return p
}

func (*If3dContext) IsIf3dContext() {}

func NewIf3dContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If3dContext {
	var p = new(If3dContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_if3d

	return p
}

func (s *If3dContext) GetParser() antlr.Parser { return s.parser }

func (s *If3dContext) Get_operacion() IOperacionContext { return s._operacion }

func (s *If3dContext) GetGo1() ISaltoContext { return s.go1 }

func (s *If3dContext) GetGo2() ISaltoContext { return s.go2 }

func (s *If3dContext) Get_etiqueta() IEtiquetaContext { return s._etiqueta }

func (s *If3dContext) Set_operacion(v IOperacionContext) { s._operacion = v }

func (s *If3dContext) SetGo1(v ISaltoContext) { s.go1 = v }

func (s *If3dContext) SetGo2(v ISaltoContext) { s.go2 = v }

func (s *If3dContext) Set_etiqueta(v IEtiquetaContext) { s._etiqueta = v }

func (s *If3dContext) GetEx() control.IF3D { return s.ex }

func (s *If3dContext) SetEx(v control.IF3D) { s.ex = v }

func (s *If3dContext) IF() antlr.TerminalNode {
	return s.GetToken(N3parserIF, 0)
}

func (s *If3dContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(N3parserPAR_IZQ, 0)
}

func (s *If3dContext) Operacion() IOperacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperacionContext)
}

func (s *If3dContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(N3parserPAR_DER, 0)
}

func (s *If3dContext) AllPUNTOCOMA() []antlr.TerminalNode {
	return s.GetTokens(N3parserPUNTOCOMA)
}

func (s *If3dContext) PUNTOCOMA(i int) antlr.TerminalNode {
	return s.GetToken(N3parserPUNTOCOMA, i)
}

func (s *If3dContext) Etiqueta() IEtiquetaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEtiquetaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEtiquetaContext)
}

func (s *If3dContext) AllSalto() []ISaltoContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISaltoContext)(nil)).Elem())
	var tst = make([]ISaltoContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISaltoContext)
		}
	}

	return tst
}

func (s *If3dContext) Salto(i int) ISaltoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISaltoContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISaltoContext)
}

func (s *If3dContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If3dContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If3dContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterIf3d(s)
	}
}

func (s *If3dContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitIf3d(s)
	}
}

func (p *N3parser) If3d() (localctx IIf3dContext) {
	this := p
	_ = this

	localctx = NewIf3dContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, N3parserRULE_if3d)

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
		p.SetState(191)
		p.Match(N3parserIF)
	}
	{
		p.SetState(192)
		p.Match(N3parserPAR_IZQ)
	}
	{
		p.SetState(193)

		var _x = p.Operacion()

		localctx.(*If3dContext)._operacion = _x
	}
	{
		p.SetState(194)
		p.Match(N3parserPAR_DER)
	}
	{
		p.SetState(195)

		var _x = p.Salto()

		localctx.(*If3dContext).go1 = _x
	}
	{
		p.SetState(196)
		p.Match(N3parserPUNTOCOMA)
	}
	{
		p.SetState(197)

		var _x = p.Salto()

		localctx.(*If3dContext).go2 = _x
	}
	{
		p.SetState(198)
		p.Match(N3parserPUNTOCOMA)
	}
	{
		p.SetState(199)

		var _x = p.Etiqueta()

		localctx.(*If3dContext)._etiqueta = _x
	}

	localctx.(*If3dContext).ex = control.NewIF3D(localctx.(*If3dContext).Get_operacion().GetEx(), localctx.(*If3dContext).GetGo1().GetEx(), localctx.(*If3dContext).GetGo2().GetEx(), localctx.(*If3dContext).Get_etiqueta().GetEx())

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp returns the exp rule contexts.
	GetExp() IExpresionContext

	// Get_operacion returns the _operacion rule contexts.
	Get_operacion() IOperacionContext

	// GetExp2 returns the exp2 rule contexts.
	GetExp2() IExpresionContext

	// SetExp sets the exp rule contexts.
	SetExp(IExpresionContext)

	// Set_operacion sets the _operacion rule contexts.
	Set_operacion(IOperacionContext)

	// SetExp2 sets the exp2 rule contexts.
	SetExp2(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() instrucciones3d.Asignacion3D

	// SetEx sets the ex attribute.
	SetEx(instrucciones3d.Asignacion3D)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         instrucciones3d.Asignacion3D
	exp        IExpresionContext
	_operacion IOperacionContext
	exp2       IExpresionContext
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) GetExp() IExpresionContext { return s.exp }

func (s *AsignacionContext) Get_operacion() IOperacionContext { return s._operacion }

func (s *AsignacionContext) GetExp2() IExpresionContext { return s.exp2 }

func (s *AsignacionContext) SetExp(v IExpresionContext) { s.exp = v }

func (s *AsignacionContext) Set_operacion(v IOperacionContext) { s._operacion = v }

func (s *AsignacionContext) SetExp2(v IExpresionContext) { s.exp2 = v }

func (s *AsignacionContext) GetEx() instrucciones3d.Asignacion3D { return s.ex }

func (s *AsignacionContext) SetEx(v instrucciones3d.Asignacion3D) { s.ex = v }

func (s *AsignacionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(N3parserIGUAL, 0)
}

func (s *AsignacionContext) Operacion() IOperacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperacionContext)
}

func (s *AsignacionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *AsignacionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *N3parser) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, N3parserRULE_asignacion)

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

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)

			var _x = p.Expresion()

			localctx.(*AsignacionContext).exp = _x
		}
		{
			p.SetState(203)
			p.Match(N3parserIGUAL)
		}
		{
			p.SetState(204)

			var _x = p.Operacion()

			localctx.(*AsignacionContext)._operacion = _x
		}

		localctx.(*AsignacionContext).ex = instrucciones3d.NewAsignacion3D(localctx.(*AsignacionContext).GetExp().GetEx(), localctx.(*AsignacionContext).Get_operacion().GetEx())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)

			var _x = p.Expresion()

			localctx.(*AsignacionContext).exp = _x
		}
		{
			p.SetState(208)
			p.Match(N3parserIGUAL)
		}
		{
			p.SetState(209)

			var _x = p.Expresion()

			localctx.(*AsignacionContext).exp2 = _x
		}

		expr := expresiones3d.NewPrimitivo3D("", interfaces3d.NUMERO3D)
		op := expresiones3d.NewOperacion3D(localctx.(*AsignacionContext).GetExp2().GetEx(), expr, "", true)
		localctx.(*AsignacionContext).ex = instrucciones3d.NewAsignacion3D(localctx.(*AsignacionContext).GetExp().GetEx(), op)

	}

	return localctx
}

// IOperacionContext is an interface to support dynamic dispatch.
type IOperacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpIzq returns the opIzq rule contexts.
	GetOpIzq() IExpresionContext

	// Get_operador returns the _operador rule contexts.
	Get_operador() IOperadorContext

	// GetOpDer returns the opDer rule contexts.
	GetOpDer() IExpresionContext

	// SetOpIzq sets the opIzq rule contexts.
	SetOpIzq(IExpresionContext)

	// Set_operador sets the _operador rule contexts.
	Set_operador(IOperadorContext)

	// SetOpDer sets the opDer rule contexts.
	SetOpDer(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() expresiones3d.Operacion3D

	// SetEx sets the ex attribute.
	SetEx(expresiones3d.Operacion3D)

	// IsOperacionContext differentiates from other interfaces.
	IsOperacionContext()
}

type OperacionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	ex        expresiones3d.Operacion3D
	opIzq     IExpresionContext
	_operador IOperadorContext
	opDer     IExpresionContext
}

func NewEmptyOperacionContext() *OperacionContext {
	var p = new(OperacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_operacion
	return p
}

func (*OperacionContext) IsOperacionContext() {}

func NewOperacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperacionContext {
	var p = new(OperacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_operacion

	return p
}

func (s *OperacionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperacionContext) GetOpIzq() IExpresionContext { return s.opIzq }

func (s *OperacionContext) Get_operador() IOperadorContext { return s._operador }

func (s *OperacionContext) GetOpDer() IExpresionContext { return s.opDer }

func (s *OperacionContext) SetOpIzq(v IExpresionContext) { s.opIzq = v }

func (s *OperacionContext) Set_operador(v IOperadorContext) { s._operador = v }

func (s *OperacionContext) SetOpDer(v IExpresionContext) { s.opDer = v }

func (s *OperacionContext) GetEx() expresiones3d.Operacion3D { return s.ex }

func (s *OperacionContext) SetEx(v expresiones3d.Operacion3D) { s.ex = v }

func (s *OperacionContext) Operador() IOperadorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperadorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperadorContext)
}

func (s *OperacionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *OperacionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *OperacionContext) AllPAR_IZQ() []antlr.TerminalNode {
	return s.GetTokens(N3parserPAR_IZQ)
}

func (s *OperacionContext) PAR_IZQ(i int) antlr.TerminalNode {
	return s.GetToken(N3parserPAR_IZQ, i)
}

func (s *OperacionContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(N3parserINT)
}

func (s *OperacionContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(N3parserINT, i)
}

func (s *OperacionContext) AllPAR_DER() []antlr.TerminalNode {
	return s.GetTokens(N3parserPAR_DER)
}

func (s *OperacionContext) PAR_DER(i int) antlr.TerminalNode {
	return s.GetToken(N3parserPAR_DER, i)
}

func (s *OperacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterOperacion(s)
	}
}

func (s *OperacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitOperacion(s)
	}
}

func (p *N3parser) Operacion() (localctx IOperacionContext) {
	this := p
	_ = this

	localctx = NewOperacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, N3parserRULE_operacion)

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

	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opIzq = _x
		}
		{
			p.SetState(215)

			var _x = p.Operador()

			localctx.(*OperacionContext)._operador = _x
		}
		{
			p.SetState(216)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opDer = _x
		}

		localctx.(*OperacionContext).ex = expresiones3d.NewOperacion3D(localctx.(*OperacionContext).GetOpIzq().GetEx(), localctx.(*OperacionContext).GetOpDer().GetEx(), localctx.(*OperacionContext).Get_operador().GetEx(), false)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(220)
			p.Match(N3parserINT)
		}
		{
			p.SetState(221)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(222)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opIzq = _x
		}
		{
			p.SetState(223)

			var _x = p.Operador()

			localctx.(*OperacionContext)._operador = _x
		}
		{
			p.SetState(224)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(225)
			p.Match(N3parserINT)
		}
		{
			p.SetState(226)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(227)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opDer = _x
		}

		localctx.(*OperacionContext).ex = expresiones3d.NewOperacion3D(localctx.(*OperacionContext).GetOpIzq().GetEx(), localctx.(*OperacionContext).GetOpDer().GetEx(), localctx.(*OperacionContext).Get_operador().GetEx(), false)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(231)
			p.Match(N3parserINT)
		}
		{
			p.SetState(232)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(233)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opIzq = _x
		}
		{
			p.SetState(234)

			var _x = p.Operador()

			localctx.(*OperacionContext)._operador = _x
		}
		{
			p.SetState(235)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opDer = _x
		}

		localctx.(*OperacionContext).ex = expresiones3d.NewOperacion3D(localctx.(*OperacionContext).GetOpIzq().GetEx(), localctx.(*OperacionContext).GetOpDer().GetEx(), localctx.(*OperacionContext).Get_operador().GetEx(), false)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(238)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opIzq = _x
		}
		{
			p.SetState(239)

			var _x = p.Operador()

			localctx.(*OperacionContext)._operador = _x
		}
		{
			p.SetState(240)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(241)
			p.Match(N3parserINT)
		}
		{
			p.SetState(242)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(243)

			var _x = p.Expresion()

			localctx.(*OperacionContext).opDer = _x
		}

		localctx.(*OperacionContext).ex = expresiones3d.NewOperacion3D(localctx.(*OperacionContext).GetOpIzq().GetEx(), localctx.(*OperacionContext).GetOpDer().GetEx(), localctx.(*OperacionContext).Get_operador().GetEx(), false)

	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_P returns the _P token.
	Get_P() antlr.Token

	// Get_TEMPORAL returns the _TEMPORAL token.
	Get_TEMPORAL() antlr.Token

	// Get_H returns the _H token.
	Get_H() antlr.Token

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Set_P sets the _P token.
	Set_P(antlr.Token)

	// Set_TEMPORAL sets the _TEMPORAL token.
	Set_TEMPORAL(antlr.Token)

	// Set_H sets the _H token.
	Set_H(antlr.Token)

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() expresiones3d.Primitivo3D

	// SetEx sets the ex attribute.
	SetEx(expresiones3d.Primitivo3D)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	ex        expresiones3d.Primitivo3D
	_P        antlr.Token
	_TEMPORAL antlr.Token
	_H        antlr.Token
	_NUMERO   antlr.Token
	_DECIMAL  antlr.Token
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) Get_P() antlr.Token { return s._P }

func (s *ExpresionContext) Get_TEMPORAL() antlr.Token { return s._TEMPORAL }

func (s *ExpresionContext) Get_H() antlr.Token { return s._H }

func (s *ExpresionContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *ExpresionContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *ExpresionContext) Set_P(v antlr.Token) { s._P = v }

func (s *ExpresionContext) Set_TEMPORAL(v antlr.Token) { s._TEMPORAL = v }

func (s *ExpresionContext) Set_H(v antlr.Token) { s._H = v }

func (s *ExpresionContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *ExpresionContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *ExpresionContext) GetEx() expresiones3d.Primitivo3D { return s.ex }

func (s *ExpresionContext) SetEx(v expresiones3d.Primitivo3D) { s.ex = v }

func (s *ExpresionContext) P() antlr.TerminalNode {
	return s.GetToken(N3parserP, 0)
}

func (s *ExpresionContext) TEMPORAL() antlr.TerminalNode {
	return s.GetToken(N3parserTEMPORAL, 0)
}

func (s *ExpresionContext) H() antlr.TerminalNode {
	return s.GetToken(N3parserH, 0)
}

func (s *ExpresionContext) RESTA() antlr.TerminalNode {
	return s.GetToken(N3parserRESTA, 0)
}

func (s *ExpresionContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(N3parserNUMERO, 0)
}

func (s *ExpresionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(N3parserDECIMAL, 0)
}

func (s *ExpresionContext) HEAP() antlr.TerminalNode {
	return s.GetToken(N3parserHEAP, 0)
}

func (s *ExpresionContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(N3parserCORCHETE_IZQ, 0)
}

func (s *ExpresionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(N3parserPAR_IZQ, 0)
}

func (s *ExpresionContext) INT() antlr.TerminalNode {
	return s.GetToken(N3parserINT, 0)
}

func (s *ExpresionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(N3parserPAR_DER, 0)
}

func (s *ExpresionContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(N3parserCORCHETE_DER, 0)
}

func (s *ExpresionContext) STACK() antlr.TerminalNode {
	return s.GetToken(N3parserSTACK, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (p *N3parser) Expresion() (localctx IExpresionContext) {
	this := p
	_ = this

	localctx = NewExpresionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, N3parserRULE_expresion)

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

	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)

			var _m = p.Match(N3parserP)

			localctx.(*ExpresionContext)._P = _m
		}

		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D((func() string {
			if localctx.(*ExpresionContext).Get_P() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_P().GetText()
			}
		}()), interfaces3d.PUNTERO_STACK)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)

			var _m = p.Match(N3parserTEMPORAL)

			localctx.(*ExpresionContext)._TEMPORAL = _m
		}

		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D((func() string {
			if localctx.(*ExpresionContext).Get_TEMPORAL() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_TEMPORAL().GetText()
			}
		}()), interfaces3d.PUNTERO_STACK)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(252)

			var _m = p.Match(N3parserH)

			localctx.(*ExpresionContext)._H = _m
		}

		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D((func() string {
			if localctx.(*ExpresionContext).Get_H() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_H().GetText()
			}
		}()), interfaces3d.PUNTERO_HEAP)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(254)
			p.Match(N3parserRESTA)
		}
		{
			p.SetState(255)

			var _m = p.Match(N3parserNUMERO)

			localctx.(*ExpresionContext)._NUMERO = _m
		}

		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D("-"+(func() string {
			if localctx.(*ExpresionContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_NUMERO().GetText()
			}
		}()), interfaces3d.NUMERO3D)

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(257)

			var _m = p.Match(N3parserNUMERO)

			localctx.(*ExpresionContext)._NUMERO = _m
		}

		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D((func() string {
			if localctx.(*ExpresionContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_NUMERO().GetText()
			}
		}()), interfaces3d.NUMERO3D)

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(259)

			var _m = p.Match(N3parserDECIMAL)

			localctx.(*ExpresionContext)._DECIMAL = _m
		}

		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D((func() string {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetText()
			}
		}()), interfaces3d.DECIMAL3D)

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(261)
			p.Match(N3parserHEAP)
		}
		{
			p.SetState(262)
			p.Match(N3parserCORCHETE_IZQ)
		}
		{
			p.SetState(263)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(264)
			p.Match(N3parserINT)
		}
		{
			p.SetState(265)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(266)

			var _m = p.Match(N3parserTEMPORAL)

			localctx.(*ExpresionContext)._TEMPORAL = _m
		}
		{
			p.SetState(267)
			p.Match(N3parserCORCHETE_DER)
		}

		valor := "heap [(int)" + (func() string {
			if localctx.(*ExpresionContext).Get_TEMPORAL() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_TEMPORAL().GetText()
			}
		}()) + "]"
		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D(valor, interfaces3d.ACCESO_HEAP)

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(269)
			p.Match(N3parserHEAP)
		}
		{
			p.SetState(270)
			p.Match(N3parserCORCHETE_IZQ)
		}
		{
			p.SetState(271)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(272)
			p.Match(N3parserINT)
		}
		{
			p.SetState(273)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(274)

			var _m = p.Match(N3parserH)

			localctx.(*ExpresionContext)._H = _m
		}
		{
			p.SetState(275)
			p.Match(N3parserCORCHETE_DER)
		}

		valor := "heap [(int)" + (func() string {
			if localctx.(*ExpresionContext).Get_H() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_H().GetText()
			}
		}()) + "]"
		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D(valor, interfaces3d.ACCESO_HEAP)

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(277)
			p.Match(N3parserSTACK)
		}
		{
			p.SetState(278)
			p.Match(N3parserCORCHETE_IZQ)
		}
		{
			p.SetState(279)
			p.Match(N3parserPAR_IZQ)
		}
		{
			p.SetState(280)
			p.Match(N3parserINT)
		}
		{
			p.SetState(281)
			p.Match(N3parserPAR_DER)
		}
		{
			p.SetState(282)

			var _m = p.Match(N3parserTEMPORAL)

			localctx.(*ExpresionContext)._TEMPORAL = _m
		}
		{
			p.SetState(283)
			p.Match(N3parserCORCHETE_DER)
		}

		valor := "stack [(int)" + (func() string {
			if localctx.(*ExpresionContext).Get_TEMPORAL() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_TEMPORAL().GetText()
			}
		}()) + "]"
		localctx.(*ExpresionContext).ex = expresiones3d.NewPrimitivo3D(valor, interfaces3d.ACCESO_STACK)

	}

	return localctx
}

// IEtiquetaContext is an interface to support dynamic dispatch.
type IEtiquetaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LABEL returns the _LABEL token.
	Get_LABEL() antlr.Token

	// Set_LABEL sets the _LABEL token.
	Set_LABEL(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() instrucciones3d.Salto3D

	// SetEx sets the ex attribute.
	SetEx(instrucciones3d.Salto3D)

	// IsEtiquetaContext differentiates from other interfaces.
	IsEtiquetaContext()
}

type EtiquetaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ex     instrucciones3d.Salto3D
	_LABEL antlr.Token
}

func NewEmptyEtiquetaContext() *EtiquetaContext {
	var p = new(EtiquetaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_etiqueta
	return p
}

func (*EtiquetaContext) IsEtiquetaContext() {}

func NewEtiquetaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EtiquetaContext {
	var p = new(EtiquetaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_etiqueta

	return p
}

func (s *EtiquetaContext) GetParser() antlr.Parser { return s.parser }

func (s *EtiquetaContext) Get_LABEL() antlr.Token { return s._LABEL }

func (s *EtiquetaContext) Set_LABEL(v antlr.Token) { s._LABEL = v }

func (s *EtiquetaContext) GetEx() instrucciones3d.Salto3D { return s.ex }

func (s *EtiquetaContext) SetEx(v instrucciones3d.Salto3D) { s.ex = v }

func (s *EtiquetaContext) LABEL() antlr.TerminalNode {
	return s.GetToken(N3parserLABEL, 0)
}

func (s *EtiquetaContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(N3parserDOSPUNTOS, 0)
}

func (s *EtiquetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EtiquetaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EtiquetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterEtiqueta(s)
	}
}

func (s *EtiquetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitEtiqueta(s)
	}
}

func (p *N3parser) Etiqueta() (localctx IEtiquetaContext) {
	this := p
	_ = this

	localctx = NewEtiquetaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, N3parserRULE_etiqueta)

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
		p.SetState(287)

		var _m = p.Match(N3parserLABEL)

		localctx.(*EtiquetaContext)._LABEL = _m
	}
	{
		p.SetState(288)
		p.Match(N3parserDOSPUNTOS)
	}

	localctx.(*EtiquetaContext).ex = instrucciones3d.NewSalto3D((func() string {
		if localctx.(*EtiquetaContext).Get_LABEL() == nil {
			return ""
		} else {
			return localctx.(*EtiquetaContext).Get_LABEL().GetText()
		}
	}()))

	return localctx
}

// ISaltoContext is an interface to support dynamic dispatch.
type ISaltoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LABEL returns the _LABEL token.
	Get_LABEL() antlr.Token

	// Set_LABEL sets the _LABEL token.
	Set_LABEL(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() instrucciones3d.Goto3D

	// SetEx sets the ex attribute.
	SetEx(instrucciones3d.Goto3D)

	// IsSaltoContext differentiates from other interfaces.
	IsSaltoContext()
}

type SaltoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ex     instrucciones3d.Goto3D
	_LABEL antlr.Token
}

func NewEmptySaltoContext() *SaltoContext {
	var p = new(SaltoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_salto
	return p
}

func (*SaltoContext) IsSaltoContext() {}

func NewSaltoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SaltoContext {
	var p = new(SaltoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_salto

	return p
}

func (s *SaltoContext) GetParser() antlr.Parser { return s.parser }

func (s *SaltoContext) Get_LABEL() antlr.Token { return s._LABEL }

func (s *SaltoContext) Set_LABEL(v antlr.Token) { s._LABEL = v }

func (s *SaltoContext) GetEx() instrucciones3d.Goto3D { return s.ex }

func (s *SaltoContext) SetEx(v instrucciones3d.Goto3D) { s.ex = v }

func (s *SaltoContext) GOTO() antlr.TerminalNode {
	return s.GetToken(N3parserGOTO, 0)
}

func (s *SaltoContext) LABEL() antlr.TerminalNode {
	return s.GetToken(N3parserLABEL, 0)
}

func (s *SaltoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SaltoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SaltoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterSalto(s)
	}
}

func (s *SaltoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitSalto(s)
	}
}

func (p *N3parser) Salto() (localctx ISaltoContext) {
	this := p
	_ = this

	localctx = NewSaltoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, N3parserRULE_salto)

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
		p.SetState(291)
		p.Match(N3parserGOTO)
	}
	{
		p.SetState(292)

		var _m = p.Match(N3parserLABEL)

		localctx.(*SaltoContext)._LABEL = _m
	}

	localctx.(*SaltoContext).ex = instrucciones3d.NewGoto3D("goto " + (func() string {
		if localctx.(*SaltoContext).Get_LABEL() == nil {
			return ""
		} else {
			return localctx.(*SaltoContext).Get_LABEL().GetText()
		}
	}()))

	return localctx
}

// IRetornoContext is an interface to support dynamic dispatch.
type IRetornoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() instrucciones3d.Return3D

	// SetEx sets the ex attribute.
	SetEx(instrucciones3d.Return3D)

	// IsRetornoContext differentiates from other interfaces.
	IsRetornoContext()
}

type RetornoContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	ex      instrucciones3d.Return3D
	_RETURN antlr.Token
	_NUMERO antlr.Token
}

func NewEmptyRetornoContext() *RetornoContext {
	var p = new(RetornoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_retorno
	return p
}

func (*RetornoContext) IsRetornoContext() {}

func NewRetornoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetornoContext {
	var p = new(RetornoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_retorno

	return p
}

func (s *RetornoContext) GetParser() antlr.Parser { return s.parser }

func (s *RetornoContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *RetornoContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *RetornoContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *RetornoContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *RetornoContext) GetEx() instrucciones3d.Return3D { return s.ex }

func (s *RetornoContext) SetEx(v instrucciones3d.Return3D) { s.ex = v }

func (s *RetornoContext) RETURN() antlr.TerminalNode {
	return s.GetToken(N3parserRETURN, 0)
}

func (s *RetornoContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(N3parserNUMERO, 0)
}

func (s *RetornoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetornoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetornoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterRetorno(s)
	}
}

func (s *RetornoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitRetorno(s)
	}
}

func (p *N3parser) Retorno() (localctx IRetornoContext) {
	this := p
	_ = this

	localctx = NewRetornoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, N3parserRULE_retorno)

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

	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(295)

			var _m = p.Match(N3parserRETURN)

			localctx.(*RetornoContext)._RETURN = _m
		}

		localctx.(*RetornoContext).ex = instrucciones3d.NewReturn3D((func() string {
			if localctx.(*RetornoContext).Get_RETURN() == nil {
				return ""
			} else {
				return localctx.(*RetornoContext).Get_RETURN().GetText()
			}
		}()))

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(297)

			var _m = p.Match(N3parserRETURN)

			localctx.(*RetornoContext)._RETURN = _m
		}
		{
			p.SetState(298)

			var _m = p.Match(N3parserNUMERO)

			localctx.(*RetornoContext)._NUMERO = _m
		}

		localctx.(*RetornoContext).ex = instrucciones3d.NewReturn3D((func() string {
			if localctx.(*RetornoContext).Get_RETURN() == nil {
				return ""
			} else {
				return localctx.(*RetornoContext).Get_RETURN().GetText()
			}
		}()) + " " + (func() string {
			if localctx.(*RetornoContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*RetornoContext).Get_NUMERO().GetText()
			}
		}()))

	}

	return localctx
}

// IIncludeContext is an interface to support dynamic dispatch.
type IIncludeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_INCLUDE returns the _INCLUDE token.
	Get_INCLUDE() antlr.Token

	// Get_STDIO returns the _STDIO token.
	Get_STDIO() antlr.Token

	// Set_INCLUDE sets the _INCLUDE token.
	Set_INCLUDE(antlr.Token)

	// Set_STDIO sets the _STDIO token.
	Set_STDIO(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() headers3d.Include3D

	// SetEx sets the ex attribute.
	SetEx(headers3d.Include3D)

	// IsIncludeContext differentiates from other interfaces.
	IsIncludeContext()
}

type IncludeContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	ex       headers3d.Include3D
	_INCLUDE antlr.Token
	_STDIO   antlr.Token
}

func NewEmptyIncludeContext() *IncludeContext {
	var p = new(IncludeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_include
	return p
}

func (*IncludeContext) IsIncludeContext() {}

func NewIncludeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncludeContext {
	var p = new(IncludeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_include

	return p
}

func (s *IncludeContext) GetParser() antlr.Parser { return s.parser }

func (s *IncludeContext) Get_INCLUDE() antlr.Token { return s._INCLUDE }

func (s *IncludeContext) Get_STDIO() antlr.Token { return s._STDIO }

func (s *IncludeContext) Set_INCLUDE(v antlr.Token) { s._INCLUDE = v }

func (s *IncludeContext) Set_STDIO(v antlr.Token) { s._STDIO = v }

func (s *IncludeContext) GetEx() headers3d.Include3D { return s.ex }

func (s *IncludeContext) SetEx(v headers3d.Include3D) { s.ex = v }

func (s *IncludeContext) INCLUDE() antlr.TerminalNode {
	return s.GetToken(N3parserINCLUDE, 0)
}

func (s *IncludeContext) STDIO() antlr.TerminalNode {
	return s.GetToken(N3parserSTDIO, 0)
}

func (s *IncludeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncludeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncludeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterInclude(s)
	}
}

func (s *IncludeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitInclude(s)
	}
}

func (p *N3parser) Include() (localctx IIncludeContext) {
	this := p
	_ = this

	localctx = NewIncludeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, N3parserRULE_include)

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
		p.SetState(302)

		var _m = p.Match(N3parserINCLUDE)

		localctx.(*IncludeContext)._INCLUDE = _m
	}
	{
		p.SetState(303)

		var _m = p.Match(N3parserSTDIO)

		localctx.(*IncludeContext)._STDIO = _m
	}

	localctx.(*IncludeContext).ex = headers3d.NewInclude3D((func() string {
		if localctx.(*IncludeContext).Get_INCLUDE() == nil {
			return ""
		} else {
			return localctx.(*IncludeContext).Get_INCLUDE().GetText()
		}
	}()) + " " + (func() string {
		if localctx.(*IncludeContext).Get_STDIO() == nil {
			return ""
		} else {
			return localctx.(*IncludeContext).Get_STDIO().GetText()
		}
	}()))

	return localctx
}

// IDeclaracionesContext is an interface to support dynamic dispatch.
type IDeclaracionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// GetLista returns the lista rule context list.
	GetLista() []IDeclaracionContext

	// SetLista sets the lista rule context list.
	SetLista([]IDeclaracionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsDeclaracionesContext differentiates from other interfaces.
	IsDeclaracionesContext()
}

type DeclaracionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	list         *arraylist.List
	_declaracion IDeclaracionContext
	lista        []IDeclaracionContext
}

func NewEmptyDeclaracionesContext() *DeclaracionesContext {
	var p = new(DeclaracionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_declaraciones
	return p
}

func (*DeclaracionesContext) IsDeclaracionesContext() {}

func NewDeclaracionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionesContext {
	var p = new(DeclaracionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_declaraciones

	return p
}

func (s *DeclaracionesContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionesContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *DeclaracionesContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *DeclaracionesContext) GetLista() []IDeclaracionContext { return s.lista }

func (s *DeclaracionesContext) SetLista(v []IDeclaracionContext) { s.lista = v }

func (s *DeclaracionesContext) GetList() *arraylist.List { return s.list }

func (s *DeclaracionesContext) SetList(v *arraylist.List) { s.list = v }

func (s *DeclaracionesContext) AllDeclaracion() []IDeclaracionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem())
	var tst = make([]IDeclaracionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclaracionContext)
		}
	}

	return tst
}

func (s *DeclaracionesContext) Declaracion(i int) IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *DeclaracionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterDeclaraciones(s)
	}
}

func (s *DeclaracionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitDeclaraciones(s)
	}
}

func (p *N3parser) Declaraciones() (localctx IDeclaracionesContext) {
	this := p
	_ = this

	localctx = NewDeclaracionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, N3parserRULE_declaraciones)
	localctx.(*DeclaracionesContext).list = arraylist.New()
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
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == N3parserFLOAT {
		{
			p.SetState(306)

			var _x = p.Declaracion()

			localctx.(*DeclaracionesContext)._declaracion = _x
		}
		localctx.(*DeclaracionesContext).lista = append(localctx.(*DeclaracionesContext).lista, localctx.(*DeclaracionesContext)._declaracion)

		p.SetState(309)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listas := localctx.(*DeclaracionesContext).GetLista()
	for _, e := range listas {
		localctx.(*DeclaracionesContext).list.Add(e.GetEx())
	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Get_temporalesTexto returns the _temporalesTexto rule contexts.
	Get_temporalesTexto() ITemporalesTextoContext

	// Set_temporalesTexto sets the _temporalesTexto rule contexts.
	Set_temporalesTexto(ITemporalesTextoContext)

	// GetEx returns the ex attribute.
	GetEx() headers3d.Declaracion3D

	// SetEx sets the ex attribute.
	SetEx(headers3d.Declaracion3D)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	ex               headers3d.Declaracion3D
	_NUMERO          antlr.Token
	_temporalesTexto ITemporalesTextoContext
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *DeclaracionContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *DeclaracionContext) Get_temporalesTexto() ITemporalesTextoContext { return s._temporalesTexto }

func (s *DeclaracionContext) Set_temporalesTexto(v ITemporalesTextoContext) { s._temporalesTexto = v }

func (s *DeclaracionContext) GetEx() headers3d.Declaracion3D { return s.ex }

func (s *DeclaracionContext) SetEx(v headers3d.Declaracion3D) { s.ex = v }

func (s *DeclaracionContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(N3parserFLOAT, 0)
}

func (s *DeclaracionContext) STACK() antlr.TerminalNode {
	return s.GetToken(N3parserSTACK, 0)
}

func (s *DeclaracionContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(N3parserCORCHETE_IZQ, 0)
}

func (s *DeclaracionContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(N3parserNUMERO, 0)
}

func (s *DeclaracionContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(N3parserCORCHETE_DER, 0)
}

func (s *DeclaracionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(N3parserPUNTOCOMA, 0)
}

func (s *DeclaracionContext) HEAP() antlr.TerminalNode {
	return s.GetToken(N3parserHEAP, 0)
}

func (s *DeclaracionContext) P() antlr.TerminalNode {
	return s.GetToken(N3parserP, 0)
}

func (s *DeclaracionContext) H() antlr.TerminalNode {
	return s.GetToken(N3parserH, 0)
}

func (s *DeclaracionContext) TemporalesTexto() ITemporalesTextoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalesTextoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemporalesTextoContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *N3parser) Declaracion() (localctx IDeclaracionContext) {
	this := p
	_ = this

	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, N3parserRULE_declaracion)

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

	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.Match(N3parserFLOAT)
		}
		{
			p.SetState(314)
			p.Match(N3parserSTACK)
		}
		{
			p.SetState(315)
			p.Match(N3parserCORCHETE_IZQ)
		}
		{
			p.SetState(316)

			var _m = p.Match(N3parserNUMERO)

			localctx.(*DeclaracionContext)._NUMERO = _m
		}
		{
			p.SetState(317)
			p.Match(N3parserCORCHETE_DER)
		}
		{
			p.SetState(318)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*DeclaracionContext).ex = headers3d.NewDeclaracion3D("float stack[" + (func() string {
			if localctx.(*DeclaracionContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_NUMERO().GetText()
			}
		}()) + "]")

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(N3parserFLOAT)
		}
		{
			p.SetState(321)
			p.Match(N3parserHEAP)
		}
		{
			p.SetState(322)
			p.Match(N3parserCORCHETE_IZQ)
		}
		{
			p.SetState(323)

			var _m = p.Match(N3parserNUMERO)

			localctx.(*DeclaracionContext)._NUMERO = _m
		}
		{
			p.SetState(324)
			p.Match(N3parserCORCHETE_DER)
		}
		{
			p.SetState(325)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*DeclaracionContext).ex = headers3d.NewDeclaracion3D("float heap[" + (func() string {
			if localctx.(*DeclaracionContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_NUMERO().GetText()
			}
		}()) + "]")

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(327)
			p.Match(N3parserFLOAT)
		}
		{
			p.SetState(328)
			p.Match(N3parserP)
		}
		{
			p.SetState(329)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*DeclaracionContext).ex = headers3d.NewDeclaracion3D("float P")

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(331)
			p.Match(N3parserFLOAT)
		}
		{
			p.SetState(332)
			p.Match(N3parserH)
		}
		{
			p.SetState(333)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*DeclaracionContext).ex = headers3d.NewDeclaracion3D("float H")

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(335)
			p.Match(N3parserFLOAT)
		}
		{
			p.SetState(336)

			var _x = p.TemporalesTexto()

			localctx.(*DeclaracionContext)._temporalesTexto = _x
		}
		{
			p.SetState(337)
			p.Match(N3parserPUNTOCOMA)
		}

		localctx.(*DeclaracionContext).ex = headers3d.NewDeclaracion3D("float " + localctx.(*DeclaracionContext).Get_temporalesTexto().GetEx())

	}

	return localctx
}

// ITemporalesTextoContext is an interface to support dynamic dispatch.
type ITemporalesTextoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_temporalesLista returns the _temporalesLista rule contexts.
	Get_temporalesLista() ITemporalesListaContext

	// Set_temporalesLista sets the _temporalesLista rule contexts.
	Set_temporalesLista(ITemporalesListaContext)

	// GetEx returns the ex attribute.
	GetEx() string

	// SetEx sets the ex attribute.
	SetEx(string)

	// IsTemporalesTextoContext differentiates from other interfaces.
	IsTemporalesTextoContext()
}

type TemporalesTextoContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	ex               string
	_temporalesLista ITemporalesListaContext
}

func NewEmptyTemporalesTextoContext() *TemporalesTextoContext {
	var p = new(TemporalesTextoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_temporalesTexto
	return p
}

func (*TemporalesTextoContext) IsTemporalesTextoContext() {}

func NewTemporalesTextoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalesTextoContext {
	var p = new(TemporalesTextoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_temporalesTexto

	return p
}

func (s *TemporalesTextoContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalesTextoContext) Get_temporalesLista() ITemporalesListaContext {
	return s._temporalesLista
}

func (s *TemporalesTextoContext) Set_temporalesLista(v ITemporalesListaContext) {
	s._temporalesLista = v
}

func (s *TemporalesTextoContext) GetEx() string { return s.ex }

func (s *TemporalesTextoContext) SetEx(v string) { s.ex = v }

func (s *TemporalesTextoContext) TemporalesLista() ITemporalesListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalesListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemporalesListaContext)
}

func (s *TemporalesTextoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalesTextoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalesTextoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterTemporalesTexto(s)
	}
}

func (s *TemporalesTextoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitTemporalesTexto(s)
	}
}

func (p *N3parser) TemporalesTexto() (localctx ITemporalesTextoContext) {
	this := p
	_ = this

	localctx = NewTemporalesTextoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, N3parserRULE_temporalesTexto)

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
		p.SetState(342)

		var _x = p.temporalesLista(0)

		localctx.(*TemporalesTextoContext)._temporalesLista = _x
	}

	temporales := ""
	lista := localctx.(*TemporalesTextoContext).Get_temporalesLista().GetList()
	for i := 0; i < lista.Len(); i++ {
		if i != lista.Len()-1 {
			temporales += lista.GetValue(i).(string) + ","
		} else {
			temporales += lista.GetValue(i).(string)
		}

	}
	localctx.(*TemporalesTextoContext).ex = temporales

	return localctx
}

// ITemporalesListaContext is an interface to support dynamic dispatch.
type ITemporalesListaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTemp returns the temp token.
	GetTemp() antlr.Token

	// SetTemp sets the temp token.
	SetTemp(antlr.Token)

	// GetLista_elementos returns the lista_elementos rule contexts.
	GetLista_elementos() ITemporalesListaContext

	// SetLista_elementos sets the lista_elementos rule contexts.
	SetLista_elementos(ITemporalesListaContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsTemporalesListaContext differentiates from other interfaces.
	IsTemporalesListaContext()
}

type TemporalesListaContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	list            *arraylist.List
	lista_elementos ITemporalesListaContext
	temp            antlr.Token
}

func NewEmptyTemporalesListaContext() *TemporalesListaContext {
	var p = new(TemporalesListaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_temporalesLista
	return p
}

func (*TemporalesListaContext) IsTemporalesListaContext() {}

func NewTemporalesListaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalesListaContext {
	var p = new(TemporalesListaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_temporalesLista

	return p
}

func (s *TemporalesListaContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalesListaContext) GetTemp() antlr.Token { return s.temp }

func (s *TemporalesListaContext) SetTemp(v antlr.Token) { s.temp = v }

func (s *TemporalesListaContext) GetLista_elementos() ITemporalesListaContext {
	return s.lista_elementos
}

func (s *TemporalesListaContext) SetLista_elementos(v ITemporalesListaContext) { s.lista_elementos = v }

func (s *TemporalesListaContext) GetList() *arraylist.List { return s.list }

func (s *TemporalesListaContext) SetList(v *arraylist.List) { s.list = v }

func (s *TemporalesListaContext) TEMPORAL() antlr.TerminalNode {
	return s.GetToken(N3parserTEMPORAL, 0)
}

func (s *TemporalesListaContext) COMA() antlr.TerminalNode {
	return s.GetToken(N3parserCOMA, 0)
}

func (s *TemporalesListaContext) TemporalesLista() ITemporalesListaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalesListaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemporalesListaContext)
}

func (s *TemporalesListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalesListaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalesListaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterTemporalesLista(s)
	}
}

func (s *TemporalesListaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitTemporalesLista(s)
	}
}

func (p *N3parser) TemporalesLista() (localctx ITemporalesListaContext) {
	return p.temporalesLista(0)
}

func (p *N3parser) temporalesLista(_p int) (localctx ITemporalesListaContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTemporalesListaContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITemporalesListaContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, N3parserRULE_temporalesLista, _p)
	localctx.(*TemporalesListaContext).list = arraylist.New()

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)

		var _m = p.Match(N3parserTEMPORAL)

		localctx.(*TemporalesListaContext).temp = _m
	}

	localctx.(*TemporalesListaContext).list.Add((func() string {
		if localctx.(*TemporalesListaContext).GetTemp() == nil {
			return ""
		} else {
			return localctx.(*TemporalesListaContext).GetTemp().GetText()
		}
	}()))

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTemporalesListaContext(p, _parentctx, _parentState)
			localctx.(*TemporalesListaContext).lista_elementos = _prevctx
			p.PushNewRecursionContext(localctx, _startState, N3parserRULE_temporalesLista)
			p.SetState(349)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(350)
				p.Match(N3parserCOMA)
			}
			{
				p.SetState(351)

				var _m = p.Match(N3parserTEMPORAL)

				localctx.(*TemporalesListaContext).temp = _m
			}

			localctx.(*TemporalesListaContext).GetLista_elementos().GetList().Add((func() string {
				if localctx.(*TemporalesListaContext).GetTemp() == nil {
					return ""
				} else {
					return localctx.(*TemporalesListaContext).GetTemp().GetText()
				}
			}()))
			localctx.(*TemporalesListaContext).list = localctx.(*TemporalesListaContext).GetLista_elementos().GetList()

		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IOperadorContext is an interface to support dynamic dispatch.
type IOperadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MAYOR_I returns the _MAYOR_I token.
	Get_MAYOR_I() antlr.Token

	// Get_MAYOR returns the _MAYOR token.
	Get_MAYOR() antlr.Token

	// Get_MENOR_I returns the _MENOR_I token.
	Get_MENOR_I() antlr.Token

	// Get_MENOR returns the _MENOR token.
	Get_MENOR() antlr.Token

	// Get_IGUALDAD returns the _IGUALDAD token.
	Get_IGUALDAD() antlr.Token

	// Get_DISTINTO returns the _DISTINTO token.
	Get_DISTINTO() antlr.Token

	// Get_MODULO returns the _MODULO token.
	Get_MODULO() antlr.Token

	// Get_MULTIPLICACION returns the _MULTIPLICACION token.
	Get_MULTIPLICACION() antlr.Token

	// Get_DIVISION returns the _DIVISION token.
	Get_DIVISION() antlr.Token

	// Get_RESTA returns the _RESTA token.
	Get_RESTA() antlr.Token

	// Get_SUMA returns the _SUMA token.
	Get_SUMA() antlr.Token

	// Set_MAYOR_I sets the _MAYOR_I token.
	Set_MAYOR_I(antlr.Token)

	// Set_MAYOR sets the _MAYOR token.
	Set_MAYOR(antlr.Token)

	// Set_MENOR_I sets the _MENOR_I token.
	Set_MENOR_I(antlr.Token)

	// Set_MENOR sets the _MENOR token.
	Set_MENOR(antlr.Token)

	// Set_IGUALDAD sets the _IGUALDAD token.
	Set_IGUALDAD(antlr.Token)

	// Set_DISTINTO sets the _DISTINTO token.
	Set_DISTINTO(antlr.Token)

	// Set_MODULO sets the _MODULO token.
	Set_MODULO(antlr.Token)

	// Set_MULTIPLICACION sets the _MULTIPLICACION token.
	Set_MULTIPLICACION(antlr.Token)

	// Set_DIVISION sets the _DIVISION token.
	Set_DIVISION(antlr.Token)

	// Set_RESTA sets the _RESTA token.
	Set_RESTA(antlr.Token)

	// Set_SUMA sets the _SUMA token.
	Set_SUMA(antlr.Token)

	// GetEx returns the ex attribute.
	GetEx() string

	// SetEx sets the ex attribute.
	SetEx(string)

	// IsOperadorContext differentiates from other interfaces.
	IsOperadorContext()
}

type OperadorContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	ex              string
	_MAYOR_I        antlr.Token
	_MAYOR          antlr.Token
	_MENOR_I        antlr.Token
	_MENOR          antlr.Token
	_IGUALDAD       antlr.Token
	_DISTINTO       antlr.Token
	_MODULO         antlr.Token
	_MULTIPLICACION antlr.Token
	_DIVISION       antlr.Token
	_RESTA          antlr.Token
	_SUMA           antlr.Token
}

func NewEmptyOperadorContext() *OperadorContext {
	var p = new(OperadorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = N3parserRULE_operador
	return p
}

func (*OperadorContext) IsOperadorContext() {}

func NewOperadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperadorContext {
	var p = new(OperadorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = N3parserRULE_operador

	return p
}

func (s *OperadorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperadorContext) Get_MAYOR_I() antlr.Token { return s._MAYOR_I }

func (s *OperadorContext) Get_MAYOR() antlr.Token { return s._MAYOR }

func (s *OperadorContext) Get_MENOR_I() antlr.Token { return s._MENOR_I }

func (s *OperadorContext) Get_MENOR() antlr.Token { return s._MENOR }

func (s *OperadorContext) Get_IGUALDAD() antlr.Token { return s._IGUALDAD }

func (s *OperadorContext) Get_DISTINTO() antlr.Token { return s._DISTINTO }

func (s *OperadorContext) Get_MODULO() antlr.Token { return s._MODULO }

func (s *OperadorContext) Get_MULTIPLICACION() antlr.Token { return s._MULTIPLICACION }

func (s *OperadorContext) Get_DIVISION() antlr.Token { return s._DIVISION }

func (s *OperadorContext) Get_RESTA() antlr.Token { return s._RESTA }

func (s *OperadorContext) Get_SUMA() antlr.Token { return s._SUMA }

func (s *OperadorContext) Set_MAYOR_I(v antlr.Token) { s._MAYOR_I = v }

func (s *OperadorContext) Set_MAYOR(v antlr.Token) { s._MAYOR = v }

func (s *OperadorContext) Set_MENOR_I(v antlr.Token) { s._MENOR_I = v }

func (s *OperadorContext) Set_MENOR(v antlr.Token) { s._MENOR = v }

func (s *OperadorContext) Set_IGUALDAD(v antlr.Token) { s._IGUALDAD = v }

func (s *OperadorContext) Set_DISTINTO(v antlr.Token) { s._DISTINTO = v }

func (s *OperadorContext) Set_MODULO(v antlr.Token) { s._MODULO = v }

func (s *OperadorContext) Set_MULTIPLICACION(v antlr.Token) { s._MULTIPLICACION = v }

func (s *OperadorContext) Set_DIVISION(v antlr.Token) { s._DIVISION = v }

func (s *OperadorContext) Set_RESTA(v antlr.Token) { s._RESTA = v }

func (s *OperadorContext) Set_SUMA(v antlr.Token) { s._SUMA = v }

func (s *OperadorContext) GetEx() string { return s.ex }

func (s *OperadorContext) SetEx(v string) { s.ex = v }

func (s *OperadorContext) MAYOR_I() antlr.TerminalNode {
	return s.GetToken(N3parserMAYOR_I, 0)
}

func (s *OperadorContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(N3parserMAYOR, 0)
}

func (s *OperadorContext) MENOR_I() antlr.TerminalNode {
	return s.GetToken(N3parserMENOR_I, 0)
}

func (s *OperadorContext) MENOR() antlr.TerminalNode {
	return s.GetToken(N3parserMENOR, 0)
}

func (s *OperadorContext) IGUALDAD() antlr.TerminalNode {
	return s.GetToken(N3parserIGUALDAD, 0)
}

func (s *OperadorContext) DISTINTO() antlr.TerminalNode {
	return s.GetToken(N3parserDISTINTO, 0)
}

func (s *OperadorContext) MODULO() antlr.TerminalNode {
	return s.GetToken(N3parserMODULO, 0)
}

func (s *OperadorContext) MULTIPLICACION() antlr.TerminalNode {
	return s.GetToken(N3parserMULTIPLICACION, 0)
}

func (s *OperadorContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(N3parserDIVISION, 0)
}

func (s *OperadorContext) RESTA() antlr.TerminalNode {
	return s.GetToken(N3parserRESTA, 0)
}

func (s *OperadorContext) SUMA() antlr.TerminalNode {
	return s.GetToken(N3parserSUMA, 0)
}

func (s *OperadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.EnterOperador(s)
	}
}

func (s *OperadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(N3parserListener); ok {
		listenerT.ExitOperador(s)
	}
}

func (p *N3parser) Operador() (localctx IOperadorContext) {
	this := p
	_ = this

	localctx = NewOperadorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, N3parserRULE_operador)

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

	p.SetState(380)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case N3parserMAYOR_I:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)

			var _m = p.Match(N3parserMAYOR_I)

			localctx.(*OperadorContext)._MAYOR_I = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_MAYOR_I() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_MAYOR_I().GetText()
			}
		}())

	case N3parserMAYOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(360)

			var _m = p.Match(N3parserMAYOR)

			localctx.(*OperadorContext)._MAYOR = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_MAYOR() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_MAYOR().GetText()
			}
		}())

	case N3parserMENOR_I:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(362)

			var _m = p.Match(N3parserMENOR_I)

			localctx.(*OperadorContext)._MENOR_I = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_MENOR_I() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_MENOR_I().GetText()
			}
		}())

	case N3parserMENOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(364)

			var _m = p.Match(N3parserMENOR)

			localctx.(*OperadorContext)._MENOR = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_MENOR() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_MENOR().GetText()
			}
		}())

	case N3parserIGUALDAD:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(366)

			var _m = p.Match(N3parserIGUALDAD)

			localctx.(*OperadorContext)._IGUALDAD = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_IGUALDAD() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_IGUALDAD().GetText()
			}
		}())

	case N3parserDISTINTO:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(368)

			var _m = p.Match(N3parserDISTINTO)

			localctx.(*OperadorContext)._DISTINTO = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_DISTINTO() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_DISTINTO().GetText()
			}
		}())

	case N3parserMODULO:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(370)

			var _m = p.Match(N3parserMODULO)

			localctx.(*OperadorContext)._MODULO = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_MODULO() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_MODULO().GetText()
			}
		}())

	case N3parserMULTIPLICACION:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(372)

			var _m = p.Match(N3parserMULTIPLICACION)

			localctx.(*OperadorContext)._MULTIPLICACION = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_MULTIPLICACION() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_MULTIPLICACION().GetText()
			}
		}())

	case N3parserDIVISION:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(374)

			var _m = p.Match(N3parserDIVISION)

			localctx.(*OperadorContext)._DIVISION = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_DIVISION() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_DIVISION().GetText()
			}
		}())

	case N3parserRESTA:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(376)

			var _m = p.Match(N3parserRESTA)

			localctx.(*OperadorContext)._RESTA = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_RESTA() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_RESTA().GetText()
			}
		}())

	case N3parserSUMA:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(378)

			var _m = p.Match(N3parserSUMA)

			localctx.(*OperadorContext)._SUMA = _m
		}

		localctx.(*OperadorContext).ex = (func() string {
			if localctx.(*OperadorContext).Get_SUMA() == nil {
				return ""
			} else {
				return localctx.(*OperadorContext).Get_SUMA().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *N3parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *TemporalesListaContext = nil
		if localctx != nil {
			t = localctx.(*TemporalesListaContext)
		}
		return p.TemporalesLista_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *N3parser) TemporalesLista_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
