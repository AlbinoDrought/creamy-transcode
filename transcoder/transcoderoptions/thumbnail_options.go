package transcoderoptions

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ThumbnailOptions struct {
	Number  int
	Offsets []int
	Every   int
	Square  bool
	Crop    bool
	Sprite  bool
	VTT     bool
}

func ParseThumbnailOptions(options map[string]string) (ThumbnailOptions, error) {
	thumbnailOptions := ThumbnailOptions{
		Offsets: []int{},
	}

	hasGenerationMethod := false

	rawNumber, ok := options["number"]
	if ok {
		number, err := strconv.Atoi(rawNumber)
		if err != nil {
			return ThumbnailOptions{}, errors.Wrap(err, "invalid number")
		}
		thumbnailOptions.Number = number
		hasGenerationMethod = true
	}

	rawOffsets, ok := options["offsets"]
	if ok {
		rawSplitOffsets := strings.Split(rawOffsets, ",")
		for _, rawSplitOffset := range rawSplitOffsets {
			splitOffset, err := strconv.Atoi(rawSplitOffset)
			if err != nil {
				return ThumbnailOptions{}, errors.Wrap(err, "invalid offset")
			}
			thumbnailOptions.Offsets = append(thumbnailOptions.Offsets, splitOffset)
		}

		if hasGenerationMethod {
			return ThumbnailOptions{}, errors.New("only able to choose one generation method")
		}

		hasGenerationMethod = true
	}

	rawEvery, ok := options["every"]
	if ok {
		every, err := strconv.Atoi(rawEvery)
		if err != nil {
			return ThumbnailOptions{}, errors.Wrap(err, "invalid every")
		}

		if hasGenerationMethod {
			return ThumbnailOptions{}, errors.New("only able to choose one generation method")
		}

		if every <= 0 {
			every = -1 // imply automatic every
		}

		thumbnailOptions.Every = every
		hasGenerationMethod = true
	}

	square, ok := options["square"]
	thumbnailOptions.Square = ok && square == "true"

	fit, ok := options["fit"]
	thumbnailOptions.Crop = ok && fit == "crop"

	sprite, ok := options["sprite"]
	thumbnailOptions.Sprite = ok && sprite == "true"

	vtt, ok := options["vtt"]
	thumbnailOptions.VTT = ok && vtt == "true"

	if !hasGenerationMethod {
		// if no other generation options set, default to 1 thumbnail
		thumbnailOptions.Number = 1
	}

	return thumbnailOptions, nil
}
