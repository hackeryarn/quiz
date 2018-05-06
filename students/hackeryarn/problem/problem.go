package problem

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Problem is the basic structure for each quiz question
type Problem struct {
	question string
	answer   string
}

// Check validates the answer against a problem's answer
func (p Problem) Check(answer string) bool {
	return p.answer == answer
}

// Ask asks the question and returns in the answer
func (p Problem) Ask() string {
	fmt.Printf("%s: ", p.question)

	reader := bufio.NewReader(os.Stdin)
	answer := readAnswer(reader)

	return cleanString(answer)
}

func cleanString(s string) string {
	return strings.TrimSpace(s)
}

func readAnswer(r *bufio.Reader) string {
	answer, err := r.ReadString('\n')
	if err != nil {
		log.Fatalln("Could not read user input:", err)
	}

	return answer
}

// New creates a new problem from a csv based record
func New(record []string) Problem {
	question := record[0]
	answer := record[1]

	return Problem{
		question: question,
		answer:   answer,
	}
}
