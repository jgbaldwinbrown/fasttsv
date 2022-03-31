package fasttsv

import (
	"fmt"
	"io"
	"bufio"
	"strings"
)

type Scanner struct {
	InScanner *bufio.Scanner
	LineBuffer []string
	Builder strings.Builder
	Separator byte
	Escape byte
}

func NewScanner(inreader io.Reader) *Scanner {
	line := make([]string, 0, 1000)
	scanner := bufio.NewScanner(inreader)
	scanner.Buffer(make([]byte, 0), 10e12)
	out := Scanner {
		InScanner: scanner,
		LineBuffer: line,
		Builder: strings.Builder{},
		Separator: '\t',
		Escape: '\\',
	}
	return &out
}

func (s *Scanner) Scan() bool {
	out := s.InScanner.Scan()
	if !out {
		return out
	}
	s.LineBuffer = BsSplit(s.LineBuffer, &s.Builder, s.InScanner.Text(), s.Separator, s.Escape)
	return out
}

func (s *Scanner) Line() []string {
	return s.LineBuffer
}

func Fprint(w io.Writer, line []string) {
	fmt.Fprintln(w, strings.Join(line, "\t"))
}

func Fprintln(w io.Writer, line []string) {
	Fprint(w, line)
	fmt.Fprint(w, "\n")
}

func Split(instring string, sep byte, outslice []string) []string {
	outslice = outslice[:0]
	start := 0
	end := 0
	for end = 0; end < len(instring); end++ {
		if instring[end] == sep {
			outslice = append(outslice, instring[start:end])
			start = end+1
		}
	}
	if start < len(instring) {
		outslice = append(outslice, instring[start:end])
	}
	return outslice
}

func BsSplitOne(ret *strings.Builder, s string, sep byte, bs byte) (first, rest string) {
	ret.Reset()
	for i:=0; i<len(s); i++ {
		if s[i] == bs {
			i++
			ret.WriteByte(s[i])
		} else if s[i] == sep {
			return ret.String(), s[i+1:]
		} else {
			ret.WriteByte(s[i])
		}

	}
	return ret.String(), ""
}

func BsSplit(ret []string, build *strings.Builder, s string, sep byte, bs byte) []string {
	ret = ret[:0]
	for len(s) > 0 {
		var tok string
		tok, s = BsSplitOne(build, s, sep, bs)
		ret = append(ret, tok)
	}
	return ret
}

