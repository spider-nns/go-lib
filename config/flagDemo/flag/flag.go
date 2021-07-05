package main

import (
	"flag"
	"fmt"
)

var name string

func main() {
	flag.Parse()
	//fmt.Printf("hello,%s!\n", name)
}

func init() {
	//flag.StringVar(&name, "name", "everyone", "greeting !")
	var s = flag.String("name", "name", "greeting")
	fmt.Printf("hello,%s!\n", *s)
}
