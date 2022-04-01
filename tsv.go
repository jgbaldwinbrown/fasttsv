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
	Builder []byte
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
		Builder: []byte{},
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

func FprintEscapeEntry(build *[]byte, entry string, sep byte, bs byte) {
	for i:=0; i<len(entry); i++ {
		b := entry[i]
		if b == sep || b == bs {
			*build = append(*build, bs)
		}
		*build = append(*build, b)
	}
}

func FprintEscape(w io.Writer, buf *[]byte, line []string, sep byte, bs byte) {
	*buf = (*buf)[:0]
	if len(line) > 0 {
		FprintEscapeEntry(buf, line[0], sep, bs)
	}
	for _, entry := range line[1:] {
		*buf = append(*buf, sep)
		FprintEscapeEntry(buf, entry, sep, bs)
	}
	w.Write(*buf)
}

func FprintlnEscape(w io.Writer, buf *[]byte, line []string, sep byte, bs byte) {
	FprintEscape(w, buf, line, sep, bs)
	w.Write([]byte("\n"))
}

func FprintSep(w io.Writer, line []string, sep string) {
	fmt.Fprintln(w, strings.Join(line, sep))
}

func Fprint(w io.Writer, line []string) {
	FprintSep(w, line, "\t")
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

func BsSplitOne(ret *[]byte, s string, sep byte, bs byte) (first, rest string) {
	*ret = (*ret)[:0]
	for i:=0; i<len(s); i++ {
		if s[i] == bs {
			i++
			*ret = append(*ret, s[i])
		} else if s[i] == sep {
			return string(*ret), s[i+1:]
		} else {
			*ret = append(*ret, s[i])
		}

	}
	return string(*ret), ""
}

func BsSplit(ret []string, build *[]byte, s string, sep byte, bs byte) []string {
	ret = ret[:0]
	for len(s) > 0 {
		var tok string
		tok, s = BsSplitOne(build, s, sep, bs)
		ret = append(ret, tok)
	}
	return ret
}

