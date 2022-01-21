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
}

func NewScanner(inreader io.Reader) *Scanner {
	line := make([]string, 0, 1000)
	scanner := bufio.NewScanner(inreader)
	scanner.Buffer(make([]byte, 0), 10e9)
	out := Scanner {
		InScanner: scanner,
		LineBuffer: line,
	}
	return &out
}

func (s *Scanner) Scan() bool {
	out := s.InScanner.Scan()
	if !out {
		return out
	}
	s.LineBuffer = Split(s.InScanner.Text(), '\t', s.Line())
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
