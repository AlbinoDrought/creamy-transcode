package format

type VideoSpecs struct {
	Disabled         bool
	ResolutionHeight int
	ResolutionWidth  int
	BitrateKbps      int
	Codec            string
	FPS              string
}

type AudioSpecs struct {
	Disabled     bool
	Codec        string
	BitrateKbps  int
	SampleRateHz int
	AudioChannel string
}

type FormatOptions struct {
	TwoPass bool
}

type Format struct {
	Container     string
	VideoSpecs    VideoSpecs
	AudioSpecs    AudioSpecs
	FormatOptions FormatOptions
}
