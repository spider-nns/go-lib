package main
import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

// go get github.com/spf13/viper
func main() {
	pflag.Parse()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	//搜索路径可以多个，依照顺序查找
	viper.AddConfigPath(".")

	viper.SetDefault("redis.port", 6381)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	fmt.Println(viper.Get("app_name"))
	set := viper.IsSet("app_name")
	if set {
		//viper.set 优先级最高
		viper.Set("app_name", "coGo")
		//根据类型获取
		fmt.Println(viper.GetString("app_name"))
	}

	fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	fmt.Println("mysql port: ", viper.Get("mysql.port"))
	fmt.Println("mysql user: ", viper.Get("mysql.user"))
	fmt.Println("mysql password: ", viper.Get("mysql.password"))
	fmt.Println("mysql database: ", viper.Get("mysql.database"))

	fmt.Println("redis ip: ", viper.Get("redis.ip"))
	fmt.Println("redis port: ", viper.Get("redis.port"))
}

func init() {
	//如果键没有通过viper.set 显示设置值 获取时将尝试从命令行选项中读取 ，如果有，优先读取
	//viper 使用 pflag 库来解析选项
	pflag.Int("redis.port", 8381, "Redis port to connect")
	//绑定命令行
	viper.BindPFlags(pflag.CommandLine)

	//绑定环境变量
	viper.AutomaticEnv()//绑定全部
	//单独绑定
	//viper.BindEnv("redis.port")
	//设置环境变量前缀 Get 的时候，viper 会自动加上前缀在查找
	//viper.SetEnvPrefix("")
	println("GOPATH: ",viper.Get("GOPATH"))
}
