package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName  string
	LogLevel string
	MySQL    MySQLConfig
	Redis    RedisConfig
}
type MySQLConfig struct {
}
type RedisConfig struct {
	Ip       string
	Port     int
	Password string
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	var c Config
	viper.Unmarshal(&c)
	fmt.Println(c.Redis)
}
