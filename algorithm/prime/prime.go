package prime

// BruteForcePrime 하나씩 대입 해보며 소수 판단 O(n log n)
func BruteForcePrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Eratosthenes 에라토노테스의 체를 이용한 풀이
// 일종의 DP, 한번에 여러 숫자의 소수 판단을 할때 유효
// https://ko.wikipedia.org/wiki/%EC%97%90%EB%9D%BC%ED%86%A0%EC%8A%A4%ED%85%8C%EB%84%A4%EC%8A%A4%EC%9D%98_%EC%B2%B4
func Eratosthenes(n int) bool {
	if n < 2 {
		return false
	}

	arr := make([]bool, n+1)
	for i := range arr {
		arr[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if arr[i] {
			for j := i * i; j <= n; j += i {
				arr[j] = false
			}
		}
	}

	return arr[n]
}
