## Description

*setfiletime* is a tiny command line utility that takes a file name, date and time as parameters,
and set the file's access and modification times to the given date and time.

Date must be given in the format yyyymmdd, time in the format hhmmss (24 hour clock).


## Usage

```
setfiletime -f "file name" -d 20180615 -t 190500
```


## Install

Clone the repository into your GOPATH somewhere and do a **go install**.