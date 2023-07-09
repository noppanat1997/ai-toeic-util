package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/noppanat1997/ai-toeic-util/handler/importExam"
)

func main() {
	var importExamInput []importExam.Exam

	b, _ := ioutil.ReadFile("./input/exam1.json")

	json.Unmarshal(b, &importExamInput)

	output := importExam.ExecuteMapping(importExamInput)
	fmt.Println(output)

	file, _ := json.MarshalIndent(output, "", "  ")
	_ = ioutil.WriteFile("./output/exam1.json", file, 0644)
}
