package libs

import (
	"errors"
	"math/big"
)

func IsPrime(s string) (bool, error) {
	b := new(big.Int)
	b, ok := b.SetString(s, 10)

	if !ok {
		return false, errors.New("value must be an integer")
	}

	result := b.ProbablyPrime(50)
	return result, nil
}
