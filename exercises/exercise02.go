package exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ProcessNumber() string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string

	fmt.Println("Insert a number:")
	if scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Please insert a valid number")
			ProcessNumber()
			return text
		}

		for i := 0; i < 11; i++ {
			text += fmt.Sprintf("%d * %d = %d\n", num, i, num*i)
		}
	}
	return text
}
