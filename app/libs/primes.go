package libs

import (
	"app/cache"
	"errors"
	"math/big"
	"strconv"
	"time"
)

func IsPrime(s string) (bool, error) {
	if cachedValue, _ := cache.Redis.Get(s); cachedValue != nil {
		result := string(cachedValue) == "true"
		return result, nil
	}

	b := new(big.Int)
	b, ok := b.SetString(s, 10)

	if !ok {
		return false, errors.New("value must be an integer")
	}

	result := b.ProbablyPrime(50)
	resultInBytes := []byte(strconv.FormatBool(result))
	duration := 86400 * time.Second
	cache.Redis.Set(s, resultInBytes, duration)

	return result, nil
}
