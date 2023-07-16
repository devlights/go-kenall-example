package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

const (
	KEN_ALL = "utf_all.csv"
)

var (
	appLog = log.New(os.Stderr, "", 0)
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Fprintln(os.Stderr, "./kenall 検索文字列")
		os.Exit(-1)
	}

	re, err := regexp.Compile(os.Args[1])
	if err != nil {
		panic(err)
	}

	if err := run(re); err != nil {
		panic(err)
	}
}

func run(re *regexp.Regexp) error {
	var (
		file *os.File
		err  error
	)

	file, err = os.Open(KEN_ALL)
	if err != nil {
		return fmt.Errorf("err: os.Open (%w)", err)
	}
	defer file.Close()

	var (
		bufReader = bufio.NewReader(file)
		csvReader = csv.NewReader(bufReader)
		recCh     = make(chan []string)
		addrCh    = make(chan string)
		resultCh  = make(chan string)
		errCh     = make(chan error, 1)
	)

	go read(csvReader, recCh, errCh)
	go concat(recCh, addrCh, errCh)
	go find(re, addrCh, resultCh, errCh)

LOOP:
	for {
		select {
		case v, ok := <-resultCh:
			if !ok {
				break LOOP
			}
			appLog.Println(v)
		case e := <-errCh:
			return e
		}
	}

	return nil
}

func read(in *csv.Reader, out chan<- []string, errCh chan<- error) {
	defer close(out)
	for {
		record, err := in.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			errCh <- fmt.Errorf("err: csv.Read (%w)", err)
			return
		}

		out <- record
	}
}

func concat(in <-chan []string, out chan<- string, errCh chan<- error) {
	defer close(out)
	for v := range in {
		out <- fmt.Sprintf("%s%s%s", v[6], v[7], v[8])
	}
}

func find(re *regexp.Regexp, in <-chan string, out chan<- string, errCh chan<- error) {
	defer close(out)
	for v := range in {
		if re.FindString(v) != "" {
			out <- v
		}
	}
}
