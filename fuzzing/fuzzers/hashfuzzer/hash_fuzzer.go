package hashfuzzer

import (
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
)

type HashFuzzer struct {
	MinMutation int
	MaxMutation int
}

func New() HashFuzzer {
	return HashFuzzer{
		MinMutation: 2,
		MaxMutation: 10,
	}
}

func RandomHash() []byte {
	hash := make([]byte, 32)
	for i := 0; i < len(hash); i++ {
		hash[i] = byte(rand.Intn(256))
	}
	return hash
}

func Fuzz(hash_fuzzer HashFuzzer, hash []byte) common.Hash {

	mutations := rand.Intn(hash_fuzzer.MaxMutation-hash_fuzzer.MinMutation) + hash_fuzzer.MinMutation

	for i := 0; i <= mutations; i++ {
		random_index := rand.Intn(5)
		hash_index := rand.Intn(len(hash))
		switch random_index {
		case 0:
			// Replace a hash byte with a random byte
			hash[hash_index] = byte(rand.Intn(256))
		case 1:
			// Replace a hash byte with a null byte
			hash[hash_index] = byte(0x00)
		case 2:
			// Reverse the order of the hash bytes in groups of 1
			hash = append(hash[1:], hash[:1]...)
		case 3:
			// Reverse the order of the hash bytes in groups of 2
			hash = append(hash[4:], hash[:4]...)
		case 4:
			// Reverse the order of the hash bytes in groups of 4
			hash = append(hash[8:], hash[:8]...)
		case 5:
			// Reverse the order of the hash bytes in groups of 8
			hash = append(hash[16:], hash[:16]...)
		}
	}

	return common.Hash(hash)
}
