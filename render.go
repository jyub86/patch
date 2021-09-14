package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"time"
)

func (job *Job) Render(ocio string) {
	job.Start = time.Now()
	dir := filepath.Dir(job.OutPath)
	if !Exists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			job.Errorlog = append(job.Errorlog, err.Error())
		}
	}
	log.Println(job.Cmd)
	cmd := exec.Command(job.Cmd[0], job.Cmd[1:]...)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "OCIO="+ocio)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		job.Errorlog = append(job.Errorlog, stderr.String())
		log.Println(stderr.String())
	}
	if out.String() != "" {
		log.Println(out.String())
	}

	job.End = time.Now()
	duration := time.Since(job.Start)
	job.Duration = fmt.Sprint(duration)
	SaveLog(job)
}

// save log data
func SaveLog(job *Job) {
	usr, err := user.Current()
	if err != nil {
		log.Println(err)
		return
	}
	logFile := filepath.Join(usr.HomeDir, "patch", "log", time.Now().Format(time.RFC3339Nano))
	// create directory if it doesn't
	dir := filepath.Dir(logFile)
	if !Exists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}
	}
	files, err := json.MarshalIndent(job, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	err = ioutil.WriteFile(logFile, files, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}
