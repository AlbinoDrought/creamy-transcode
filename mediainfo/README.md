# creamy-transcode/mediainfo

Extract media information from locally-stored videos.

## Usage

```golang
package main

import (
	"github.com/AlbinoDrought/creamy-transcode/mediainfo"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	result, _ := mediainfo.IdentifyFile("./mediainfo/ffprobe/test_fixtures/doggo_waddling.mp4")

	spew.Dump(result)
}

// output:
/*
(mediainfo.MediaInfo) {
 MimeType: (string) (len=9) "video/mp4",
 Size: (int64) 204141,
 Width: (int) 640,
 Height: (int) 360,
 AspectRatio: (string) (len=4) "16:9",
 FPS: (string) (len=14) "3823232/127989",
 VideoBitrate: (string) (len=6) "276883",
 IsHD: (bool) false,
 IsAudioOnly: (bool) false,
 Duration: (string) (len=8) "4.285011",
 VideoCodec: (string) (len=4) "h264",
 AudioCodec: (string) (len=3) "aac",
 Channels: (int) 2,
 AudioBitrate: (string) (len=5) "95999",
 SampleRate: (string) (len=5) "44100"
}
*/
```
