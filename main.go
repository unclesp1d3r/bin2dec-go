package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/erikgeiser/promptkit/textinput"
)

func main() {
	input := textinput.New("Please enter a binary string, up to one byte long.")
	input.InitialValue = "00000001"
	input.Placeholder = "The value cannot be empty"
	input.Validate = func(v string) bool {
		matched, err := regexp.MatchString("^[01]+$", v)

		return err != nil || matched
	}

	input.CharLimit = 8

	value, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}

	if i, err := strconv.ParseInt(value, 2, 9); err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(i)
		os.Exit(0)
	}
}
