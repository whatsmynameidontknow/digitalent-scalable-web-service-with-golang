package generator

import (
	"math/rand/v2"
	"strings"
)

var huruf []byte = []byte{
	'a', 'b', 'c', 'd', 'e',
	'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o',
	'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y',
	'z', 'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H', 'I',
	'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S',
	'T', 'U', 'V', 'W', 'X',
	'Y', 'Z',
}

var R = rand.New(rand.NewPCG(0, 5))

func GenerateRandomString(numCh int, numWords int) string {
	var generated strings.Builder

	for i := 0; i < numWords; i++ {
		usedCh := numCh / numWords
		for j := 0; j < usedCh; j++ {
			generated.WriteByte(huruf[R.IntN(len(huruf))])
		}
		if i < numWords-1 {
			generated.WriteByte(' ')
		}
	}

	return generated.String()
}
