package main

import (
    "fmt"
    "io"
    "bufio"
    "os"
    "strings"
    "github.com/pkg/profile"
)

type Scanner struct {
    InScanner bufio.Scanner
    LineBuffer []string
}

func NewScanner(inreader io.Reader) *Reader {
    line := make([]string, 0, 1000)
    scanner := bufio.NewScanner(inconn)
    scanner.Buffer(make([]byte, 0, 10e9))
    out := Reader {
        InScanner,
        LineBuffer: line,
    }
    return &out
}

func (s *Scanner) Scan() bool {
    out := s.InScanner.Scan()
    if !out {
        return out
    }
    s.Line = manual_split(s.InScanner.Text(), '\t', s.Line)
    return out
}

func (s *Scanner) Line() []slice {
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
    if len(s.LinesBuffer) > 0 {
        fmt.Fprintln(w.OutWriter, strings.Join(s.LinesBuffer, "\n"))
        s.LinesBuffer = s.LinesBuffer[:0]
    }
}

func (w *Writer) Write(line []string) {
    s.LinesBuffer = append(s.LinesBuffer, strings.Join(line, "\t"))
    if len(s.LinesBuffer) >= 10000 {
        s.Flush()
    }
}

func rearrange_col(split_line []string, cols []int, old_rcols []string) []string {
    out := old_rcols[:0]
    for _, col := range cols {
        out = append(out, split_line[col])
    }
    return out
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

func rearrange_cols(inconn io.Reader, outconn io.Writer, cols []int) {
    scanner := bufio.NewScanner(inconn)
    scanner.Buffer(make([]byte, 0), 10e9)
    var b strings.Builder
    b.Grow(1e6)
    rcols := make([]string, 0, 1000)
    out_lines := make([]string, 0, 20000)
    split_line := make([]string, 0, 1000)
    for scanner.Scan() {
        split_line := manual_split(scanner.Text(), '\t', split_line)
        rcols = rearrange_col(split_line, cols, rcols)
        // fmt.Fprintln(&b, strings.Join(rcols, "\t"))
        // b.WriteString(strings.Join(rcols, "\t"))
        out_lines = append(out_lines, strings.Join(rcols, "\t"))
        if len(out_lines) >= 10000 {
            fmt.Fprintln(outconn, strings.Join(out_lines, "\n"))
            // b.Reset()
            out_lines = out_lines[:0]
        }
    }
    if len(out_lines) > 0 {
        fmt.Fprintln(outconn, strings.Join(out_lines, "\n"))
    }
}

func main() {
    defer profile.Start().Stop()
    rearrange_cols(os.Stdin, os.Stdout, []int{3,4,1})
}
