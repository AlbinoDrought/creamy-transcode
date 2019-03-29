package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/imroc/req"

	"github.com/AlbinoDrought/creamy-transcode/files/drivers/driver"
	"github.com/stretchr/testify/assert"
)

func TestHandles(t *testing.T) {
	cases := []struct {
		path    string
		handles bool
	}{
		{
			"",
			false,
		},
		{
			"C:\\Windows\\System32\\drivers\\etc\\hosts",
			false,
		},
		{
			"ftp://albinodrought.com/hueg.jpg",
			false,
		},
		{
			"http://albinodrought.com/",
			true,
		},
		{
			"https://say.hi.to.albinodrought.com/some?good=stuff%20here",
			true,
		},
	}

	for _, c := range cases {
		var name string
		if c.handles {
			name = fmt.Sprintf("should handle '%v'", c.path)
		} else {
			name = fmt.Sprintf("shouldn't handle '%v'", c.path)
		}

		t.Run(name, func(t *testing.T) {
			driver := Driver{}
			assert.Equal(t, c.handles, driver.Handles(c.path))
		})
	}
}

// should be a DownloadDriver
var _ driver.DownloadDriver = Driver{}

func TestDownload(t *testing.T) {
	expected := []byte("hello world")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expected)
	}))
	defer ts.Close()

	dir, err := ioutil.TempDir("", "httpDriver_TestDownload")
	assert.Nil(t, err)

	defer os.RemoveAll(dir)

	parsedURL, _ := url.Parse(ts.URL)
	parsedURL.User = url.UserPassword("foo", "bar")
	parsedURL.Path = "cream.txt?X-Access-Token=double-deluxe"

	source := parsedURL.String()
	dest := path.Join(dir, "foo.txt")

	driver := Driver{}
	err = driver.Download(source, dest)
	assert.Nil(t, err)

	contents, err := ioutil.ReadFile(dest)
	assert.Nil(t, err)

	assert.Equal(t, expected, contents)
}

// should be an UploadDriver
var _ driver.UploadDriver = Driver{}

func TestUpload(t *testing.T) {
	req.Debug = true
	cases := []struct {
		bytes       []byte
		filename    string
		httpPath    string
		formFileKey string
	}{
		{
			[]byte("i am a talking computer"),
			"sentience.txt",
			"foo",
			"file",
		},
		{
			[]byte("i am a walking computer"),
			"revolution.txt",
			"bar?http_file_name=thestuff",
			"thestuff",
		},
	}

	for i, c := range cases {
		si := strconv.Itoa(i)
		t.Run("upload case #"+si, func(t *testing.T) {
			// hack to wait for dummy server to receive request before leaving testcase
			receivedRequest := make(chan bool, 1)

			// boot up a thicc dummy server to receive and actually check uploads
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer func() {
					// allow testcase to finish
					receivedRequest <- true
				}()

				// make sure the path+query is correct
				assert.Truef(
					t,
					strings.HasSuffix(r.URL.String(), c.httpPath),
					"URL should end with %v",
					c.httpPath,
				)

				// check if the file is actually posted
				file, fileHeader, err := r.FormFile(c.formFileKey)
				assert.Nil(t, err)
				if err != nil {
					// prevent test from complete explosion
					return
				}

				// prevent leaving files on disk
				defer r.MultipartForm.RemoveAll()

				// ensure proper filename was sent
				assert.Equal(t, c.filename, fileHeader.Filename)

				// ensure proper bytes were sent
				receivedBytes, err := ioutil.ReadAll(file)
				assert.Nil(t, err)
				assert.Equal(t, c.bytes, receivedBytes)
			}))
			defer ts.Close()

			// create dummy folder to store fixtures
			dir, err := ioutil.TempDir("", "httpDriver_TestUpload_"+si)
			assert.Nil(t, err)

			defer os.RemoveAll(dir)

			source := path.Join(dir, c.filename)
			dest := ts.URL + "/" + c.httpPath

			// dump the actual fixture to disk
			err = ioutil.WriteFile(source, c.bytes, os.ModePerm)
			assert.Nil(t, err)

			// attempt the upload
			driver := Driver{}
			err = driver.Upload(source, dest)
			assert.Nil(t, err)

			select {
			case <-receivedRequest:
				return
			case <-time.After(3 * time.Second):
				t.Fatal("no upload request received after 3s")
			}
		})
	}
}
