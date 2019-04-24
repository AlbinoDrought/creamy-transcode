package transcoder

// ApostleTranscoder determines the proper transcoder to use based on the TranscodeRequest
type ApostleTranscoder struct {
	ContainerTranscoders map[string]Transcoder
	DefaultTranscoder    Transcoder
}

// Transcode the TranscodeRequest based on config
func (transcoder ApostleTranscoder) Transcode(request *TranscodeRequest) TranscodeResult {
	container := request.Format.Container
	containerTranscoder, ok := transcoder.ContainerTranscoders[container]
	if ok {
		return containerTranscoder.Transcode(request)
	}

	return transcoder.DefaultTranscoder.Transcode(request)
}
