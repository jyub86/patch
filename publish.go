package main

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"trimmer.io/go-timecode/timecode"
)

type Item struct {
	ID          int     `json:"id"`
	Path        string  `json:"path"`
	FrameIn     int     `json:"framein"`
	FrameOut    int     `json:"frameout"`
	FrameRange  int     `json:"framerange"`
	Seqs        []int   `json:"seqs"`
	TimecodeIn  string  `json:"timecodein"`
	TimecodeOut string  `json:"timecodeout"`
	Pad         int     `json:"pad"`
	Width       int     `json:"width"`
	Height      int     `json:"height"`
	Ext         string  `json:"ext"`
	Fps         float64 `json:"fps"`
	Codec       string  `json:"codec"`
	Shotname    string  `json:"shotname"`
	TrimIn      int     `json:"trimin"`
	TrimOut     int     `json:"trimout"`
	TrimInTc    string  `json:"trimintc"`
	TrimOutTc   string  `json:"trimouttc"`
	ColorIn     string  `json:"colorin"`
	ColorOut    string  `json:"colorout"`
	ReWidth     int     `json:"rewidth"`
	ReHeight    int     `json:"reheight"`
	Pub         bool    `json:"pub"`
	Log         string  `json:"log"`
}

type Job struct {
	Shotname string    `json:"shotname"`
	Task     string    `json:"task"`
	SubTask  string    `json:"subtask"`
	OutPath  string    `json:"outpath"`
	Cmd      []string  `json:"cmd"`
	Cmdstr   string    `json:"cmdstr"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Duration string    `json:"duration"`
	Errorlog []string  `json:"errorlog"`
}

func thumbnailJob(initData *InitData, item Item) []Job {
	jobs := make([]Job, 0)
	shotname := item.Shotname
	thumbPath := filepath.Join(initData.ThumbDir, initData.ThumbName+initData.ThumbExt)
	thumbPath = replaceName(initData, shotname, thumbPath)
	if strContains(item.Ext, []string{".mp4", ".mov"}) {
		// video input
		tmpPath := filepath.Join(initData.ThumbDir, initData.ThumbName+"_tmp"+initData.ThumbExt)
		tmpPath = replaceName(initData, shotname, tmpPath)
		// make tmp
		cmd := []string{}
		cmd = append(cmd, initData.Ffmpeg, "-loglevel", "error", "-y",
			"-r", strconv.FormatFloat(initData.VideoFps, 'f', 5, 64),
			"-i", item.Path, "-vframes", "1", tmpPath)
		job := Job{Shotname: shotname, Task: "thumbnail", SubTask: "thumbnail_tmp", OutPath: tmpPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
		// convert to thumbnail
		cmd = []string{}
		cmd = append(cmd, initData.OiioTool, tmpPath)
		if initData.ThumbColorspaceIn != "" && initData.ThumbColorspaceOut != "" {
			cmd = append(cmd, "--colorconvert", initData.ThumbColorspaceIn, initData.ThumbColorspaceOut)
		}
		if initData.VideoWidth != 0 && initData.VideoHeight != 0 {
			cmd = append(cmd, "-resize", fmt.Sprintf("%sx%s", strconv.Itoa(initData.VideoWidth), strconv.Itoa(initData.VideoHeight)))
		}
		cmd = append(cmd, "-o", thumbPath)
		job = Job{Shotname: shotname, Task: "thumbnail", SubTask: "thumbnail", OutPath: thumbPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
	} else {
		// sequence input
		cmd := []string{}
		cmd = append(cmd, initData.OiioTool, item.Path, "--frames", StrPad(item.FrameIn, item.Pad))
		if initData.ThumbColorspaceIn != "" && initData.ThumbColorspaceOut != "" {
			cmd = append(cmd, "--colorconvert", initData.ThumbColorspaceIn, initData.ThumbColorspaceOut)
		}
		if initData.VideoWidth != 0 && initData.VideoHeight != 0 {
			cmd = append(cmd, "-resize", fmt.Sprintf("%sx%s", strconv.Itoa(initData.VideoWidth), strconv.Itoa(initData.VideoHeight)))
		}
		cmd = append(cmd, "-o", thumbPath)
		job := Job{Shotname: shotname, Task: "thumbnail", SubTask: "thumbnail", OutPath: thumbPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
	}
	return jobs
}

func plateJob(initData *InitData, item Item) []Job {
	jobs := make([]Job, 0)
	shotname := item.Shotname
	// set in, out frame
	item.timecodeToFrame()
	in := item.FrameIn
	out := item.FrameOut
	if item.TrimIn != 0 && item.TrimOut != 0 {
		in = item.TrimIn
		out = item.TrimOut
	}
	// make pad
	pad := initData.SeqSplitter
	if initData.StartAt != 0 {
		pad = pad + strconv.Itoa(initData.StartAt) + "-" + strconv.Itoa(initData.StartAt+out-in) + "%0" + strconv.Itoa(initData.SeqPad) + "d" // .1001-1010%04d
	} else {
		pad = pad + "%0" + strconv.Itoa(initData.SeqPad) + "d" //.%04d
	}
	platePath := filepath.Join(initData.PlateDir, initData.PlateName+pad+initData.PlateExt)
	platePath = replaceName(initData, shotname, platePath)

	if strContains(item.Ext, []string{".mp4", ".mov"}) { // video input
		// make tmp path
		tmpPath := filepath.Join(initData.PlateDir+"_tmp", initData.PlateName+initData.SeqSplitter+"%0"+strconv.Itoa(initData.SeqPad)+"d"+".png")
		tmpPath = replaceName(initData, shotname, tmpPath)
		cmd := []string{}
		cmd = append(cmd, initData.Ffmpeg, "-loglevel", "error", "-y",
			"-r", strconv.FormatFloat(initData.VideoFps, 'f', 5, 64),
			"-i", item.Path, tmpPath)
		job := Job{Shotname: shotname, Task: "plate", SubTask: "plate_tmp", OutPath: tmpPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
		// convert to plate
		cmd = []string{}
		cmd = append(cmd, initData.OiioTool, tmpPath)
		if item.TrimIn != 0 && item.TrimOut != 0 {
			cmd = append(cmd, "--frames", fmt.Sprintf("%s-%s", StrPad(item.TrimIn, initData.SeqPad), StrPad(item.TrimOut, initData.SeqPad)))
		}
		if item.ColorIn != "" && item.ColorOut != "" {
			cmd = append(cmd, "--colorconvert", item.ColorIn, item.ColorOut)
		}
		if item.ReWidth != 0 && item.ReHeight != 0 {
			cmd = append(cmd, "-resize", fmt.Sprintf("%sx%s", strconv.Itoa(item.ReWidth), strconv.Itoa(item.ReHeight)))
		}
		cmd = append(cmd, "-o", platePath)
		job = Job{Shotname: shotname, Task: "plate", SubTask: "plate", OutPath: platePath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
	} else { // sequence input
		cmd := []string{}
		cmd = append(cmd, initData.OiioTool, item.Path)
		if item.TrimIn != 0 && item.TrimOut != 0 {
			cmd = append(cmd, "--frames", fmt.Sprintf("%s-%s", StrPad(item.TrimIn, item.Pad), StrPad(item.TrimOut, item.Pad)))
		}
		if item.ColorIn != "" && item.ColorOut != "" {
			cmd = append(cmd, "--colorconvert", item.ColorIn, item.ColorOut)
		}
		if item.ReWidth != 0 && item.ReHeight != 0 {
			cmd = append(cmd, "-resize", fmt.Sprintf("%sx%s", strconv.Itoa(item.ReWidth), strconv.Itoa(item.ReHeight)))
		}
		cmd = append(cmd, "-o", platePath)
		job := Job{Shotname: shotname, Task: "plate", SubTask: "plate", OutPath: platePath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
	}
	return jobs
}

func videoJob(initData *InitData, item Item) []Job {
	jobs := make([]Job, 0)
	shotname := item.Shotname
	// set in, out frame
	item.timecodeToFrame()
	in := item.FrameIn
	out := item.FrameOut
	if item.TrimIn != 0 && item.TrimOut != 0 {
		in = item.TrimIn
		out = item.TrimOut
	}
	pad := initData.SeqSplitter
	if initData.StartAt != 0 {
		pad = pad + strconv.Itoa(initData.StartAt) + "-" + strconv.Itoa(initData.StartAt+out-in) + "%0" + strconv.Itoa(initData.SeqPad) + "d"
	} else {
		pad = pad + "%0" + strconv.Itoa(initData.SeqPad) + "d"
	}
	proxyPath := filepath.Join(initData.ProxyDir, initData.ProxyName+pad+initData.ProxyExt)
	proxyPath = replaceName(initData, shotname, proxyPath)
	proxySimplePath := filepath.Join(initData.ProxyDir, initData.ProxyName+initData.SeqSplitter+"%0"+strconv.Itoa(initData.SeqPad)+"d"+initData.ProxyExt)
	proxySimplePath = replaceName(initData, shotname, proxySimplePath)
	videoPath := filepath.Join(initData.VideoDir, initData.VideoName+initData.VideoExt)
	videoPath = replaceName(initData, shotname, videoPath)

	if strContains(item.Ext, []string{".mp4", ".mov"}) { // video input
		tmpPath := filepath.Join(initData.PlateDir+"_tmp", initData.PlateName+initData.SeqSplitter+"%0"+strconv.Itoa(initData.SeqPad)+"d"+".png")
		tmpPath = replaceName(initData, shotname, tmpPath)
		if !initData.PlateCheck { // plate tmp files not exist
			// make tmp
			cmd := []string{}
			cmd = append(cmd, initData.Ffmpeg, "-loglevel", "error", "-y",
				"-r", strconv.FormatFloat(initData.VideoFps, 'f', 5, 64), "-i", item.Path, tmpPath)
			job := Job{Shotname: shotname, Task: "video", SubTask: "plate_tmp", OutPath: tmpPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
			jobs = append(jobs, job)
		}
		// proxy image
		start := StrPad(item.FrameIn, initData.SeqPad)
		cmd := []string{}
		cmd = append(cmd, initData.OiioTool, tmpPath)
		if item.TrimIn != 0 && item.TrimOut != 0 {
			start = StrPad(item.TrimIn, initData.SeqPad)
			cmd = append(cmd, "--frames", fmt.Sprintf("%s-%s", StrPad(item.TrimIn, initData.SeqPad), StrPad(item.TrimOut, initData.SeqPad)))
		}
		if initData.StartAt != 0 {
			start = StrPad(initData.StartAt, initData.SeqPad)
		}
		if initData.VideoColorspaceIn != "" && initData.VideoColorspaceOut != "" {
			cmd = append(cmd, "--colorconvert", initData.VideoColorspaceIn, initData.VideoColorspaceOut)
		}
		if initData.VideoWidth != 0 && initData.VideoHeight != 0 {
			cmd = append(cmd, "-resize", fmt.Sprintf("%sx%s", strconv.Itoa(initData.VideoWidth), strconv.Itoa(initData.VideoHeight)))
		}
		cmd = append(cmd, "-o", proxyPath)
		job := Job{Shotname: shotname, Task: "video", SubTask: "proxy", OutPath: proxyPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
		// Video
		cmd = []string{}
		cmd = append(cmd, initData.Ffmpeg, "-loglevel", "error", "-y",
			"-r", strconv.FormatFloat(initData.VideoFps, 'f', 5, 64),
			"-start_number", start, "-i", proxySimplePath)
		switch initData.VideoCodec {
		case "h264":
			cmd = append(cmd, "-c:v", "libx264", "-pix_fmt", "yuv420p")
		case "proresLT":
			cmd = append(cmd, "-c:v", "prores_ks", "-profile:v", "1", "-vendor", "ap10", "pix_fmt", "yuv422p10le")
		case "proresHQ":
			cmd = append(cmd, "-c:v", "prores_ks", "-profile:v", "3", "-vendor", "ap10", "pix_fmt", "yuv422p10le")
		case "proresXQ":
			cmd = append(cmd, "-c:v", "prores_ks", "-profile:v", "5", "-vendor", "ap10", "-bits_per_mb", "8000", "pix_fmt", "yuva444p10le")
		}
		cmd = append(cmd, "-vf", "pad=ceil(iw/2)*2:ceil(ih/2)*2", videoPath)
		job = Job{Shotname: shotname, Task: "video", SubTask: "video", OutPath: videoPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
	} else { // sequence input
		start := StrPad(item.FrameIn, item.Pad)
		cmd := []string{}
		cmd = append(cmd, initData.OiioTool, item.Path)
		if item.TrimIn != 0 && item.TrimOut != 0 {
			start = StrPad(item.TrimIn, item.Pad)
			cmd = append(cmd, "--frames", fmt.Sprintf("%s-%s", StrPad(item.TrimIn, item.Pad), StrPad(item.TrimOut, item.Pad)))
		}
		if initData.StartAt != 0 {
			start = StrPad(initData.StartAt, initData.SeqPad)
		}
		if initData.VideoColorspaceIn != "" && initData.VideoColorspaceOut != "" {
			cmd = append(cmd, "--colorconvert", initData.VideoColorspaceIn, initData.VideoColorspaceOut)
		}
		if initData.VideoWidth != 0 && initData.VideoHeight != 0 {
			cmd = append(cmd, "-resize", fmt.Sprintf("%sx%s", strconv.Itoa(initData.VideoWidth), strconv.Itoa(initData.VideoHeight)))
		}
		cmd = append(cmd, "-o", proxyPath)
		job := Job{Shotname: shotname, Task: "video", SubTask: "proxy", OutPath: proxyPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
		// make Video
		cmd = []string{}
		cmd = append(cmd, initData.Ffmpeg, "-loglevel", "error", "-y",
			"-r", strconv.FormatFloat(initData.VideoFps, 'f', 5, 64),
			"-start_number", start, "-i", proxySimplePath)
		switch initData.VideoCodec {
		case "h264":
			cmd = append(cmd, "-c:v", "libx264", "-pix_fmt", "yuv420p")
		case "proresLT":
			cmd = append(cmd, "-c:v", "prores_ks", "-profile:v", "1", "-vendor", "ap10", "pix_fmt", "yuv422p10le")
		case "proresHQ":
			cmd = append(cmd, "-c:v", "prores_ks", "-profile:v", "3", "-vendor", "ap10", "pix_fmt", "yuv422p10le")
		case "proresXQ":
			cmd = append(cmd, "-c:v", "prores_ks", "-profile:v", "5", "-vendor", "ap10", "-bits_per_mb", "8000", "pix_fmt", "yuva444p10le")
		}
		cmd = append(cmd, "-vf", "pad=ceil(iw/2)*2:ceil(ih/2)*2", videoPath)
		job = Job{Shotname: shotname, Task: "video", SubTask: "video", OutPath: videoPath, Cmd: cmd, Cmdstr: strings.Join(cmd, " ")}
		jobs = append(jobs, job)
	}
	return jobs
}

func Publish(data []Item) ([]Item, error) {
	initData, err := LoadInit()
	if err != nil {
		return data, err
	}
	for i := range data {
		if !data[i].Pub {
			continue
		}
		jobArray := make([]Job, 0)
		if initData.ThumbCheck {
			jobs := thumbnailJob(initData, data[i])
			jobArray = append(jobArray, jobs...)
		}
		if initData.PlateCheck {
			jobs := plateJob(initData, data[i])
			jobArray = append(jobArray, jobs...)
		}
		if initData.VideoCheck {
			jobs := videoJob(initData, data[i])
			jobArray = append(jobArray, jobs...)
		}

		// render by shot
		errCheck := false
		for _, job := range jobArray {
			job.Render(initData.Ocio)
			if strings.Join(job.Errorlog, ",") != "" {
				errCheck = true
			}
		}

		now := time.Now().Format("2006-01-02 15:04:05")
		if errCheck {
			data[i].Log = "Error : " + now
		} else {
			data[i].Log = "Done : " + now
		}
	}
	return data, nil
}

// replace <SHOTNAME>, <PREFIX>, <SUFFIX> to real name
func replaceName(i *InitData, shotname, name string) string {
	if i.Splitter == "" {
		return name
	}
	prefix := strings.SplitN(shotname, i.Splitter, 2)[0]
	suffix := strings.SplitN(shotname, i.Splitter, 2)[1]
	reName := strings.Replace(name, "<SHOTNAME>", shotname, -1)
	reName = strings.Replace(reName, "<PREFIX>", prefix, -1)
	reName = strings.Replace(reName, "<SUFFIX>", suffix, -1)
	return reName
}

func StrPad(num, pad int) string {
	input := strconv.Itoa(num)
	output := strings.Repeat("0", pad) + input
	return output[len(output)-pad:]
}

// Check if item is in string array
func strContains(i string, s []string) bool {
	for _, a := range s {
		if a == i {
			return true
		}
	}
	return false
}

// calculate FrameIn & FrameOut from Trim Timecode
func (i *Item) timecodeToFrame() {
	if i.TrimInTc == "" && i.TrimOutTc == "" {
		return
	}
	// Trim in
	rate := timecode.NewFloatRate(float32(i.Fps)) // timecode rate
	tcin, _ := timecode.Parse(i.TimecodeIn)       // convert TimecodeIn
	tcin = tcin.SetRate(rate)                     // set rate to tcin
	titc, _ := timecode.Parse(i.TrimInTc)         // convert TrimTimecodeIn
	titc = titc.SetRate(rate)                     // set rate in titc
	duration := titc.Sub(tcin)                    // get duration
	trimInframes := rate.Frames(duration)         // get frames
	i.TrimIn = int(1 + trimInframes)
	// Trim out
	totc, _ := timecode.Parse(i.TrimOutTc) //convert TrimTimecodeOut
	totc = totc.SetRate(rate)              // set rate in totc
	duration = totc.Sub(tcin)              // get duration
	trimOutframes := rate.Frames(duration) // get frames
	i.TrimOut = int(1 + trimOutframes)
}
