package main

import (
	"fmt"
	"math"
	"testing"
)

// TestSumOfFirstNIntegers() tests the SumOfFirstNIntegers function on several
// different integer inputs.
func TestSumOfFirstNIntegers(t *testing.T) {
	r := SumOfFirstNIntegers(100)
	t.Logf("SumOfFirstIntegers(100) = %d", r)
	if r != 5050 {
		t.Errorf("SumOfFirstNIntegers(100) should equal 5050")
	}

	r = SumOfFirstNIntegers(2456)
	t.Logf("SumOfFirstIntegers(2456) = %d", r)
	if r != 3017196 {
		t.Errorf("SumOfFirstNIntegers(2456) should equal 3017196")
	}

	r = SumOfFirstNIntegers(10101)
	t.Logf("SumOfFirstIntegers(10101) = %d", r)
	if r != 51020151 {
		t.Errorf("SumOfFirstNIntegers(10101) should equal 51020151")
	}
}

// TestGenFib() tests the generalized fibbonacci function using a few different
// initial conditions.
func TestTimeToRun(t *testing.T) {
    r := TimeToRun(3.1, 23.2, 107.1)
    if math.Abs(r - 0.5938) > 1e-4 {
        t.Errorf("TimeToRun(3.1, 23.2, 107.1) should equal 0.5938")
    }

    r = TimeToRun(5.0, 7.0, 1000.3)
    if math.Abs(r - 8.1396) > 1e-4 {
        t.Errorf("TimeToRun(5.0, 7.0, 1000.3) should equal 8.1396")
    }

    r = TimeToRun(2.0, 4.3, 1000.3)
    if math.Abs(r - 3.2956) > 1e-4 {
        t.Errorf("TimeToRun(2.0, 4.3, 1000.3) should equal 3.2956")
    }

}

// TestGenFib() tests the generalized fibbonacci function using a few different
// initial conditions.
func TestGenFib(t *testing.T) {
	r := GenFib(1, 1, 10)
	t.Logf("GenFib(1,1,10) = %d", r)
	if r != 89 {
		t.Errorf("GenFib(1,1,10) should equal 89")
	}

	r = GenFib(1, 2, 10)
	t.Logf("GenFib(1,2,10) = %d", r)
	if r != 144 {
		t.Errorf("GenFib(1,2,10) should equal 144")
	}

	r = GenFib(3, 2, 20)
	t.Logf("GenFib(3,2,20) = %d", r)
	if r != 26073 {
		t.Errorf("GenFib(3,2,20) should equal 26073")
	}
}

// TestReverseIntegers runs ReverseIntegers()
func TestReverseIntegers(t *testing.T) {
    r := ReverseInteger(1234)
    if r != 4321 {
        t.Errorf("ReverseInteger(1234) should equal 4321")
    }

    r = ReverseInteger(200000)
    if r != 2 {
        t.Errorf("ReverseInteger(200000) should equal 2")
    }

    r = ReverseInteger(-2277)
    if r != -7722 {
        t.Errorf("ReverseInteger(-2277) should equal -7722")
    }
}

// TestPopSize() tests the modeling of population sizes by checking that the
// final population size is correct. Note: this function does *not* test that
// the intermediate values are printed out correctly.
func TestPopSize(t *testing.T) {
    fmt.Println("************ Population Size *************")
	fmt.Println("==== a=2 ====")
	r := PopSize(2.0, 0.1, 20.0)
	if math.Abs(r-0.5) > 1e-8 {
		t.Errorf("PopSize(2.0, 0.1, 20.0) should equal 0.5")
	}

	fmt.Println("==== a=2.9 ====")
	r = PopSize(2.9, 0.1, 20.0)
	if math.Abs(r-0.645941819757292) > 1e-8 {
		t.Errorf("PopSize(2.9, 0.1, 20.0) should equal 0.645941819757292")
	}

	fmt.Println("==== a=4 ====")
	r = PopSize(4.0, 0.1, 20.0)
	if math.Abs(r-0.8200138733909665) > 1e-8 {
		t.Errorf("PopSize(4.0, 0.1, 20.0) should equal 0.8200138733909665")
	}
}

// TestHailstone() checks that HailstoneReturnsTo1() on a few initial values.
func TestHailstone(t *testing.T) {
	r := HailstoneReturnsTo1(54)
	t.Logf("HailstoneReturnsTo1(54) after %d steps", r)
	if r != 112 {
		t.Errorf("HailstoneReturnsTo1(54) should be 112")
	}

	r = HailstoneReturnsTo1(649)
	t.Logf("HailstoneReturnsTo1(649) after %d steps", r)
	if r != 144 {
		t.Errorf("HailstoneReturnsTo1(54) should be 144")
	}
}

// TestMaxHailstone() checks that MaxHailstoneValue() returns the correct 
// value for a few examples
func TestMaxHailstone(t *testing.T) {
	r := MaxHailstoneValue(54)
	t.Logf("Max Hailstone value starting at 54 = %d", r)
	if r != 9232 {
		t.Errorf("MaxHAilstoneValue(54) should be 9232")
	}

	r = MaxHailstoneValue(101)
	t.Logf("Max Hailstone value starting at 101 = %d", r)
	if r != 304 {
		t.Errorf("MaxHAilstoneValue(101) should be 304")
	}

	r = MaxHailstoneValue(649)
	t.Logf("Max Hailstone value starting at 649 = %d", r)
	if r != 9232 {
		t.Errorf("MaxHAilstoneValue(649) should be 9232")
	}
}

// TestKthDigit() runs KthDigit()
func TestKthDigit(t *testing.T) {
    r := KthDigit(3344556, 1)
    if r != 6 {
        t.Errorf("KthDigit(3344556, 1) should = 1")
    }

    r = KthDigit(-23478, 3)
    if r != 4 {
        t.Errorf("KthDigit(-23478, 3) should = 4")
    }

    r = KthDigit(24, 3)
    if r != 0 {
        t.Errorf("KthDigit(24, 3) should = 0")
    }
}

// TestHypergeometric() runs Hypergeometric()
func TestHypergeometric(t *testing.T) {
    r := Hypergeometric(20, 50, 25, 10)
    if math.Abs(r - 0.064415) > 1e-3 {
        t.Errorf("Hypergeometric(20,50, 25, 10) should equal 0.064415")
    }

    r = Hypergeometric(5000, 5000, 25, 15)
    if math.Abs(r - 0.09741) > 1e-4 {
        t.Errorf("Hypergeometric(5000,5000, 25, 15) should equal 0.09741")
    }

    r = Hypergeometric(5000, 5000, 50, 15)
    if math.Abs(r - 0.00196) > 1e-3 {
        t.Errorf("Hypergeometric(5000,5000, 50, 15) should equal 0.00196")
    }
}
