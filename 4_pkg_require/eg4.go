package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
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
	var t *testing.T

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
	err1 := json.Unmarshal(content_a, &resultFileAJson)
	if err1 != nil {
		return
	}
	err2 := json.Unmarshal(content_b, &resultFileBJson)
	if err2 != nil {
		return
	}

	// Compare the unmarshaled objects using reflect.DeepEqual()
	require.JSONEq(t, resultFileAJson, resultFileBJson)
}

// This is not helpful, checking other packages.
