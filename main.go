package main

import (
	"fmt"
	"home/lautaro/dev/ecommerce-golang/variables"
)

func main() {
	status, text := variables.ParseToText(23)
	fmt.Println("🚀 ~ file: main.go:10 ~ funcmain ~ status:", status)
	fmt.Println("🚀 ~ file: main.go:10 ~ funcmain ~ text:", text)
	fmt.Println("Hello world")
}
