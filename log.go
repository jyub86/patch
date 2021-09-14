package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
)

func logDir() (targetDir string, err error) {
	usr, err := user.Current()
	if err != nil {
		return targetDir, err
	}
	targetDir = filepath.Join(usr.HomeDir, "patch", "log")
	return targetDir, err
}

func Readlog(logFile string) (job *Job, err error) {
	data := Job{}
	files, err := ioutil.ReadFile(logFile)
	if err != nil {
		return &data, err
	}
	err = json.Unmarshal([]byte(files), &data)
	return &data, err
}

func LogData() ([]Job, error) {
	logs := make([]Job, 0)
	targetDir, err := logDir()
	if err != nil {
		return logs, err
	}
	if !Exists(targetDir) {
		return logs, err
	}
	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return logs, err
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		logFile := filepath.Join(targetDir, file.Name())
		data, err := Readlog(logFile)
		if err != nil {
			return logs, err
		}
		logs = append(logs, *data)
	}
	return logs, err
}
