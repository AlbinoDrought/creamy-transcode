package http

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/imroc/req"
	"github.com/pkg/errors"
)

// uploadFormDataName is the default key used when no other is given
const uploadFormDataName = "file"

// uploadParameterFormDataName is the name of the query parameter that
// overrides the default form data key.
//
// example url:
//
// 	https://albinodrought.com/something?http_file_name=foo
//  - file would be uploaded under the key "foo"
//
// https://albinodrought.com/something
// - file would be uploaded under the default key (uploadFormDataName const)
const uploadParameterFormDataName = "http_file_name"

// Driver handles downloading of files from HTTP sources
type Driver struct{}

// Handles returns true if the path is an HTTP resource
func (driver Driver) Handles(path string) bool {
	return strings.Index(path, "http://") == 0 || strings.Index(path, "https://") == 0
}

// Download the file from some HTTP source to a local destination
func (driver Driver) Download(source string, dest string) error {
	resp, err := http.Get(source)
	if err != nil {
		return errors.Wrapf(err, "error getting source %v", source)
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return errors.Wrapf(err, "error creating local file %v", dest)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return errors.Wrap(err, "error copying response to created local file")
	}

	return nil
}

// Upload the file from a local source to some HTTP destination
func (driver Driver) Upload(source string, dest string) error {
	parsedURL, err := url.Parse(dest)
	if err != nil {
		return errors.Wrapf(err, "unable to parse destination url %v", dest)
	}

	formDataName := uploadFormDataName

	// allow orverriding formdata name with "http_file_name" param in url
	if override := parsedURL.Query().Get(uploadParameterFormDataName); override != "" {
		formDataName = override
	}

	file, err := os.Open(source)
	if err != nil {
		return errors.Wrapf(err, "unable to open local source %v", source)
	}
	defer file.Close()

	// actually upload our file
	_, err = req.Post(
		dest,
		req.FileUpload{
			File:      file,
			FieldName: formDataName,
			FileName:  path.Base(source),
		},
	)

	if err != nil {
		return errors.Wrap(err, "unable to upload file to destination")
	}

	return nil
}
