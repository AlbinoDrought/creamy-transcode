package format

import (
	"errors"
	"fmt"
	"strings"
)

type unparsedPieces struct {
	container     string
	video         string
	audio         string
	formatOptions string
}

const splitFieldCharacter = ":"
const splitOptionCharacter = "_"

const splitFieldContainer = 0
const splitFieldVideo = 1
const splitFieldAudio = 2
const splitFieldFormatOptions = 3
const splitMaxFields = 4

// splitPieces maps an input string like mp4:720p:x:2pass to unparsedPieces
func splitPieces(input string) (unparsedPieces, error) {
	if input == "" {
		return unparsedPieces{}, errors.New("empty input string")
	}

	rawPieces := strings.Split(input, splitFieldCharacter)
	length := len(rawPieces)
	pieces := unparsedPieces{}

	if length > splitFieldContainer {
		pieces.container = rawPieces[splitFieldContainer]
	}

	if isAudioContainer(pieces.container) {
		// insert an empty value where the video data would normally be.
		// also force empty format options to trigger "too many input fields" error if >2 fields
		// i know its spooky, please forgive me. its tested.

		// things we want to allow:
		// - mp3
		// - mp3:mono
		// things we dont want to allow:
		// - mp3:hevc:mono
		// - mp3:mono:2pass
		// - mp3:x:mono:2pass

		// assuming we have "mp3:mono", rawPieces will be      --->      // ["mp3", "mono"]
		rawPieces = append(rawPieces, "", "")                            // ["mp3", "mono", "", ""]
		length = len(rawPieces)                                          // length=4
		copy(rawPieces[splitFieldVideo+1:], rawPieces[splitFieldVideo:]) // ["mp3", "mono", "mono", ""]
		rawPieces[splitFieldVideo] = ""                                  // ["mp3", "", "mono", ""]
	}

	if length > splitMaxFields {
		return unparsedPieces{}, errors.New("too many input fields")
	}

	if length > splitFieldVideo {
		pieces.video = rawPieces[splitFieldVideo]
	}

	if length > splitFieldAudio {
		pieces.audio = rawPieces[splitFieldAudio]
	}

	if length > splitFieldFormatOptions {
		pieces.formatOptions = rawPieces[splitFieldFormatOptions]
	}

	return pieces, nil
}

func getOptions(input string) []string {
	return strings.Split(input, splitOptionCharacter)
}

func parseVideoSpec(input string) (VideoSpecs, error) {
	videoSpecs := VideoSpecs{}

	if input == "x" {
		videoSpecs.Disabled = true
		return videoSpecs, nil
	}

	for _, option := range getOptions(input) {
		handled := false
		var err error
		for _, optionHandler := range videoSpecOptions {
			handled, err = optionHandler(option, &videoSpecs)

			if err != nil {
				return VideoSpecs{}, err
			}

			if handled {
				break
			}
		}

		if !handled {
			return VideoSpecs{}, fmt.Errorf("unhandled video option: %+v", option)
		}
	}

	return videoSpecs, nil
}

func parseFormatOptions(input string) (FormatOptions, error) {
	formatOptions := FormatOptions{}

	for _, option := range getOptions(input) {
		if option == formatOptionTwoPass {
			formatOptions.TwoPass = true
			continue
		}

		return FormatOptions{}, fmt.Errorf("unsupported option: %+v", option)
	}

	return formatOptions, nil
}

func Parse(input string) (FormatOptions, error) {
	/*
		pieces, err := splitPieces(input)
		if err != nil {
			return FormatOptions{}, err
		}
	*/
	return FormatOptions{}, errors.New("foo")
}
