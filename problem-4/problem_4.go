package main

func sum_to_n_a(n int) int {
	// your code here
	sum := 0;
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum;
}

func sum_to_n_b(n int) int {
	// your code here
	// summation formula: (n(n + 1)) / 2
	x := (n * (n + 1)) / 2
	return x
}

func sum_to_n_c(n int) int {
	// your code here
	// recursion 
	// base case 
	if n ==	0 {
		return 0
	}
	return n + sum_to_n_c(n - 1)
}


func main() {
	test1 := sum_to_n_a(10)
	test2 := sum_to_n_b(10)
	test3 := sum_to_n_c(10)

	println(test1)
	println(test2)
	println(test3)
}



// alternative solution
// using go routine -> staring a new thread in Java
// make use of channel when doing 

// make more base cases, then call 

// alternative 
// func sum_to_n_d(n int) int {
// 	subSum1 := make(chan int)
// 	subSum2 := make(chan int)
// 	subSum3 := make(chan int)
   
// 	// Function to compute partial sums in goroutines
// 	partSum := func(start, end int, subCh chan int) {
// 	 sum := 0
// 	 for num := start; num <= end; num++ {
// 	  sum += num
// 	 }
// 	 subCh <- sum // Send the sum to the channel
// 	}
   
// 	// Divide n into 3 parts and compute sums concurrently
// 	go partSum(1, n/3, subSum1)
// 	go partSum(n/3+1, 2*n/3, subSum2)
// 	go partSum(2*n/3+1, n, subSum3)
   
// 	// Collect results from channels
// 	result := <-subSum1 + <-subSum2 + <-subSum3
// 	return result
//    }


// sumsum1, subsum2, subsum3 must execute finish in order to get a result. else it will wait indefinitely