package ffprobe

import (
	"encoding/json"
	"os/exec"

	"github.com/pkg/errors"
)

// Stream is a golang-compatible representation of the JSON stream output from FFProbe
type Stream struct {
	Index              int    `json:"index"`
	CodecName          string `json:"codec_name"`
	CodecLongName      string `json:"codec_long_name"`
	Profile            string `json:"profile"`
	CodecType          string `json:"codec_type"`
	CodecTimeBase      string `json:"codec_time_base"`
	CodecTagString     string `json:"codec_tag_string"`
	CodecTag           string `json:"codec_tag"`
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	CodedWidth         int    `json:"coded_width,omitempty"`
	CodedHeight        int    `json:"coded_height,omitempty"`
	HasBFrames         int    `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
	PixFmt             string `json:"pix_fmt,omitempty"`
	Level              int    `json:"level,omitempty"`
	ChromaLocation     string `json:"chroma_location,omitempty"`
	Refs               int    `json:"refs,omitempty"`
	IsAvc              string `json:"is_avc,omitempty"`
	NalLengthSize      string `json:"nal_length_size,omitempty"`
	RFrameRate         string `json:"r_frame_rate"`
	AvgFrameRate       string `json:"avg_frame_rate"`
	TimeBase           string `json:"time_base"`
	StartPts           int    `json:"start_pts"`
	StartTime          string `json:"start_time"`
	DurationTs         int    `json:"duration_ts"`
	Duration           string `json:"duration"`
	BitRate            string `json:"bit_rate"`
	BitsPerRawSample   string `json:"bits_per_raw_sample,omitempty"`
	NbFrames           string `json:"nb_frames"`
	Disposition        struct {
		Default         int `json:"default"`
		Dub             int `json:"dub"`
		Original        int `json:"original"`
		Comment         int `json:"comment"`
		Lyrics          int `json:"lyrics"`
		Karaoke         int `json:"karaoke"`
		Forced          int `json:"forced"`
		HearingImpaired int `json:"hearing_impaired"`
		VisualImpaired  int `json:"visual_impaired"`
		CleanEffects    int `json:"clean_effects"`
		AttachedPic     int `json:"attached_pic"`
		TimedThumbnails int `json:"timed_thumbnails"`
	} `json:"disposition"`
	SampleFmt     string `json:"sample_fmt,omitempty"`
	SampleRate    string `json:"sample_rate,omitempty"`
	Channels      int    `json:"channels,omitempty"`
	ChannelLayout string `json:"channel_layout,omitempty"`
	BitsPerSample int    `json:"bits_per_sample,omitempty"`
}

// Format is a golang-consumable representation of the JSON format output from FFProbe
type Format struct {
	Filename       string `json:"filename"`
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      string `json:"start_time"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
}

// Result is a golang-consumable representation of the JSON output from FFProbe
type Result struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

func decode(command *exec.Cmd) (Result, error) {
	output, err := command.Output()
	if err != nil {
		return Result{}, errors.Wrap(err, "unable to run ffprobe and get output")
	}

	result := Result{}
	err = json.Unmarshal(output, &result)
	if err != nil {
		return Result{}, errors.Wrap(err, "unable to decode ffprobe result")
	}

	return result, nil
}

/*
// ProbeStream uses ffprobe to fetch media info from a stream
func ProbeStream(stream io.Reader) (Result, error) {
	cmd := exec.Command("ffprobe", "-", "-hide_banner", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams")
	cmd.Stdin = stream

	return decode(cmd)
}
*/

// ProbeFile uses ffprobe to fetch media info from an ffprobe-compatible path
func ProbeFile(path string) (Result, error) {
	cmd := exec.Command("ffprobe", path, "-hide_banner", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams")

	return decode(cmd)
}
