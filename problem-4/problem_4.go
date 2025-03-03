package main

func sum_to_n_a(n int) int {
	// your code here
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// time complexity : o(n) -> loop iterates n times
// efficiency : moderate. Works well for moderate values of n. Inefficient for very large values of n

func sum_to_n_b(n int) int {
	// your code here
	// summation formula: (n(n + 1)) / 2
	x := (n * (n + 1)) / 2
	return x
}

// time complexity: O(1) (constant) – fixed number of operations regardless of n
// efficiency: very efficient. performs the same for large n

func sum_to_n_c(n int) int {
	// your code here
	// recursion
	// base case
	if n == 0 {
		return 0
	}
	return n + sum_to_n_c(n-1)
}

// time complexity: O(n) – Makes n recursive calls
// efficiency: not efficient for large n. As it utilizes recursion, can result in stack overflow when n is large

func main() {
	test1 := sum_to_n_a(10)
	test2 := sum_to_n_b(10)
	test3 := sum_to_n_c(10)

	println(test1)
	println(test2)
	println(test3)
}
