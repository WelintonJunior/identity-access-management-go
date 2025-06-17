package infraestructure

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisDb *redis.Client

func InitRedis() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		fmt.Println("Erro ao converter REDIS_DB para inteiro:", err)
		db = 0
	}

	rdb := redis.NewClient(&redis.Options{
		Username: os.Getenv("REDIS_USERNAME"),
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	RedisDb = rdb

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	//     panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	//     panic(err)
	// }
	// fmt.Println("key", val)

	// val2, err := rdb.Get(ctx, "key2").Result()
	// if err == redis.Nil {
	//     fmt.Println("key2 does not exist")
	// } else if err != nil {
	//     panic(err)
	// } else {
	//     fmt.Println("key2", val2)
	// }

	return rdb
}
