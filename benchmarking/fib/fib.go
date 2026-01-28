package fib

var memo = make(map[int]int)

func Fib(n int) int {
	if n < 2 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = Fib(n-2) + Fib(n-1)

	return memo[n]
}
