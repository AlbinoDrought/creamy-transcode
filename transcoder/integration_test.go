package transcoder

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlbinoDrought/creamy-transcode/format"
)

var inputFiles = []string{
	"./../mediainfo/ffprobe/test_fixtures/doggo_waddling.mp4",
}

var formatStrings = []string{
	"mp4",
	"webm",
	"avi",
	"asf",
	"mpegts",
	"mov",
	"flv",
	"mkv",
	"3gp:352x288",
	"ogv",
	"ogg",
	"mp3",
	"jpg",
	"png",
}

func TestTranscodeDoesNotError(t *testing.T) {
	for i, file := range inputFiles {
		dir, err := ioutil.TempDir("", fmt.Sprintf("transcode_integration_%v", i))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(dir)

		defer os.RemoveAll(dir)

		for _, formatString := range formatStrings {
			parsed, err := format.Parse(formatString)
			assert.Nil(t, err)

			err = TranscodeFormat(
				file,
				path.Join(dir, "output."+parsed.Container),
				&parsed,
			)

			assert.Nil(t, err, "transcode", formatString)
		}
	}
}
