package fibonacci

// FiboIter 반복문
func FiboIter(cnt int) int {
	var result int
	for i, first, second := 0, 0, 1; i <= cnt; i, first, second = i+1, first+second, first {
		if i == cnt {
			result = first
		}
	}

	return result
}

// FiboRecur 재귀 함수
func FiboRecur(cnt int) int {
	if cnt < 2 {
		return cnt
	}

	return FiboRecur(cnt-1) + FiboRecur(cnt-2)
}

// FiboRecurTail 재귀함수 (꼬리)
func FiboRecurTail(cnt int) int {
	return fibonacciRecursiveTail(cnt, 0, 1)
}

func fibonacciRecursiveTail(cnt, first, second int) int {
	if cnt == 0 {
		return first
	}

	return fibonacciRecursiveTail(cnt-1, second, first+second)
}
