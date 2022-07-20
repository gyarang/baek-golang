package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cnt, require int
	fmt.Fscanln(reader, &cnt, &require)

	arr := make([]int, cnt)
	top := 0
	for i := 0; i < cnt; i++ {
		fmt.Fscan(reader, &arr[i])
		if arr[i] > top {
			top = arr[i]
		}
	}

	start := 0
	end := top
	result := 0
	for start <= end {
		middle := (start + end) / 2
		sum := 0
		for _, height := range arr {
			cut := middle + height - top
			if cut > 0 {
				sum += cut
			}
		}

		if sum >= require {
			result = middle
			end = middle - 1
			continue
		} else {
			start = middle + 1
			continue
		}
	}
	fmt.Println(top - result)
}
