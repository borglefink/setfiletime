// Copyright 2018 Erlend Johannessen.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	filename        = flag.String("f", "", "name of file to set date and time for")
	newdate         = flag.String("d", "", "date, given as yyyymmdd")
	newtime         = flag.String("t", "", "time, given as hhmmss")
	sourcefilename  = flag.String("sfn", "", "source file name, file to get the date and time from")
	newModifiedTime time.Time
)

// init
func init() {
	flag.Usage = usage
	flag.Parse()

	if len(*filename) == 0 {
		exit("File must be given")
	}

	if _, err := os.Stat(*filename); os.IsNotExist(err) {
		exit(fmt.Sprintf("File [%s] not found.\n", *filename))
	}

	if len(*sourcefilename) > 0 {
		if _, err := os.Stat(*sourcefilename); errors.Is(err, os.ErrNotExist) {
			exit(fmt.Sprintf("Source file [%s] does not exist", *sourcefilename))
		}
		if finfo, ferr := os.Stat(*sourcefilename); ferr != nil {
			exit(fmt.Sprintf("Error [%s] while getting source time for file [%s]", ferr.Error(), *sourcefilename))
		} else {
			newModifiedTime = finfo.ModTime().Local()
		}
		return
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
	var year = time.Now().Year()
	var yearString = fmt.Sprintf("-%v", year)
	fmt.Printf("\n%s (C) Copyright 2018%s Erlend Johannessen\n", strings.ToUpper(executableName), yearString)
	fmt.Printf("\nUsage: %s [options]", executableName)
	fmt.Printf("\n       NB: Option -sfn overrides all other parameters\n\n")
	flag.PrintDefaults()
}

// main
func main() {
	setDateTime(*filename, *newdate, *newtime, newModifiedTime)
}
