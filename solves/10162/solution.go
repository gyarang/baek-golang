package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	if t%10 != 0 {
		fmt.Println(-1)
		return
	}

	var a, b, c int

	for t != 0 {
		if t%300 != 0 {
			if t%60 != 0 {
				t -= 10
				c++
				continue
			}
			t -= 60
			b++
			continue
		}
		a = t / 300
		break
	}

	fmt.Println(a, b, c)
}
