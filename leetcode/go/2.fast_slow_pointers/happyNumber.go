// Problem statement 
// Any number will be called a happy number, if after repeatedly 
// replacing it with a number equal to the sum of the square 
// of all digits, leads us to number 1. All other (not-happy) numbers 
// will never reach 1. Instead, they will be stuck in a cycle of 
// numbers which does not include 1.
// Input 23 
// Output: true (23 is a happy number) 
// Explanation: Here are the steps to find out that 23 is a happy number 
// 2**2 + 3**2 = 4 + 12 = 13 
// 1**2 + 3 **2 = 1 + 9 = 10 
// 1**2 + 0**2 = 1 + 0 = 1

// Example 2: 
// Input: 12 
// Output: false (12 is not a happy number) 
// Explanation: Here are the steps: 
// 1**2 + 2**2 = 1+ 4= 5 
// 5**2 = 25 
// 2**2 + 5**2 = 29 
// 2**2 + 9**2 = 85 
// 8**2 + 5**2 = 145 
// 1**2 + 4**2 + 5**2 = 1 + 16 + 25 = 42 
// 4**2 + 2**2 = 16 + 4 = 20 
// 2**2 + 0**2 = 4 + 0 = 4 
// 4**2 = 16 
// 1**2 + 6**2 = 1 + 36 = 37 
// 3**2 + 7**2 = 9 + 49 = 58 
// 5**2 + 8**2 = 25 + 64 = 89 
// The last step leads us back to step  as the number becomes equal to 89, this means 
// that we can never reach 1, therefore, 12 is not a happy number 

// Solution 
// The process defined above to find out if a number is a happy number or not, always ends in a cycle
// If the number is a happy number, the process will be stuck in a cycle on number '1', and if the 
// number is not a happy number then the process will be stuck in a cycle with a set of numbers. 
// We can use the fast and slow pointer method to find a cycle among a set of elements. As we 
// have described above, each number will definitely have a cycle. 
// Therefore, we will use the same fast & slow pointer strategy to find the cycle and once
// the cycle is found we will se if the cycle is stuck on number '1' to find out if the number 
// is happy or not.
package main 

import "fmt" 

// Time Complexity 
// The time compelxity of the algorithm is difficult to determine. 
// However we know the fact that all unhappy numbers eventually 
// get stuck in the cycle 4 -> 16 -> 37 -> 58 -> 89 -> 145 -> 42 -> 20 -> 4 
// This sequence's behavior tell us two things:
// 
// 1. If the numbr N is less than or equal to 1000, the we reach the cycle or 1 
//  at most 1001 steps. 
// For N > 1000, suppose the number has M digits and the next number is N1. From 
// the 

func IsHappy(num int)bool{
	slow, fast := num, num 

	for{
		slow = findSquareSum(slow) 
		fast = findSquareSum(findSquareSum(fast)) 
		if !(slow != fast){
			break 
		} 
	} 
	return slow == 1; 
} 

func findSquareSum(num int) int{
	sum := 0 
	var digit int 
	for num > 0{
		digit = num % 10 
		sum += digit * digit 
		num /= 10 
	} 
	return sum 
}

func main(){
	fmt.Printf("Is 23 a happy number? %v: \n", IsHappy(23)) 
	fmt.Printf("Is 12 a happy number? %v: \n", IsHappy(12))
} 

