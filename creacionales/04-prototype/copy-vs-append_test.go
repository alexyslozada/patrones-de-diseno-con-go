package main_test

import (
	"testing"
)

func TestCopy(t *testing.T) {
	y := doCopy(true, false)
	if len(y) != 1000 {
		t.Fatalf("Expected len(y) to be 1000 but was %d", len(y))
	}
}

func TestAppend(t *testing.T) {
	y := doCopy(false, false)
	if len(y) != 1000 {
		t.Fatalf("Expected len(y) to be 1000 but was %d", len(y))
	}
}

func TestAppendAlloc(t *testing.T) {
	y := doCopy(false, true)
	if len(y) != 1000 {
		t.Fatalf("Expected len(y) to be 1000 but was %d", len(y))
	}
}

func doCopy(useCopy bool, preAlloc bool) []int64 {
	existing := make([]int64, 1000, 1000)
	if useCopy {
		y := make([]int64, 1000, 1000)
		copy(y, existing)
		return y
	}

	var init []int64
	if preAlloc {
		init = make([]int64, 0, 1000)
	} else {
		init = []int64{}
	}
	y := append(init, existing...)

	return y
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doCopy(false, false)
	}
}

func BenchmarkAppendAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doCopy(false, true)
	}
}

func BenchmarkAppendAllocInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		existing := make([]int64, 1000, 1000)
		var init []int64

		init = make([]int64, 0, 1000)
		_ = append(init, existing...)
	}
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doCopy(true, true)
	}
}

