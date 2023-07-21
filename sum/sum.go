package sum

import "strconv"

// create a package sum
// create add func that accepts two arguments, string, and string
// convert string to int and do the addition of both the numbers and send the result or error back to the main function

func Add(num1, num2 string) (res int, err error) {
	no1, err := strconv.Atoi(num1)

	if err == nil {
		no2, err := strconv.Atoi(num2)
		if err == nil {
			res = no1 + no2
		}

	}
	return
}
