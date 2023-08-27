package randomfuzzer

import (
	"fmt"
	"math/rand"
)

type RandomFuzzer struct {
	MinLength int
	MaxLength int
	CharStart int
	CharRange int
}

func New() RandomFuzzer {
	return RandomFuzzer{
		MinLength: 10,
		MaxLength: 20,
		CharStart: 32,
		CharRange: 32,
	}
}

func Fuzz(random_fuzzer RandomFuzzer) string {

	string_length := rand.Intn(random_fuzzer.MaxLength-random_fuzzer.MinLength+1) + random_fuzzer.MinLength
	out := ""
	for i := 0; i <= string_length; i++ {
		out += string(rune(rand.Intn(random_fuzzer.CharRange+1) + random_fuzzer.CharStart))
	}
	fmt.Println(out)
	return out
}
