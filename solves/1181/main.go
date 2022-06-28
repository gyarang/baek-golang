package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l int
	fmt.Fscanln(reader, &l)

	arr := make([]string, l)
	for i := 0; i < l; i++ {
		fmt.Fscanln(reader, &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i]) < len(arr[j]) {
			return true
		} else if len(arr[i]) == len(arr[j]) {
			return arr[i] < arr[j]
		} else {
			return false
		}
	})

	for i, v := range arr {
		if i > 0 && arr[i-1] == v {
			continue
		}
		fmt.Fprintln(writer, v)
	}
}
