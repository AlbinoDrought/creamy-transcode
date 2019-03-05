package format

import (
	"fmt"
	"strconv"
)

type audioSpecOption func(input string, option *AudioSpecs) (bool, error)

var audioSpecOptions = []audioSpecOption{
	parseAudioCodec,
	parseAudioBitrate,
	parseAudioSampleRate,
	parseAudioChannel,
}

func parseAudioCodec(input string, option *AudioSpecs) (bool, error) {
	if !audioCodecRegex.MatchString(input) {
		return false, nil
	}

	option.Codec = input

	return true, nil
}

func parseAudioBitrate(input string, option *AudioSpecs) (bool, error) {
	match := bitrateRegex.FindStringSubmatch(input)
	if match == nil {
		return false, nil
	}

	asInt, err := strconv.Atoi(match[1])

	if err != nil {
		return true, err
	}

	if asInt > audioBitrateMaximum {
		return true, fmt.Errorf("maximum audio bitrate is %+v", audioBitrateMaximum)
	}

	option.BitrateKbps = asInt

	return true, nil
}

func parseAudioSampleRate(input string, option *AudioSpecs) (bool, error) {
	match := sampleRateRegex.FindStringSubmatch(input)
	if match == nil {
		return false, nil
	}

	asInt, err := strconv.Atoi(match[1])

	if err != nil {
		return true, err
	}

	option.SampleRateHz = asInt

	return true, nil
}

func parseAudioChannel(input string, option *AudioSpecs) (bool, error) {
	match := audioChannelRegex.FindStringSubmatch(input)
	if match == nil {
		return false, nil
	}

	option.AudioChannel = match[1]

	return true, nil
}
