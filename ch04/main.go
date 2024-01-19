// UTF8 aware charcounter
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflens [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar {
			invalid++
			continue
		}
		counts[r]++
		utflens[n]++
	}
	fmt.Printf("rune\tcount\n")
	for r, n := range counts {
		fmt.Printf("%-5q %3d\n", r, n)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflens {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
