package format

const formatMP3 = "mp3"
const formatOGG = "ogg"
const formatAAC = "aac"
const formatWAV = "wav"

var audioContainers = []string{
	formatMP3,
	formatOGG,
	formatAAC,
	formatWAV,
}

func isAudioContainer(container string) bool {
	for _, format := range audioContainers {
		if container == format {
			return true
		}
	}

	return false
}
