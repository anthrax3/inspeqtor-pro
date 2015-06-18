package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"check",
		"jobs",
		"with",
		"name",
		",",
		"hour",
		"hours",
		"minute",
		"minutes",
		"day",
		"days",
		"happens",
		"every",
	},

	idMap: map[string]Type{
		"INVALID": 0,
		"$":       1,
		"check":   2,
		"jobs":    3,
		"with":    4,
		"name":    5,
		",":       6,
		"hour":    7,
		"hours":   8,
		"minute":  9,
		"minutes": 10,
		"day":     11,
		"days":    12,
		"happens": 13,
		"every":   14,
	},
}
