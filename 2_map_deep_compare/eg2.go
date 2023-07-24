package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func readJsonFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic("Unable to Open Json file")
	}
	return ioutil.ReadAll(file)
	// content, err := ioutil.ReadAll(file)
	// return content, err

}

func main() {
	var result1, result2 map[string]interface{}
	file_a, err_a := readJsonFile("resources/file_a.json")
	// fmt.Println(string(file_a))
	if err_a != nil {
		fmt.Println("Reading File A failed")
		return
	}
	file_b, err_b := readJsonFile("resources/file_a.json")
	if err_b != nil {
		fmt.Println("Reading File B failed")
		return
	}
	// Now comparison part after Unmarshalling
	err1 := json.Unmarshal(file_a, &result1)
	if err1 != nil {
		return
	}

	err2 := json.Unmarshal(file_b, &result2)
	if err2 != nil {
		return
	}
	fmt.Println(reflect.DeepEqual(result1, result2))
	// ans := string(result1) == string(result2)
}

// This is confusing as this example is same as 1_refect_eg/reflect_eg.go
// BUt this gives different result and other gives different.
// Same json files are used.
