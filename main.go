package main

import (
	"fmt"
	"home/lautaro/dev/ecommerce-golang/exercises"
)

func main() {
	num, str := exercises.IsGreatherThan("10")
	fmt.Printf("%d, %s \n", num, str)
}
