package main

import "testing"

func BenchmarkCopySlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copySlice()
	}
}

func BenchmarkMergeSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeSlice()
	}
}
