package main

import "testing"

func testPopCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if (x>>i)&1 == 1 {
			count += 1
		}
	}
	return count
}

func TestPopCount(t *testing.T) {
	for x := uint64(0); x < 10_000_000; x++ {
		if got, want := popCount(x), testPopCount(x); got != want {
			t.Fatalf("popCount(%d): %d != %d\n", x, got, want)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for x := uint64(0); x < uint64(b.N); x++ {
		popCount(x)
	}
}

func BenchmarkTestPopCount(b *testing.B) {
	for x := uint64(0); x < uint64(b.N); x++ {
		testPopCount(x)
	}
}
