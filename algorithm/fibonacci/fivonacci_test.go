package fibonacci

import "testing"

var fiboTC []struct{ input, result int } = []struct{ input, result int }{
	{1, 1},
	{2, 1},
	{6, 8},
	{50, 12586269025},
}

func TestFiboIter(t *testing.T) {
	for _, v := range fiboTC {
		result := FiboIter(v.input)
		if result != v.result {
			t.Errorf("fibo iter failed at %d, expected: %d, get: %d", v.input, v.result, result)
		}
	}
}

func TestFiboDynamicBTU(t *testing.T) {
	for _, v := range fiboTC {
		result := FiboDynamicBTU(v.input)
		if result != v.result {
			t.Errorf("fibo iter failed at %d, expected: %d, get: %d", v.input, v.result, result)
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
