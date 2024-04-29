package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	logfile *os.File
}

func NewLogger(path string, target string) (Logger, error) {
	newPath := filepath.Join(".", path)
	err := os.MkdirAll(newPath, os.ModePerm)
	if err != nil {
		return Logger{}, err
	}

	path = fmt.Sprintf("%s/%s.log", path, target)
	logfile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return Logger{}, err
	}
	log.SetOutput(logfile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	return Logger{logfile: logfile}, err
}

func (l *Logger) GetLogfile() *os.File {
	return l.logfile
}
