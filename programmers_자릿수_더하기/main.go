package main

func main(int n) int {
	answer := 0
	for {
		answer += n % 10
		n /= 10
		if n < 10 {
			answer += n
			break
		}
	}
	return answer
}
