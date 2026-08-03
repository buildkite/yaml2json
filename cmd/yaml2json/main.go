package main

import (
	"fmt"
	"os"

	"github.com/buildkite/yaml2json"
	"github.com/ghodss/yaml"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("yaml2json v%s\nUsage: yaml2json [file]\n", yaml2json.Version())
		os.Exit(1)
	}

	file := os.Args[1]
	input, err := os.ReadFile(file)
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
