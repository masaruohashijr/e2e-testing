package api

import (
	"io/ioutil"
	"os"
)

func beginLog(level string) int64 {
	var filePath = NameLogFile(level, *logPtr)
	f, _ := os.Open(filePath)
	stat, _ := f.Stat()
	filesize := stat.Size()
	return filesize
}

func endLog(level string) int64 {
	var filePath = NameLogFile(level, *logPtr)
	f, _ := os.Open(filePath)
	stat, _ := f.Stat()
	filesize := stat.Size()
	return filesize
}

func getLogs(level string, _beginMark, _endMark int64) string {
	var filePath = NameLogFile(level, *logPtr)
	file, _ := os.Open(filePath)
	data, _ := ioutil.ReadAll(file)
	return string(data[_beginMark:_endMark])
}
