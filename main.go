package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mikunup/go-createfile/term"
)

const FileNameLayout = "20060102"

var (
	MovingDate     string
	CreateCount    int
	CreateFilePath string
)

const InitCount = 10

func main() {
	os.Exit(run())
}

func run() int {
	err := initArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: initArgs, %s", err)
		return 1
	}
	err = createFiles()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: createFiles, %s", err)
		return 1
	}
	return 0
}

func initArgs() error {
	flag.StringVar(&MovingDate, "t", term.Daily, "term")
	flag.StringVar(&MovingDate, "term", term.Daily, "term")
	flag.IntVar(&CreateCount, "c", InitCount, "count")
	flag.IntVar(&CreateCount, "count", InitCount, "count")
	flag.StringVar(&CreateFilePath, "p", "", "filepath")
	flag.StringVar(&CreateFilePath, "path", "", "filepath")
	flag.Parse()

	if CreateFilePath == "" {
		var err error
		CreateFilePath, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("todo context :%s", err)
		}
	}

	if !fileTermIs() {
		return errors.New("term Please choose from one of 'd', 'm', 'y'")
	}

	_, err := os.Stat(CreateFilePath)
	if err != nil {
		return fmt.Errorf("todo context :%s", err)
	}

	return nil
}

func createFiles() error {
	for date := range nextTerms(0, CreateCount) {
		name := fmt.Sprintf("%s.txt", date.Format(FileNameLayout))
		fp := filepath.Join(CreateFilePath, name)
		file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return fmt.Errorf("todo context :%s", err)
		}
		defer file.Close()
		os.Chtimes(fp, date, date)
	}
	return nil
}

// nextTermsは次に生成するファイルのtime.Timeを返却する
func nextTerms(start int, end int) <-chan time.Time {
	tt := term.NewTimeTerm(time.Now(), MovingDate)
	ch := make(chan time.Time)
	go func(ch chan time.Time) {
		for i := start; i <= end; i++ {
			ch <- tt.NextTerm()
		}
		close(ch)
	}(ch)
	return ch
}

func fileTermIs() bool {
	for _, t := range term.FileTerms {
		if MovingDate == t {
			return true
		}
	}
	return false
}
