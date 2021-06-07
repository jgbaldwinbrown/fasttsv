package main

import (
    "fmt"
    "io"
    "bufio"
    "os"
    "strings"
    "github.com/pkg/profile"
)

func rearrange_col(split_line []string, cols []int) []string {
    out := make([]string, len(cols))
    for _, col := range cols {
        out = append(out, split_line[col])
    }
    return out
}

func rearrange_cols(inconn io.Reader, outconn io.Writer, cols []int) {
    scanner := bufio.NewScanner(inconn)
    scanner.Buffer(make([]byte, 0), 10e9)
    for scanner.Scan() {
        split_line := strings.Split(scanner.Text(), "\t")
        rcols := rearrange_col(split_line, cols)
        fmt.Fprintln(outconn, strings.Join(rcols, "\t"))
    }
}

func main() {
    defer profile.Start().Stop()
    rearrange_cols(os.Stdin, os.Stdout, []int{3,4,1})
}
