package exercises

import "strconv"

func IsGreatherThan(value string) (int, string) {
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, "error " + err.Error()
	}

	if num > 100 {
		return num, "Is greater than 100"
	} else {
		return num, "Is less than 100"
	}
}
