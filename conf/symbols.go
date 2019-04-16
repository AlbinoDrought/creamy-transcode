package conf

import (
	"errors"
	"regexp"
	"strings"
)

type symbolParser func(line string, parseContext *parseContext) (bool, error)

var symbolParsers = []symbolParser{
	parseVar,
	parseSet,
	parseOutput,
}

var parseVarRegex = regexp.MustCompile("^var +([^ ]+) += +(.+)$")

func parseVar(line string, parseContext *parseContext) (bool, error) {
	matches := parseVarRegex.FindStringSubmatch(line)
	if matches == nil {
		return false, nil
	}

	name := matches[1]
	value := matches[2]
	value = parseContext.expand(value) // something something billion laughs attack

	parseContext.setVariable(name, value)

	return true, nil
}

const setSource = "source"
const setWebhook = "webhook"

var parseSetRegex = regexp.MustCompile("^set +([^ ]+) += +([^#\n]+)")

func parseSet(line string, parseContext *parseContext) (bool, error) {
	matches := parseSetRegex.FindStringSubmatch(line)
	if matches == nil {
		return false, nil
	}

	name := matches[1]
	value := matches[2]
	value = parseContext.expand(value)

	if name == setSource {
		parseContext.conf.SourceURL = value
	} else if name == setWebhook {
		parseContext.conf.WebhookURL = value
	} else {
		return true, errors.New("unhandled set name: " + name)
	}

	return true, nil
}

var parseOutputRegex = regexp.MustCompile("^\\-\\> +([^ ]+) += +(.+)$")

func parseOutput(line string, parseContext *parseContext) (bool, error) {
	matches := parseOutputRegex.FindStringSubmatch(line)
	if matches == nil {
		return false, nil
	}

	name := matches[1]
	value := matches[2]
	value = parseContext.expand(value)

	if _, alreadyExists := parseContext.conf.Outputs[name]; alreadyExists {
		return true, errors.New("output already exists: " + name)
	}

	parsedOutput := ParsedOutput{
		URL:     value,
		Options: map[string]string{},
	}

	// some sources can look like this:
	// -> mp4 = http://foo.bar, keep=video_bitrate, if=$source_width >= 1280
	// we want to end up with something like this:
	/*
		ParsedOutput{
			URL: "http://foo.bar",
			Options: map[string]string{
				"keep": "video_bitrate",
				"if": "$source_width >= 1280",
			}
		}
	*/
	// no validation on proper options, nothing - just parse them for later use

	//
	splitByCommaValue := strings.Split(value, ", ")
	if len(splitByCommaValue) > 1 {
		parsedOutput.URL = splitByCommaValue[0]
		for i := 1; i < len(splitByCommaValue); i++ {
			rawOptionStatement := splitByCommaValue[i]

			splitByEqualsStatement := strings.Split(rawOptionStatement, "=")

			var optionName string
			var optionValue string

			optionName = splitByEqualsStatement[0]
			if len(splitByEqualsStatement) > 1 {
				optionValue = strings.Join(splitByEqualsStatement[1:], "=")
			} else {
				optionValue = ""
			}

			if _, alreadyExists := parsedOutput.Options[optionName]; alreadyExists {
				return true, errors.New("output option already exists: " + name)
			}

			parsedOutput.Options[optionName] = optionValue
		}
	}

	parseContext.conf.Outputs[name] = parsedOutput

	return true, nil
}
