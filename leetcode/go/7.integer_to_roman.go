package main

import (
	"fmt"
	"strings"
	"testing"
)

func intToRoman(num int) string{
	integer := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var sb strings.Builder 
	for i := 0; i < 13; i++{
		if num/integer[i] > 0{
			count := num / integer[i] 
			for count > 0{
				sb.WriteString(roman[i])
				count--
			}
			num = num % integer[i]
		}
	}
	return sb.String()
}

func TestIntToRoman(t testing.T){
	expected := "XCD"
	intNum := 390 
	actual := intToRoman(intNum) 
	if actual != expected{
		t.Errorf("intToRoman function result expecting %s but got %s", expected, actual)
	}
}

func main(){
	expected := "XCD" 
	intNum := 390 
	actual := intToRoman(intNum) 
	if actual != expected{
		fmt.Printf("intToRoman function result expecting %s but got %s\n", expected, actual)
	} else{
		fmt.Println("Test passed")
	}
}