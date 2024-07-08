package main

import "fmt"

func isValid(s string) bool {
	char := []string{"(", ")", "[", "]", "{", "}"}
	i, j := 0, 0
	for i < len(s) && j < len(char) {
		if string(s[i]) == char[j] && string(s[i+1]) == char[j+1] {
			return true
			i++
		}
		j++
	}
	return false
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("])"))
}
