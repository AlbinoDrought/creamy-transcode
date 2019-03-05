package format

import (
	"errors"
	"fmt"
	"strconv"
)

type videoSpecOption func(input string, option *VideoSpecs) (bool, error)

var videoSpecOptions = []videoSpecOption{
	parseVideoDefinition,
	parseVideoResolution,
	parseVideoCodec,
	parseVideoBitrate,
	parseVideoFPS,
}

func parseVideoDefinition(input string, option *VideoSpecs) (bool, error) {
	matches := definitionRegex.FindStringSubmatch(input)
	if matches == nil {
		return false, nil
	}
	match := matches[1]

	switch match {
	case "240":
		option.ResolutionWidth = 0
		option.ResolutionHeight = 240
		option.BitrateKbps = 500
		break
	case "360":
		option.ResolutionWidth = 0
		option.ResolutionHeight = 360
		option.BitrateKbps = 800
		break
	case "480":
		option.ResolutionWidth = 0
		option.ResolutionHeight = 480
		option.BitrateKbps = 1000
		break
	case "720":
		option.ResolutionWidth = 1280
		option.ResolutionHeight = 720
		option.BitrateKbps = 2000
		break
	case "1080":
		option.ResolutionWidth = 1920
		option.ResolutionHeight = 1080
		option.BitrateKbps = 4000
		break
	default:
		return true, errors.New("unhandled *p notation")
	}

	return true, nil
}

func parseVideoResolution(input string, option *VideoSpecs) (bool, error) {
	matches := resolutionRegex.FindStringSubmatch(input)
	if matches == nil {
		return false, nil
	}

	width, err := strconv.Atoi(matches[1])
	if err != nil {
		return true, err
	}

	height, err := strconv.Atoi(matches[2])
	if err != nil {
		return true, err
	}

	option.ResolutionWidth = width
	option.ResolutionHeight = height

	return true, nil
}

func parseVideoCodec(input string, option *VideoSpecs) (bool, error) {
	if !videoCodecRegex.MatchString(input) {
		return false, nil
	}

	option.Codec = input

	return true, nil
}

func parseVideoBitrate(input string, option *VideoSpecs) (bool, error) {
	match := bitrateRegex.FindStringSubmatch(input)
	if match == nil {
		return false, nil
	}

	asInt, err := strconv.Atoi(match[1])

	if err != nil {
		return true, err
	}

	if asInt >= bitrateMaximum {
		return true, fmt.Errorf("maximum bitrate is %+v", bitrateMaximum)
	}

	option.BitrateKbps = asInt

	return true, nil
}

func parseVideoFPS(input string, option *VideoSpecs) (bool, error) {
	match := fpsRegex.FindStringSubmatch(input)
	if match == nil {
		return false, nil
	}

	option.FPS = match[1]

	return true, nil
}
