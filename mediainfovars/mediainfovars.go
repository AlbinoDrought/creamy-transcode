package mediainfovars

import (
	"fmt"
	"strings"

	"github.com/AlbinoDrought/creamy-transcode/mediainfo"
)

type mappedSourceVar struct {
	name  string
	value interface{}
}

func ConvertToConfigurationVariables(info mediainfo.MediaInfo) string {
	sourceVars := []mappedSourceVar{
		{"mime_type", info.MimeType},
		{"size", info.Size},
		{"width", info.Width},
		{"height", info.Height},
		{"aspect_ratio", info.AspectRatio},
		{"fps", info.FPS},
		{"video_bitrate", info.VideoBitrate},
		{"is_hd", info.IsHD},
		{"is_audio_only", info.IsAudioOnly},
		{"duration", info.Duration},
		{"video_codec", info.VideoCodec},
		{"audio_codec", info.AudioCodec},
		{"channels", info.Channels},
		{"audio_bitrate", info.AudioBitrate},
		{"sample_rate", info.SampleRate},
	}

	var sb strings.Builder

	for _, sourceVar := range sourceVars {
		sb.WriteString(fmt.Sprintf("var source_%v = %v\n", sourceVar.name, sourceVar.value))
	}

	return sb.String()
}
