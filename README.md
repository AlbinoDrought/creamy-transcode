# Creamy Transcode

<a href="https://travis-ci.org/AlbinoDrought/creamy-transcode"><img alt="Travis build status" src="https://img.shields.io/travis/AlbinoDrought/creamy-transcode.svg?maxAge=360"></a>
<a href="https://github.com/AlbinoDrought/creamy-transcode/blob/master/LICENSE"><img alt="AGPL-3.0 License" src="https://img.shields.io/github/license/AlbinoDrought/creamy-transcode"></a>

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
- [x] when given an HTTP(S) source video url, download it
- [x] when given a source video path, extract metadata like width, length, etc: [mediainfo](mediainfo)
- [x] parse a config file into something machine-readable
- [x] shove source video metadata into config file vars
- [x] somehow parse "secondary url options" like `, metadata=true`, `, number=6`, etc
- [ ] actually handle these parsed "secondary url options"
- [ ] working `if` statements in config file
- [ ] output thumbnails using format `-> jpg:300x = $base_s3/thumbnail_small_#num#.jpg, number=6`
- [ ] post with-metadata and without-metadata webhooks on state changes
- [ ] upload to S3 when given a file and an S3 url like `s3://access:secret@bucket/video.mp4`
- [ ] also handle "unofficial" S3-compatible services like Minio
- [ ] download from and upload to FTP, SFTP locations
- [ ] WebVTT thumbnails/metadata
