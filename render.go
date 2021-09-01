package main

import (
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

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		job.Errorlog = append(job.Errorlog, err.Error())
	}
	if err := cmd.Wait(); err != nil {
		job.Errorlog = append(job.Errorlog, err.Error())
	}
	job.End = time.Now()
	duration := time.Since(job.Start)
	job.Duration = fmt.Sprint(duration)
}
