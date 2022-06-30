package fibonacci

import "testing"

func TestFiboIter(t *testing.T) {
	tc := []int{0, 1, 1, 2, 3, 5, 8, 13}
	for i, v := range tc {
		result := FiboIter(i)
		if result != v {
			t.Errorf("fibo iter failed at %d, expected: %d, get: %d", i, v, result)
		}
	}
}

func TestFiboDynamicBTU(t *testing.T) {
	tc := []int{0, 1, 1, 2, 3, 5, 8, 13}
	for i, v := range tc {
		result := FiboDynamicBTU(i)
		if result != v {
			t.Errorf("fibo dynanic bottom up failed at %d, expected: %d, get: %d", i, v, result)
		}
	}
}

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

func BenchmarkFiboDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboDynamicTD(10)
	}
}

func BenchmarkFiboDynamicBTU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboDynamicBTU(10)
	}
}
