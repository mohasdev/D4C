package stringfuzzer

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
)

type AddrFuzzer struct {
	MinMutation int
	MaxMutation int
}

func New() AddrFuzzer {
	return AddrFuzzer{
		MinMutation: 2,
		MaxMutation: 10,
	}
}

func Fuzz(addr_fuzzer AddrFuzzer, s string) (net.IP, int) {

	sep := strings.Index(s, ":")
	ip := net.ParseIP(s[:sep])
	port, err := strconv.Atoi(s[sep+1:])
	if err != nil {
		fmt.Println(err)
	}

	mutations := rand.Intn(addr_fuzzer.MaxMutation-addr_fuzzer.MinMutation) + addr_fuzzer.MinMutation

	for i := 0; i <= mutations; i++ {
		random_index := rand.Intn(4)
		switch random_index {
		case 0:
			for i := 0; i < len(ip); i++ {
				ip[i] = ip[i] + 1
			}
		case 1:
			port = port + 1
			if port < 0 || port > 65535 {
				port = 80
			}
		case 2:
			random := rand.Intn(len(ip))
			ip[random] = ^ip[random]
		case 3:
			random := rand.Intn(len(ip))
			ip[random] = byte(rand.Intn(256))
		case 4:
			port = rand.Intn(65536)
		}

	}

	fmt.Println(ip)
	fmt.Println(port)
	return ip, port
}
