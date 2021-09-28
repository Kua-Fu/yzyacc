// Code generated by goyacc -o echo.go -p Input parser.y. DO NOT EDIT.

//line parser.y:4

package main

import __yyfmt__ "fmt"

//line parser.y:5

import (
	"fmt"
)

//line parser.y:14
type InputSymType struct {
	yys int
	val string
}

const CHARACTER = 57346

var InputToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"CHARACTER",
	"'\\n'",
}

var InputStatenames = [...]string{}

const InputEofCode = 1
const InputErrCode = 2
const InputInitialStackSize = 16

//line parser.y:37

//line yacctab:1
var InputExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const InputPrivate = 57344

const InputLast = 5

var InputAct = [...]int{
	5, 4, 3, 1, 2,
}

var InputPact = [...]int{
	-1000, -2, -4, -1000, -1000, -1000,
}

var InputPgo = [...]int{
	0, 4, 3,
}

var InputR1 = [...]int{
	0, 2, 2, 1, 1,
}

var InputR2 = [...]int{
	0, 0, 3, 1, 2,
}

var InputChk = [...]int{
	-1000, -2, -1, 4, 5, 4,
}

var InputDef = [...]int{
	1, -2, 0, 3, 2, 4,
}

var InputTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5,
}

var InputTok2 = [...]int{
	2, 3, 4,
}

var InputTok3 = [...]int{
	0,
}

var InputErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	InputDebug        = 0
	InputErrorVerbose = false
)

type InputLexer interface {
	Lex(lval *InputSymType) int
	Error(s string)
}

type InputParser interface {
	Parse(InputLexer) int
	Lookahead() int
}

type InputParserImpl struct {
	lval  InputSymType
	stack [InputInitialStackSize]InputSymType
	char  int
}

func (p *InputParserImpl) Lookahead() int {
	return p.char
}

func InputNewParser() InputParser {
	return &InputParserImpl{}
}

const InputFlag = -1000

