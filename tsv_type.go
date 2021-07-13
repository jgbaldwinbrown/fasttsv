package fasttsv

import (
	"io"
)

type Tsvi interface {
	AddLine([]string)
	SetLine([]string, int)
	GetLine(int) []string
	SetHeader([]string)
	GetHeader() []string
	UnsetHeader()
	NumLines() int
	NameIndex(string) int
}

type Tsv struct {
	Header map[string]int
	Lines [][]string
}

func NewTsv() Tsv {
	var out Tsv
	out.Header = make(map[string]int)
	out.Lines = make([][]string, 0)
	return out
}

func (t *Tsv) AddLine(line []string) {
	t.Lines = append(t.Lines, line)
}

func (t *Tsv) SetLine(line []string, linenum int) {
	t.Lines[linenum] = line
}

func (t *Tsv) GetLine(linenum int) []string {
	return t.Lines[linenum]
}

func (t *Tsv) SetHeader(line []string) {
	if t.Header == nil {
		t.Header = make(map[string]int)
	}
	for i, s := range line {
		t.Header[s] = i
	}
}

func (t *Tsv) GetHeader() []string {
	if t.Header == nil {
		return nil
	} else {
		out := make([]string, len(t.Header))
		for k, v := range t.Header {
			out[v] = k
		}
		return out
	}
}

func (t *Tsv) NameIndex(s string) int {
	return t.Header[s]
}

func (t *Tsv) UnsetHeader() {
	t.Header = nil
}

func ReadTsv(r io.Reader, Header int) Tsv {
	out := NewTsv()
	scanner := NewScanner(r)
	i := 0
	for scanner.Scan() {
		if i == Header {
			out.SetHeader(scanner.Line())
		} else {
			line := make([]string, len(scanner.Line()))
			copy(line, scanner.Line())
			(&out).AddLine(line)
		}
		i++
	}
	return out
}

func WriteTsv(t Tsvi, iw io.Writer) {
	w := NewWriter(iw)
	if t.GetHeader() != nil {
		hline := make([]string, len(t.GetHeader()))
		w.Write(hline)
	}
	for i:=0; i<t.NumLines(); i++ {
		l := t.GetLine(i)
		w.Write(l)
	}
	w.Flush()
}
