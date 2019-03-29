package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/AlbinoDrought/creamy-transcode/conf"
	"github.com/AlbinoDrought/creamy-transcode/files"
	"github.com/AlbinoDrought/creamy-transcode/format"
	"github.com/AlbinoDrought/creamy-transcode/mediainfo"
	"github.com/AlbinoDrought/creamy-transcode/mediainfovars"
	"github.com/AlbinoDrought/creamy-transcode/transcoder"
	"github.com/kennygrant/sanitize"
	"github.com/pkg/errors"
)

type transcodeResult struct {
	formatName string
	outputURLs []string
	err        error
}

func main() {
	if len(os.Args) < 2 {
		log.Println("usage: creamy-transcode [path to config file]")
		os.Exit(0)
	}

	// get a stream to the config file
	configFilePath := os.Args[1]

	configFileBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	configFileText := string(configFileBytes)
	configFileBytes = nil // is this allowed

	// actually read and parse the config file
	parsedConfiguration, err := conf.ParseString(configFileText)
	if err != nil {
		log.Fatalf("error parsing configuration: %v", err)
	}

	// make temp dir
	tempDir, err := ioutil.TempDir("", "creamy-transcode")
	if err != nil {
		log.Fatalf("error making temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// download the source material
	sourceFilename := path.Base(parsedConfiguration.SourceURL)
	if sourceFilename == "" {
		sourceFilename = "video"
	}
	tempDownloadPath := path.Join(tempDir, sourceFilename)

	log.Printf("downloading %v to %v", parsedConfiguration.SourceURL, tempDownloadPath)
	err = files.Download(parsedConfiguration.SourceURL, tempDownloadPath)
	if err != nil {
		log.Fatalf("error downloading %v to %v: %v", parsedConfiguration.SourceURL, tempDownloadPath, err)
	}

	// examine it
	log.Printf("examining %v", tempDownloadPath)
	examinedInfo, err := mediainfo.IdentifyFile(tempDownloadPath)
	if err != nil {
		log.Fatalf("error examining %v: %v", tempDownloadPath, err)
	}

	// convert to configuration vars
	examinedInfoAsConfigurationScript := mediainfovars.ConvertToConfigurationVariables(examinedInfo)

	monsterConfigFile := examinedInfoAsConfigurationScript + configFileText

	// re-parse monster config file
	parsedConfiguration, err = conf.ParseString(monsterConfigFile)
	if err != nil {
		log.Fatalf("error parsing monster config: %v", err)
	}

	transcodeResultChannel := make(chan transcodeResult, len(parsedConfiguration.Outputs))

	for formatName, outputURL := range parsedConfiguration.Outputs {
		go func(formatName string, outputURL string) {
			result := transcodeResult{
				formatName: formatName,
				outputURLs: []string{},
			}

			// always shove the result into the result channel
			defer func() {
				transcodeResultChannel <- result
			}()

			// parse the simple mp4:360p format
			parsedFormat, err := format.Parse(formatName)
			if err != nil {
				result.err = errors.Wrap(err, "format parse error")
				return
			}

			// shove output file somewhere in the existing tempdir
			tempOutputPath := path.Join(tempDir, sanitize.Path(formatName))

			err = transcoder.TranscodeFormat(tempDownloadPath, tempOutputPath, &parsedFormat)
			if err != nil {
				result.err = errors.Wrap(err, "transcode format error")
				return
			}

			// dump it to output url
			err = files.Upload(tempOutputPath, outputURL)
			if err != nil {
				result.err = errors.Wrap(err, "upload error")
				return
			}

			// alrighty, everything maybe worked
			result.outputURLs = []string{outputURL}
		}(formatName, outputURL)
	}

	passed := 0
	failed := 0

	transcodeResults := make([]transcodeResult, len(parsedConfiguration.Outputs))
	for i := 0; i < len(parsedConfiguration.Outputs); i++ {
		// not sure of the proper way to do this
		// i want to loop over the channel len(parsedConfiguration.Outputs) times (blocking),
		// and then close it.
		result := <-transcodeResultChannel
		log.Printf("this just in: -> %v = %v | %v", result.formatName, result.outputURLs, result.err)
		transcodeResults = append(transcodeResults, result)

		if result.err == nil {
			passed++
		} else {
			failed++
		}
	}
	close(transcodeResultChannel)

	// todo: webhook here
	log.Printf("passed transcodes: %v failed transcodes: %v", passed, failed)
}
