package log

import (
	"log"
	"os"
)

var CommonLog *log.Logger
var ErrorLog *log.Logger

func GenerateLog() {
	f, err := os.OpenFile("wedding.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	CommonLog = log.New(f, "Common Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(f, "Error Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
