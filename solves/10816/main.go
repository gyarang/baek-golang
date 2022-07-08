package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var inCnt, outCnt, input int
	fmt.Fscanln(reader, &inCnt)

	dict := make(map[int]int)
	for i := 0; i < inCnt; i++ {
		fmt.Fscan(reader, &input)
		if _, ok := dict[input]; !ok {
			dict[input] = 1
		} else {
			dict[input]++
		}
	}

	fmt.Fscanf(reader, "\n%d\n", &outCnt)
	for i := 0; i < outCnt; i++ {
		fmt.Fscan(reader, &input)
		if _, ok := dict[input]; !ok {
			fmt.Fprint(writer, "0 ")
		} else {
			fmt.Fprintf(writer, "%d ", dict[input])
		}
	}
}
