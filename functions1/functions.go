package main

import (
    "fmt"
    "math"
)

/*==============================================================================
 * 1. Gauss' formula. (Lecture 2)
 *
 * Implement a function that takes a positive integer n, and returns
 * the sum of the integers from 1 to n.
 *
 *============================================================================*/

func SumOfFirstNIntegers(n int) int {
    return n * (n+1) /2 
}


/*==============================================================================
 * 2. Speed of a Marathoner (lecture 2)
 *
 * Write a function TimeToRun(marathonHours, marathonMinutes, miles) that
 * takes: the time a runner ran a marathon in possibly fractional hours and
 * possibly fractional minutes and a possibly fractional number of miles and
 * return the time in DAYS it should take the runner to run miles if he or she
 * runs at the same pace as they did in the marathon.
 * 
 * For example: TimeToRun(3.1, 23.2, 107.1) should return 0.5938.
 *
 * Your function should also print out the answer in the format:
 *
 *     You could run 107.1 miles in 0.5938 days.
 *
 * Recall that there are 26.2 miles in a marathon.
 *============================================================================*/

func TimeToRun(marathonHours, marathonMinutes, miles float64) float64 {
    var marathonDays float64 = (marathonHours + marathonMinutes/ 60.0) / 24.0 
    const marathonMiles = 26.2
    var marathonSpeed = marathonMiles / marathonDays
    var time = miles / marathonSpeed
    fmt.Println(" You could run ", miles, " miles in ", time, " days.")
    return time
}

/*==============================================================================
 * 3. Generalized Fibonacci sequences (Lecture 3)
 *
 * Implement a function GenFib(a0,a1,n) that takes: two positive integers a0,
 * a1 and a positive integer n, and returns the nth item in the sequence
 * defined by the rule:
 *      a_n = a_{n-1} + a_{n-2}.
 *============================================================================*/

func GenFib(a0, a1, n int) int {
	if n==0 {
		return a0
	}
	if n==1 {
		return a1
	}

	return GenFib(a0, a1, n-1) + GenFib(a0, a1, n-2)
}

/*==============================================================================
 * 4. Reversing Integers (Lecture 3)
 *
 * Write a function ReverseInteger(n) that takes an integer, and returns 
 * the integer  formed by reversing the decimal digits of n. For example:
 *      1234 -> 4321
 *      20000 -> 2
 *      1331  -> 1331
 *      -60 -> -6
 *===========================================================================*/

func ReverseInteger(n int) int {
	var sum int = 0
	for; n != n % 10; {
		sum = (n % 10 + sum) * 10
		n = n / 10
	}
	sum = sum + n
	return sum
}

/*==============================================================================
 * 5. Growth of a Population (Lecture 3)
 *
 * The size at time t of a population with a birth rate r can be modeled as:
 *
 *      x_t = r*x_{t-1}(1 - x_{t-1})
 *
 * Write a function PopSize(r, x_0, max_t) that prints out the size of the
 * population (x_t) for t=0...max_t, where x_0 is the initial population size.
 * Assume population size and the birth rate parameter r are real numbers; t is
 * an integer.
 *
 * Your function should also return the final population size.
 *============================================================================*/
func PopSize(r, x0 float64, max_t int) float64 {
	var x float64 = x0
	for i:=1; i <= max_t; i ++ {
		x = r * x * (1.0 - x)
		if x < 0.0 {
			x = 0.0
		} 
		if x > 1.0 {
			x = 1.0
		}
		fmt.Println(x)
	}
	return x
}

/*==============================================================================
 * 6. Hailstone function (Lecture 3)
 *
 * The Hailstone function h(n) is defined as n/2 if n is even or 3n+1 if n is
 * odd.  The Hailstone sequence for n is defined by repeatedly applying this
 * function:
 *
 *      h(n),  h(h(n)),  h(h(h(n))), ...
 *
 * It's conjectured that for all n, this sequence eventually returns to 1.
 * Write a function HailstoneReturnsTo1(n) to compute the smallest number of times h
 * must be applied to n before the sequence returns to 1.
 *============================================================================*/

func HailstoneReturnsTo1(n int) int {
	if n == 1 {
		return 0
	}
	if n % 2 == 0 {
		n = n/2
	} else {
		n =  3 * n + 1
	}
	return HailstoneReturnsTo1(n) + 1
}

/*==============================================================================
 * 7. Hailstone function maximum (Lecture 3)
 *
 * Write a function MaxHailstoneValue(n) that takes an integer, and returns the
 * maximum value that the Hailstone sequence:
 *
 *      h(n),  h(h(n)),  h(h(h(n))), ...
 *
 * achieves before it returns to 1.
 *============================================================================*/

func MaxHailstoneValue(n int) int {
	if n == 1 {
		return 1
	}
	if n % 2 == 0 {
		n = n/2
	} else {
		n =  3 * n + 1
	}
	var temp = MaxHailstoneValue(n)
 	if n < temp {
 		return temp
 	} else {
 		return n
 	}
}

/*==============================================================================
 * 8. Find the kth digit of an integer n (Lecture 4)
 *
 * Implement a function that takes an integer n, and a positive integer k and
 * returns the k-th decimal digit of n, with digit 1 being the rightmost (least
 * significant) digit.
 *
 *============================================================================*/

func KthDigit(n, k int) int {
	temp := int(math.Pow10(k))
	if n < 0 {
		n = -n
	}
	n = n % temp
	n = n / (temp / 10)
	return n
}

/*===========================================================================
 * 9. Hypergeometric distribution (Lecture 4)
 *
 * Write a function Hypergeometric(M,N,n,k) that takes 4 integers and returns
 * a float64 which is the value of the hypergeometric distribution 
 *   Pr[red = k] = {M choose k}{N choose n-k} / {M+N choose n}
 *
 * Be careful about overflow: Your funciton should be able to compute
 *      Hypergeometric(5000, 5000, 25, 15)
 *      Hypergeometric(5000, 5000, 50, 15)
 * but not necessarily:
 *      Hypergeometric(5000, 5000, 100, 15)
 *===========================================================================*/
func Combination(M, k int) float64 {
	var result float64 = 1.0
	for ; k != 0; {
		result = result * (float64 (M) / float64 (k) )
		M = M - 1
		k = k - 1
	}
	return result
}
func Hypergeometric(M, N, n, k int) float64 {
	Mk := Combination(M, k)
	Nk := Combination(N, n - k)
	MN := Combination(M + N, n)
	return Mk / MN * Nk
}

