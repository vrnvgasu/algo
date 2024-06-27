package dynamic_programming

var m map[int]int

func Fib(n int) int {
	m = make(map[int]int)

	return fib(n)
}

func fib(n int) int {
	if res, ok := m[n]; ok {
		return res
	}

	if n <= 2 {
		m[n] = 1

		return 1
	}

	m[n] = fib(n-1) + fib(n-2)

	return m[n]
}
