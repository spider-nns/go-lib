package main

import (
	"fmt"
	"github.com/apcera/termtables"
)

func main(){
	t := termtables.CreateTable()
	t.AddHeaders("User","Age")
	t.AddRow("spider",24)
	t.AddRow("nns",24)
	fmt.Println("Terminal:")
	fmt.Println(t.Render())

	fmt.Println("HTML:")
	t.SetModeHTML()
	fmt.Println(t.Render())

	fmt.Println("Markdown:")
	t.SetModeMarkdown()
	fmt.Println(t.Render())
}
