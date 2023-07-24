package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// "github.com/wI2L/jsondiff"
	"github.com/hgsgtk/jsoncmp"
)

func get_file_info(filepath string) (string, error) {

	file, err := os.Open(filepath)
	if err != nil {
		panic("Error while opening file.")
	}
	fileBytes, err := ioutil.ReadAll(file)
	return string(fileBytes), err

}

type stubT struct{}

func (t *stubT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func main() {

	// var resultFileAJson map[string]interface{}
	// var resultFileBJson map[string]interface{}

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
	// err1 := json.Unmarshal(content_a, &resultFileAJson)
	// if err1 != nil {
	// 	return
	// }
	// err2 := json.Unmarshal(content_b, &resultFileBJson)
	// if err2 != nil {
	// 	return
	// }

	// Compare the unmarshaled objects using reflect.DeepEqual()
	// patch, err := jsondiff.Compare(resultFileAJson, resultFileBJson)
	// ans := jsoncmp.Diff(string(content_a), string(content_b))
	var t stubT
	if diff := jsoncmp.Diff(string(content_a), string(content_b)); diff != "" {
		t.Errorf("jsoncmp.Diff() got differs: (-got +want)\n%s", diff)
	}

	// if err != nil {
	// 	panic("This is also failing..")
	// }
	// fmt.Println(ans)
	// b, err := json.MarshalIndent(patch, "", "    ")
	// if err != nil {
	// 	panic("This is also failing..")
	// }
	// os.Stdout.Write(b)

}
