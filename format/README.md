# creamy-transcode/format

Parse human-readable format strings into well-defined machine-readable objects.

## Usage

```golang
package main

import (
	"encoding/json"
	"log"

	"github.com/AlbinoDrought/creamy-transcode/format"
)

func main() {
  // parse format
	parsed, err := format.Parse("mp4:hevc_720p_1500k")
	if err != nil {
		log.Fatal(err)
	}

  // convert to nicely-readable thing for dumping
	json, _ := json.MarshalIndent(parsed, "", "  ")
	log.Printf("%+v\n", string(json))
}

// output:
/*
{
  "Container": "mp4",
  "VideoSpecs": {
    "Disabled": false,
    "ResolutionHeight": 720,
    "ResolutionWidth": 1280,
    "BitrateKbps": 1500,
    "Codec": "hevc",
    "FPS": ""
  },
  "AudioSpecs": {
    "Disabled": false,
    "Codec": "aac",
    "BitrateKbps": 128,
    "SampleRateHz": 44100,
    "AudioChannel": "stereo"
  },
  "FormatOptions": {
    "TwoPass": false
  }
}
*/
```
