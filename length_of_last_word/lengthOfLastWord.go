package main

import (
	"strings"
	"fmt"
)

func lengthOfLastWord(s string) int {
	count := 0
	str := strings.TrimSpace(s)

	for i:=len(str)-1; i>=0; i-- {
		if str[i] == ' '{
			break
		}
		count++
	}
	return count
}


func main() {
	fmt.Println(lengthOfLastWord("Hello World          Welcome to the new world order!     "))
}


// Trimming function
// func TrimSpace(s string) string {
// 	left := 0
// 	right := len(s)-1

// 	for _,char := range s {
// 		if char == ' '{
// 			left++
// 		} else {
// 			break
// 		}
// 	}
// 	for i := right; i>=0; i-- {
// 		if s[i] == ' ' {
// 			right--
// 		} else {
// 			break
// 		}
// 	}

// 	results := ""
// 	for i:=left; i<=right; i++ {
// 		results += string(s[i])
// 	}
// 	return results

	// var result []string
	// for i:=left; i<=right; i++ {
	// 	result = append(result, string(s[i]))
	// }
	// return string(result)
	// }func TrimSpace(s string) string {
		// 	left := 0
		// 	right := len(s)-1
		
		// 	for _,char := range s {
		// 		if char == ' '{
		// 			left++
		// 		} else {
		// 			break
		// 		}
		// 	}
		// 	for i := right; i>=0; i-- {
		// 		if s[i] == ' ' {
		// 			right--
		// 		} else {
		// 			break
		// 		}
		// 	}
		
		// 	results := ""
		// 	for i:=left; i<=right; i++ {
		// 		results += string(s[i])
		// 	}
		// 	return results
		
