package mediainfo

import (
	"regexp"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
)

var inputFiles = []string{
	"./ffprobe/test_fixtures/doggo_waddling.mp4",
	"./ffprobe/test_fixtures/test_chairsqueak.wav",
}
var sanitize = regexp.MustCompile("[^A-Za-z0-9_]")

func TestSnapshotIdentifyFile(t *testing.T) {
	for _, file := range inputFiles {
		niceName := sanitize.ReplaceAllString(file, "_")
		t.Run("IdentifyFile snapshot "+niceName, func(t *testing.T) {
			actual, err := IdentifyFile(file)
			cupaloy.SnapshotT(t, actual, err)
		})
	}
}
