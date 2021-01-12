package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

//Input ... structure for data input to the program
type Input struct {
	Environment string `json:"environment"`
	App         string `json:"app"`
}

func main() {
	var input Input
	fmt.Println("Args:", os.Args[1])
	//var temp string = `{"environment":"development","app":"golang-tests"}`

	err := json.Unmarshal([]byte(os.Args[1]), &input)
	if err != nil {
		fmt.Println("Unmarshal ERROR:", err)
	}
	fmt.Println("Input:", input)

	output, err := ingestInput(input)
	if err != nil {
		fmt.Println("main() Ending Error:", err)
	}

	fmt.Println(output)
} //End main()

func ingestInput(input Input) (string, error) {
	var err error
	var output string
	env := input.Environment
	app := input.App

	//Check Environment
	if len(env) <= 0 {
		err = errors.New("Please provide a valid environment value")
	} else {
		output = output + "Current Environment is " + env + "\n"
	}

	//Check App
	if len(app) <= 0 {
		err = errors.New("Please provide a valid application value")
	} else {
		output = output + "Current Environment is " + app
	}

	return output, err
} //End ingestInput()
