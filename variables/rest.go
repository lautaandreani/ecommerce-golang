package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Date time.Time

func Rest() {
	Name = "Lautaro"
	Status = true
	Salary = 1800.88
	Date = time.Now()

	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ParseToText(number int) (bool, string) {
	name := strconv.Itoa(number)
	return true, name
}
