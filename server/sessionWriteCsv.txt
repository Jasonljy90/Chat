package main

import (
	"encoding/csv"
	"os"
)

// Write unique session id and save to csv file
func sessionWriteCsv() {
	pairs := [][]string{}
	for key, value := range dbSessions {
		pairs = append(pairs, []string{key, value})
	}

	var records = make([][]string, len(pairs))
	for i := 0; i < len(records); i++ {
		var row = pairs[i]
		var copiedRow = make([]string, len(row))
		for j := 0; j < len(copiedRow); j++ {
			copiedRow[j] = row[j]
		}
		records[i] = copiedRow
	}

	file, err := os.Create("session.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range records {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

/*
https://golangcode.com/write-data-to-a-csv-file/
https://zetcode.com/golang/csv/
https://golang.org/src/encoding/csv/example_test.go
https://www.dotnetperls.com/convert-map-slice-go
*/
