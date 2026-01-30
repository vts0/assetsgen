package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := Run(Generate); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

func Run(generate func(source, output string) error) error {
	var (
		output  string
		source  string
		verbose bool
		help    bool
	)

	flag.StringVar(&output, "output", "", "Output file")
	flag.StringVar(&output, "o", "", "Output file (short)")

	flag.StringVar(&source, "source", "", "Source file or directory")
	flag.StringVar(&source, "s", "", "Source file (short)")

	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.BoolVar(&verbose, "verbose", false, "Verbose")

	flag.BoolVar(&help, "help", false, "Show help")

	flag.Parse()

	if help {
		flag.Usage()
		return nil
	}

	if output == "" {
		output = "output_gen.go"
	}

	var files []string
	if source == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		files, err = filepath.Glob(filepath.Join(cwd, "*"))
		if err != nil {
			return err
		}
	} else {
		files = []string{source}
	}

	if len(files) == 0 {
		if verbose {
			fmt.Println("No files to process")
		}
		return nil
	}

	for _, f := range files {
		if strings.HasSuffix(f, "_gen.go") {
			continue
		}

		if verbose {
			fmt.Printf("Processing %s -> %s\n", f, output)
		}

		if err := generate(f, output); err != nil {
			return fmt.Errorf("processing %s: %w", f, err)
		}
	}

	return nil
}
