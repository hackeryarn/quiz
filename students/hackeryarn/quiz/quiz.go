package quiz

import (
	"encoding/csv"
	"fmt"
	"gophercises/quiz/students/hackeryarn/problem"
	"io"
)

// Quiz is used to run the quiz exercise
type Quiz struct {
	problems  []problem.Problem
	correct   int
	incorrect int
}

func (q *Quiz) updateResults(correct bool) {
	if correct {
		q.correct = q.correct + 1
	} else {
		q.incorrect = q.incorrect + 1
	}

}

// Run executes the quiz tracking the score
func (q *Quiz) Run() {
	for _, problem := range q.problems {
		answer := problem.Ask()
		correct := problem.Check(answer)
		q.updateResults(correct)
	}
}

// PrintResults outputs the results of the quiz
func (q *Quiz) PrintResults() {
	fmt.Println("Congratulations on completing the quiz!")
	fmt.Printf("You got %d answers correct, and %d answers incorrect\n",
		q.correct, q.incorrect)

}

func readProblem(reader *csv.Reader) (p problem.Problem, err error) {
	record, err := reader.Read()
	if err != nil {
		return p, err
	}
	p = problem.New(record)
	return p, err
}

func (q *Quiz) addProblem(p problem.Problem) {
	q.problems = append(q.problems, p)
}

// New takes a problem file and returns a new quiz
func New(problemsFile io.Reader) (quiz Quiz, err error) {
	reader := csv.NewReader(problemsFile)
	for {
		p, err := readProblem(reader)
		if err == io.EOF {
			return quiz, nil
		} else if err != nil {
			return quiz, err
		}

		quiz.addProblem(p)
	}
}
