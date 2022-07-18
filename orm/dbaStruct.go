package orm

import (
	"log"

	"github.com/go-redis/redis"
)

/*
	For now the database will be updated like so

	1. if post command is inputed ->
	2. Get redis dataset with RedisJSONGet
	3. append redis dataset with new information
	4. delete the oldest information from the dataset (rolling window)
	5. change the whole JSON data with RedisJSONSet

	It is really inefficient, but it'll do.
	We use this method because
	1. we might run out of memory because it's a free tier
	2. JSON data cannot be appended. JSON data must be updated.
	  - might as well do it as a rolling window while we update it.
	3. The original JSON file will still hold the information.
*/

func RedisJSONGet(redisdb *redis.Client, key1, key2 string) *redis.StringCmd {
	log.Printf("%v : %v JSON.GET from Redis Complete\n", key1, key2)
	cmd := redis.NewStringCmd("JSON.GET", key1, key2)
	redisdb.Process(cmd)
	return cmd
}

func RedisJSONSet(redisdb *redis.Client, key1, key2, value string) *redis.StringCmd {
	cmd := redis.NewStringCmd("JSON.SET", key1, key2, value)
	redisdb.Process(cmd)
	return cmd
}
