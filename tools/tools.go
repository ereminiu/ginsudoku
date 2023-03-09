package tools

import (
	"fmt"
	"strconv"
)

func ToInt(s string) int {
	if s == "" {
		return 0
	}

	x, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err.Error())
	}

	return x
}

func ToString(x int) string {
	return strconv.Itoa(x)
}
