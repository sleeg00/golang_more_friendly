package main

import (
	"strconv"
	"strings"
)

func main(s string) int {
	str := s[0:len(s)]
	number, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return number
	if strings.HasPrefix(s, "+") {
		return number
	}
	return -number
}
