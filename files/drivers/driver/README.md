# creamy-transcode/files/drivers/driver

Driver interfaces that allow downloading from some source to a local destination,
and uploading from some local source to some destination.

If a Driver can `Handle` a path, it can be used to `Upload` to it or `Download` from it,
depending if it is a `DownloadDriver`, and `UploadDriver`, or both.
