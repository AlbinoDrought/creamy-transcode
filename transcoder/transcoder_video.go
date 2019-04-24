package transcoder

import (
	"fmt"
	"io/ioutil"
	"path"

	cfmt "github.com/AlbinoDrought/creamy-transcode/format"
	"github.com/AlbinoDrought/creamy-transcode/transcoder/ffmpeg"
	"github.com/pkg/errors"
)

type VideoTranscoder struct{}

func (transcoder VideoTranscoder) Transcode(request *TranscodeRequest) TranscodeResult {
	format := request.Format

	result := []string{"-i", request.SourceLocalPath, "-strict", "-2"}

	result = append(result, "-y", "-hide_banner", "-f", getFormat(format.Container))

	if format.VideoSpecs.Disabled {
		result = append(result, "-vn")
	} else {
		// video resolution
		if format.VideoSpecs.ResolutionHeight != 0 || format.VideoSpecs.ResolutionWidth != 0 {
			height := format.VideoSpecs.ResolutionHeight
			width := format.VideoSpecs.ResolutionWidth

			// -1 signifies automatic in ffmpeg
			if height == 0 {
				height = -1
			} else if width == 0 {
				width = -1
			}

			result = append(result, "-vf", fmt.Sprintf("scale=%d:%d", width, height))
		}

		// video bitrate
		if format.VideoSpecs.BitrateKbps != 0 {
			bitrate := fmt.Sprintf("%dk", format.VideoSpecs.BitrateKbps)
			result = append(result, "-b:v", bitrate, "-bufsize", bitrate, "-maxrate", bitrate)
		}

		// video codec
		if format.VideoSpecs.Codec != "" {
			result = append(result, "-codec:v", format.VideoSpecs.Codec)
		}

		// video fps
		if format.VideoSpecs.FPS != "" && format.VideoSpecs.FPS != cfmt.FPSAutomatic {
			result = append(result, "-r", format.VideoSpecs.FPS)
		}
	}

	if format.AudioSpecs.Disabled {
		result = append(result, "-an")
	} else {
		// audio cdoec
		if format.AudioSpecs.Codec != "" {
			result = append(result, "-codec:a", format.AudioSpecs.Codec)
		}

		// audio bitrate
		if format.AudioSpecs.BitrateKbps != 0 {
			result = append(result, "-b:a", fmt.Sprintf("%dk", format.AudioSpecs.BitrateKbps))
		}

		// audio sample rate
		if format.AudioSpecs.SampleRateHz != 0 {
			result = append(result, "-ar", fmt.Sprintf("%d", format.AudioSpecs.SampleRateHz))
		}

		// audio channel (mono, stereo)
		if format.AudioSpecs.AudioChannel == cfmt.AudioChannelMono {
			result = append(result, "-ac", "1")
		} else if format.AudioSpecs.AudioChannel == cfmt.AudioChannelStereo {
			result = append(result, "-ac", "2")
		}
	}

	transcodeResult := TranscodeResult{
		Request:        request,
		ResultingFiles: []string{},
		Error:          nil,
	}

	outputTempDir, err := ioutil.TempDir(request.TemporaryLocalPath, "tv")
	if err != nil {
		transcodeResult.Error = errors.Wrapf(err, "error creating tempdir %v", request.TemporaryLocalPath)
		return transcodeResult
	}

	outputPath := path.Join(outputTempDir, "video")

	var commands [][]string

	// experimental
	if format.FormatOptions.TwoPass {
		commands = [][]string{
			append(result, "-pass", "1", outputPath),
			append(result, "-pass", "2", outputPath),
		}
	} else {
		commands = [][]string{
			append(result, outputPath),
		}
	}

	err = ffmpeg.TranscodeRawAll(commands)
	if err != nil {
		transcodeResult.Error = errors.Wrapf(err, "error running commands %+v", commands)
		return transcodeResult
	}

	transcodeResult.ResultingFiles = []string{
		outputPath,
	}

	return transcodeResult
}
