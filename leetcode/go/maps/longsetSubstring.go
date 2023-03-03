package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	left := 0
	res := 0

	for right, _ := range s {
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}
		charSet[s[right]] = true
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	string1 := "aaabbceeeeekdiloman"
	string2 := "aabkk"
	fmt.Printf("The lenght of the longest substring for the string %s is %d\n", string1, lengthOfLongestSubstring(string1)) 
	fmt.Printf("The lenght of the longest substring for the string %s is %d\n", string2, lengthOfLongestSubstring(string2)) 
}