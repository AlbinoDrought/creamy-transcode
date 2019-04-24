package transcoder

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/AlbinoDrought/creamy-transcode/transcoder/ffmpeg"
	"github.com/pkg/errors"
)

type ThumbnailTranscoder struct{}

func (transcoder ThumbnailTranscoder) Transcode(request *TranscodeRequest) TranscodeResult {
	transcodeResult := TranscodeResult{
		Request:        request,
		ResultingFiles: []string{},
		Error:          nil,
	}

	baseCommand := []string{"-i", request.SourceLocalPath, "-strict", "-2", "-y", "-hide_banner"}
	videoFilters := []string{}

	height := request.Format.VideoSpecs.ResolutionHeight
	width := request.Format.VideoSpecs.ResolutionWidth

	if request.ThumbnailOptions.Square {
		if width == 0 {
			transcodeResult.Error = errors.New("thumbnail width required when making square thumbnails")
			return transcodeResult
		}

		height = width
	}

	// -2 signifies automatic in ffmpeg
	hasHeight := true
	if height == 0 {
		hasHeight = false
		height = -2
	}

	hasWidth := true
	if width == 0 {
		hasWidth = false
		width = -2
	}

	hasAnySize := hasHeight || hasWidth
	hasAllSize := hasHeight && hasWidth

	if !hasAnySize {
		transcodeResult.Error = errors.New("width or height must be set when making thumbnails")
		return transcodeResult
	}

	if request.ThumbnailOptions.Crop {
		// crop since it was explicitly specified.
		if !hasAllSize {
			transcodeResult.Error = errors.New("width and height must be set when cropping thumbnails")
			return transcodeResult
		}

		videoFilters = append(
			videoFilters,
			fmt.Sprintf(
				"scale=%d:%d:force_original_aspect_ratio=increase",
				width,
				height,
			),
			fmt.Sprintf(
				"crop=%d:%d",
				width,
				height,
			),
		)
	} else if hasAllSize {
		// pad since we have both width and height
		videoFilters = append(
			videoFilters,
			fmt.Sprintf(
				"scale=%d:%d:force_original_aspect_ratio=decrease",
				width,
				height,
			),
			fmt.Sprintf(
				"pad=%d:%d:(ow-iw)/2:(oh-ih)/2",
				width,
				height,
			),
		)
	} else {
		// resize and keep aspect ratio since we only have width or height
		videoFilters = append(
			videoFilters,
			fmt.Sprintf(
				"scale=%d:%d:force_original_aspect_ratio=decrease",
				width,
				height,
			),
		)
	}

	if request.ThumbnailOptions.Sprite {
		transcodeResult.Error = errors.New("sprite not implemented")
		return transcodeResult
	}

	if request.ThumbnailOptions.VTT {
		transcodeResult.Error = errors.New("vtt not implemented")
		return transcodeResult
	}

	// shove video filters into base command
	// baseCommand = append(baseCommand, "-vf", strings.Join(videoFilters, ","))

	durationFloat, err := strconv.ParseFloat(request.SourceMediaInfo.Duration, 64)
	if err != nil {
		transcodeResult.Error = errors.Wrapf(err, "unable to parse duration %v", request.SourceMediaInfo.Duration)
		return transcodeResult
	}
	duration := time.Duration(durationFloat) * time.Second

	outputTempDir, err := ioutil.TempDir(request.TemporaryLocalPath, "tt")
	if err != nil {
		transcodeResult.Error = errors.Wrapf(err, "error creating tempdir %v", request.TemporaryLocalPath)
		return transcodeResult
	}

	outputPath := path.Join(outputTempDir, fmt.Sprintf("thumb%%d.%v", request.Format.Container))

	var commands [][]string

	if request.ThumbnailOptions.Number != 0 {
		if request.ThumbnailOptions.Number == 1 {
			baseCommand = append(baseCommand, "-frames:v", strconv.Itoa(request.ThumbnailOptions.Number))
			videoFilters = append(videoFilters, "thumbnail")
		} else {
			fps := float64(request.ThumbnailOptions.Number) / durationFloat
			videoFilters = append(videoFilters, fmt.Sprintf("fps=fps=%f", fps))
		}

		commands = [][]string{
			append(
				baseCommand,
				"-vf",
				strings.Join(videoFilters, ","),
				outputPath,
			),
		}
	} else if request.ThumbnailOptions.Every != 0 {
		every := request.ThumbnailOptions.Every

		if every < 0 {
			if duration <= 2*time.Minute {
				every = 2
			} else if duration <= 10*time.Minute {
				every = 5
			} else if duration <= 30*time.Minute {
				every = 10
			} else if duration <= 60*time.Minute {
				every = 20
			} else {
				every = 30
			}
		}

		videoFilters = append(videoFilters, fmt.Sprintf("fps=%d", every))

		commands = [][]string{
			append(
				baseCommand,
				"-vf",
				strings.Join(videoFilters, ","),
				outputPath,
			),
		}
	} else if len(request.ThumbnailOptions.Offsets) != 0 {
		for i, offset := range request.ThumbnailOptions.Offsets {
			commands = append(
				commands,
				append(
					baseCommand,
					"-vf",
					strings.Join(videoFilters, ","),
					"-ss",
					strconv.Itoa(offset),
					fmt.Sprintf(outputPath, i),
				),
			)
		}
	} else {
		transcodeResult.Error = errors.New("must specify at least one thumbnail generation method")
		return transcodeResult
	}

	err = ffmpeg.TranscodeRawAll(commands)
	if err != nil {
		transcodeResult.Error = errors.Wrapf(err, "error running commands %+v", commands)
		return transcodeResult
	}

	files, err := ioutil.ReadDir(outputTempDir)
	if err != nil {
		transcodeResult.Error = errors.Wrapf(err, "error reading tempdir %v", outputTempDir)
		return transcodeResult
	}

	resultingFiles := []string{}
	for _, file := range files {
		if !file.IsDir() {
			resultingFiles = append(resultingFiles, path.Join(outputTempDir, file.Name()))
		}
	}

	transcodeResult.ResultingFiles = resultingFiles

	return transcodeResult
}
