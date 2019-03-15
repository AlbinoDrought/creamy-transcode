# creamy-transcode/mediainfovars

Convert `mediainfo.MediaInfo` structs into a config script that can be parsed by `conf.Parse`.

## Usage

```golang
package main

import (
	"fmt"

	"github.com/AlbinoDrought/creamy-transcode/conf"
	"github.com/AlbinoDrought/creamy-transcode/mediainfo"
	"github.com/AlbinoDrought/creamy-transcode/mediainfovars"
	"github.com/davecgh/go-spew/spew"
)

func main() {
  // get mediainfo.MediaInfo from a local file
	result, _ := mediainfo.IdentifyFile("./mediainfo/ffprobe/test_fixtures/doggo_waddling.mp4")

  // convert it to a config script
	script := mediainfovars.ConvertToConfigurationVariables(result)
	fmt.Println(script)

  // our script which uses some converted variables like $source_width
	myScript := `
	set source = https://creamy-videos.r.albinodrought.com/static/videos/1.mov

	var output_file_name = transcoded_$source_widthx$source_height.mp4

	-> mp4 = https://creamy-videos.r.albinodrought.com/api/upload?size=$source_size&name=$output_file_name
	`

  // parse our merged script with conf.ParseString
	parsedScript, _ := conf.ParseString(script + myScript)
	spew.Dump(parsedScript)
}

// output:
/*
var source_mime_type = video/mp4
var source_size = 204141
var source_width = 640
var source_height = 360
var source_aspect_ratio = 16:9
var source_fps = 3823232/127989
var source_video_bitrate = 276883
var source_is_hd = false
var source_is_audio_only = false
var source_duration = 4.285011
var source_video_codec = h264
var source_audio_codec = aac
var source_channels = 2
var source_audio_bitrate = 95999
var source_sample_rate = 44100

(conf.ParsedConf) {
 SourceURL: (string) (len=61) "https://creamy-videos.r.albinodrought.com/static/videos/1.mov",
 WebhookURL: (string) "",
 Outputs: (map[string]string) (len=1) {
  (string) (len=3) "mp4": (string) (len=92) "https://creamy-videos.r.albinodrought.com/api/upload?size=204141&name=transcoded_640x360.mp4"
 }
}
*/
```
