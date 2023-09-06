package mutationfuzzer

import (
	"fmt"
	"math/rand"
)

type MutationFuzzer struct {
	MinMutation int
	MaxMutation int
}

func New() MutationFuzzer {
	return MutationFuzzer{
		MinMutation: 2,
		MaxMutation: 10,
	}
}

func delChar(s []rune, index int) string {
	return string(append(s[0:index], s[index+1:]...))
}

func deleteRandomCharacter(s string) string {

	if s == "" {
		return ""
	}

	sample := []rune(s)
	return delChar(sample, rand.Intn(len(s)))
}

func insertRandomCharacter(s string) string {

	pos := rand.Intn(len(s))
	random_character := string(rune(rand.Intn(127) + 32))
	return s[:pos] + random_character + s[pos:]

}

func flipRandomCharacter(s string) string {
	if s == "" {
		return ""
	}

	pos := rand.Intn(len(s))
	random_character := string(rune(rand.Intn(127) + 32))

	return s[:pos] + random_character + s[pos+1:]
}

func Mutate(mutation_fuzzer MutationFuzzer, s string) string {
	random_index := rand.Intn(3)
	mutations := rand.Intn(mutation_fuzzer.MaxMutation-mutation_fuzzer.MinMutation+1) + mutation_fuzzer.MinMutation

	result := s
	for i := 0; i <= mutations; i++ {
		switch random_index {
		case 0:
			result = deleteRandomCharacter(result)
		case 1:
			result = insertRandomCharacter(result)
		case 2:
			result = flipRandomCharacter(result)
		}
	}

	fmt.Println(result)

	return result

}
