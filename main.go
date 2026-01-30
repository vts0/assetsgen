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
	output := flag.String("output", "output_gen.go", "Output file")
	outputShort := flag.String("o", "output_gen.go", "Output file (short)")
	source := flag.String("source", "", "Source file or directory")
	sourceShort := flag.String("s", "", "Source file (short)")
	verbose := flag.Bool("v", false, "Verbose")
	verboseLong := flag.Bool("verbose", false, "Verbose")
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if *help {
		flag.Usage()
		return nil
	}

	out := firstNonEmpty(*output, *outputShort)
	src := firstNonEmpty(*source, *sourceShort)
	verb := *verbose || *verboseLong

	var files []string
	if src == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		files, err = filepath.Glob(filepath.Join(cwd, "*"))
		if err != nil {
			return err
		}
	} else {
		files = []string{src}
	}

	if len(files) == 0 {
		if verb {
			fmt.Println("No files to process")
		}
		return nil
	}

	for _, f := range files {
		if strings.HasSuffix(f, "_gen.go") {
			continue
		}

		if verb {
			fmt.Printf("Processing %s -> %s\n", f, out)
		}

		if err := generate(f, out); err != nil {
			return fmt.Errorf("processing %s: %w", f, err)
		}
	}

	return nil
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}
