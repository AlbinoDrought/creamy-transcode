package transcoder

var formatOverrides = map[string]string{
	"mkv": "matroska",
}

func getFormat(container string) string {
	override, ok := formatOverrides[container]

	if ok {
		return override
	}

	return container
}
