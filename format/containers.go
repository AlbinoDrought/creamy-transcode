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

var containerAliases = map[string]string{
	"divx":   "avi",
	"xvid":   "avi",
	"wmv":    "asf",
	"hls":    "mpegts",
	"flash":  "flv",
	"theora": "ogv",
}

func unaliasContainer(container string) string {
	unalias, ok := containerAliases[container]

	if ok {
		return unalias
	}

	return container
}

var defaultOptionsByContainer = map[string]Format{
	"mp4": Format{
		Container: "mp4",
		VideoSpecs: VideoSpecs{
			Codec:       "h264",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "aac",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"webm": Format{
		Container: "webm",
		VideoSpecs: VideoSpecs{
			Codec:       "vp8",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "vorbis",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"avi": Format{
		Container: "avi",
		VideoSpecs: VideoSpecs{
			Codec:       "mpeg4",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "mp3",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"asf": Format{
		Container: "asf",
		VideoSpecs: VideoSpecs{
			Codec:       "wmv2",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "wmav2",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"mpegts": Format{
		Container: "mpegts",
		VideoSpecs: VideoSpecs{
			Codec:       "h264",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "aac",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"mov": Format{
		Container: "mov",
		VideoSpecs: VideoSpecs{
			Codec:       "h264",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "aac",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"flv": Format{
		Container: "flv",
		VideoSpecs: VideoSpecs{
			Codec:       "flv",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "mp3",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"mkv": Format{
		Container: "mkv",
		VideoSpecs: VideoSpecs{
			Codec:       "h264",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "aac",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"3gp": Format{
		Container: "3gp",
		VideoSpecs: VideoSpecs{
			Codec:       "h263",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "aac",
			SampleRateHz: 44100,
			BitrateKbps:  32,
			AudioChannel: audioChannelStereo,
		},
	},
	"ogv": Format{
		Container: "ogv",
		VideoSpecs: VideoSpecs{
			Codec:       "theora",
			BitrateKbps: 1000,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "vorbis",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"ogg": Format{
		Container: "ogg",
		VideoSpecs: VideoSpecs{
			Disabled: true,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "vorbis",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"mp3": Format{
		Container: "mp3",
		VideoSpecs: VideoSpecs{
			Disabled: true,
		},
		AudioSpecs: AudioSpecs{
			Codec:        "mp3",
			SampleRateHz: 44100,
			BitrateKbps:  128,
			AudioChannel: audioChannelStereo,
		},
	},
	"jpg": Format{
		Container: "jpg",
		VideoSpecs: VideoSpecs{
			Codec: "jpg",
		},
		AudioSpecs: AudioSpecs{
			Disabled: true,
		},
	},
	"png": Format{
		Container: "png",
		VideoSpecs: VideoSpecs{
			Codec: "png",
		},
		AudioSpecs: AudioSpecs{
			Disabled: true,
		},
	},
	"gif": Format{
		Container: "gif",
		VideoSpecs: VideoSpecs{
			Codec: "gif",
		},
		AudioSpecs: AudioSpecs{
			Disabled: true,
		},
	},
	"storyboard": Format{
		Container: "storyboard",
		VideoSpecs: VideoSpecs{
			Codec: "png",
		},
		AudioSpecs: AudioSpecs{
			Disabled: true,
		},
	},
}

func getDefaultOptionsForContainer(container string) Format {
	options, ok := defaultOptionsByContainer[container]

	if !ok {
		options = Format{
			Container: container,
		}
	}

	return options
}
