package main

import (
	"gophercises/quiz/students/hackeryarn/quiz"
	"io"
	"log"
	"os"
)

func openFile(fileName string) io.ReadCloser {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln("Could not read file:", err)
	}

	return file
}

func main() {
	file := openFile("problems.csv")
	defer file.Close()

	quiz, err := quiz.New(file)
	if err != nil {
		log.Fatalln("Could not populate quiz:", err)
	}

	quiz.Run()
	quiz.PrintResults()
}
