package mediainfovars

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlbinoDrought/creamy-transcode/conf"

	"github.com/AlbinoDrought/creamy-transcode/mediainfo"

	"github.com/bradleyjkemp/cupaloy"
)

var mediaInfos = []mediainfo.MediaInfo{
	mediainfo.MediaInfo{
		MimeType:     "video/mp4",
		Size:         204141,
		Width:        640,
		Height:       360,
		AspectRatio:  "16:9",
		FPS:          "3823232/127989",
		VideoBitrate: "276883",
		IsHD:         false,
		IsAudioOnly:  false,
		Duration:     "4.285011",
		VideoCodec:   "h264",
		AudioCodec:   "aac",
		Channels:     2,
		AudioBitrate: "95999",
		SampleRate:   "44100",
	},
}

func TestSnapshotConvertToConfigurationVariables(t *testing.T) {
	for i, mediaInfo := range mediaInfos {
		t.Run("ConvertToConfigurationVariables snapshot "+strconv.Itoa(i), func(t *testing.T) {
			script := ConvertToConfigurationVariables(mediaInfo)
			cupaloy.SnapshotT(t, script)
		})
	}
}

func TestParsesSuccessfully(t *testing.T) {
	for i, mediaInfo := range mediaInfos {
		t.Run("ConvertToConfigurationVariables parse "+strconv.Itoa(i), func(t *testing.T) {
			script := ConvertToConfigurationVariables(mediaInfo)

			_, err := conf.ParseString(script)
			assert.Nil(t, err)
		})
	}
}
