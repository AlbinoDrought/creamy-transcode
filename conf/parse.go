package conf

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

type ParsedConf struct {
	SourceURL  string
	WebhookURL string
	Outputs    map[string]string
}

type parseContext struct {
	conf                ParsedConf
	variables           map[string]string
	sortedVariableNames []string
}

func (parseContext *parseContext) setVariable(name string, value string) {
	_, alreadyExists := parseContext.variables[name]
	parseContext.variables[name] = value

	if !alreadyExists {
		parseContext.sortedVariableNames = append(parseContext.sortedVariableNames, name)
		sort.Sort(variablesByLength(parseContext.sortedVariableNames))
	}
}

func (parseContext *parseContext) expand(input string) string {
	for _, name := range parseContext.sortedVariableNames {
		value := parseContext.variables[name]
		input = strings.Replace(input, "$"+name, value, -1)
	}

	return input
}

// comments must be at the beginning line or have a space before them.
var removeCommentsRegex = regexp.MustCompile("((^#)|( #)).*$")

func removeComments(line string) string {
	return removeCommentsRegex.ReplaceAllString(line, "")
}

var trimWhitespaceRegex = regexp.MustCompile("(^\\s+)|(\\s+$)")

func trimWhitespace(line string) string {
	return trimWhitespaceRegex.ReplaceAllString(line, "")
}

// Parse a configuration file stream into a configuration file
func Parse(reader io.Reader) (ParsedConf, error) {
	parseContext := &parseContext{
		conf: ParsedConf{
			Outputs: make(map[string]string),
		},
		variables:           make(map[string]string),
		sortedVariableNames: []string{},
	}

	scanner := bufio.NewScanner(reader)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		rawLine := scanner.Text()
		line := rawLine
		line = trimWhitespace(line)
		line = removeComments(line)
		line = trimWhitespace(line)

		if line == "" {
			continue
		}

		var handled bool
		var err error

		for _, symbolParser := range symbolParsers {
			handled, err = symbolParser(line, parseContext)

			if err != nil {
				return ParsedConf{}, err
			}

			if handled {
				break
			}
		}

		if !handled {
			return ParsedConf{}, fmt.Errorf("unhandled configuration expression on line %+v: %+v (parsed as %+v)", lineNumber, rawLine, line)
		}
	}

	return parseContext.conf, nil
}

// ParseString parses a configuration file string into a configuration file
func ParseString(input string) (ParsedConf, error) {
	reader := strings.NewReader(input)
	return Parse(reader)
}
