package cache

import (
	"app/utils"
	"encoding/json"
	"reflect"
	"runtime"
	"time"

	"github.com/gofiber/storage/redis/v3"
)

var Redis *redis.Storage

func LoadRedis() {
	if Redis != nil {
		return
	}

	Redis = redis.New(redis.Config{
		URL: utils.Env["REDIS_URL"],
	})
}

func Memoize[P any, T any](
	callback func(P) (T, error),
	exp time.Duration,
) func(P) (T, error) {
	p := reflect.ValueOf(callback).Pointer()
	namespace := runtime.FuncForPC(p).Name()

	return func(args P) (T, error) {
		marshalledArgs, _ := json.Marshal(args)

		key := namespace + ":" + string(marshalledArgs)

		if cachedValue, _ := Redis.Get(key); cachedValue != nil {
			result := new(T)

			if err := json.Unmarshal(cachedValue, result); err != nil {
				return *result, err
			}

			return *result, nil
		}

		result, _ := callback(args)
		resultInBytes, err := json.Marshal(result)

		if err != nil {
			return result, err
		}

		Redis.Set(key, resultInBytes, exp)
		return result, nil
	}
}
