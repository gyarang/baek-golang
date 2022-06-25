package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var l int
	fmt.Fscanln(rd, &l)
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Fscanf(rd, "%d ", &arr[i])
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	fmt.Fprintf(wr, "%d %d", arr[0], arr[len(arr)-1])
}
