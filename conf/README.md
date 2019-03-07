# creamy-transcode/conf

Parse human-readable configuration files into machine-readable objects.

## Usage

```golang
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/AlbinoDrought/creamy-transcode/conf"
)

func main() {
	someResourceID := 5678
	s3Key := "some-s3-key"
	s3Secret := "some-s3-secret"
	s3Bucket := "some-s3-bucket"

	confText := fmt.Sprintf(`
	# sample valid creamy-transcode conf file
	var resource_id = %+v

	var s3_key = %+v
	var s3_secret = %+v
	var s3_bucket = %+v

	var base_s3 = s3://$s3_key:$s3_secret@$s3_bucket/videos/$resource_id/transcode

	set source = https://creamy-videos.internal.albinodrought.com/static/videos/24/video
	set webhook = https://creamy-videos.internal.albinodrought.com/api/video/$resource_id/transcoded, metadata=true

	-> webm = $base_s3/video.webm
	-> mp4 = $base_s3/video.mp4
	-> mp4:720p = $base_s3/video_720p.mp4, if=$source_width >= 1280

	# thumbnails
	-> jpg:300x = $base_s3/thumbnail_small_#num#.jpg, number=6
	-> jpg:160x = $base_s3/sprite.jpg, every=5, sprite=yes, vtt=yes
	`, someResourceID, s3Key, s3Secret, s3Bucket)

	parsed, err := conf.ParseString(confText)
	if err != nil {
		log.Fatal(err)
	}

	jsonForHumans, _ := json.MarshalIndent(parsed, "", "  ")
	log.Printf("%+v\n", string(jsonForHumans))

	// alternatively, to parse directly from a file

	// dump generated conf to disk for example
	file, err := os.Create("streamable_video_webvtt.conf")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(confText)
	file.Close()

	// read conf from disk
	file, err = os.Open("streamable_video_webvtt.conf")
	defer file.Close()

	parsed, err = conf.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	jsonForHumans, _ = json.MarshalIndent(parsed, "", "  ")
	log.Printf("%+v\n", string(jsonForHumans))
}

// output:
/*
{
  "SourceURL": "https://creamy-videos.internal.albinodrought.com/static/videos/24/video",
  "WebhookURL": "https://creamy-videos.internal.albinodrought.com/api/video/5678/transcoded, metadata=true",
  "Outputs": {
    "jpg:160x": "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/sprite.jpg, every=5, sprite=yes, vtt=yes",
    "jpg:300x": "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/thumbnail_small_#num#.jpg, number=6",
    "mp4": "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/video.mp4",
    "mp4:720p": "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/video_720p.mp4, if=$source_width >= 1280",
    "webm": "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/video.webm"
  }
}
*/
```