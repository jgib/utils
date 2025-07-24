package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"time"
	"unicode/utf8"
)

func Er(err error) {
	var now = time.Now()
	pc, file, line, ok := runtime.Caller(1)

	if !ok {
		log.Fatal("error getting caller function\n")
		os.Exit(3)
	}

	if err != nil {
		//log.Fatalf("| %v | %s:%d | Error: %v\n", runtime.FuncForPC(pc).Name(), path.Base(file), line, err)
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%02d:%02d:%02d.%03d | %v | %s:%d | Error: %v\n", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000, runtime.FuncForPC(pc).Name(), path.Base(file), line, err))
		os.Exit(2)
	}
}

func Debug(text string, verbose ...bool) {
	now := time.Now()
	if len(verbose) != 0 && verbose[0] {
		pc, file, line, ok := runtime.Caller(1)

		if !ok {
			log.Fatalf("error getting caller function\n")
			os.Exit(4)
		}

		//fmt.Fprintf(os.Stderr, "%s | %v | %s:%d | Debug: %s\n", currTime.Format("2006/01/02 15:04:05"), runtime.FuncForPC(pc).Name(), path.Base(file), line, text)
		fmt.Fprintf(os.Stderr, fmt.Sprintf("%02d:%02d:%02d.%03d | %v | %s:%d | Debug: %v\n", now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1000000, runtime.FuncForPC(pc).Name(), path.Base(file), line, text))
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
	//output += fmt.Sprint("\n")

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
	} else {
		output = fmt.Sprintf("(%T) %v", *input, *input)
	}

	return output
}

func ValidateIP(input string) bool {
	regex := regexp.MustCompile(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) == 5 {
		for i := 1; i < 5; i++ {
			octet, err := strconv.Atoi(matches[i])
			if err != nil || octet < 0 || octet > 255 {
				return false
			}
		}
		return true
	}
	return false
}

func Ip2Uint32(input string) (uint32, error) {
	if ValidateIP(input) {
		ip := uint32(0)
		regex := regexp.MustCompile(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`)
		matches := regex.FindStringSubmatch(input)

		if len(matches) == 5 {
			for i := 1; i < 5; i++ {
				octet, err := strconv.ParseUint(matches[i], 10, 32)
				Er(err)

				ip <<= 8
				ip += uint32(octet)
			}
			return ip, nil
		}
	}
	return 0, fmt.Errorf("Invalid IP Address.")
}

func ValidatePort(input string) bool {
	regex := regexp.MustCompile(`^(\d{1,5})$`)
	matches := regex.FindStringSubmatch(input)

	if len(matches) == 2 {
		_, err := strconv.ParseUint(input, 10, 16)
		if err != nil {
			return false
		} else {
			return true
		}
	}
	return false
}

func Port2Uint16(input string) (uint16, error) {
	if ValidatePort(input) {
		regex := regexp.MustCompile(`^(\d{1,5})$`)
		matches := regex.FindStringSubmatch(input)

		if len(matches) == 2 {
			port, err := strconv.ParseUint(input, 10, 16)
			Er(err)
			return uint16(port), nil
		}
	}
	return 0, fmt.Errorf("Invalid Port.")
}
