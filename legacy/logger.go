package main

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	headlessMode bool
	logFileMode  bool
	logFile      *os.File
}

func (l Logger) Setup() error {
	writers := []io.Writer{}

	if l.headlessMode {
		writers = append(writers, os.Stdout)
	}

	if l.logFileMode {
		file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}

		l.logFile = file
		writers = append(writers, l.logFile)
	}

	writer := io.MultiWriter(writers...)
	log.SetOutput(writer)

	return nil
}

func (l Logger) Close() {
	l.logFile.Close()
}
