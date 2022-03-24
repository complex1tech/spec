// Code generated by goyacc -l -v grammar.out -o grammar.go grammar.y. DO NOT EDIT.
package parser

import __yyfmt__ "fmt"

import "fmt"

type yySymType struct {
	yys int
	// tokens
	ident   string
	integer int
	string  string

	// import
	import_ *Import
	imports []*Import

	// option
	option  *Option
	options []*Option

	// definition
	definition  *Definition
	definitions []*Definition

	// enum
	enum_value  *EnumValue
	enum_values []*EnumValue

	// message
	message_field  *MessageField
	message_fields []*MessageField

	// struct
	struct_field  *StructField
	struct_fields []*StructField

	// type
	type_ *Type
}

const ENUM = 57346
const IMPORT = 57347
const MESSAGE = 57348
const OPTIONS = 57349
const STRUCT = 57350
const IDENT = 57351
const INTEGER = 57352
const STRING = 57353

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ENUM",
	"IMPORT",
	"MESSAGE",
	"OPTIONS",
	"STRUCT",
	"IDENT",
	"INTEGER",
	"STRING",
	"'('",
	"')'",
	"'='",
	"'{'",
	"'}'",
	"';'",
	"'['",
	"']'",
	"'.'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 61

var yyAct = [...]int{
	48, 47, 55, 54, 50, 58, 57, 56, 45, 42,
	39, 31, 30, 49, 8, 43, 40, 37, 29, 21,
	46, 20, 32, 19, 27, 6, 60, 36, 25, 28,
	53, 52, 14, 50, 15, 24, 16, 3, 23, 22,
	5, 1, 35, 44, 13, 34, 41, 51, 12, 33,
	38, 11, 7, 10, 4, 59, 17, 26, 2, 9,
	18,
}

var yyPact = [...]int{
	32, -1000, 33, 13, -1000, 2, -1000, 28, -1000, 10,
	-1000, -1000, -1000, -1000, 30, 29, 26, 15, -1000, -1000,
	-1000, 18, 3, -3, -4, -1000, -1000, 8, -1000, -1000,
	-1000, -1000, 16, 1, 0, -1, -1000, -1000, -1000, 6,
	-1000, -1000, -5, -1000, -1000, -5, 21, 20, -1000, -16,
	-18, -10, -11, -12, 24, 17, -1000, -1000, -1000, -1000,
	-1000,
}

var yyPgo = [...]int{
	0, 60, 59, 58, 57, 56, 54, 53, 52, 51,
	50, 49, 48, 46, 45, 44, 43, 42, 1, 0,
	41,
}

var yyR1 = [...]int{
	0, 20, 1, 1, 2, 2, 3, 3, 6, 6,
	5, 5, 4, 7, 7, 7, 8, 8, 9, 10,
	11, 11, 12, 13, 14, 14, 15, 16, 17, 17,
	18, 18, 19, 19,
}

var yyR2 = [...]int{
	0, 3, 1, 2, 0, 2, 0, 4, 0, 4,
	0, 2, 3, 1, 1, 1, 0, 2, 5, 4,
	0, 2, 5, 4, 0, 2, 5, 3, 0, 2,
	1, 3, 1, 3,
}

var yyChk = [...]int{
	-1000, -20, -3, 5, -6, 7, 12, -8, 12, -2,
	-7, -9, -12, -15, 4, 6, 8, -5, -1, 13,
	11, 9, 9, 9, 9, 13, -4, 9, 11, 15,
	15, 15, 14, -11, -14, -17, 11, 16, -10, 9,
	16, -13, 9, 16, -16, 9, 14, -18, -19, 18,
	9, -18, 10, 10, 19, 20, 17, 17, 17, -19,
	9,
}

var yyDef = [...]int{
	6, -2, 8, 0, 16, 0, 4, 1, 10, 0,
	17, 13, 14, 15, 0, 0, 0, 0, 5, 7,
	2, 0, 0, 0, 0, 9, 11, 0, 3, 20,
	24, 28, 0, 0, 0, 0, 12, 18, 21, 0,
	22, 25, 0, 26, 29, 0, 0, 0, 30, 0,
	32, 0, 0, 0, 0, 0, 27, 19, 23, 31,
	33,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	12, 13, 3, 3, 3, 3, 20, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 17,
	3, 14, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 18, 3, 19, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 15, 3, 16,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
}

