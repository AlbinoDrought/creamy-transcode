package transcoder

import (
	"os/exec"

	cfmt "github.com/AlbinoDrought/creamy-transcode/format"
	"github.com/pkg/errors"
)

// TranscodeFormat takes the file at `input` and transcodes it using the given Format,
// saving it at `output`
func TranscodeFormat(input string, output string, format *cfmt.Format) error {
	commands := FormatToFFMPEG(format)
	return TranscodeRaw(input, output, commands)
}

// TranscodeRaw takes the file at `input` and transcodes it using the given raw commands,
// saving it at `output`
func TranscodeRaw(input string, output string, commands [][]string) error {
	for _, command := range commands {
		fullArgs := []string{
			"-i",
			input,
		}

		fullArgs = append(fullArgs, command...)
		fullArgs = append(fullArgs, "-strict", "-2", output)

		cmd := exec.Command("ffmpeg", fullArgs...)
		output, err := cmd.Output()

		if err != nil {
			exitError := err.(*exec.ExitError)
			errorOutput := ""
			if exitError != nil {
				errorOutput = string(exitError.Stderr)
			}

			return errors.Wrapf(
				err,
				"input: %+v\noutput: %+v\n command: %+v\n output: %+v\nerror output: %+v\n",
				input,
				output,
				command,
				string(output),
				errorOutput,
			)
		}
	}

	return nil
}
