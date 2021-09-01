package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type InitData struct {
	AutosaveCheck      bool    `json:"autosavecheck"`
	Ocio               string  `json:"ocio"`
	OiioTool           string  `json:"oiiotool"`
	Ffmpeg             string  `json:"ffmpeg"`
	Ffprobe            string  `json:"ffprobe"`
	Splitter           string  `json:"splitter"`
	Shotname           string  `json:"shotname"`
	SeqSplitter        string  `json:"seqsplitter"`
	SeqPad             int     `json:"seqpad"`
	StartAt            int     `json:"startat"`
	ThumbCheck         bool    `json:"thumbcheck"`
	ThumbColorspaceIn  string  `json:"thumbcolorspacein"`
	ThumbColorspaceOut string  `json:"thumbcolorspaceout"`
	ThumbWidth         int     `json:"thumbwidth"`
	ThumbHeight        int     `json:"thumbheight"`
	ThumbDir           string  `json:"thumbdir"`
	ThumbName          string  `json:"thumbname"`
	ThumbExt           string  `json:"thumbext"`
	ThumbPath          string  `json:"thumbpath"`
	PlateCheck         bool    `json:"platecheck"`
	PlateColorspaceIn  string  `json:"platecolorspacein"`
	PlateColorspaceOut string  `json:"platecolorspaceout"`
	PlateWidth         int     `json:"platewidth"`
	PlateHeight        int     `json:"plateheight"`
	PlateDir           string  `json:"platedir"`
	PlateName          string  `json:"platename"`
	PlateExt           string  `json:"plateext"`
	PlatePath          string  `json:"platepath"`
	VideoCheck         bool    `json:"videocheck"`
	VideoColorspaceIn  string  `json:"videocolorspacein"`
	VideoColorspaceOut string  `json:"videocolorspaceout"`
	VideoWidth         int     `json:"videowidth"`
	VideoHeight        int     `json:"videoheight"`
	VideoFps           float64 `json:"videofps"`
	VideoCodec         string  `json:"videocodec"`
	VideoDir           string  `json:"videodir"`
	VideoName          string  `json:"videoname"`
	VideoExt           string  `json:"videoext"`
	VideoPath          string  `json:"videopath"`
	ProxyDir           string  `json:"proxydir"`
	ProxyName          string  `json:"proxyname"`
	ProxyExt           string  `json:"proxyext"`
	ProxyPath          string  `json:"proxypath"`
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

// loadInit is get init data from userfolder.
func LoadInit() (data *InitData, err error) {
	initFile, err := initFilePath()
	if err != nil {
		return data, err
	}
	if Exists(initFile) {
		data, err = readInit(initFile)
		return data, err
	}
	// create directory if it doesn't
	err = os.MkdirAll(filepath.Dir(initFile), os.ModePerm)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read data from userfolder
func readInit(initFile string) (*InitData, error) {
	data := InitData{}
	files, err := ioutil.ReadFile(initFile)
	if err != nil {
		return &data, err
	}
	err = json.Unmarshal([]byte(files), &data)
	return &data, err
}

// save json data
func MakeInit(data *InitData) error {
	initFile, err := initFilePath()
	if err != nil {
		return err
	}
	files, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(initFile, files, 0644)
	return err
}
