package dependencies

import (
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

type ServerDep struct {
	Redis *redis.Client
}

func Initialize() (*ServerDep, []func(*sync.WaitGroup), error) {
	redisClient, redisCleanupFunc, redisErr := initRedis()
	if redisErr != nil {
		return nil, nil, redisErr
	}

	deps := &ServerDep{
		Redis: redisClient,
	}

	cleanupFuncs := []func(*sync.WaitGroup) {
		redisCleanupFunc,
		func(wg *sync.WaitGroup) {
			time.AfterFunc(6 * time.Second, func() {
				log.Println("Closed Module that ends in 6 seconds")
				wg.Done()
			})
		},
		func(wg *sync.WaitGroup) {
			time.AfterFunc(3 * time.Second, func() {
				log.Println("Closed Module that ends in 3 seconds")
				wg.Done()
			})
		},
	}

	return deps, cleanupFuncs, nil
}