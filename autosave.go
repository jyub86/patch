package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func saveFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	saveFile := filepath.Join(usr.HomeDir, "patch", "autosave.json")
	return saveFile, nil
}

func Autosave(data *[]Item) error {
	saveFile, err := saveFilePath()
	if err != nil {
		return err
	}
	// create directory if it doesn't
	dir := filepath.Dir(saveFile)
	if !Exists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	files, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(saveFile, files, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadData() (*[]Item, error) {
	data := []Item{}
	saveFile, err := saveFilePath()
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadFile(saveFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(files), &data)
	return &data, err
}
