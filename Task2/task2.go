package main

import (
	"fmt"
	"strings"
	"unicode"
)
func clean_str(str string) string{
	cleaned := ""
	lower := strings.ToLower((str))
	for _ , char := range(lower){
				if unicode.IsDigit(char) || unicode.IsLetter(char)|| unicode.IsSpace(char){
				cleaned +=string(char)
				}

	}
	return cleaned
}
func wordCount(word string) map[string] int{
	new := clean_str(word)
	split := strings.Split(new, " ")
	dic := map[string] int{}

	for _ , word := range(split){
		dic[word]++
	}
	return dic

}

func palindrome(str string) bool{
	temp := clean_str(str)
	new_str := strings.ReplaceAll(temp, " ", "")

	n := len(new_str)
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