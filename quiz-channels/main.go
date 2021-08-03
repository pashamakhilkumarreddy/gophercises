package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fmt.Println("Welcome to Quiz App")
	csvFileName := flag.String("csv", "problems.csv", "a .csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		logError(fmt.Sprintf("Failed to open the .csv file %v %s", *csvFileName, err))
	}
	data := csv.NewReader(file)
	lines, err := data.ReadAll()
	if err != nil {
		logError(fmt.Sprintf("Failed to parse the given .csv file %s", err))
	}
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	totalProblems := len(problems)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, totalProblems)
			os.Exit(0)
		case answer := <-answerCh:
			if answer == p.answer {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, totalProblems)
	fmt.Println(problems)
}

func parseLines(lines [][]string) []problem {
	var res []problem
	for _, line := range lines {
		res = append(res, problem{
			question: line[0],
			answer:   line[1],
		})
	}
	return res
}

func logError(msg string) {
	log.Fatal(msg)
	os.Exit(1)
}
