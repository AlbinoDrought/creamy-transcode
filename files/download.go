package files

import (
	"fmt"

	"github.com/AlbinoDrought/creamy-transcode/files/drivers/driver"
	"github.com/AlbinoDrought/creamy-transcode/files/drivers/http"
)

var downloadDrivers = []driver.DownloadDriver{
	http.Driver{},
}

// Download a file from the remote source to the local destination
func Download(source string, dest string) error {
	for _, downloadDriver := range downloadDrivers {
		if downloadDriver.Handles(source) {
			return downloadDriver.Download(source, dest)
		}
	}

	return fmt.Errorf("unhandled source type: %v", source)
}
