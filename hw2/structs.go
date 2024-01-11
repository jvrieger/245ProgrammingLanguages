//Julia Rieger
//HW2
//Sept 24 2023
package main

import (
	"encoding/csv"
	"fmt"
	"os"
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

func main() {
	//initialize dates
	jan1157 := date{1, 11, 1957}
	mar2449 := date{3, 24, 1949}
	mar2618 := date{3, 26, 2018}

	//initialize launches
	polaris := launch{"1957-A166",2435849.50,jan1157,"Polaris FTV-4","","Thrust reversal test","","PM","",10,"Test"}
	aerobee := launch{"1949-A07",2433000.13,mar2449,"Aerobee RTV-N-8","","A12","","USN49A","",5,"Ionosphe"}
	nudol := launch{"2018-S18",2458203.68,mar2618,"Nudol","","","GVM","GIK-1","39A",500,""}

	//test if three instances are equal
	fmt.Println(polaris == aerobee)
	fmt.Println(aerobee == nudol)
	fmt.Println(polaris == nudol)

	//tests for copy
	copy := polaris
	fmt.Println(polaris == copy)
	copy.launchdate = mar2449
	fmt.Println(copy)
	fmt.Println(polaris == copy)
}
