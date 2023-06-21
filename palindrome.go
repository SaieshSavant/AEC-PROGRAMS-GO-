package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}
	return true
}

func largestpalindromeproduct() (int, int, int) {
	largestpalindrome := 0
	var mul1, mul2 int
	for i := 999; i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if product < largestpalindrome {
				break
			}
			if isPalindrome(product) && product > largestpalindrome {
				largestpalindrome = product
				mul1 = i
				mul2 = j
			}
		}
	}

	return largestpalindrome, mul1, mul2

}

func main() {
	result, mul1, mul2 := largestpalindromeproduct()
	fmt.Printf("Largest palindrome number is :%d\n", result)
	fmt.Printf("Multiplicands are %d and %d \n", mul1, mul2)
	//Largest palindrome number is :906609
	//Multiplicands are 993 and 913 

}
