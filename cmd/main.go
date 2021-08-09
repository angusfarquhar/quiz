package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

type QuizData struct {
	question string
	answer   string
}

type Score struct {
	amount    int
	correct   int
	incorrect int
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
	f, err := os.Open("../problems.csv")
	fmt.Println("Opening CSV File...")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	// defer csv.Close()

	// get contents of CSV
	csvLines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	problems := []QuizData{}
	for i, line := range csvLines {
		problems = append(problems, QuizData{line[0], line[1]})
		fmt.Println(i, problems[i].question, problems[i].answer)
	}

	quizScore := Score{0, 0, 0}

	for _, problem := range problems {
		fmt.Println("Question: ", problem.question)
		quizScore.amount++
		input := myScanner()
		if problem.answer == input {
			fmt.Println("Correct!")
			quizScore.correct++
			continue
		} else {
			fmt.Println("Incorrect!")
			quizScore.incorrect++
		}
	}

	fmt.Println("End Quiz!")
	fmt.Println("Final Score:", quizScore.correct, "/", quizScore.amount)

}
