package main

import "strconv"
import "testing"

// isPalindrome checks if a number 
// is palindromic or not. 
func isPalindrome(x int) bool {
	s := strconv.Itoa(x) 
	r := []rune(s) 
	for i, j := 0, len(r)-1; i <j; i,j = i+1, j-1{
		if r[i] != r[j]{
			return false
		}
	}
	return true 
}

func TestIsPalindrome(t testing.T){
	pal_1 := 22122 
	pal_2 := 22022022 
	result_1 :=  isPalindrome(pal_1) 
	result_2 := isPalindrome(pal_2) 
	if result_1 != true && result_2 != true{
		t.Errorf("Expected %d and %d to be true but got %c and %c", pal_1, pal_2, result_1, result_2)
	}
}