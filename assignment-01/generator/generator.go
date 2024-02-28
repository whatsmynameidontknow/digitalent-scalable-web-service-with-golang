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
	generated.Grow(numCh)

	numCh -= numWords - 1
	eachWordChNum := numCh / numWords

	for numCh > 0 {
		if numWords == 1 {
			eachWordChNum = numCh
		}
		for j := 0; j < eachWordChNum; j++ {
			generated.WriteByte(huruf[R.IntN(len(huruf))])
		}
		numCh -= eachWordChNum
		if numWords > 1 {
			generated.WriteByte(' ')
			numWords--
		}
	}

	return generated.String()
}
