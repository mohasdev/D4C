package stringfuzzer

import (
	"fmt"
	"math/rand"
	"strings"
)

type StringFuzzer struct {
	MinMutation int
	MaxMutation int
}

func New() StringFuzzer {
	return StringFuzzer{
		MinMutation: 2,
		MaxMutation: 10,
	}
}

func Fuzz(string_fuzzer StringFuzzer, s string) string {

	mutations := rand.Intn(string_fuzzer.MaxMutation-string_fuzzer.MinMutation+1) + string_fuzzer.MinMutation

	result := s
	for i := 0; i <= mutations; i++ {
		random_index := rand.Intn(3)
		switch random_index {
		case 0:
			// Add a random character
			pos := rand.Intn(len(result))
			random_character := string(rune(rand.Intn(127) + 32))
			result = result[:pos] + random_character + result[pos:]
		case 1:
			// Delete a random character
			if result != "" {
				sample := []rune(result)
				result = delChar(sample, rand.Intn(len(result)))
			}
		case 2:
			// flip random character
			if result != "" {
				pos := rand.Intn(len(s))
				random_character := string(rune(rand.Intn(127) + 32))
				result = result[:pos] + random_character + result[pos+1:]
			}
		case 3:
			// get the first half of the string
			result = result[:len(result)/2]
		case 4:
			//get the second half of the string
			result = result[len(result)/2:]
		case 5:
			// get a random number of character from the start of the string
			result = result[:rand.Intn(len(result))]
		case 6:
			// get a random number of character starting the end of the string
			result = result[rand.Intn(len(result)):]
		case 7:
			// gets all characters in even order
			var str string
			for i := 0; i < len(result); i++ {
				if i%2 == 0 {
					str += string(result[i])
				}
			}
			result = str
		case 8:
			// gets all characters in odd order
			var str string
			for i := 0; i < len(result); i++ {
				if i%2 != 0 {
					str += string(result[i])
				}
			}
			result = str
		case 9:
			// double spaces
			result = strings.ReplaceAll(result, " ", "    ")
		case 10:
			// replace spaces by line break
			result = strings.ReplaceAll(result, " ", "\n")
		case 11:
			//replace spaces by tab
			result = strings.ReplaceAll(result, " ", "\t")
		case 12:
			// replace spaces by \x00
			result = strings.ReplaceAll(result, " ", "\x00")
		case 13:
			// replace spaces by \u0000
			result = strings.ReplaceAll(result, " ", "\u0000")
		case 14:
			// replace spaces by _
			result = strings.ReplaceAll(result, " ", "_")
		}
	}

	fmt.Println(result)
	return result
}

func delChar(s []rune, index int) string {
	return string(append(s[0:index], s[index+1:]...))
}
