# creamy-transcode/files/drivers/http

HTTP Driver that allows downloading from remote HTTP sources and uploading
to remote HTTP destinations.

## Usage

### Download

```golang
err := driver.Download("https://albinodrought.com/", "hi.html")
```

### Upload

Files are POSTed to the given URL under the FormData key `file`.

To change this, use the `http_file_name` query parameter.

```golang
// default formdata key
err := driver.Upload("hi.html", "https://albinodrought.com/")

// custom formdata key
err := driver.Upload("hi.html", "https://albinodrought.com/?http_file_name=encoded_video")
```

## Example

```golang
package main

import (
	"log"

	"github.com/AlbinoDrought/creamy-transcode/files/drivers/http"
	"github.com/imroc/req"
)

func main() {
	req.Debug = true

	driver := http.Driver{}

	err := driver.Download("https://albinodrought.com/", "hi.html")
	if err != nil {
		log.Fatalf("failed to download: %v", err)
	}

	err = driver.Upload("hi.html", "https://albinodrought.com/")
	if err != nil {
		log.Fatalf("failed to upload: %v", err)
	}

	err = driver.Upload("hi.html", "https://albinodrought.com/?http_file_name=honey")
	if err != nil {
		log.Fatalf("failed to upload with different key: %v", err)
	}
}

// Output: (the 405 responses are because albinodrought.com doesn't accept POST requests)
/*
POST / HTTP/1.1
Host: albinodrought.com
User-Agent: Go-http-client/1.1
Transfer-Encoding: chunked
Content-Type: multipart/form-data; boundary=15c3dc206cd19609cb2c69111093795ae2ea8264593ec8685f62972b4d16
Accept-Encoding: gzip

--1ba46da54b9bdfa9c4ff22b04cb1d5058c13decedddf416ca11f301e1c20
Content-Disposition: form-data; name="file"; filename="hi.html"
Content-Type: application/octet-stream

******
--1ba46da54b9bdfa9c4ff22b04cb1d5058c13decedddf416ca11f301e1c20--


=================================

HTTP/1.1 405 Not Allowed
Content-Length: 157
Content-Type: text/html
Date: Thu, 28 Mar 2019 05:15:23 GMT
Server: nginx/1.15.9
Strict-Transport-Security: max-age=16000000

<html>
<head><title>405 Not Allowed</title></head>
<body>
<center><h1>405 Not Allowed</h1></center>
<hr><center>nginx/1.15.9</center>
</body>
</html>

POST /?http_file_name=honey HTTP/1.1
Host: albinodrought.com
User-Agent: Go-http-client/1.1
Transfer-Encoding: chunked
Content-Type: multipart/form-data; boundary=1dd1a1cf77ff01136a184a55982a5603ffd8c7cffa202b616a47f3515995
Accept-Encoding: gzip

--2b37ce6cd7bcebacfed0a5c84b0d3a512a6516bd379db42656806a7c6dfb
Content-Disposition: form-data; name="honey"; filename="hi.html"
Content-Type: application/octet-stream

******
--2b37ce6cd7bcebacfed0a5c84b0d3a512a6516bd379db42656806a7c6dfb--


=================================

HTTP/1.1 405 Not Allowed
Content-Length: 157
Content-Type: text/html
Date: Thu, 28 Mar 2019 05:15:23 GMT
Server: nginx/1.15.9
Strict-Transport-Security: max-age=16000000

<html>
<head><title>405 Not Allowed</title></head>
<body>
<center><h1>405 Not Allowed</h1></center>
<hr><center>nginx/1.15.9</center>
</body>
</html>
*/
```