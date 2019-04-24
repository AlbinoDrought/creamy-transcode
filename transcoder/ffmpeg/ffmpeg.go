package ffmpeg

import (
	"os/exec"

	"github.com/pkg/errors"
)

// TranscodeRaw shells out a raw arg array to ffmpeg
func TranscodeRaw(command []string) error {
	cmd := exec.Command("ffmpeg", command...)
	output, err := cmd.Output()

	if err != nil {
		exitError := err.(*exec.ExitError)
		errorOutput := ""
		if exitError != nil {
			errorOutput = string(exitError.Stderr)
		}

		return errors.Wrapf(
			err,
			"command: %+v\n output: %+v\nerror output: %+v\n",
			command,
			string(output),
			errorOutput,
		)
	}

	return nil
}

// TranscodeRawAll shells out multiple raw arg arrays to ffmpeg
func TranscodeRawAll(commands [][]string) error {
	for _, command := range commands {
		err := TranscodeRaw(command)

		if err != nil {
			return err
		}
	}

	return nil
}
