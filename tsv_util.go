package fasttsv

import (
	"strconv"
	"math"
	"fmt"
)

func Col(t Tsvi, cnum int) (col []string) {
	for i:=0; i<t.NumLines(); i++ {
		line := t.GetLine(i)
		if len(line) <= cnum {
			col = append(col, "")
		} else {
			col = append(col, line[cnum])
		}
	}
	return col
}

func ToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		f = math.NaN()
	}
	return f
}

func ToFloats(ss []string) (f []float64) {
	for _, s := range ss {
		f = append(f, ToFloat(s))
	}
	return f
}

func RemoveNaNs(fs []float64) (ofs []float64) {
	for _, f := range fs {
		if ! math.IsNaN(f) {
			ofs = append(ofs, f)
		}
	}
	return ofs
}

func ToStrings(fs []float64) (ss []string) {
	for _, f := range fs {
		ss = append(ss, fmt.Sprintf("%g", f))
	}
	return ss
}
