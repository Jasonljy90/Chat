package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	Trace   *log.Logger // just about anything
	Info    *log.Logger // important information
	Warning *log.Logger // be concerned
	Error   *log.Logger // Critical problems
)

// Initialist logger
func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Error.Println("Error opening amyQueue.csv")
		fmt.Println("Failed to open log file: ", err)
	}

	// customized loggers
	Trace = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info :", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}
