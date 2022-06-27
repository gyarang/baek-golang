package fibonacci

import "testing"

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboIter(10)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRecur(10)
	}
}

func BenchmarkFibonacciRecursiveTail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRecurTail(10)
	}
}
