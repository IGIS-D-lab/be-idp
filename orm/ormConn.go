package orm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-redis/redis"
)

type DataBaseConfig struct {
	Address  string `json:"db"`
	Password string `json:"pw"`
}

func processKey(keyAddr string) (DataBaseConfig, error) {
	var c DataBaseConfig
	file, err := os.Open(keyAddr)
	if err != nil {
		return c, err
	}
	byteVal, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteVal, &c)
	return c, err
}

func Conn(keyAddr string) (*redis.Client, error) {
	loginInfo, err := processKey(keyAddr)
	fmt.Println("go redis tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     loginInfo.Address,
		Password: loginInfo.Password,
		DB:       0,
	})

	// pong to check if server is alive
	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(pong, "connection successful")
	return client, nil
}
