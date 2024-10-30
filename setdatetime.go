package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

// setDateTime sets given date and time on file with given filename
// if file is not found, the program will abort
func setDateTime(filename, newdate, newtime string, newModifiedTime time.Time) {
	var newdatetime time.Time

	if newModifiedTime.IsZero() {
		newdatetime = getDateTime(newdate, newtime)
	} else {
		newdatetime = newModifiedTime
	}

	if err := os.Chtimes(filename, newdatetime, newdatetime); err != nil {
		exit(fmt.Sprintf("Could not set datetime [%v] for file [%s].\n", newdatetime, filename))
	}

	if err := setFileCreationTime(filename, newdatetime); err != nil {
		exit(fmt.Sprintf("Could not set creation datetime [%v] for file [%s].\n", newdatetime, filename))
	}
}

// setFileCreationTime sets the creation time of the specified file.
func setFileCreationTime(path string, creationTime time.Time) error {
	// Open the file handle with FILE_WRITE_ATTRIBUTES access
	fileHandle, err := syscall.CreateFile(
		syscall.StringToUTF16Ptr(path),
		syscall.FILE_WRITE_ATTRIBUTES,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer syscall.CloseHandle(fileHandle)

	// Convert Go time to Windows FILETIME
	ft := syscall.NsecToFiletime(creationTime.UnixNano())

	// Set file time attributes
	err = syscall.SetFileTime(fileHandle, &ft, nil, nil)
	if err != nil {
		return fmt.Errorf("could not set file creation time: %w", err)
	}
	return nil
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
