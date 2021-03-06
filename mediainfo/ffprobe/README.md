# creamy-transcode/mediainfo/ffprobe

Run `ffprobe` on local files and decode the output.

## Usage

```golang
package main

import (
	"fmt"

	"github.com/AlbinoDrought/creamy-transcode/mediainfo/ffprobe"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	result, _ := ffprobe.ProbeFile("./mediainfo/ffprobe/test_fixtures/doggo_waddling.mp4")

	for i, stream := range result.Streams {
		fmt.Printf("stream %v:\n", i)
		fmt.Printf("audio    [%v]\n", ffprobe.IsAudioStream(&stream))
		fmt.Printf("video    [%v]\n", ffprobe.IsVideoStream(&stream))
		fmt.Printf("HD video [%v]\n\n", ffprobe.IsHDVideoStream(&stream))
	}

	spew.Dump(result)
}


// output:
/*
stream 0:
audio    [false]
video    [true]
HD video [false]

stream 1:
audio    [true]
video    [false]
HD video [false]

(ffprobe.Result) {
 Streams: ([]ffprobe.Stream) (len=2 cap=4) {
  (ffprobe.Stream) {
   Index: (int) 0,
   CodecName: (string) (len=4) "h264",
   CodecLongName: (string) (len=41) "H.264 / AVC / MPEG-4 AVC / MPEG-4 part 10",
   Profile: (string) (len=20) "Constrained Baseline",
   CodecType: (string) (len=5) "video",
   CodecTimeBase: (string) (len=14) "127989/7646464",
   CodecTagString: (string) (len=4) "avc1",
   CodecTag: (string) (len=10) "0x31637661",
   Width: (int) 640,
   Height: (int) 360,
   CodedWidth: (int) 640,
   CodedHeight: (int) 368,
   HasBFrames: (int) 0,
   SampleAspectRatio: (string) (len=3) "1:1",
   DisplayAspectRatio: (string) (len=4) "16:9",
   PixFmt: (string) (len=7) "yuv420p",
   Level: (int) 30,
   ChromaLocation: (string) (len=4) "left",
   Refs: (int) 1,
   IsAvc: (string) (len=4) "true",
   NalLengthSize: (string) (len=1) "4",
   RFrameRate: (string) (len=10) "29869/1000",
   AvgFrameRate: (string) (len=14) "3823232/127989",
   TimeBase: (string) (len=7) "1/29869",
   StartPts: (int) 0,
   StartTime: (string) (len=8) "0.000000",
   DurationTs: (int) 127989,
   Duration: (string) (len=8) "4.285011",
   BitRate: (string) (len=6) "276883",
   BitsPerRawSample: (string) (len=1) "8",
   NbFrames: (string) (len=3) "128",
   Disposition: (struct { Default int "json:\"default\""; Dub int "json:\"dub\""; Original int "json:\"original\""; Comment int "json:\"comment\""; Lyrics int "json:\"lyrics\""; Karaoke int "json:\"karaoke\""; Forced int "json:\"forced\""; HearingImpaired int "json:\"hearing_impaired\""; VisualImpaired int "json:\"visual_impaired\""; CleanEffects int "json:\"clean_effects\""; AttachedPic int "json:\"attached_pic\""; TimedThumbnails int "json:\"timed_thumbnails\"" }) {
    Default: (int) 1,
    Dub: (int) 0,
    Original: (int) 0,
    Comment: (int) 0,
    Lyrics: (int) 0,
    Karaoke: (int) 0,
    Forced: (int) 0,
    HearingImpaired: (int) 0,
    VisualImpaired: (int) 0,
    CleanEffects: (int) 0,
    AttachedPic: (int) 0,
    TimedThumbnails: (int) 0
   },
   SampleFmt: (string) "",
   SampleRate: (string) "",
   Channels: (int) 0,
   ChannelLayout: (string) "",
   BitsPerSample: (int) 0
  },
  (ffprobe.Stream) {
   Index: (int) 1,
   CodecName: (string) (len=3) "aac",
   CodecLongName: (string) (len=27) "AAC (Advanced Audio Coding)",
   Profile: (string) (len=2) "LC",
   CodecType: (string) (len=5) "audio",
   CodecTimeBase: (string) (len=7) "1/44100",
   CodecTagString: (string) (len=4) "mp4a",
   CodecTag: (string) (len=10) "0x6134706d",
   Width: (int) 0,
   Height: (int) 0,
   CodedWidth: (int) 0,
   CodedHeight: (int) 0,
   HasBFrames: (int) 0,
   SampleAspectRatio: (string) "",
   DisplayAspectRatio: (string) "",
   PixFmt: (string) "",
   Level: (int) 0,
   ChromaLocation: (string) "",
   Refs: (int) 0,
   IsAvc: (string) "",
   NalLengthSize: (string) "",
   RFrameRate: (string) (len=3) "0/0",
   AvgFrameRate: (string) (len=3) "0/0",
   TimeBase: (string) (len=7) "1/44100",
   StartPts: (int) 0,
   StartTime: (string) (len=8) "0.000000",
   DurationTs: (int) 195584,
   Duration: (string) (len=8) "4.435011",
   BitRate: (string) (len=5) "95999",
   BitsPerRawSample: (string) "",
   NbFrames: (string) (len=3) "191",
   Disposition: (struct { Default int "json:\"default\""; Dub int "json:\"dub\""; Original int "json:\"original\""; Comment int "json:\"comment\""; Lyrics int "json:\"lyrics\""; Karaoke int "json:\"karaoke\""; Forced int "json:\"forced\""; HearingImpaired int "json:\"hearing_impaired\""; VisualImpaired int "json:\"visual_impaired\""; CleanEffects int "json:\"clean_effects\""; AttachedPic int "json:\"attached_pic\""; TimedThumbnails int "json:\"timed_thumbnails\"" }) {
    Default: (int) 1,
    Dub: (int) 0,
    Original: (int) 0,
    Comment: (int) 0,
    Lyrics: (int) 0,
    Karaoke: (int) 0,
    Forced: (int) 0,
    HearingImpaired: (int) 0,
    VisualImpaired: (int) 0,
    CleanEffects: (int) 0,
    AttachedPic: (int) 0,
    TimedThumbnails: (int) 0
   },
   SampleFmt: (string) (len=4) "fltp",
   SampleRate: (string) (len=5) "44100",
   Channels: (int) 2,
   ChannelLayout: (string) (len=6) "stereo",
   BitsPerSample: (int) 0
  }
 },
 Format: (ffprobe.Format) {
  Filename: (string) (len=52) "./mediainfo/ffprobe/test_fixtures/doggo_waddling.mp4",
  NbStreams: (int) 2,
  NbPrograms: (int) 0,
  FormatName: (string) (len=23) "mov,mp4,m4a,3gp,3g2,mj2",
  FormatLongName: (string) (len=15) "QuickTime / MOV",
  StartTime: (string) (len=8) "0.000000",
  Duration: (string) (len=8) "4.435000",
  Size: (string) (len=6) "204141",
  BitRate: (string) (len=6) "368236",
  ProbeScore: (int) 100
 }
}
*/
```
