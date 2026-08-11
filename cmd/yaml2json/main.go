package main

import (
	"fmt"
	"os"

	"github.com/buildkite/yaml2json"
	"github.com/ghodss/yaml"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "yaml2json v%s\nUsage: yaml2json [file]\n", yaml2json.Version())
		os.Exit(1)
	}

	file := os.Args[1]
	if file == "--version" {
		fmt.Printf("yaml2json v%s\n", yaml2json.Version())
		return
	}

	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
		os.Exit(1)
	}

	json, err := yaml.YAMLToJSON(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(json))
}
