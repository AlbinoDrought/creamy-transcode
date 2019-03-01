package format

import (
	"testing"

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
