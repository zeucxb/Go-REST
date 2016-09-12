package logger

import (
	"io"
	"log"
	"os"
)

var (
	// Info - Info logger
	Info *log.Logger

	// Warning - Warning errors logger
	Warning *log.Logger

	// Error - Critical errors logger
	Error *log.Logger
)

func init() {
	fileErr, err := os.OpenFile("errors.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalln("Failed to open error log")
	}

	fileInfo, err := os.OpenFile("debug.txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalln("Failed to open debug log")
	}

	dfLog := log.Ldate | log.Ltime | log.Lshortfile

	Info = log.New(io.MultiWriter(fileInfo, os.Stdout),
		"INFO: ",
		dfLog)

	Warning = log.New(io.MultiWriter(fileErr, os.Stderr),
		"WARNING: ",
		dfLog)

	Error = log.New(io.MultiWriter(fileErr, os.Stderr),
		"ERROR: ",
		dfLog)
}
