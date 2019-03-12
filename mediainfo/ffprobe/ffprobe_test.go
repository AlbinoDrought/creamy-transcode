package ffprobe

import (
	"regexp"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

var inputFiles = []string{
	"./test_fixtures/doggo_waddling.mp4",
	"./test_fixtures/test_chairsqueak.wav",
}
var sanitize = regexp.MustCompile("[^A-Za-z0-9_]")

func TestSnapshotProbeFile(t *testing.T) {
	for _, file := range inputFiles {
		niceName := sanitize.ReplaceAllString(file, "_")
		t.Run("ProbeFile snapshot "+niceName, func(t *testing.T) {
			actual, err := ProbeFile(file)
			cupaloy.SnapshotT(t, actual, err)
		})
	}
}

/*
func TestSnapshotProbeStream(t *testing.T) {
	for _, file := range inputFiles {
		niceName := sanitize.ReplaceAllString(file, "_")
		t.Run("ProbeStream snapshot "+niceName, func(t *testing.T) {
			stream, err := os.Open(file)
			defer stream.Close()
			assert.Nil(t, err, "sanity check: input file should exist")
			actual, err := ProbeStream(stream)
			cupaloy.SnapshotT(t, actual, err)
		})
	}
}

func TestProbeStreamProbeFileSameStreams(t *testing.T) {
	// assert retrieved stream information is the same across ProbeFile/ProbeStream.
	// `format` will be different since `ProbeStream` does not have access to this
	// information (bitrate, size, filename)
	for _, file := range inputFiles {
		niceName := sanitize.ReplaceAllString(file, "_")
		t.Run("ProbeStream snapshot "+niceName, func(t *testing.T) {
			stream, err := os.Open(file)
			defer stream.Close()
			assert.Nil(t, err, "sanity check: input file should exist")

			streamActual, streamError := ProbeStream(stream)
			fileActual, fileError := ProbeFile(file)

			assert.Equal(t, streamActual.Streams, fileActual.Streams)
			assert.Equal(t, streamError, fileError)
		})
	}
}
*/
