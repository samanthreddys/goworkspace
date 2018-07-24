package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "csv file in format of Q & A")
	timeLimit := flag.Int("limit", 30, "Time limit for quiz in seconds")
	flag.Parse()
	//_ = csvFileName
	file, err := os.Open(*csvFileName)
	if err != nil {
		Exit(fmt.Sprintf("Failed to open csv file:%s \n", *csvFileName))

	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		Exit("Failed to parse the csv file")
	}
	problems := ParseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	//fmt.Println(problems)
problemLoop:
	for i, p := range problems {
		fmt.Printf("Problem #%d :%s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string

			fmt.Scanf("%s \n", &answer)
			answerCh <- answer

		}()
		select {
		case <-timer.C:
			fmt.Printf("\n")
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
				//fmt.Println("Answer is correct!")
			}

		}

	}
	fmt.Printf("you scored %d out of %d: \n", correct, len(problems))
}

// ParseLines function
func ParseLines(lines [][]string) []Problem {

	ret := make([]Problem, len(lines))
	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret

}

//Problem Struct
type Problem struct {
	q string
	a string
}

//Exit function to exit
func Exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
