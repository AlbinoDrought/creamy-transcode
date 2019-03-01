package format

import (
	"errors"
	"strings"
)

type unparsedPieces struct {
	container     string
	video         string
	audio         string
	formatOptions string
}

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

	rawPieces := strings.Split(input, ":")
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

func Parse(input string) (FormatOptions, error) {
	/*
		pieces, err := splitPieces(input)
		if err != nil {
			return FormatOptions{}, err
		}
	*/

	return FormatOptions{}, errors.New("foo")
}
