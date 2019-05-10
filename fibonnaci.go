// Package fibonacci Fibonacci number
// f(0) = 0
// f(1) = 1
// f(n) = f(n-1) + f(n-2) [n > 1]
package fibonacci

// FibR recursion implementation
func FibR(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return FibR(n-1) + FibR(n-2)
}

// FibRM recursion implementation + memoize
func FibRM(n int) int {
	memo := make(map[int]int, n)
	memo[0] = 0
	memo[1] = 1

	var fib func(int) int

	fib = func(p int) int {
		if result, ok := memo[p]; ok {
			return result
		}

		result := fib(p-1) + fib(p-2)
		memo[p] = result

		return result
	}

	return fib(n)
}

// FibL loop implementation
func FibL(n int) int {
	if n < 2 {
		return n
	}

	memo := make([]int, n+1)
	memo[1] = 1

	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
