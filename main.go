package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
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

func processUrl(url string) error {
	h := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := h.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("%v %s\n", resp.Proto, resp.Status)

	keys := make([]string, 0, len(resp.Header))
	for key := range resp.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val := resp.Header[key][0] // Only take first value of each header
		fmt.Printf("%s: %s\n", key, val)
	}

	return nil
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
		var err error
		hadError := false
		showsHeader := len(args) > 1
		for idx, filepath := range args {
			if showsHeader {
				if idx >= 1 {
					fmt.Println()
				}
				fmt.Printf("==> %s <==\n", filepath)
			}

			if strings.HasPrefix(filepath, "http://") || strings.HasPrefix(filepath, "https://") {
				err = processUrl(filepath)
			} else {
				err = processFile(filepath, *maxLines)
			}

			if err != nil {
				hadError = true
				printError(err)
			}
		}
		if hadError {
			os.Exit(1)
		}
	}
}
