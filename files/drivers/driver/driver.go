package driver

// A PathHandler determines if it can handle some arbitrary path
type PathHandler interface {
	Handles(path string) bool
}

// A DownloadDriver does the heavy lifting of actually downloading things
type DownloadDriver interface {
	PathHandler
	Download(source string, dest string) error
}

// An UploadDriver does the heavy lifting of actually uploading things
type UploadDriver interface {
	PathHandler
	Upload(source string, dest string) error
}
