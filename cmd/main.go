package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type quizData struct {
	question string
	answer   string
}

func myScanner() string {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		return input
	}
	return "No input given!"
}

func main() {
	// open CSV
	csvFile, err := os.Open("../problems.csv")
	fmt.Println("Opening CSV File...")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	fmt.Println(csvFile)
	// defer csv.Close()

	// get contents of CSV
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("Question %v: %v= ?", 1, csvLines[0])
	// input := myScanner()
	// fmt.Println(input)
	problems := []quizData{}
	for i, line := range csvLines {
		problems = append(problems, quizData{line[0], line[1]})
		fmt.Println(i, problems[i].question)
		// problems[i].answer = line[1]
		// fmt.Println(line[1])
	}

	input := myScanner()
	if input == problems[0].answer {
		fmt.Println("YES")
	}
	// fmt.Scanf(x[0], input)
	// if input == x[1] {
	// 	fmt.Println("Correct!")
	// }

	// print out questions

	// CSV to struct array
	// Get Stdin input and compare to each Answer

}
