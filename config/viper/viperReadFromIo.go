package main

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigType("toml")
	tomlConfig := []byte(`app_name = "viper read from io"
	log_level = "DEBUG"
	[mysql]
	ip="127.0.0.1"
	port=3306
	[redis]
	ip="127.0.0.1"
	password="123456"
    database=5
	`)
	err := viper.ReadConfig(bytes.NewBuffer(tomlConfig))
	if err != nil {
		log.Fatalf("read config failed: %v",err.Error())
	}
	fmt.Println("redis database: ",viper.GetInt("redis.database"))
}
