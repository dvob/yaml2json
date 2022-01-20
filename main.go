package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "no files provided")
		os.Exit(0)
	}

	data := map[string]interface{}{}

	for _, file := range files {
		var (
			err      error
			input    io.Reader
			filename string
		)
		if file == "-" {
			filename = "standard input"
			input = os.Stdin
		} else {
			filename = file
			input, err = os.Open(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
		content, err := io.ReadAll(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read '%s': %s\n", filename, err)
			os.Exit(1)
		}

		err = yaml.Unmarshal(content, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to parse '%s': %s\n", filename, err)
			os.Exit(1)

		}
	}

	output, err := json.MarshalIndent(&data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode: %s", err)
		os.Exit(1)
	}
	_, _ = os.Stdout.Write(output)
	fmt.Println()
}
