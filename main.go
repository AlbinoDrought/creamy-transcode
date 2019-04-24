package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/AlbinoDrought/creamy-transcode/transcoder/transcoderoptions"

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

	apostle := transcoder.ApostleTranscoder{
		ContainerTranscoders: map[string]transcoder.Transcoder{
			"jpg": transcoder.ThumbnailTranscoder{},
			"png": transcoder.ThumbnailTranscoder{},
		},
		DefaultTranscoder: transcoder.VideoTranscoder{},
	}

	for formatName, outputURL := range parsedConfiguration.Outputs {
		go func(formatName string, outputURL conf.ParsedOutput) {
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

			thumbnailOptions, err := transcoderoptions.ParseThumbnailOptions(outputURL.Options)
			if err != nil {
				result.err = errors.Wrap(err, "thumbnail option parse error")
				return
			}

			// shove output file somewhere in the existing tempdir
			tempOutputPath, err := ioutil.TempDir(tempDir, sanitize.Path(formatName))
			if err != nil {
				result.err = errors.Wrap(err, "unable to create temp dir")
				return
			}

			transcodeRequest := &transcoder.TranscodeRequest{
				Format:             &parsedFormat,
				SourceLocalPath:    tempDownloadPath,
				SourceMediaInfo:    &examinedInfo,
				TemporaryLocalPath: tempOutputPath,
				ParsedOutput:       &outputURL,
				ThumbnailOptions:   &thumbnailOptions,
			}

			transcodeResult := apostle.Transcode(transcodeRequest)
			if transcodeResult.Error != nil {
				result.err = transcodeResult.Error
				return
			}

			// dump it to output url
			numlessOutputURL := strings.Replace(outputURL.URL, "#num#", "%.2d", 1)
			for i, resultingFile := range transcodeResult.ResultingFiles {
				resultingFileOutputURL := fmt.Sprintf(numlessOutputURL, i)
				if strings.Contains(resultingFileOutputURL, "%!") {
					// assume there was an sprintf issue...
					// todo: somehow check before shoving into sprintf?
					resultingFileOutputURL = numlessOutputURL
				}

				err = files.Upload(resultingFile, resultingFileOutputURL)
				if err != nil {
					result.err = errors.Wrapf(err, "unable to upload result #%d from %v", i, resultingFile)
					return
				}
				result.outputURLs = append(result.outputURLs, resultingFileOutputURL)
			}
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
