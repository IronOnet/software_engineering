package main 

import "testing"

func longestPalindrome(s string) string{
	res, resLen := "", 0 

	for i:= 0; i < len(s); i++{
		left, right := i, i 
		for left >= 0 && right < len(s) && s[left] == s[right]{
			if (right - left + 1) > resLen{
				res = s[left : right + 1] 
				resLen = right - left + 1 
			}
			left++ 
			right++
		}
		left, right = i, i+1 
		for left >= 0 && right < len(s) && s[left] == s[right]{
			if (right - left + 1) > resLen{
				res = s[left : right+1] 
				resLen = right - left + 1 
			}
			left--
			right++
		}
	}
	return res 
}

func TestLogestPalindrome(t *testing.T){
	string_a := "aaabbbaaa" 
	string_b := "aabaaazezeze" 
	result_a := longestPalindrome(string_a) 
	expected_a := "aaa" 
	result_b := longestPalindrome(string_b) 
	expected_b := "zezeze" 
	if result_a != expected_a && result_b != expected_b{
		t.Errorf("Expected %s and %s but got %s and %s", expected_a, expected_b, result_a, result_b)
	}
}