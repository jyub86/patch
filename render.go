package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
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
}
