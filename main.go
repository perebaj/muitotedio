package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type Problem struct {
	Question string
	Answer   int
}

var fName string

func main() {

	flag.StringVar(&fName, "csv", "problem.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	if fName == "" {
		fmt.Println("Please provide a csv file")
		os.Exit(1)
	}

	problems, err := parseCSV(fName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var userAnswer int
	for _, v := range problems {
		fmt.Println("What is the result of", v.Question, "?")
		fmt.Scan(&userAnswer)
		if userAnswer != v.Answer {
			fmt.Println("Incorrect")
		} else {
			fmt.Println("Correct")
		}
	}
	fmt.Println("You have finished the quiz")
}

func parseCSV(fName string) (problems []Problem, err error) {
	file, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, v := range records {
		i, err := strconv.Atoi(v[1])
		if err != nil {
			return nil, err
		}
		problems = append(problems, Problem{v[0], i})
	}

	return problems, nil
}
