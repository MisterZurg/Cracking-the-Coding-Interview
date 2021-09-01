package String_Compression

import "fmt"

func StringCompression(word string) string{
	counter := 0
	answer := ""
	for i:=1; i < len(word); i++ {
		if word[i - 1] == word[i]{
			counter++
		} else {
			answer += string(word[i - 1])
			answer += fmt.Sprint(counter + 1)
			counter = 0
		}
	}
	answer += string(word[len(word) -1])
	answer += fmt.Sprint(counter + 1)
	return answer
}