func InputTokname(c int) string {
	if c >= 1 && c-1 < len(InputToknames) {
		if InputToknames[c-1] != "" {
			return InputToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func InputStatname(s int) string {
	if s >= 0 && s < len(InputStatenames) {
		if InputStatenames[s] != "" {
			return InputStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func InputErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !InputErrorVerbose {
		return "syntax error"
	}

	for _, e := range InputErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + InputTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := InputPact[state]
	for tok := TOKSTART; tok-1 < len(InputToknames); tok++ {
		if n := base + tok; n >= 0 && n < InputLast && InputChk[InputAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if InputDef[state] == -2 {
		i := 0
		for InputExca[i] != -1 || InputExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; InputExca[i] >= 0; i += 2 {
			tok := InputExca[i]
			if tok < TOKSTART || InputExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if InputExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += InputTokname(tok)
	}
	return res
}

func Inputlex1(lex InputLexer, lval *InputSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = InputTok1[0]
		goto out
	}
	if char < len(InputTok1) {
		token = InputTok1[char]
		goto out
	}
	if char >= InputPrivate {
		if char < InputPrivate+len(InputTok2) {
			token = InputTok2[char-InputPrivate]
			goto out
		}
	}
	for i := 0; i < len(InputTok3); i += 2 {
		token = InputTok3[i+0]
		if token == char {
			token = InputTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = InputTok2[1] /* unknown char */
	}
	if InputDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", InputTokname(token), uint(char))
	}
	return char, token
}

func InputParse(Inputlex InputLexer) int {
	return InputNewParser().Parse(Inputlex)
}

func (Inputrcvr *InputParserImpl) Parse(Inputlex InputLexer) int {
	var Inputn int
	var InputVAL InputSymType
	var InputDollar []InputSymType
	_ = InputDollar // silence set and not used
	InputS := Inputrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Inputstate := 0
	Inputrcvr.char = -1
	Inputtoken := -1 // Inputrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Inputstate = -1
		Inputrcvr.char = -1
		Inputtoken = -1
	}()
	Inputp := -1
	goto Inputstack

ret0:
	return 0

ret1:
	return 1

Inputstack:
	/* put a state and value onto the stack */
	if InputDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", InputTokname(Inputtoken), InputStatname(Inputstate))
	}

	Inputp++
	if Inputp >= len(InputS) {
		nyys := make([]InputSymType, len(InputS)*2)
		copy(nyys, InputS)
		InputS = nyys
	}
	InputS[Inputp] = InputVAL
	InputS[Inputp].yys = Inputstate

Inputnewstate:
	Inputn = InputPact[Inputstate]
	if Inputn <= InputFlag {
		goto Inputdefault /* simple state */
	}
	if Inputrcvr.char < 0 {
		Inputrcvr.char, Inputtoken = Inputlex1(Inputlex, &Inputrcvr.lval)
	}
	Inputn += Inputtoken
	if Inputn < 0 || Inputn >= InputLast {
		goto Inputdefault
	}
	Inputn = InputAct[Inputn]
	if InputChk[Inputn] == Inputtoken { /* valid shift */
		Inputrcvr.char = -1
		Inputtoken = -1
		InputVAL = Inputrcvr.lval
		Inputstate = Inputn
		if Errflag > 0 {
			Errflag--
		}
		goto Inputstack
	}

Inputdefault:
	/* default state action */
	Inputn = InputDef[Inputstate]
	if Inputn == -2 {
		if Inputrcvr.char < 0 {
			Inputrcvr.char, Inputtoken = Inputlex1(Inputlex, &Inputrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if InputExca[xi+0] == -1 && InputExca[xi+1] == Inputstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Inputn = InputExca[xi+0]
			if Inputn < 0 || Inputn == Inputtoken {
				break
			}
		}
		Inputn = InputExca[xi+1]
		if Inputn < 0 {
			goto ret0
		}
	}
	if Inputn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Inputlex.Error(InputErrorMessage(Inputstate, Inputtoken))
			Nerrs++
			if InputDebug >= 1 {
				__yyfmt__.Printf("%s", InputStatname(Inputstate))
				__yyfmt__.Printf(" saw %s\n", InputTokname(Inputtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Inputp >= 0 {
				Inputn = InputPact[InputS[Inputp].yys] + InputErrCode
				if Inputn >= 0 && Inputn < InputLast {
					Inputstate = InputAct[Inputn] /* simulate a shift of "error" */
					if InputChk[Inputstate] == InputErrCode {
						goto Inputstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if InputDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", InputS[Inputp].yys)
				}
				Inputp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if InputDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", InputTokname(Inputtoken))
			}
			if Inputtoken == InputEofCode {
				goto ret1
			}
			Inputrcvr.char = -1
			Inputtoken = -1
			goto Inputnewstate /* try again in the same state */
		}
	}

	/* reduction by production Inputn */
	if InputDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Inputn, InputStatname(Inputstate))
	}

	Inputnt := Inputn
	Inputpt := Inputp
	_ = Inputpt // guard against "declared and not used"

	Inputp -= InputR2[Inputn]
	// Inputp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Inputp+1 >= len(InputS) {
		nyys := make([]InputSymType, len(InputS)*2)
		copy(nyys, InputS)
		InputS = nyys
	}
	InputVAL = InputS[Inputp+1]

	/* consult goto table to find next state */
	Inputn = InputR1[Inputn]
	Inputg := InputPgo[Inputn]
	Inputj := Inputg + InputS[Inputp].yys + 1

	if Inputj >= InputLast {
		Inputstate = InputAct[Inputg]
	} else {
		Inputstate = InputAct[Inputj]
		if InputChk[Inputstate] != -Inputn {
			Inputstate = InputAct[Inputg]
		}
	}
	// dummy call; replaced with literal code
	switch Inputnt {

	case 2:
		InputDollar = InputS[Inputpt-3 : Inputpt+1]
//line parser.y:29
		{
			fmt.Printf("Read character: %s\n", InputDollar[2].val)
		}
	case 4:
		InputDollar = InputS[Inputpt-2 : Inputpt+1]
//line parser.y:34
		{
			InputVAL.val = InputDollar[1].val + InputDollar[2].val
		}
	}
	goto Inputstack /* stack new state and value */
}
