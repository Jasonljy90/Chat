package main

import (
	"encoding/csv"
	"log"
	"os"
)

var tableUser [4][8]string

// Write patient's bookings and save to csv file
func writeCsv() {
	var records = make([][]string, len(tableUser))
	for i := 0; i < len(records); i++ {
		var row = tableUser[i]
		var copiedRow = make([]string, len(row))
		for j := 0; j < len(copiedRow); j++ {
			copiedRow[j] = row[j]
		}
		records[i] = copiedRow
	}

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range records {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

/*
https://golangcode.com/write-data-to-a-csv-file/
https://zetcode.com/golang/csv/
https://golang.org/src/encoding/csv/example_test.go
*/
