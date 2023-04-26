package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s1 := "X"
	s2 := "x"
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))
	fmt.Printf(
		"%c: 0x%x(%08b)\n%c: 0x%x(%08b)\n%x\n%x\n%t\n%T\n",
		s1[0], s1, s1[0], s2[0], s2, s2[0], c1, c2, c1 == c2, c1,
	)

	zero(&c1)
	fmt.Printf("%x\n", c1)
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
