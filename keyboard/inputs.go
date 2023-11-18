package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var number1 int
var number2 int
var text string
var err error

func UserInput() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese numero 1: ")
	if scanner.Scan() {
		number1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Please type a valid number 1" + err.Error())
		}
	}

	fmt.Println("Ingrese numero 2: ")
	if scanner.Scan() {
		number2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("Please type a valid number 2" + err.Error())
		}
	}

	fmt.Println("Type a description")
	if scanner.Scan() {
		text = scanner.Text()
	}

	fmt.Println(text, number1*number2)
}
