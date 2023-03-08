// Given a string find, the length of  the longest substring in it
// with no more than K distinct characters
// Example 1:
// Input: String="araaci", k=2
// Explanation: The longest substring with no more than '2' distinct characters
// is "araa"
//
// Example 2:
// Input: string="araaci", K=1
// Output: 2
// Explanation: THe longest substring with no more than '1' distintc character
// is "aa"
//
// Example 3
// Input: string="cbbebi", K=3
// Ouput: 5
// Explanation: THe longest substrings with no more than 3 distinct characters are
// "cbbeb" & "bbebi".

// Solution
// This problem follows the sliding window pattern and we can use a similar
// dynamic sliding window strategy as discuessed in the smallest subarray with
// a given sum. We can use a Hashmap to remember the frequency of each characgter
// we have processed.

// 1. First we will insert characters from the begining of the string until we have
// K distinct characters in the HashMap

// 2. These characters will constitute our sliding windw. We are asked to find the
// longest such window having no more than K distinct characters. We will remember
// the length of this window as the longest window so far.

// 3. After this we will keep adding one character in the sliding window, in a stepwise
// fashion.

// 4. IN each step, we will try to shrink the window from the beginning if the count
// of distinct characters in the Hashmap is larger than 'K'. We will shrink the window
// until we have no more than 'K' distinct charactrers in teh Hashmap.

// 5. While shrinking, we'll decreemnt the frequencey of the character going out
// of the window and remove it from the Hashmap if its frequency becomes zero

// 6. At the end of each step, we'll check if the current window length is the longest so far
// and if so remember its length.
package main

import (
	"fmt"
	"log"
) 

func LongestSubstringWithKdistinct(s string, K int) int{
	windowStart := 0 
	charFrequency := make(map[string]int, 0)
	maxLength := 0 

	for windowEnd := 0; windowEnd < len(s); windowEnd++{
		rightChar := string(s[windowEnd])
		charFrequency[rightChar]++

		for len(charFrequency) > K{
			leftChar := string(s[windowStart])
			charFrequency[leftChar]-- 
			if charFrequency[leftChar] == 0{
				delete(charFrequency, leftChar)
			}
			windowStart++
		}
		maxLength = Max(maxLength, windowEnd - windowStart + 1)
	}

	return maxLength
}

func LongestSubstringWithKdistinct2(s string, K int) int{
	windowStart := 0 
	charFrequency := make(map[string]int, 0) 
	maxLength := 0 
	for windowEnd := 0; windowEnd < len(s); windowEnd++{
		rightChar := string(s[windowEnd])
		charFrequency[rightChar]++ 

		for len(charFrequency) > K{
			leftChar := string(s[windowStart]) 
			charFrequency[leftChar]-- 
			if charFrequency[leftChar] == 0{
				delete(charFrequency, leftChar)
			}
			windowStart++
		}
		maxLength = Max(maxLength	, windowEnd - windowStart + 1)
	}
	return maxLength
}

func LongestSubstringWithKdistinct3(s string, K int) int{
	windowStart := 0 
	maxLen := 0 
	charFrequency := make(map[string]int, 0)

	for windowEnd := 0; windowEnd < len(s); windowEnd++{
		rightChar := string(s[windowEnd])
		charFrequency[rightChar]++ 

		for len(charFrequency) > K{
			leftChar := string(s[windowStart]) 
			charFrequency[leftChar]-- 
			if charFrequency[leftChar] == 0{
				delete(charFrequency, leftChar)
			}
			windowStart++
		}
		maxLen = Max(maxLen, windowEnd - windowStart + 1)
	}
	return maxLen
}

func LongestSubstringWithKdistinct4(s string, K int) int{
	charFrequency := make(map[string]int, 0) 
	windowStart := 0 
	maxLen := 0 

	for windowEnd := 0; windowEnd < len(s); windowEnd++{
		rightChar := string(s[windowEnd]) 
		charFrequency[rightChar]++ 
		for len(charFrequency) > K{
			leftChar := string(s[windowStart]) 
			charFrequency[leftChar]-- 
			if charFrequency[leftChar] == 0{
				delete(charFrequency, leftChar)
			}
			windowStart++
		}
		maxLen  = Max(maxLen, windowEnd - windowStart + 1)
	}

	return maxLen 
}

func LongestSubstringWithKdistinct5(s string, K int) int{
	charFrequency := make(map[string]int, 0) 
	windowStart := 0 
	maxLen := 0 

	for windowEnd := 0; windowEnd < len(s); windowEnd++{
		rightChar := string(s[windowEnd]) 
		charFrequency[rightChar]++ 

		for len(charFrequency) > K{
			leftChar := string(s[windowStart]) 
			charFrequency[leftChar]-- 
			if charFrequency[leftChar] == 0{
				delete(charFrequency, leftChar)
			}
			windowStart++
		}
		maxLen = Max(maxLen, windowEnd - windowStart + 1)
	}

	return maxLen
}

func LongestSubstringWithKdistinct6(s string, K int) int{
	charFrequency := make(map[string]int, 0) 
	windowStart := 0 
	maxLen := 0 
	for windowEnd := 0; windowEnd < len(s); windowEnd++{
		rightChar := string(s[windowEnd])
		charFrequency[rightChar]++ 
		
		for len(charFrequency) > K{
			leftChar := string(s[windowStart])
			charFrequency[leftChar]-- 
			if charFrequency[leftChar] == 0{
				delete(charFrequency, leftChar)
			}
			windowStart++
		}
		maxLen = Max(maxLen, windowEnd - windowStart + 1)
	}
	return maxLen
}

func Max(a, b int) int{
	if a > b{
		return a 
	}
	return b 
}



func computeLenKeys(dict map[byte]int) int{
	numKeys := 0 
	for key, _ := range dict{
		numKeys += 1 
		log.Print(key)
	}
	return numKeys
}

func main(){
	s1 := "araaci" 
	d1 := 2 
	s2 := "araaci" 
	d2 := 1 
	s3 := "cbbebi" 
	d3 := 3
	fmt.Println(len(s1))
	fmt.Println(string(s1[0]))
	fmt.Printf("The length of the longest substring for string %s with window %d is :%d\n", s1, d1, LongestSubstringWithKdistinct4(s1, d1))
	fmt.Printf("The length of the longest substring for string %s with window %d is :%d\n", s2, d2, LongestSubstringWithKdistinct3(s2, d2))
	fmt.Printf("The length of the longest substring for string %s with window %d is :%d\n", s3, d3, LongestSubstringWithKdistinct5(s3, d3))
}