package ffprobe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var audioStreams = []Stream{
	Stream{
		CodecType: "audio",
	},
}

var videoStreams = []Stream{
	Stream{
		CodecType: "video",
	},
}

var sdVideoStreams = []Stream{
	Stream{
		CodecType: "video",
		Height:    360,
		Width:     480,
	},
	Stream{
		CodecType: "video",
		Height:    240,
		Width:     426,
	},
}

var hdVideoStreams = []Stream{
	Stream{
		CodecType: "video",
		Height:    720,
		Width:     1280,
	},
	Stream{
		CodecType: "video",
		Height:    1080,
		Width:     1920,
	},
	Stream{
		CodecType: "video",
		Height:    2160,
		Width:     3840,
	},
}

func TestIsAudioStream(t *testing.T) {
	for i, audioStream := range audioStreams {
		assert.True(t, IsAudioStream(&audioStream), "audio stream", i)
	}

	for i, videoStream := range videoStreams {
		assert.False(t, IsAudioStream(&videoStream), "video stream", i)
	}
}

func TestIsVideoStream(t *testing.T) {
	for i, audioStream := range audioStreams {
		assert.False(t, IsVideoStream(&audioStream), "audio stream", i)
	}

	for i, videoStream := range videoStreams {
		assert.True(t, IsVideoStream(&videoStream), "video stream", i)
	}
}

func TestIsHDVideoStream(t *testing.T) {
	for i, hdStream := range hdVideoStreams {
		assert.True(t, IsHDVideoStream(&hdStream), "hd stream", i)
	}

	for i, sdStream := range sdVideoStreams {
		assert.False(t, IsHDVideoStream(&sdStream), "sd stream", i)
	}
}
