package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"
	"unicode/utf8"
)

func Er(err error) {
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("error getting caller function\n")
		os.Exit(3)
	}

	if err != nil {
		log.Fatalf("| %v | %s:%d | Error: %v\n", runtime.FuncForPC(pc).Name(), path.Base(file), line, err)
	}
}

func Debug(text string, verbose ...bool) {
	currTime := time.Now()
	if len(verbose) != 0 && verbose[0] {
		pc, file, line, ok := runtime.Caller(1)

		if !ok {
			log.Fatalf("error getting caller function\n")
			os.Exit(4)
		}

		fmt.Fprintf(os.Stderr, "%s | %v | %s:%d | Debug: %s\n", currTime.Format("2006/01/02 15:04:05"), runtime.FuncForPC(pc).Name(), path.Base(file), line, text)
	}
}

func GetArgs(n int) ([]string, error) {
	// n is minimum number of args required, 0 is no minimum
	// Returns string slice of arguments
	args := os.Args[1:]

	if n != 0 && len(args) < n {
		return nil, fmt.Errorf("incorrect arguments given, expecting %d and received %d", n, len(args))
	}

	return args, nil
}

func WalkByteSlice(input []byte) string {
	var output string

	if len(input) == 0 {
		return ""
	}

	for n := 0; n < len(input); n++ {
		if n%8 == 0 {
			output += fmt.Sprint("    ")
		}

		output += fmt.Sprintf("%02X ", input[n])

		if (n+1)%32 == 0 {
			output += fmt.Sprint("\n")
		}
	}
	output += fmt.Sprint("\n")

	return output
}

func IsPipe() (bool, error) {
	// Return true if pipe
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return fileInfo.Mode()&os.ModeCharDevice == 0, nil
}

func PP[T any](input *T) string {
	var output string
	dataType := fmt.Sprintf("%T", *input)

	//	fmt.Printf("TYPE: %s\n", dataType)
	//	fmt.Printf("LEN: %d, %v\n", len(dataType), *input)

	if utf8.RuneCountInString(dataType) >= 2 && dataType[:2] == "[]" {
		output = fmt.Sprintf("(%T) %v", *input, *input)
	} else if utf8.RuneCountInString(dataType) >= 3 && dataType[:3] == "map" {
		output = fmt.Sprintf("(%T) %v", *input, *input)
	} else if (utf8.RuneCountInString(dataType) >= 3 && dataType[:3] == "int") || (utf8.RuneCountInString(dataType) >= 4 && dataType[:4] == "uint") || (utf8.RuneCountInString(dataType) >= 5 && dataType[:5] == "float") || (utf8.RuneCountInString(dataType) >= 7 && dataType[:7] == "complex") {
		output = fmt.Sprintf("(%T) %d", *input, *input)
	} else {

	}

	return output
}
