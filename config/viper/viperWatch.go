package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	//自动监听配置修改，如果修改重新加载配置
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
	viper.WatchConfig()
	//添加回调,fsnotify 库实现 监听修改
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config file: %s op:%s\n", e.Name, e.Op)
	})
	fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
	time.Sleep(time.Second * 5)
	fmt.Println("redis port after sleep: ", viper.Get("redis.port"))

}
