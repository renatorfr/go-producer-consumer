package main

import "testing"

func BenchmarkRun(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Run(5, 5)
	}
}
