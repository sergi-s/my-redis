package main

import (
	"bufio"
	"os"
)

var aofFile *os.File

func initAOF(filename string) error {
	var err error
	aofFile, err = os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return err
}

func appendToAOF(data interface{}) error {
	if aofFile == nil {
		return nil
	}

	writer := bufio.NewWriter(aofFile)
	if err := writeResp(writer, data); err != nil {
		return err
	}
	return writer.Flush()
}
