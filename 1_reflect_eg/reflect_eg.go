package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func get_file_info(filepath string) ([]byte, error) {

	file, err := os.Open(filepath)
	if err != nil {
		panic("Error while opening file.")
	}
	fileBytes, err := ioutil.ReadAll(file)
	return fileBytes, err

}

func main() {

	var resultFileAJson map[string]interface{}
	var resultFileBJson map[string]interface{}

	// Read File A Contents
	content_a, err_a := get_file_info("resources/file_a.json")
	if err_a != nil {
		fmt.Println("Error seen while reading file_a.json")
		return
	}
	// Read File B
	content_b, err_b := get_file_info("resources/file_b.json")
	if err_b != nil {
		fmt.Println("Error seen while reading file_b.json")
		return
	}
	json.Unmarshal(content_a, &resultFileAJson)
	json.Unmarshal(content_b, &resultFileBJson)

	// Compare the unmarshaled objects using reflect.DeepEqual()
	if reflect.DeepEqual(resultFileAJson, resultFileBJson) {
		fmt.Println("Objects are deeply equal.")
	} else {
		fmt.Println("Objects are not deeply equal.")
	}
	// fmt.Println(string(content_a))
	// fmt.Println(string(content_b))

}
