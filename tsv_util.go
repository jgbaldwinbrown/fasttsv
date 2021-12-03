package fasttsv

import (
	"strconv"
)

type Type int

const (
	BOOL Type = iota
	INT
	FLOAT
	STRING
)

type Typed interface {
	Type() Type
}

type Int int64
func (i Int) Type() {
	return INT
}

type Bool bool
func (b Bool) Type() {
	return BOOL
}

type Float float64
func (b Float) Type() {
	return FLOAT
}

type String string
func (b String) Type() {
	return STRING
}

func Col(t Tsvi, cnum int) (col []string) {
	for i:=0; i<t.NumLines(); i++ {
		line := t.Line(i)
		if len(line) <= cnum {
			col = append(col, "")
		} else {
			col = append(col, line[i])
		}
	}
	return col
}

func Conv(s string) (t Typed) {
	i, err := strconv.Atoi(s)
	if err == nil {
		return Int(i)
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return Float(i)
	}
	if s == "true" {
		return Bool(true)
	}
	if s == "false" {
		return Bool(false)
	}
	return String(s)
}
