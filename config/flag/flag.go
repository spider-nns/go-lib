package main

import (
	"flag"
	"go-lib/config/flag/lib"
)

var name string
func main() {
	flag.Parse()
	lib.Hello(name)
}


func init(){
	flag.StringVar(&name,"name","李广昌","say hi")
}