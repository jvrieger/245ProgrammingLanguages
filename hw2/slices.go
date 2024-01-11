//Julia Rieger
//HW2
//Sept 24 2023
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type date struct {
	mon int
	day int
	yr  int
}

func (this date) String() string {
	return fmt.Sprintf("%02d-%v-%v", this.mon, this.day, this.yr)
}

type launch struct {
	identifier      string
	something        float64
	launchdate date
	vehicle string
	variant string
	flight string
	mission string
	launchsite string
	launchpad string
	apogee int
	category string
}

func (this launch) String() string {
	return fmt.Sprintf("%v %v %v %.2f", this.launchdate, this.vehicle, this.category, (this.something + float64(this.apogee)))
}

const identifierindex = 0
const somethingindex = 1
const yrindex = 2
const monindex = 3
const dayindex = 4
const vehicleindex = 5
const variantindex = 6
const flightindex = 7
const missionindex = 8
const launchsiteindex = 9
const launchpadindex = 10
const apogeeindex = 11
const categoryindex = 12

func main() {
	file, err := os.Open("hl.csv")
	if err != nil {
		fmt.Printf("Error while reading the file %v\n", err)
	}

	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	var launches []launch

	//initialize each record as a launch in launces slice 
	//i.e. populate launches after processing csv
	for _, eachrecord := range records {

		//create and initialize temp instance to add to launches
		var temp launch

		temp.identifier = eachrecord[identifierindex]
		smth, err := strconv.ParseFloat(eachrecord[somethingindex], 64)
		if err != nil {fmt.Println("Error parsing float64")}
		temp.something = smth

		mon, err := strconv.Atoi(eachrecord[monindex])
		if err != nil {fmt.Println("Error parsing int")}
		dy, err := strconv.Atoi(eachrecord[dayindex])
		if err != nil {fmt.Println("Error parsing int")}
		yr, err := strconv.Atoi(eachrecord[yrindex])
		if err != nil {fmt.Println("Error parsing int")}
		temp.launchdate = date{mon, dy, yr}

		temp.vehicle = eachrecord[vehicleindex]
		temp.variant = eachrecord[variantindex]
		temp.flight = eachrecord[flightindex]
		temp.mission = eachrecord[missionindex]
		temp.launchsite = eachrecord[launchsiteindex]
		temp.launchpad = eachrecord[launchpadindex]
		apo, err := strconv.Atoi(eachrecord[apogeeindex])
		if err != nil {fmt.Println("Error parsing int")}
		temp.apogee = apo
		temp.category = eachrecord[categoryindex]

		launches = append(launches, temp)
	}

	first5 := launches[:5]
	fmt.Printf("First 5 Launches\n")
	//print each launch
	for _, eachlaunch := range first5 {
		fmt.Println(eachlaunch)
	}

	fmt.Printf("\nLast 5 Launches\n")
	last5 := launches[len(launches)-5:]
	//print each launch
	for _, eachlaunch := range last5 {
		fmt.Println(eachlaunch)
	}

	fmt.Printf("\nFirst 3 Last 2 Launches\n")
	first3 := launches[:3]
	last2 := launches[len(launches)-2:]
	first3last2 := append(first3, last2...)
	//print each launch
	for _, eachlaunch := range first3last2 {
		fmt.Println(eachlaunch)
	}
}
