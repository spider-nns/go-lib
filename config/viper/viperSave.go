package main

import (
	"github.com/spf13/viper"
	"log"
)

//保存配置
func main() {
	viper.SetConfigName("vc")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	viper.Set("app_name", "awesome web")
	viper.Set("log_level", "DEBUG")
	viper.Set("mysql.ip", "127.0.0.1")
	viper.Set("mysql.database", "spider")

	err := viper.SafeWriteConfig()
	if err != nil {
		log.Fatalf("write config failed: %v", err)
	}

}
