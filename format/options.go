package format

import (
	"regexp"
	"strings"
)

const formatOptionTwoPass = "2pass"

// const bitrateSuffix = "k"
const bitrateMaximum = 200000

var bitrateRegex = regexp.MustCompile("^(\\d{1,6})k$")

const fpsAutomatic = "0"
const fpsSuffix = "fps"

var resolutionRegex = regexp.MustCompile("^(\\d{1,4})x(\\d{1,4})$")

var definitionRegex = regexp.MustCompile("^(\\d{3,4})p$")

var fpsValid = []string{
	fpsAutomatic,
	"15",
	"23.98",
	"25",
	"29.97",
	"30",
}
var fpsRegex = regexp.MustCompile("^(" + strings.Join(fpsValid, "|") + ")fps$")
