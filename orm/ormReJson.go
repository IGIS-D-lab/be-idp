package orm

import (
	"encoding/json"
	"log"

	"github.com/go-redis/redis"
)

/*
	JSONGet
	- T: go generics.
	- gets any type constraint, fill it up with parsed JSON
	- returns container
*/
func JSONGet[T any](database *redis.Client, key1, key2 string, container *T) (T, error) {
	log.Printf("JSON GET: %v:: %v Called\n", key1, key2)
	cmd := redis.NewStringCmd("JSON.GET", key1, key2)
	_ = database.Process(cmd)
	v, _ := cmd.Result()

	byteVal := []byte(v)
	err := json.Unmarshal(byteVal, &container)
	if err != nil {
		return *container, err
	}
	return *container, err
}

/*
	JSONSet
	- T: go generics
	- gets any type of constraint -> and parse it into string
*/
func JSONSet[T any](database *redis.Client, key1, key2 string, container *T) (string, error) {
	setVal, _ := json.Marshal(container)

	log.Printf("JSON SET : %v:: %v Setting", key1, key2)
	cmd := redis.NewStringCmd("JSON.SET", key1, key2, string(setVal))
	_ = database.Process(cmd)
	result, err := cmd.Result()
	return result, err
}
