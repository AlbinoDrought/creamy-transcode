package mediainfo

import (
	"os"

	"github.com/AlbinoDrought/creamy-transcode/mediainfo/ffprobe"
	"github.com/h2non/filetype"
	"github.com/pkg/errors"
)

// MediaInfo represents metadata of a media file and includes information about
// video and audio streams
type MediaInfo struct {
	MimeType     string
	Size         int64
	Width        int
	Height       int
	AspectRatio  string
	FPS          string
	VideoBitrate string
	IsHD         bool
	IsAudioOnly  bool
	Duration     string
	VideoCodec   string
	AudioCodec   string
	Channels     int
	AudioBitrate string
	SampleRate   string
}

func ripFromAudioStream(stream *ffprobe.Stream, mediaInfo *MediaInfo) {
	mediaInfo.AudioCodec = stream.CodecName
	mediaInfo.Channels = stream.Channels
	mediaInfo.AudioBitrate = stream.BitRate
	mediaInfo.SampleRate = stream.SampleRate

	// only set duration if it wasn't already set, prefer video duration instead
	if mediaInfo.Duration == "" {
		mediaInfo.Duration = stream.Duration
	}
}

func ripFromVideoStream(stream *ffprobe.Stream, mediaInfo *MediaInfo) {
	mediaInfo.Width = stream.Width
	mediaInfo.Height = stream.Height
	mediaInfo.AspectRatio = stream.DisplayAspectRatio
	mediaInfo.FPS = stream.AvgFrameRate
	mediaInfo.VideoBitrate = stream.BitRate
	mediaInfo.IsHD = ffprobe.IsHDVideoStream(stream)
	mediaInfo.Duration = stream.Duration
	mediaInfo.VideoCodec = stream.CodecName
}

// IdentifyFile parses metadata from a local file into easy-to-consume MediaInfo
func IdentifyFile(path string) (MediaInfo, error) {
	mediaInfo := MediaInfo{}

	statResult, err := os.Stat(path)
	if err != nil {
		return mediaInfo, errors.Wrap(err, "unable to stat file")
	}

	mediaInfo.Size = statResult.Size()

	matchedType, err := filetype.MatchFile(path)
	if err != nil {
		return mediaInfo, errors.Wrap(err, "unable to match mimetype")
	}

	mediaInfo.MimeType = matchedType.MIME.Value

	probeResult, err := ffprobe.ProbeFile(path)
	if err != nil {
		return mediaInfo, errors.Wrap(err, "unable to probe file")
	}

	foundVideoStream := false
	foundAudioStream := false

	for _, stream := range probeResult.Streams {
		if ffprobe.IsVideoStream(&stream) {
			if foundVideoStream {
				// already have a video stream, continue
				continue
			}
			foundVideoStream = true
			ripFromVideoStream(&stream, &mediaInfo)
		} else if ffprobe.IsAudioStream(&stream) {
			if foundAudioStream {
				// already have an audio stream, continue
				continue
			}
			foundAudioStream = true
			ripFromAudioStream(&stream, &mediaInfo)
		}
	}

	mediaInfo.IsAudioOnly = foundAudioStream && !foundVideoStream

	return mediaInfo, err
}
