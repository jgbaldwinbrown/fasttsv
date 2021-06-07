package main

import (
    "fmt"
    "io"
    "bufio"
    "os"
    "strings"
    "github.com/pkg/profile"
)

func rearrange_col(split_line []string, cols []int, old_rcols []string) []string {
    out := old_rcols[:0]
    for _, col := range cols {
        out = append(out, split_line[col])
    }
    return out
}

func rearrange_cols(inconn io.Reader, outconn io.Writer, cols []int) {
    scanner := bufio.NewScanner(inconn)
    scanner.Buffer(make([]byte, 0), 10e9)
    var b strings.Builder
    b.Grow(1e6)
    rcols := make([]string, 1000)
    for scanner.Scan() {
        split_line := strings.Split(scanner.Text(), "\t")
        rcols = rearrange_col(split_line, cols, rcols)
        // fmt.Fprintln(&b, strings.Join(rcols, "\t"))
        b.WriteString(strings.Join(rcols, "\t"))
        if b.Len() > 100000 {
            fmt.Fprintf(outconn, "%s", b)
            b.Reset()
        }
    }
}

func main() {
    defer profile.Start().Stop()
    rearrange_cols(os.Stdin, os.Stdout, []int{3,4,1})
}