var yyTok3 = [...]int{
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
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
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
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
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
	yyn = yyPact[yystate]
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
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
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
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
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
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
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

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			file := &File{
				Imports:     yyDollar[1].imports,
				Options:     yyDollar[2].options,
				Definitions: yyDollar[3].definitions,
			}
			setLexerResult(yylex, file)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("import ", yyDollar[1].string)
			}
			yyVAL.import_ = &Import{
				ID: trimString(yyDollar[1].string),
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("import ", yyDollar[1].ident, yyDollar[2].string)
			}
			yyVAL.import_ = &Import{
				Alias: yyDollar[1].ident,
				ID:    trimString(yyDollar[2].string),
			}
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.imports = nil
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("import_list", yyDollar[1].imports, yyDollar[2].import_)
			}
			yyVAL.imports = append(yyVAL.imports, yyDollar[2].import_)
		}
	case 6:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.imports = nil
		}
	case 7:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("imports", yyDollar[3].imports)
			}
			yyVAL.imports = append(yyVAL.imports, yyDollar[3].imports...)
		}
	case 8:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.options = nil
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("options", yyDollar[3].options)
			}
			yyVAL.options = append(yyVAL.options, yyDollar[3].options...)
		}
	case 10:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.options = nil
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("option_list", yyDollar[1].options, yyDollar[2].option)
			}
			yyVAL.options = append(yyVAL.options, yyDollar[2].option)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("option ", yyDollar[1].ident, yyDollar[3].string)
			}
			yyVAL.option = &Option{
				Name:  yyDollar[1].ident,
				Value: trimString(yyDollar[3].string),
			}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.definitions = nil
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("definitions", yyDollar[1].definitions, yyDollar[2].definition)
			}
			yyVAL.definitions = append(yyVAL.definitions, yyDollar[2].definition)
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if debugParser {
				fmt.Println("enum", yyDollar[2].ident, yyDollar[4].enum_values)
			}
			yyVAL.definition = &Definition{
				Type: DefinitionEnum,
				Name: yyDollar[2].ident,

				Enum: &Enum{
					Values: yyDollar[4].enum_values,
				},
			}
		}
	case 19:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("enum value", yyDollar[1].ident, yyDollar[3].integer)
			}
			yyVAL.enum_value = &EnumValue{
				Name:  yyDollar[1].ident,
				Value: yyDollar[3].integer,
			}
		}
	case 20:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.enum_values = nil
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("enum values", yyDollar[1].enum_values, yyDollar[2].enum_value)
			}
			yyVAL.enum_values = append(yyVAL.enum_values, yyDollar[2].enum_value)
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if debugParser {
				fmt.Println("message", yyDollar[2].ident, yyDollar[4].message_fields)
			}
			yyVAL.definition = &Definition{
				Type: DefinitionMessage,
				Name: yyDollar[2].ident,

				Message: &Message{
					Fields: yyDollar[4].message_fields,
				},
			}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		{
			if debugParser {
				fmt.Println("message field", yyDollar[1].ident, yyDollar[2].type_, yyDollar[3].integer)
			}
			yyVAL.message_field = &MessageField{
				Name: yyDollar[1].ident,
				Type: yyDollar[2].type_,
				Tag:  yyDollar[3].integer,
			}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.message_fields = nil
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("message fields", yyDollar[1].message_fields, yyDollar[2].message_field)
			}
			yyVAL.message_fields = append(yyVAL.message_fields, yyDollar[2].message_field)
		}
	case 26:
		yyDollar = yyS[yypt-5 : yypt+1]
		{
			if debugParser {
				fmt.Println("struct", yyDollar[2].ident, yyDollar[4].struct_fields)
			}
			yyVAL.definition = &Definition{
				Type: DefinitionStruct,
				Name: yyDollar[2].ident,

				Struct: &Struct{
					Fields: yyDollar[4].struct_fields,
				},
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Println("struct field", yyDollar[1].ident, yyDollar[2].type_)
			}
			yyVAL.struct_field = &StructField{
				Name: yyDollar[1].ident,
				Type: yyDollar[2].type_,
			}
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
		{
			yyVAL.struct_fields = nil
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		{
			if debugParser {
				fmt.Println("struct fields", yyDollar[1].struct_fields, yyDollar[2].struct_field)
			}
			yyVAL.struct_fields = append(yyVAL.struct_fields, yyDollar[2].struct_field)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Printf("type *%v\n", yyDollar[1].type_)
			}
			yyVAL.type_ = yyDollar[1].type_
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Printf("type []%v\n", yyDollar[3].type_)
			}
			yyVAL.type_ = &Type{
				Kind:    KindList,
				Element: yyDollar[3].type_,
			}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		{
			if debugParser {
				fmt.Println("base type", yyDollar[1].ident)
			}
			yyVAL.type_ = &Type{
				Kind: getKind(yyDollar[1].ident),
				Name: yyDollar[1].ident,
			}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		{
			if debugParser {
				fmt.Printf("base type %v.%v\n", yyDollar[1].ident, yyDollar[3].ident)
			}
			yyVAL.type_ = &Type{
				Kind:   KindReference,
				Name:   yyDollar[3].ident,
				Import: yyDollar[1].ident,
			}
		}
	}
	goto yystack /* stack new state and value */
}