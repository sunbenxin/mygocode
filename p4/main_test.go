package main

import "testing"

var result []T

const size = 10000

type T int

func BenchmarkCopy(b *testing.B) {
	orig := make([]T, size)

	for n := 0; n < b.N; n++ {
		cpy := make([]T, len(orig))
		copy(cpy, orig)
		orig = cpy
	}
	result = orig
}

func BenchmarkAppend(b *testing.B) {
	orig := make([]T, size)

	for n := 0; n < b.N; n++ {
		cpy := append([]T{}, orig...)
		orig = cpy
	}
	result = orig
}

func BenchmarkAppendPreCapped(b *testing.B) {
	orig := make([]T, size)

	for n := 0; n < b.N; n++ {
		cpy := append(make([]T, 0, len(orig)), orig...)
		orig = cpy
	}
	result = orig
}

func BenchmarkAppendZeroCapped(b *testing.B) {
	orig := make([]T, size)

	for n := 0; n < b.N; n++ {
		cpy := append(make([]T, 0), orig...)
		orig = cpy
	}
	result = orig
}
