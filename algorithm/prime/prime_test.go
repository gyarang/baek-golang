package prime

import "testing"

var primeTestCase = []struct {
	input  int
	result bool
}{
	{1, false},
	{4652507, true},
	{82589933, true},
}

func TestBruteForcePrime(t *testing.T) {
	for _, v := range primeTestCase {
		result := BruteForcePrime(v.input)
		if result != v.result {
			t.Errorf("brute force prime check failed at %d, expected: %t, get: %t", v.input, v.result, result)
		}
	}
}

func TestEratosthenes(t *testing.T) {
	for _, v := range primeTestCase {
		result := Eratosthenes(v.input)
		if result != v.result {
			t.Errorf("eratosthenes prime check failed at %d, expected: %t, get: %t", v.input, v.result, result)
		}
	}
}
