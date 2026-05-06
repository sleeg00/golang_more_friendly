package main

// s의 알파벳을 index만큼 뒤의 알파벳으로 바꿔준다
// abc, 3 => bcd
// if 각 알파벳이 z를 넘어갈 경우 다시 a로 돌아가야한다.
// skip에 있는 알파벳이 등장하면 등장횟수 만큼 더한다.

// golang string immutable
func main(s string, skip string, index int) string {
	skipped := make(map[byte]bool)

	for i := range skip {
		skipped[skip[i]] = true
	}

	str := []byte(s)

	for i := range str {
		moved := 0
		for moved < index {
			str[i]++

			if str[i] > 'z' {
				str[i] = 'a'
			}

			if skipped[str[i]] {
				continue
			}

			moved++
		}
	}

	return string(str)
}
