package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	c, err := ini.Load("/Users/spider/golangProjects/go-lib/config/ini/my.ini")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}
	fmt.Println(c.Section("").Key("lib_name"))
	fmt.Println(c.Section("").Key("get"))
	fmt.Println(c.Section("date").Key("year"), "-", c.Section("date").Key("month"),
		"-", c.Section("date").Key("day"))
	fmt.Println(c.Section("d").Key("di").String())
	i, _ := c.Section("ssh").Key("port").Int()
	fmt.Println(i)
	mustInt := c.Section("ssh").Key("port").MustInt(6379)
	fmt.Println(mustInt)
	//获取信息
	sections := c.Sections()

	fmt.Println("sections: ", sections)
	//分区信息
	names := c.SectionStrings()
	fmt.Println("names: ", names)

	newSection := c.Section("name")
	fmt.Println("section 获取不到会自动创建", newSection)
	section, err := c.NewSection("lib_name")
	if err != nil {
		fmt.Println("create new section if already  exists return err")
		return
	} else {
		fmt.Println("create new section is : ", section)
	}
	strings := c.SectionStrings()
	fmt.Println(strings)

	//	父子分区
	val := c.Section("foo.bar").Key("name")
	fmt.Println("val is : ", val)
	fmt.Println("子分区没有找父分区: ", c.Section("foo.bar").Key("aka"))
}
