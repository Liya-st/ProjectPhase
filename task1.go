package main

import (
	"fmt"
)

type SubjectScore struct {
	name  string
	score float64
}

func getInputs() (string, int, []SubjectScore) {
	fmt.Print("Enter your name: ")
	name := ""
	fmt.Scanln(&name)

	numOfSub := 0
	fmt.Print("Enter the number of subjects you took: ")
	fmt.Scanln(&numOfSub)

	subjects := make([]SubjectScore, 0, numOfSub)

	for i := 0; i < numOfSub; i++ {
		sub_name := ""
		score := 0.0

		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scanln(&sub_name)

		fmt.Printf("Enter the score for subject %d: ", i+1)
		fmt.Scanln(&score)

		if score < 0 || score > 100 {
			fmt.Printf("Invalid entry. The score must be between 0 and 100. Your entry was %v\n", score)
			i--
			continue
		}
		subjects = append(subjects, SubjectScore{sub_name, score})
	}
	return name, numOfSub, subjects
}

func getAverage(subjects []SubjectScore, numOfSub int) float64 {
	sum := 0.0
	for _, sub := range subjects {
		sum += sub.score
	}
	return sum / float64(numOfSub)
}

func main() {
	name, numOfSub, subjects := getInputs()
	avg := getAverage(subjects, numOfSub)
	fmt.Println()
	fmt.Printf("\t %s's Score \n", name)
	fmt.Println("---------------------------------")
	fmt.Printf("| Subject \t | Score\t|\n")
		fmt.Println("---------------------------------")

	for _, sub := range subjects{
		fmt.Printf("| %s\t\t | %v\t\t|\n", sub.name, sub.score)
	}
		fmt.Println("---------------------------------")

	fmt.Printf("| Total \t | %.2f\t|\n", avg)
	fmt.Println("---------------------------------")
	fmt.Printf("%s's average score is %.2f\n", name, avg)
}
