// Copyright 2018 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	filename = flag.String("f", "", "name of file to set date and time for")
	newdate  = flag.String("d", "", "date, given as yyyymmdd")
	newtime  = flag.String("t", "", "time, given as hhmmss")
)

// init
func init() {
	flag.Usage = usage
	flag.Parse()

	if len(*filename) == 0 {
		exit("File must be given")
	}

	if len(*newdate) != 8 {
		exit("Date should be formatted as yyyymmdd")
	}

	if len(*newtime) != 6 {
		exit("Time should be formatted as hhmmss")
	}

	if len(*newdate) == 0 || len(*newtime) == 0 {
		exit("Date and time should both be given")
	}
}

// exit
func exit(msg string) {
	fmt.Println(msg)
	//usage()
	os.Exit(1)
}

// usage
func usage() {
	var executableName = filepath.Base(os.Args[0])
	fmt.Printf("\n%s (C) Copyright 2018 Erlend Johannessen\n", strings.ToUpper(executableName))
	fmt.Printf("\nUsage: %s [options]  \n", executableName)
	flag.PrintDefaults()
}

// main
func main() {
	setDateTime(*filename, *newdate, *newtime)
}
