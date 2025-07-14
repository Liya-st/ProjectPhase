package main

import (
	"fmt"
	"strings"
)

func wordCount(word string) map[string] int{
	new := strings.ToLower((word))
	split := strings.Split(new, " ")
	dic := map[string] int{}

	for _ , word := range(split){
		dic[word]++
	}
	return dic

}

func palindrome(str string) bool{
	new_str := strings.ToLower(str)
	n := len(str)
	for i:= 0; i < n/2; i++ {
		if new_str[i] != new_str[n-i-1] {
			return false
		}


	}
	return true
}

func main(){
	str1 := "kayak"
	str2 := "racecar"
	str3 := "two man cell"
	str4 := "Everywhere that I go, everywhere that I be"
	str5 := "toronto"
	fmt.Printf("The word count of the string %s is %v\n",str4, wordCount(str4))
	fmt.Printf("The word count of the string %s is %v\n", str3, wordCount(str3))
	fmt.Printf("Is the word  %s a palindrome? %v\n", str1, palindrome(str1))
	fmt.Printf("Is the word  %s a palindrome? %v\n", str2, palindrome((str2)))
	fmt.Printf("Is the word  %s a palindrome? %v\n", str5, palindrome((str5)))


}