package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
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

func printHeadLines(r io.Reader, maxLines int) error {
	scanner := bufio.NewScanner(r)

	for i := 0; i < maxLines && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}

func processFile(filepath string, maxLines int) error {
	fi, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return errors.New(fmt.Sprintf("Must be a file: '%v'", filepath))
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	return printHeadLines(f, maxLines)
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
		err := printHeadLines(os.Stdin, *maxLines)
		if err != nil {
			die(err)
		}
	} else {
		filepath := args[0]
		err := processFile(filepath, *maxLines)
		if err != nil {
			die(err)
		}
	}
}
