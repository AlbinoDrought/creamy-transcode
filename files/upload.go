package files

import (
	"fmt"

	"github.com/AlbinoDrought/creamy-transcode/files/drivers/driver"
	"github.com/AlbinoDrought/creamy-transcode/files/drivers/http"
)

var uploadDrivers = []driver.UploadDriver{
	http.Driver{},
}

// Upload a file from the local source to the remote destination
func Upload(source string, dest string) error {
	for _, uploadDriver := range uploadDrivers {
		if uploadDriver.Handles(dest) {
			return uploadDriver.Upload(source, dest)
		}
	}

	return fmt.Errorf("unhandled destination type: %v", source)
}
