// Code generated by goyacc -l -v grammar.out -o grammar.go grammar.y. DO NOT EDIT.
package parser

import __yyfmt__ "fmt"

import (
	"fmt"

	"github.com/basecomplextech/spec/internal/lang/ast"
)

type yySymType struct {
	yys int
	// Tokens
	ident   string
	integer int
	string  string

	// Type
	type_ *ast.Type

	// Import
	import_ *ast.Import
	imports []*ast.Import

	// Option
	option  *ast.Option
	options []*ast.Option

	// Definition
	definition  *ast.Definition
	definitions []*ast.Definition

	// Enum
	enum_value  *ast.EnumValue
	enum_values []*ast.EnumValue

	// Message
	message_field  *ast.MessageField
	message_fields []*ast.MessageField

	// Struct
	struct_field  *ast.StructField
	struct_fields []*ast.StructField

	// Service
	service        *ast.Service
	method         *ast.Method
	methods        []*ast.Method
	method_result  *ast.MethodResult
	method_results []*ast.MethodResult
	method_field   *ast.MethodField
	method_fields  []*ast.MethodField
}

const ANY = 57346
const ENUM = 57347
const IMPORT = 57348
const MESSAGE = 57349
const OPTIONS = 57350
const STRUCT = 57351
const SERVICE = 57352
const SUBSERVICE = 57353
const IDENT = 57354
const INTEGER = 57355
const STRING = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ANY",
	"ENUM",
	"IMPORT",
	"MESSAGE",
	"OPTIONS",
	"STRUCT",
	"SERVICE",
	"SUBSERVICE",
	"IDENT",
	"INTEGER",
	"STRING",
	"'('",
	"')'",
	"'='",
	"'['",
	"']'",
	"'.'",
	"'{'",
	"'}'",
	"';'",
	"','",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 146

var yyAct = [...]int8{
	44, 97, 92, 63, 43, 64, 55, 47, 100, 48,
	49, 50, 51, 52, 53, 45, 47, 95, 48, 49,
	50, 51, 52, 53, 45, 75, 80, 94, 88, 62,
	77, 39, 38, 47, 72, 48, 49, 50, 51, 52,
	53, 45, 60, 37, 102, 67, 56, 36, 68, 76,
	81, 69, 103, 66, 35, 71, 74, 74, 47, 65,
	48, 49, 50, 51, 52, 53, 45, 78, 25, 40,
	24, 105, 23, 33, 90, 82, 58, 31, 84, 8,
	6, 67, 57, 34, 68, 93, 86, 107, 16, 66,
	17, 98, 18, 19, 20, 85, 87, 101, 79, 30,
	29, 93, 104, 106, 98, 108, 47, 28, 48, 49,
	50, 51, 52, 53, 45, 27, 26, 5, 3, 99,
	61, 1, 91, 96, 89, 83, 73, 15, 14, 54,
	70, 13, 42, 12, 41, 59, 11, 7, 10, 4,
	21, 32, 2, 9, 22, 46,
}

var yyPact = [...]int16{
	112, -1000, 109, 65, -1000, 64, -1000, 83, -1000, 56,
	-1000, -1000, -1000, -1000, -1000, -1000, 104, 103, 95, 88,
	87, 61, -1000, -1000, -1000, 69, 33, 26, 22, 11,
	10, -1000, -1000, 52, -1000, -1000, 102, -1000, -1000, -1000,
	68, 54, 6, -1000, 41, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 29, 12, 3, -1000, -1000, -1000,
	32, 8, 102, 85, -1000, 7, 30, -1000, -1000, -1000,
	-1000, 41, -1000, -1000, 63, -1000, 82, -1000, -1000, -1000,
	77, 84, 5, 59, 102, 4, -1000, -1000, -1000, -6,
	102, -16, -1000, 41, -1000, -1000, 28, -1000, 41, 55,
	102, 74, -1000, 102, -1000, -1000, -1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 145, 0, 144, 143, 142, 141, 140, 139, 3,
	5, 138, 137, 136, 135, 134, 133, 4, 132, 131,
	130, 129, 128, 127, 6, 126, 125, 124, 123, 1,
	2, 122, 121, 120, 119,
}

var yyR1 = [...]int8{
	0, 2, 2, 1, 1, 1, 1, 1, 1, 1,
	32, 3, 3, 4, 4, 5, 5, 8, 8, 7,
	7, 6, 9, 9, 10, 10, 10, 10, 11, 11,
	11, 11, 11, 12, 12, 13, 14, 15, 15, 16,
	17, 18, 18, 18, 19, 20, 21, 21, 22, 23,
	24, 24, 25, 26, 31, 31, 31, 30, 27, 27,
	28, 28, 28, 29, 34, 34, 33, 33,
}

var yyR2 = [...]int8{
	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 1, 2, 0, 2, 0, 4, 0, 4, 0,
	2, 3, 1, 3, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 0, 2, 5, 4, 0, 2, 6,
	3, 0, 1, 3, 5, 3, 0, 2, 5, 5,
	0, 2, 4, 4, 0, 1, 3, 3, 0, 3,
	0, 1, 3, 2, 0, 1, 0, 1,
}

