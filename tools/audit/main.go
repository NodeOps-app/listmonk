package main

import (
	"fmt"

	"golang.org/x/text/language"
)

func main() {
	t, err := language.Parse("en-US")
	fmt.Println(t, err)
}
