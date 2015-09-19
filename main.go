package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: yaml2json [file]\n")
		os.Exit(1)
	}

	file := os.Args[1]
	input, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	json, err := yaml.YAMLToJSON(input)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(json))
}
