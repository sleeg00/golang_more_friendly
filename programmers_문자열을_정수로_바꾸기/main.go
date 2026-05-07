package main

import (
	"strconv"
)

// strconv.Atoi는 앞자리가 +든 -든 숫자면 숫자로 변환해준다..... (수정))
func main(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}
