package orm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

type DataBaseConfig struct {
	Address  string `json:"db"`
	Password string `json:"pw"`
}

func processKey(keyAddr string) (DataBaseConfig, error) {
	file, err := os.Open(keyAddr)
	if err != nil {
		log.Println("No Config file")
	}
	byteVal, _ := ioutil.ReadAll(file)
	var c DataBaseConfig
	err = json.Unmarshal(byteVal, &c)
	return c, err
}

func CreateDatabaseObject(keyAddr string) (*redis.Client, error) {
	loginInfo, err := processKey(keyAddr)
	if err != nil {
		log.Println("Login fail")
	}
	client := redis.NewClient(&redis.Options{
		Addr:     loginInfo.Address,
		Password: loginInfo.Password,
		DB:       0, // use default db
	})

	// test if the client is responding
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong, "connection successful")
	return client, nil
}

func RedisJSONGet(redisdb *redis.Client, key1, key2 string) *redis.StringCmd {
	cmd := redis.NewStringCmd("JSON.GET", key1, key2)
	redisdb.Process(cmd)
	return cmd
}

func RedisJSONSet(redisdb *redis.Client, key1, key2, value string) *redis.StringCmd {
	cmd := redis.NewStringCmd("JSON.SET", key1, key2, value)
	redisdb.Process(cmd)
	return cmd
}
