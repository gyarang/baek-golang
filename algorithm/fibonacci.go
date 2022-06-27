package main

func fibonacci(cnt int) int {
	if cnt == 0 {
		return 0
	}
	if cnt == 1 || cnt == 2 {
		return 1
	}

	return fibonacci(cnt-1) + fibonacci(cnt-2)
}
