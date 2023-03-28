// A duplicate counter.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)

	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		counts[reader.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
