package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	text := "Hello, OTUS!"
	reversedText := reverse.String(text)
	fmt.Println(reversedText)
}
