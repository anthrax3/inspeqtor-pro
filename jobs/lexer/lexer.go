package lexer

import (

	// "fmt"
	// "github.com/mperham/inspeqtor-pro/jobs/util"

	"io/ioutil"
	"unicode/utf8"

	"github.com/mperham/inspeqtor-pro/jobs/token"
)

const (
	NoState    = -1
	NumStates  = 53
	NumSymbols = 103
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {

	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)

	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {

		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)

		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: 'c'
1: 'h'
2: 'e'
3: 'c'
4: 'k'
5: 'j'
6: 'o'
7: 'b'
8: 's'
9: 'w'
10: 'i'
11: 't'
12: 'h'
13: ','
14: 'h'
15: 'o'
16: 'u'
17: 'r'
18: 'h'
19: 'o'
20: 'u'
21: 'r'
22: 's'
23: 'm'
24: 'i'
25: 'n'
26: 'u'
27: 't'
28: 'e'
29: 'm'
30: 'i'
31: 'n'
32: 'u'
33: 't'
34: 'e'
35: 's'
36: 'd'
37: 'a'
38: 'y'
39: 'd'
40: 'a'
41: 'y'
42: 's'
43: 'h'
44: 'a'
45: 'p'
46: 'p'
47: 'e'
48: 'n'
49: 's'
50: 'e'
51: 'v'
52: 'e'
53: 'r'
54: 'y'
55: '#'
56: '\n'
57: '_'
58: '-'
59: '.'
60: '/'
61: 'k'
62: 'm'
63: 'g'
64: 't'
65: 'p'
66: '%'
67: '!'
68: '#'
69: '$'
70: '%'
71: '&'
72: '''
73: '*'
74: '+'
75: '-'
76: '/'
77: '='
78: '?'
79: '^'
80: '_'
81: '`'
82: '{'
83: '|'
84: '}'
85: '~'
86: '.'
87: '@'
88: '\'
89: '"'
90: '"'
91: ' '
92: '\t'
93: '\n'
94: '\r'
95: 'a'-'z'
96: 'A'-'Z'
97: '0'-'9'
98: 'A'-'Z'
99: 'a'-'z'
100: '0'-'9'
101: \u0100-\U0010ffff
102: .

*/
