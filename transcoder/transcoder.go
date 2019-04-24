package transcoder

import (
	"github.com/AlbinoDrought/creamy-transcode/conf"
	cfmt "github.com/AlbinoDrought/creamy-transcode/format"
	"github.com/AlbinoDrought/creamy-transcode/mediainfo"
	"github.com/AlbinoDrought/creamy-transcode/transcoder/transcoderoptions"
)

type TranscodeRequest struct {
	Format             *cfmt.Format
	SourceLocalPath    string
	SourceMediaInfo    *mediainfo.MediaInfo
	TemporaryLocalPath string
	ParsedOutput       *conf.ParsedOutput
	ThumbnailOptions   *transcoderoptions.ThumbnailOptions
}

type TranscodeResult struct {
	Request        *TranscodeRequest
	ResultingFiles []string
	Error          error
}

type Transcoder interface {
	Transcode(request *TranscodeRequest) TranscodeResult
}
