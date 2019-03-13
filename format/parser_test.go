package format

import (
	"regexp"
	"testing"

	"github.com/bradleyjkemp/cupaloy"

	"github.com/stretchr/testify/assert"
)

func TestSplitPieces(t *testing.T) {
	cases := []struct {
		input     string
		expected  unparsedPieces
		errorText string
	}{
		{
			input:     "",
			expected:  unparsedPieces{},
			errorText: "empty input string",
		},
		{
			input:     "0:1:2:3:4",
			expected:  unparsedPieces{},
			errorText: "too many input fields",
		},
		{
			input: "mp4:720p:x:2pass",
			expected: unparsedPieces{
				container:     "mp4",
				video:         "720p",
				audio:         "x",
				formatOptions: "2pass",
			},
			errorText: "",
		},
		{
			input: "jpg",
			expected: unparsedPieces{
				container:     "jpg",
				video:         "",
				audio:         "",
				formatOptions: "",
			},
			errorText: "",
		},
		{
			input: "mp4:hevc_720p_1500k",
			expected: unparsedPieces{
				container:     "mp4",
				video:         "hevc_720p_1500k",
				audio:         "",
				formatOptions: "",
			},
			errorText: "",
		},
		// audio testcase, assume audio specs instead of video specs
		{
			input: "mp3:64k_22050hz_mono",
			expected: unparsedPieces{
				container:     "mp3",
				video:         "",
				audio:         "64k_22050hz_mono",
				formatOptions: "",
			},
			errorText: "",
		},
		// audio testcase, max of 2 fields
		{
			input:     "mp3:64k_22050hz_mono:2pass",
			expected:  unparsedPieces{},
			errorText: "too many input fields",
		},
	}

	for _, c := range cases {
		t.Run("input: "+c.input, func(t *testing.T) {
			actual, err := splitPieces(c.input)

			assert.EqualValues(t, c.expected, actual)
			if c.errorText == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, c.errorText)
			}
		})
	}
}

func TestParseVideoSpec(t *testing.T) {
	cases := []struct {
		input     string
		expected  VideoSpecs
		errorText string
	}{
		{
			input:    "",
			expected: VideoSpecs{},
		},
		{
			input: "x",
			expected: VideoSpecs{
				Disabled: true,
			},
		},
		{
			input:     "240p_x",
			expected:  VideoSpecs{},
			errorText: "unhandled video option: x",
		},
		{
			input: "240p_400k",
			expected: VideoSpecs{
				ResolutionWidth:  0,
				ResolutionHeight: 240,
				BitrateKbps:      400,
			},
		},
		{
			input: "720p_23.98fps",
			expected: VideoSpecs{
				ResolutionWidth:  1280,
				ResolutionHeight: 720,
				BitrateKbps:      2000,
				FPS:              "23.98",
			},
		},
		{
			input: "200x",
			expected: VideoSpecs{
				ResolutionWidth:  200,
				ResolutionHeight: 0,
			},
		},
		{
			input: "600x0_800k",
			expected: VideoSpecs{
				ResolutionWidth:  600,
				ResolutionHeight: 0,
				BitrateKbps:      800,
			},
		},
		{
			input: "hevc_1080p_1500k",
			expected: VideoSpecs{
				Codec:            "hevc",
				ResolutionWidth:  1920,
				ResolutionHeight: 1080,
				BitrateKbps:      1500,
			},
		},
	}

	for _, c := range cases {
		t.Run("input: "+c.input, func(t *testing.T) {
			actual, err := parseVideoSpec(c.input, VideoSpecs{})

			assert.EqualValues(t, c.expected, actual)
			if c.errorText == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, c.errorText)
			}
		})
	}
}

func TestParseAudioSpec(t *testing.T) {
	cases := []struct {
		input     string
		expected  AudioSpecs
		errorText string
	}{
		{
			input:    "",
			expected: AudioSpecs{},
		},
		{
			input: "x",
			expected: AudioSpecs{
				Disabled: true,
			},
		},
		{
			input: "64k_22050hz_mono",
			expected: AudioSpecs{
				BitrateKbps:  64,
				SampleRateHz: 22050,
				AudioChannel: AudioChannelMono,
			},
		},
		{
			input: "vorbis_44100hz_128k_stereo",
			expected: AudioSpecs{
				Codec:        "vorbis",
				SampleRateHz: 44100,
				BitrateKbps:  128,
				AudioChannel: AudioChannelStereo,
			},
		},
	}

	for _, c := range cases {
		t.Run("input: "+c.input, func(t *testing.T) {
			actual, err := parseAudioSpec(c.input, AudioSpecs{})

			assert.EqualValues(t, c.expected, actual)
			if c.errorText == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, c.errorText)
			}
		})
	}
}

func TestParseFormatOptions(t *testing.T) {
	cases := []struct {
		input     string
		expected  FormatOptions
		errorText string
	}{
		{
			input:     "",
			expected:  FormatOptions{},
			errorText: "",
		},
		{
			input:     "3pass",
			expected:  FormatOptions{},
			errorText: "unsupported format option: 3pass",
		},
		{
			input: "2pass",
			expected: FormatOptions{
				TwoPass: true,
			},
			errorText: "",
		},
	}

	for _, c := range cases {
		t.Run("input: "+c.input, func(t *testing.T) {
			actual, err := parseFormatOptions(c.input, FormatOptions{})

			assert.EqualValues(t, c.expected, actual)
			if c.errorText == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, c.errorText)
			}
		})
	}
}

func TestSnapshotParse(t *testing.T) {
	cases := []string{
		"jpg:300x",
		"webm:400x200_600k",
		"mp4:hevc_720p",
		"ios:300x0_400k",
		"gif",
		"storyboard",
		"mp3:64k_22050hz_mono",
		"mp4:x:256k_48000hz",
		"mp4:x",
		"mp4:hevc_1080p_2000k",
		"avi:mpeg4_640x320",
		"mp4:240p_15fps",
		"mp4:720p:x",
		"mp4:hevc_1080p:x",
		"mkv:mpeg4:mp3",
		"mpegts::mp3",
		"webm:1200k:256k",
		"mp4:1080p:512k",
		"mp4:240p:64k_22050hz",
		"mp4:720p:48000hz",
		"mp4:240p:64k_mono",
		"mp4:720p::2pass",
	}

	sanitize := regexp.MustCompile("[^A-Za-z0-9_]")

	for _, c := range cases {
		niceName := sanitize.ReplaceAllString(c, "_")
		t.Run("snapshot_"+niceName, func(t *testing.T) {
			actual, err := Parse(c)
			cupaloy.SnapshotT(t, actual, err)
		})
	}
}
