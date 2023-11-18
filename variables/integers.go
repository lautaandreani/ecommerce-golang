package variables

import "fmt"

func ShowIntegers() {
	var int int
	int32 := int32(23)
	int64 := int64(100)

	fmt.Println("int", int)
	fmt.Println("int32", int32)
	fmt.Println("int64", int64)
}
