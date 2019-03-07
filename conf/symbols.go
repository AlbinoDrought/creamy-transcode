package conf

import (
	"errors"
	"regexp"
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

	parseContext.variables[name] = value

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

	parseContext.conf.Outputs[name] = value

	return true, nil
}
