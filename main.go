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

func printHeadBytes(r io.Reader, maxBytes int) error {
	// FIXME smart loop reading small buffer each time
	b := make([]byte, maxBytes)

	rb, err := r.Read(b)
	if err != nil {
		return err
	}

	os.Stdout.Write(b[:rb])

	return nil
}

func processFile(filepath string, maxCount int, headFunc HeadFunc) error {
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

	return headFunc(f, maxCount)
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

type HeadFunc func(r io.Reader, maxLines int) error

func main() {
	maxLines := flag.Int("n", 10, "Max. number of lines to display (incompatible with -c)")
	maxBytes := flag.Int("c", -1, "Max. number of bytes to display (incompatible with -n)")

	flag.Usage = printUsage

	flag.Parse()

	maxLinesWasSet := false
	maxBytesWasSet := false

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "n":
			maxLinesWasSet = true
		case "c":
			maxBytesWasSet = true
		}
	})

	if maxLinesWasSet && maxBytesWasSet {
		die(errors.New("Cannot combine -n and -c"))
	}

	var maxCount int
	var headFunc HeadFunc
	if maxBytesWasSet {
		maxCount = *maxBytes
		headFunc = printHeadBytes
	} else {
		maxCount = *maxLines
		headFunc = printHeadLines
	}

	if maxCount <= 0 {
		fmt.Fprintf(os.Stderr, "COUNT must be greater than 0!\n")
		flag.Usage()
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) <= 0 {
		err := headFunc(os.Stdin, maxCount)
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
				err = processFile(filepath, maxCount, headFunc)
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
