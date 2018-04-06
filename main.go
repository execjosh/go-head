package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

func die(err error) {
	printError(err)
	os.Exit(1)
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "USAGE: myhead [OPTION] <FILE>\n")
	flag.PrintDefaults()
}

func main() {
	maxLines := flag.Int("n", 10, "Max. number of lines to display")

	flag.Usage = printUsage

	flag.Parse()

	if *maxLines <= 0 {
		fmt.Fprintf(os.Stderr, "COUNT must be greater than 0!\n")
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	filepath := args[0]

	fi, err := os.Stat(filepath)
	if err != nil {
		die(err)
	}
	if fi.IsDir() {
		die(errors.New(fmt.Sprintf("Must be a file: '%v'", filepath)))
	}

	f, err := os.Open(filepath)
	if err != nil {
		die(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for i := 0; i < *maxLines && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		die(err)
	}
}
