package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	var err error
	var input []byte

	if len(os.Args) == 1 {
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Failed to read stdin: %v\n", err)
			os.Exit(1)
		}
	} else if len(os.Args) == 2 {
		input, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Failed to read file: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Usage: yaml2json [file]\n")
		os.Exit(1)
	}

	json, err := yaml.YAMLToJSON(input)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(json))
}
