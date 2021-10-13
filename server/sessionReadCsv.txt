package main

import (
	"encoding/csv"
	"log"
	"os"
)

// Read unique session id from saved csv file
func sessionReadCsv() [][]string {
	f, err := os.Open("session.csv")
	if err != nil {
		Error.Println("Error opening sessionRead.csv")
		log.Fatalf("Cannot open file: %s\n", err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ','
	rows, err := r.ReadAll()
	if err != nil {
		Error.Println("Cannot read CSV data")
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}
