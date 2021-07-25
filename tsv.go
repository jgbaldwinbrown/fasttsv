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
	s.LineBuffer = manual_split(s.InScanner.Text(), '\t', s.Line())
	return out
}

func (s *Scanner) Line() []string {
	return s.LineBuffer
}

type Writer struct {
	OutWriter io.Writer
	LinesBuffer []string
}

func NewWriter(outwriter io.Writer) *Writer {
	lines := make([]string, 0, 20000)
	out := Writer {
		OutWriter: outwriter,
		LinesBuffer: lines,
	}
	return &out
}

func (w *Writer) Flush() {
	if len(w.LinesBuffer) > 0 {
		fmt.Fprintln(w.OutWriter, strings.Join(w.LinesBuffer, "\n"))
		w.LinesBuffer = w.LinesBuffer[:0]
	}
}

func (w *Writer) Write(line []string) {
	w.LinesBuffer = append(w.LinesBuffer, strings.Join(line, "\t"))
	if len(w.LinesBuffer) >= 10000 {
		w.Flush()
	}
}

func manual_split(instring string, sep byte, outslice []string) []string {
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
