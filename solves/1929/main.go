package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var min, max int
	fmt.Scanln(&min, &max)
	for i := min; i <= max; i++ {
		if IsPrime(i) {
			fmt.Fprintln(writer, i)
		}
	}
}
