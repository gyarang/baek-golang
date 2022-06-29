package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	nm := make(map[int]struct{})
	for i := 1; i <= 10000; i++ {
		num, sum := i, i
		for num > 0 {
			sum += num % 10
			num /= 10
		}

		if _, ex := nm[sum]; !ex {
			nm[sum] = struct{}{}
		}
		if _, ex := nm[i]; !ex {
			fmt.Fprintln(writer, i)
		}
	}
}
