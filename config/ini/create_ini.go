package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type MysqlIni struct {
	Ip       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

func main() {
	ini := ini.Empty()
	defaultSection := ini.Section("")
	defaultSection.NewKey("app_name", "awesome web")
	defaultSection.NewKey("log_level", "DEBUG")
	mysql, err := ini.NewSection("mysql")
	if err != nil {
		fmt.Println("new mysql section err:", err)
		return
	}
	mysql.NewKey("ip", "127.0.0.1")
	mysql.NewKey("port", "3306")
	mysql.NewKey("user", "spider")
	mysql.NewKey("password", "xxx")
	mysql.NewKey("database", "spider")
	//写入
	err = ini.SaveTo("create.ini")
	if err != nil {
		fmt.Println("save to fail", err)
	}
	err = ini.SaveToIndent("indent.ini", "\t")
	if err != nil {
		fmt.Println("save to indent err:", err)
	}
	_, _ = ini.WriteTo(os.Stdout)
	fmt.Println()
	_, _ = ini.WriteToIndent(os.Stdout, "\t")

	mi := MysqlIni{}
	err = ini.Section("mysql").MapTo(&mi)
	if err != nil {
		fmt.Printf("ini map to struct err:%v", err)
	}
	fmt.Printf("%v\n", mi)
}
