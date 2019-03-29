# creamy-transcode/files

For when you just want to skeet files around :tm:

This package combines the inner `creamy-transcode/files/drivers`, automatically picking the
correct file driver to use, if any.

## Usage

### Download

Configuration can vary by driver. See [specific driver implementations](drivers) for details.

```golang
err := files.Download("https://albinodrought.com/", "hi.html")
```

### Upload

Configuration can vary by driver. See [specific driver implementations](drivers) for details.

```golang
err := files.Upload("hi.html", "https://albinodrought.com/")
```

## Example

```golang
package main

import (
	"log"

	"github.com/AlbinoDrought/creamy-transcode/files"
)

func main() {
	err := files.Download("https://albinodrought.com/", "hi.html")
	if err != nil {
		log.Fatalf("failed to download: %v", err)
	}

	err = files.Upload("hi.html", "https://albinodrought.com/")
	if err != nil {
		log.Fatalf("failed to upload: %v", err)
	}
}

// Output: Nothing, it just worked :)
```