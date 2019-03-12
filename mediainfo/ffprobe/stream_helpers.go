package ffprobe

const codecTypeAudio = "audio"
const codecTypeVideo = "video"

const minimumHDVideoHeight = 720
const minimumHDVideoWidth = 1280

// IsAudioStream returns true if the stream is an audio stream
func IsAudioStream(stream *Stream) bool {
	return stream.CodecType == codecTypeAudio
}

// IsVideoStream returns true if the stream is a video stream
func IsVideoStream(stream *Stream) bool {
	return stream.CodecType == codecTypeVideo
}

// IsHDVideoStream returns true if the stream is a high-definition video stream (min 720p)
func IsHDVideoStream(stream *Stream) bool {
	return IsVideoStream(stream) && stream.Width >= minimumHDVideoWidth && stream.Height >= minimumHDVideoHeight
}
