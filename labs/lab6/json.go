package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsonMap map[string]interface{}
type jsonSlice []jsonMap

type dob struct {
	Year  int32  `json:"year"`
	Month string `json:"month"`
	Day   int32  `json:"day"`
}

type students struct {
	FirstName          string  `json:"firstName"`
	LastName           string  `json:"lastName"`
	Id                 string  `json:"id"`
	Age                int32   `json:"age"`
	DateOfBirth        dob     `json:"dateOfBirth"`
	FavoriteRealNumber float32 `json:"favoriteRealNumber"`
}

func main() {
	//read entire contents from json file into an array of bytes
	jsonFile, err := os.Open("students.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened students.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()

	var res students
	json.Unmarshal(byteValue, &res)

	//now have read in json file

	//marshall to JSON
	stu_obj := students{FirstName: "Julia"}
	stu, _ := json.Marshal(stu_obj)
	fmt.Println(string(stu))

	//unmarshalling into structs
	bytes := []byte(stu_obj.FirstName)
	json.Unmarshal(bytes, &res)

	fmt.Println(res.FirstName)

	data, _ := json.Marshal(res)
	fmt.Println(string(data))
}
