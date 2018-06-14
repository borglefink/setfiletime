package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// setDateTime sets given date and time on file with given filename
// if file is not found, the program will abort
func setDateTime(filename, newdate, newtime string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("File [%s] not found.\n", filename)
		os.Exit(1)
	}

	var newdatetime = getDateTime(newdate, newtime)

	if err := os.Chtimes(filename, newdatetime, newdatetime); err != nil {
		fmt.Printf("Could not set datetime [%v] for file [%s].\n", newdatetime, filename)
		os.Exit(1)
	}
}

// getDateTime returns a time.Time based on given date and time parameters
// if date or time is not formatted correctly, the program will abort
func getDateTime(newdate, newtime string) time.Time {
	var year, yerr = strconv.Atoi((newdate)[0:4])
	var month, merr = strconv.Atoi((newdate)[4:6])
	var day, derr = strconv.Atoi((newdate)[6:8])

	if yerr != nil || merr != nil || derr != nil {
		fmt.Printf("Error in date [%s].\n", newdate)
		os.Exit(1)
	}

	var hour, herr = strconv.Atoi((newtime)[0:2])
	var min, mierr = strconv.Atoi((newtime)[2:4])
	var sec, serr = strconv.Atoi((newtime)[4:6])

	if herr != nil || mierr != nil || serr != nil {
		fmt.Printf("Error in time [%s].\n", newtime)
		os.Exit(1)
	}

	return time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Local)
}
