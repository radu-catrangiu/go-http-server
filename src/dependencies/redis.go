package dependencies

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/redis/go-redis/v9"
)

func getRedisOptions() (*redis.Options, error) {
	database, err := strconv.Atoi(os.Getenv("REDIS_CLIENT_DB_INDEX"))
	
	options := &redis.Options{
        Addr:     os.Getenv("REDIS_CLIENT_ADDRESS"),
		Username: os.Getenv("REDIS_CLIENT_USERNAME"),
        Password: os.Getenv("REDIS_CLIENT_PASSWORD"),
        DB:       database,
		ClientName: os.Getenv("REDIS_CLIENT_CONN_NAME"),	
    }

	return options, err;
}


func getCleanupFunc(client *redis.Client) (func(*sync.WaitGroup)) {
	return func(wg *sync.WaitGroup)  {
		err := client.Close()
		if err != nil {
			log.Println("Failed to Close Redis Connection")
		}
		log.Println("Closed Redis Connection")
		wg.Done()
	}
}

func initRedis() (*redis.Client, func(*sync.WaitGroup), error) {
	options, err := getRedisOptions()
	if err != nil {
		return nil, nil, err
	}

	client := redis.NewClient(options)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, nil, err
	}

	cleanupFunc := getCleanupFunc(client)

	log.Println("Initialized Redis connection")
	return client, cleanupFunc, nil
}