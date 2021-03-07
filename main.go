package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const filepath = "./docs/b50f.csv"

type energyConsumpton struct {
	min, max float64 //in Watt
}

func main() {
	records, err := readData(filepath)
	if err != nil {
		log.Fatalln("Cannot read csv", err)
	}

	consumption := process(records)
	// fmt.Println("min is: ", consumption.min, "max: ", consumption.max)

	consumed, invoiced := calculateInvoice(consumption)
	total := math.Round(invoiced*100) / 100

	fmt.Printf("energy consumed: %v kwh, total invoiced: %v Eur", consumed/1000, total)

}

// use total consumption to calculate invoice amount price per kwh = 2 Eur
func calculateInvoice(consumption *energyConsumpton) (float64, float64) {

	//calculate consumption in watt and round to .00 because of float64
	sumWatt := consumption.max - consumption.min
	fmt.Println("sumWatt", sumWatt)

	sumRounded := math.Round(sumWatt*100) / 100
	fmt.Println("sumRounded", sumRounded)

	//calculate price in watt
	//price per khw = 2 Euro => price per watt = 2/1000 Eur = 0,002 Euro per Watt
	price := sumWatt * 0.002

	return sumWatt, price
}

//find mindate and max date with consumption
func process(rows [][]string) *energyConsumpton {
	var max float64
	min, err := strconv.ParseFloat(rows[0][8], 64)
	if err != nil {
		log.Fatalln("Cannot retrieve consumption", err)
	}

	for i := range rows {
		if rows[i][8] == "" {
			continue
		}
		cons, err := strconv.ParseFloat(rows[i][8], 64)
		if err != nil {
			log.Fatalln("Cannot retrieve consumption", err)
		}
		// fmt.Println(cons)

		if min > cons {
			min = cons
		}
		if max < cons {
			max = cons
		}
	}

	ec := energyConsumpton{
		min: min * 1000,
		max: max * 1000,
	}
	return &ec
}

//readData read data from csv
func readData(name string) ([][]string, error) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Cannot open %s", name)
	}
	defer f.Close()

	r := csv.NewReader(f)

	//skip first line header
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	//go through row find record[8] and last record[len(r)]-1 append to []string and then append []string to [][]string
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV")
	}

	return rows, nil
}