var yyChk = [...]int16{
	-1000, -32, -5, 6, -8, 8, 15, -12, 15, -4,
	-11, -13, -16, -19, -22, -23, 5, 7, 9, 10,
	11, -7, -3, 16, 14, 12, 12, 12, 12, 12,
	12, 16, -6, 12, 14, 21, 21, 21, 21, 21,
	17, -15, -18, -17, -2, 12, -1, 4, 6, 7,
	8, 9, 10, 11, -21, -24, -24, 14, 22, -14,
	-2, -33, 23, -9, -10, 18, 12, 4, 7, 22,
	-20, -2, 22, -25, -2, 22, 17, 22, -17, 13,
	19, 20, -9, -26, 15, 13, -10, 12, 23, -27,
	15, -31, -30, -2, 23, 23, -28, -29, -2, -34,
	24, -9, 16, 24, -9, 16, -30, 13, -29,
}

var yyDef = [...]int8{
	15, -2, 17, 0, 33, 0, 13, 10, 19, 0,
	34, 28, 29, 30, 31, 32, 0, 0, 0, 0,
	0, 0, 14, 16, 11, 0, 0, 0, 0, 0,
	0, 18, 20, 0, 12, 37, 41, 46, 50, 50,
	0, 0, 66, 42, 0, 1, 2, 3, 4, 5,
	6, 7, 8, 9, 0, 0, 0, 21, 35, 38,
	0, 0, 67, 0, 22, 0, 24, 26, 27, 44,
	47, 0, 48, 51, 0, 49, 0, 39, 43, 40,
	0, 0, 0, 58, 54, 0, 23, 25, 45, 0,
	60, 64, 55, 0, 36, 52, 0, 61, 0, 0,
	65, 0, 59, 0, 63, 53, 56, 57, 62,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	15, 16, 3, 3, 24, 3, 20, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 23,
	3, 17, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 18, 3, 19, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 21, 3, 22,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = yyDollar[1].ident
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = yyDollar[1].ident
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = "any"
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = "import"
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = "message"
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = "options"
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = "struct"
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = "service"
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.ident = "subservice"
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			file := &ast.File{
				Imports:     yyDollar[1].imports,
				Options:     yyDollar[2].options,
				Definitions: yyDollar[3].definitions,
			}
			setLexerResult(yylex, file)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("import ", yyDollar[1].string)
			}
			yyVAL.import_ = &ast.Import{
				ID: trimString(yyDollar[1].string),
			}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("import ", yyDollar[1].ident, yyDollar[2].string)
			}
			yyVAL.import_ = &ast.Import{
				Alias: yyDollar[1].ident,
				ID:    trimString(yyDollar[2].string),
			}
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.imports = nil
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("import_list", yyDollar[1].imports, yyDollar[2].import_)
			}
			yyVAL.imports = append(yyVAL.imports, yyDollar[2].import_)
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.imports = nil
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("imports", yyDollar[3].imports)
			}
			yyVAL.imports = append(yyVAL.imports, yyDollar[3].imports...)
		}
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.options = nil
		}
	case 18:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("options", yyDollar[3].options)
			}
			yyVAL.options = append(yyVAL.options, yyDollar[3].options...)
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.options = nil
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("option_list", yyDollar[1].options, yyDollar[2].option)
			}
			yyVAL.options = append(yyVAL.options, yyDollar[2].option)
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("option ", yyDollar[1].ident, yyDollar[3].string)
			}
			yyVAL.option = &ast.Option{
				Name:  yyDollar[1].ident,
				Value: trimString(yyDollar[3].string),
			}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Printf("type *%v\n", yyDollar[1].type_)
			}
			yyVAL.type_ = yyDollar[1].type_
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Printf("type []%v\n", yyDollar[3].type_)
			}
			yyVAL.type_ = &ast.Type{
				Kind:    ast.KindList,
				Element: yyDollar[3].type_,
			}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("base type", yyDollar[1].ident)
			}
			yyVAL.type_ = &ast.Type{
				Kind: ast.GetKind(yyDollar[1].ident),
				Name: yyDollar[1].ident,
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Printf("base type %v.%v\n", yyDollar[1].ident, yyDollar[3].ident)
			}
			yyVAL.type_ = &ast.Type{
				Kind:   ast.KindReference,
				Name:   yyDollar[3].ident,
				Import: yyDollar[1].ident,
			}
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("base type", "any")
			}
			yyVAL.type_ = &ast.Type{
				Kind: ast.KindAny,
				Name: "any",
			}
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("base type", "message")
			}
			yyVAL.type_ = &ast.Type{
				Kind: ast.KindAnyMessage,
				Name: "message",
			}
		}
	case 33:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.definitions = nil
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("definitions", yyDollar[1].definitions, yyDollar[2].definition)
			}
			yyVAL.definitions = append(yyVAL.definitions, yyDollar[2].definition)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if debugParser {
				fmt.Println("enum", yyDollar[2].ident, yyDollar[4].enum_values)
			}
			yyVAL.definition = &ast.Definition{
				Type: ast.DefinitionEnum,
				Name: yyDollar[2].ident,

				Enum: &ast.Enum{
					Values: yyDollar[4].enum_values,
				},
			}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("enum value", yyDollar[1].ident, yyDollar[3].integer)
			}
			yyVAL.enum_value = &ast.EnumValue{
				Name:  yyDollar[1].ident,
				Value: yyDollar[3].integer,
			}
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.enum_values = nil
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("enum values", yyDollar[1].enum_values, yyDollar[2].enum_value)
			}
			yyVAL.enum_values = append(yyVAL.enum_values, yyDollar[2].enum_value)
		}
	case 39:
		yyDollar = yyS[yypt-6 : yypt+1]
		{
			if debugParser {
				fmt.Println("message", yyDollar[2].ident, yyDollar[4].message_fields)
			}
			yyVAL.definition = &ast.Definition{
				Type: ast.DefinitionMessage,
				Name: yyDollar[2].ident,

				Message: &ast.Message{
					Fields: yyDollar[4].message_fields,
				},
			}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("message field", yyDollar[1].ident, yyDollar[2].type_, yyDollar[3].integer)
			}
			yyVAL.message_field = &ast.MessageField{
				Name: yyDollar[1].ident,
				Type: yyDollar[2].type_,
				Tag:  yyDollar[3].integer,
			}
		}
	case 41:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.message_fields = nil
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("message fields", yyDollar[1].message_field)
			}
			yyVAL.message_fields = []*ast.MessageField{yyDollar[1].message_field}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("message fields", yyDollar[1].message_fields, yyDollar[3].message_field)
			}
			yyVAL.message_fields = append(yyVAL.message_fields, yyDollar[3].message_field)
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if debugParser {
				fmt.Println("struct", yyDollar[2].ident, yyDollar[4].struct_fields)
			}
			yyVAL.definition = &ast.Definition{
				Type: ast.DefinitionStruct,
				Name: yyDollar[2].ident,

				Struct: &ast.Struct{
					Fields: yyDollar[4].struct_fields,
				},
			}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("struct field", yyDollar[1].ident, yyDollar[2].type_)
			}
			yyVAL.struct_field = &ast.StructField{
				Name: yyDollar[1].ident,
				Type: yyDollar[2].type_,
			}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.struct_fields = nil
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("struct fields", yyDollar[1].struct_fields, yyDollar[2].struct_field)
			}
			yyVAL.struct_fields = append(yyVAL.struct_fields, yyDollar[2].struct_field)
		}
	case 48:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if debugParser {
				fmt.Println("service", yyDollar[2].ident, yyDollar[4].methods)
			}
			yyVAL.definition = &ast.Definition{
				Type: ast.DefinitionService,
				Name: yyDollar[2].ident,

				Service: &ast.Service{
					Methods: yyDollar[4].methods,
				},
			}
		}
	case 49:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if debugParser {
				fmt.Println("subservice", yyDollar[2].ident, yyDollar[4].methods)
			}
			yyVAL.definition = &ast.Definition{
				Type: ast.DefinitionService,
				Name: yyDollar[2].ident,

				Service: &ast.Service{
					Sub:     true,
					Methods: yyDollar[4].methods,
				},
			}
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.methods = nil
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			yyVAL.methods = append(yyDollar[1].methods, yyDollar[2].method)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("method", yyDollar[1].ident, yyDollar[2].method_fields, yyDollar[3].method_results)
			}
			yyVAL.method = &ast.Method{
				Name:    yyDollar[1].ident,
				Args:    yyDollar[2].method_fields,
				Results: yyDollar[3].method_results,
			}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("method_args", yyDollar[2].method_fields)
			}
			yyVAL.method_fields = yyDollar[2].method_fields
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.method_fields = nil
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("method fields", yyDollar[1].method_field)
			}
			yyVAL.method_fields = []*ast.MethodField{yyDollar[1].method_field}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("method fields", yyDollar[1].method_fields, yyDollar[3].method_field)
			}
			yyVAL.method_fields = append(yyDollar[1].method_fields, yyDollar[3].method_field)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("method field", yyDollar[1].ident, yyDollar[2].type_, yyDollar[3].integer)
			}
			yyVAL.method_field = &ast.MethodField{
				Name: yyDollar[1].ident,
				Type: yyDollar[2].type_,
				Tag:  yyDollar[3].integer,
			}
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.method_results = nil
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.method_results = yyDollar[2].method_results
		}
	case 60:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.method_results = nil
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			yyVAL.method_results = []*ast.MethodResult{yyDollar[1].method_result}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			yyVAL.method_results = append(yyDollar[1].method_results, yyDollar[3].method_result)
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("method result", yyDollar[1].ident, yyDollar[2].type_)
			}
			yyVAL.method_result = &ast.MethodResult{
				Name: yyDollar[1].ident,
				Type: yyDollar[2].type_,
			}
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	case 66:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
		}
	}
	goto yystack /* stack new state and value */
}
