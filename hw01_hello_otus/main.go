package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	s := "Hello, OTUS!"
	fmt.Println(revertWord(s))
}

func revertWord(s string) string {
	return reverse.String(s)
}
