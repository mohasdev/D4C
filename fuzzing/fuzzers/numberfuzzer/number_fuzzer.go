package stringfuzzer

import (
	"fmt"
	"math"
	"math/rand"
)

type NumberFuzzer struct {
	MinMutation int
	MaxMutation int
}

func New() NumberFuzzer {
	return NumberFuzzer{
		MinMutation: 2,
		MaxMutation: 10,
	}
}

func Fuzz(number_fuzzer NumberFuzzer, u int64) int64 {

	mutations := rand.Intn(number_fuzzer.MaxMutation-number_fuzzer.MinMutation+1) + number_fuzzer.MinMutation

	result := u
	for i := 0; i <= mutations; i++ {
		random_index := rand.Intn(14)
		switch random_index {
		case 0:
			// Add 1
			result++
		case 1:
			// Subtract 1
			result--
		case 2:
			// Double
			result *= 2
		case 3:
			// Divide by 2
			result /= 2
		case 4:
			// Replace the value with a random number
			result = int64(rand.Intn(math.MaxInt64))
		case 5:
			// Invert value
			result = ^result
		case 6:
			// left shift by 2
			result <<= 2
		case 7:
			// right shift by 2
			result >>= 2
		case 8:
			// get the square of the value
			result = result * result
		case 9:
			// get the square root of the value
			result = int64(math.Sqrt(float64(result)))
		}
	}

	fmt.Println(result)
	return result
}
