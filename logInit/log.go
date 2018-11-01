package logInit

import (

	//"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var Info, Error *log.Logger

var infoFile, errorFile *os.File

func init() {
	gopath := os.Getenv("GOPATH")

	infoFilePath := filepath.Join(gopath, "/src/Agenda-Golang/data/infolog.log")
	infoFile, err := os.OpenFile(infoFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {

	}

	errorFilePath := filepath.Join(gopath, "/src/Agenda-Golang/data/errorlog.log")
	errorFile, err := os.OpenFile(errorFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {

	}

	Info = log.New(infoFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Free : close log file
func Free() {
	infoFile.Close()
	errorFile.Close()
}
