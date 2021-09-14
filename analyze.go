package main

import (
	"encoding/json"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"trimmer.io/go-timecode/timecode"
)

// FileInfo는 파일 경로를 받아 ffprobe를 사용하여 데이터를 분석하고,
// width, height, fps, codec, err를 반환한다.
// ffprobe data sample
// {
//     "streams": [
//         {
//             "index": 0,
//             "codec_name": "h264",
//             "codec_long_name": "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
//             "profile": "High",
//             "codec_type": "video",
//             "codec_time_base": "1/5994",
//             "codec_tag_string": "avc1",
//             "codec_tag": "0x31637661",
//             "width": 2048,
//             "height": 1152,
//             "coded_width": 2048,
//             "coded_height": 1152,
//             "has_b_frames": 2,
//             "pix_fmt": "yuv420p",
//             "level": 62,
//             "chroma_location": "left",
//             "refs": 1,
//             "is_avc": "true",
//             "nal_length_size": "4",
//             "r_frame_rate": "2997/1",
//             "avg_frame_rate": "2997/1",
//             "time_base": "1/11988",
//             "start_pts": 0,
//             "start_time": "0.000000",
//             "duration_ts": 40,
//             "duration": "0.003337",
//             "bit_rate": "56844698",
//             "bits_per_raw_sample": "8",
//             "nb_frames": "10",
//             "disposition": {
//                 "default": 1,
//                 "dub": 0,
//                 "original": 0,
//                 "comment": 0,
//                 "lyrics": 0,
//                 "karaoke": 0,
//                 "forced": 0,
//                 "hearing_impaired": 0,
//                 "visual_impaired": 0,
//                 "clean_effects": 0,
//                 "attached_pic": 0,
//                 "timed_thumbnails": 0
//             },
//             "tags": {
//                 "handler_name": "VideoHandler",
//                 "encoder": "Lavc58.54.100 libx264"
//             }
//         }
//     ]
// }
func (i *Item) Info(ffprobe, startNumber string) error {
	cmd := exec.Command(ffprobe, "-v", "quiet", "-print_format", "json", "-show_streams", i.Path)
	if startNumber != "" {
		cmd = exec.Command(ffprobe, "-v", "quiet", "-print_format", "json", "-show_streams", "-start_number", startNumber, i.Path)
	}
	stdout, err := cmd.Output()
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal([]byte(stdout), &data)
	if err != nil {
		return err
	}
	if value, ok := data["streams"]; ok {
		streams := value.([]interface{})
		item := streams[0].(map[string]interface{})
		duration := ""
		for k, v := range item {
			switch k {
			case "width":
				i.Width = int(v.(float64))
			case "height":
				i.Height = int(v.(float64))
			case "r_frame_rate":
				slice := strings.Split(v.(string), "/")
				r1, _ := strconv.ParseFloat(slice[0], 64)
				r2, _ := strconv.ParseFloat(slice[1], 64)
				i.Fps = r1 / r2
			case "codec_name":
				i.Codec = v.(string)
			case "duration":
				duration = v.(string)
			case "tags":
				for kk, vv := range v.(map[string]interface{}) {
					tc := strings.ToLower(kk)
					switch tc {
					case "timecode":
						i.TimecodeIn = vv.(string)
					}
				}
			}
		}
		if i.TimecodeIn == "" {
			i.TimecodeIn = "00:00:00:00"
		}
		// get timecode out
		rate := timecode.NewFloatRate(float32(i.Fps))         // timecode rate
		tcin, _ := timecode.Parse(i.TimecodeIn)               // convert type
		timeDuration, _ := time.ParseDuration(duration + "s") // convert duration
		frames := rate.Frames(timeDuration)                   // get total frames
		if frames != 0 {
			tcoutRaw := tcin.AddFrames(frames - 1) // make timecode out without rate
			tcout := tcoutRaw.SetRate(rate)        // make timecode out with rate
			i.TimecodeOut = tcout.String()         // convert string
		} else {
			i.TimecodeOut = "00:00:00:00"
		}
		// set frame
		if i.FrameIn == 0 && i.FrameOut == 0 {
			i.FrameIn = 1
			i.FrameOut = int(frames)
			i.FrameRange = int(frames)
			i.Pad = 0
		}
	}
	return nil
}
