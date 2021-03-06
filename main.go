package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type str []string

const filename = "./docs/b50f.csv"

func main() {
	//Open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Unable to open csv file", err)
	}
	defer file.Close()

	//Parse
	r := csv.NewReader(file)

	//iterate records

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//print columns "i"
		fmt.Printf("%s\n", record[8])

	}
}
