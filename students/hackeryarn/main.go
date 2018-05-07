package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/hackeryarn/quiz/students/hackeryarn/quiz"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "f", "problems.csv", "CSV file with questions")
	flag.Parse()
}

func openFile(fileName string) io.ReadCloser {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Could not read file:", err)
	}

	return file
}

func main() {
	file := openFile(fileName)
	defer file.Close()

	quiz, err := quiz.New(file)
	if err != nil {
		log.Fatalln("Could not populate quiz:", err)
	}

	quiz.Run()
	quiz.PrintResults()
}
