# creamy-transcode

Simple self-hostable media transcoding SaaSS 

## Building

```sh
go get
go build
```

## Testing

```sh
# get all dependencies of all packages including test dependencies
go get -t ./...
# vet all packages
go vet ./...
# test all packages
go test ./...
```

## Roadmap

- [x] parse a format like `mp4:1080p:mp3:2pass` into something machine-readable
- [ ] when given this basic parsed format, some secondary options, and a source video URL, convert the source video into the target format
- [ ] when given a source video URL, download it and extract metadata like width, length, etc
- [x] parse a config file into something machine-readable
- [ ] shove source video metadata into config file vars
- [ ] working `if` statements in config file
- [ ] somehow handle "secondary url options" like `, metadata=true`, `, number=6`, etc
- [ ] post with-metadata and without-metadata webhooks on state changes
- [ ] upload to S3 when given a file and an S3 url like `s3://access:secret@bucket/video.mp4`
- [ ] also handle "unofficial" S3-compatible services like Minio
- [ ] output thumbnails using format `-> jpg:300x = $base_s3/thumbnail_small_#num#.jpg, number=6`
- [ ] WebVTT thumbnails/metadata
