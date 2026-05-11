package main

func main(n int64) []int {
	var arr []int

	for n > 0 {
		arr = append(arr, int(n%10))
		n /= 10
	}

	return arr
}
