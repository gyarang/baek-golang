package fibonacci

// FiboIter 반복문
func FiboIter(cnt int) int {
	if cnt == 0 {
		return 0
	}

	first, second := 0, 1
	for i := 0; i < cnt; i++ {
		first, second = second, first+second
	}

	return first
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
	return fiboRecurTail(cnt, 0, 1)
}

func fiboRecurTail(cnt, first, second int) int {
	if cnt == 0 {
		return first
	}

	return fiboRecurTail(cnt-1, second, first+second)
}

var fiboArr []int

// FiboDynamicTD 동적계획법 탑-다운
func FiboDynamicTD(cnt int) int {
	fiboArr = make([]int, cnt+1)
	return fiboDynamic(cnt)
}

func fiboDynamic(cnt int) int {
	if cnt <= 2 {
		return 1
	}
	if fiboArr[cnt] == 0 {
		fiboArr[cnt] = fiboDynamic(cnt-1) + fiboDynamic(cnt-2)
	}
	return fiboArr[cnt]
}

// FiboDynamicBTU 동적계획법 바텀 업
func FiboDynamicBTU(cnt int) int {
	fiboArr := make([]int, cnt+3)
	fiboArr[1] = 1
	fiboArr[2] = 1

	for i := 3; i <= cnt; i++ {
		fiboArr[i] = fiboArr[i-1] + fiboArr[i-2]
	}
	return fiboArr[cnt]
}
