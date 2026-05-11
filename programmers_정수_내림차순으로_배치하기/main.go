package main

import (
	"sort"
	"strconv"
	"strings"
)

func solution(n int64) int64 {
	s := strconv.FormatInt(n, 10)

	chars := strings.Split(s, "")
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] > chars[j]
	})

	result := strings.Join(chars, "")
	answer, _ := strconv.ParseInt(result, 10, 64)
	return answer
}
