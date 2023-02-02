package main

import "fmt"

func main() {
	// fmt.Println(fib(1))   // 1
	// fmt.Println(fib(3))   // 2
	// fmt.Println(fib(7))   // 13
	// fmt.Println(fib(10))  // ...
	// fmt.Println(fib(50))  // ...
	// fmt.Println(fib(100)) // ...
	// fmt.Println(fib_rec(1))  // 1
	// fmt.Println(fib_rec(3))  // 2
	// fmt.Println(fib_rec(7))  // 13
	// fmt.Println(fib_rec(10)) // ...
	// fmt.Println(fib_rec(30)) // ...

	// fmt.Println(fib_rec(50)) // ...
	// fmt.Println(fib_rec(100)) // ...
	fmt.Println(fib_rec_memo(1, map[int]int{}))   // 1
	fmt.Println(fib_rec_memo(3, map[int]int{}))   // 2
	fmt.Println(fib_rec_memo(7, map[int]int{}))   // 13
	fmt.Println(fib_rec_memo(10, map[int]int{}))  // ...
	fmt.Println(fib_rec_memo(30, map[int]int{}))  // ...
	fmt.Println(fib_rec_memo(50, map[int]int{}))  // ...
	fmt.Println(fib_rec_memo(100, map[int]int{})) // ...

	// fmt.Println(fib_rec(50)) // ...
	// fmt.Println(fib_rec(100)) // ...
}

func fib(n int) int {

	if n < 2 {
		return n
	}

	x := 0
	y := 1

	res := 0

	for i := 2; i <= n; i++ {
		res = x + y
		x, y = y, res
	}

	return res
}

func fib_rec(n int) int {
	if n < 2 {
		return n
	}
	return fib_rec(n-1) + fib_rec(n-2)
}

func fib_rec_memo(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	if n < 2 {
		return n
	}
	memo[n] = fib_rec_memo(n-1, memo) + fib_rec_memo(n-2, memo)
	return memo[n]
}
