package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"
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
	args := os.Args[1:]

	if n != 0 && len(args) < n {
		return (nil, fmt.Errorf("incorrect arguments given, expecting %d and received %d", n, len(args)))
	}

	return (args, nil)
}
