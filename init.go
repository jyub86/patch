package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"time"
)

type InitData struct {
	PresetName         string    `json:"presetname"`
	Date               time.Time `json:"date"`
	AutosaveCheck      bool      `json:"autosavecheck"`
	Ocio               string    `json:"ocio"`
	OiioTool           string    `json:"oiiotool"`
	Ffmpeg             string    `json:"ffmpeg"`
	Ffprobe            string    `json:"ffprobe"`
	Splitter           string    `json:"splitter"`
	Shotname           string    `json:"shotname"`
	SeqSplitter        string    `json:"seqsplitter"`
	SeqPad             int       `json:"seqpad"`
	StartAt            int       `json:"startat"`
	ThumbCheck         bool      `json:"thumbcheck"`
	ThumbColorspaceIn  string    `json:"thumbcolorspacein"`
	ThumbColorspaceOut string    `json:"thumbcolorspaceout"`
	ThumbWidth         int       `json:"thumbwidth"`
	ThumbHeight        int       `json:"thumbheight"`
	ThumbDir           string    `json:"thumbdir"`
	ThumbName          string    `json:"thumbname"`
	ThumbExt           string    `json:"thumbext"`
	ThumbPath          string    `json:"thumbpath"`
	PlateCheck         bool      `json:"platecheck"`
	PlateColorspaceIn  string    `json:"platecolorspacein"`
	PlateColorspaceOut string    `json:"platecolorspaceout"`
	PlateWidth         int       `json:"platewidth"`
	PlateHeight        int       `json:"plateheight"`
	PlateDir           string    `json:"platedir"`
	PlateName          string    `json:"platename"`
	PlateExt           string    `json:"plateext"`
	PlatePath          string    `json:"platepath"`
	VideoCheck         bool      `json:"videocheck"`
	VideoColorspaceIn  string    `json:"videocolorspacein"`
	VideoColorspaceOut string    `json:"videocolorspaceout"`
	VideoWidth         int       `json:"videowidth"`
	VideoHeight        int       `json:"videoheight"`
	VideoFps           float64   `json:"videofps"`
	VideoCodec         string    `json:"videocodec"`
	VideoDir           string    `json:"videodir"`
	VideoName          string    `json:"videoname"`
	VideoExt           string    `json:"videoext"`
	VideoPath          string    `json:"videopath"`
	ProxyDir           string    `json:"proxydir"`
	ProxyName          string    `json:"proxyname"`
	ProxyExt           string    `json:"proxyext"`
	ProxyPath          string    `json:"proxypath"`
}

// Return json path in userfolder
func initFilePath() (initFile string, err error) {
	usr, err := user.Current()
	initFile = filepath.Join(usr.HomeDir, "patch", "init.json")
	return initFile, err
}

// check file exists
func Exists(name string) bool {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return true
}

// loadInit is get all init data from userfolder.
func LoadInits() (data []*InitData, err error) {
	initFile, err := initFilePath()
	if err != nil {
		return data, err
	}
	if !Exists(initFile) {
		return data, err
	}
	files, err := ioutil.ReadFile(initFile)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal([]byte(files), &data)
	if err != nil {
		return data, err
	}
	return data, err
}

// loadInit is get init data from userfolder.
func LoadInit() (*InitData, error) {
	data := []InitData{}
	value := InitData{}
	initFile, err := initFilePath()
	if err != nil {
		return &value, err
	}
	if !Exists(initFile) {
		return &value, err
	}
	files, err := ioutil.ReadFile(initFile)
	if err != nil {
		return &value, err
	}
	err = json.Unmarshal([]byte(files), &data)
	if err != nil {
		return &value, err
	}
	//Sort data by date
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.After(data[j].Date)
	})
	return &data[0], err
}

// saveInit is save init data to userfolder.
func SaveInit(value *InitData) error {
	data := []InitData{}
	newData := []InitData{}
	initFile, err := initFilePath()
	if err != nil {
		return err
	}
	if Exists(initFile) {
		files, err := ioutil.ReadFile(initFile)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(files), &data)
		if err != nil {
			return err
		}
	}
	value.Date = time.Now()
	newData = append(newData, *value)
	for _, item := range data {
		if item.PresetName != value.PresetName {
			newData = append(newData, item)
		}
	}
	// create directory if it doesn't
	err = os.MkdirAll(filepath.Dir(initFile), os.ModePerm)
	if err != nil {
		return err
	}
	files, err := json.MarshalIndent(newData, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(initFile, files, 0644)
	return err
}

// delete init data.
func DeleteInit(value *InitData) error {
	data := []InitData{}
	newData := []InitData{}
	initFile, err := initFilePath()
	if err != nil {
		return err
	}
	if Exists(initFile) {
		files, err := ioutil.ReadFile(initFile)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(files), &data)
		if err != nil {
			return err
		}
	}
	for _, item := range data {
		if item.PresetName != value.PresetName {
			newData = append(newData, item)
		}
	}
	// create directory if it doesn't
	err = os.MkdirAll(filepath.Dir(initFile), os.ModePerm)
	if err != nil {
		return err
	}
	files, err := json.MarshalIndent(newData, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(initFile, files, 0644)
	return err
}
