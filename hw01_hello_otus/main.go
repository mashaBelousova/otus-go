package main

import (
	"golang.org/x/example/hello/reverse"
	"fmt"
)

func main() {
	text := "Hello, OTUS!"
	reversedText := reverse.String(text)
	fmt.Println(reversedText)
}
