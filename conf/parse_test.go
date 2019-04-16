package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {
	cases := []struct {
		input     string
		variables map[string]string
		expected  string
	}{
		{
			input:     "$nothing",
			variables: map[string]string{},
			expected:  "$nothing",
		},
		{
			input: "some$thing",
			variables: map[string]string{
				"thing": "body",
			},
			expected: "somebody",
		},
		{
			input: "$a $couple $things $here!",
			variables: map[string]string{
				"here":   "me",
				"things": "told",
				"a":      "somebody",
				"couple": "once",
			},
			expected: "somebody once told me!",
		},
		{
			input: "$foo$foo$foo$foo",
			variables: map[string]string{
				"foo": "barbar",
			},
			expected: "barbarbarbarbarbarbarbar",
		},
		{
			input: "$foobar",
			variables: map[string]string{
				"foo":    "this is not really expected but I'm not sure how it should be handled ",
				"foobar": "it's a miracle",
			},
			expected: "it's a miracle",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			parseContext := parseContext{
				variables:           map[string]string{},
				sortedVariableNames: []string{},
			}

			for name, value := range c.variables {
				parseContext.setVariable(name, value)
			}

			actual := parseContext.expand(c.input)

			assert.EqualValues(t, c.expected, actual)
		})
	}
}

func TestRemoveComments(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "$nothing",
			expected: "$nothing",
		},
		{
			input:    "var something = foo #bar",
			expected: "var something = foo",
		},
		{
			input:    "var url = http://foo.bar/index.html#about-us",
			expected: "var url = http://foo.bar/index.html#about-us",
		},
		{
			input:    "-> jpg:300x = $base_s3/thumbnail_small_#num#.jpg, number=6",
			expected: "-> jpg:300x = $base_s3/thumbnail_small_#num#.jpg, number=6",
		},
		{
			input:    "# this is a comment",
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			actual := removeComments(c.input)

			assert.EqualValues(t, c.expected, actual)
		})
	}
}

func TestParseString(t *testing.T) {
	cases := []struct {
		name      string
		input     string
		expected  ParsedConf
		errorText string
	}{
		{
			name: "unhandled expression",
			input: `# this is a line
			mama mia
			`,
			expected:  ParsedConf{},
			errorText: "unhandled configuration expression on line 2: \t\t\tmama mia (parsed as mama mia)",
		},
		{
			name: "basic variable expansion",
			input: `
			var foo = bar
			set source = http://$foo.com
			`,
			expected: ParsedConf{
				SourceURL: "http://bar.com",
				Outputs:   map[string]ParsedOutput{},
			},
		},
		{
			name: "variable-in-variable expansion",
			input: `
			var foo = bar
			var http = http://
			var com = .com
			var source = $http$foo$com
			set source = $source
			`,
			expected: ParsedConf{
				SourceURL: "http://bar.com",
				Outputs:   map[string]ParsedOutput{},
			},
		},
		{
			name: "billion laughs variable expansion",
			input: `
			var foo = bar
			var bar = $foo$foo$foo$foo
			var foobar = $bar$bar$bar$bar

			set webhook = $foobar
			`,
			expected: ParsedConf{
				WebhookURL: "barbarbarbarbarbarbarbarbarbarbarbarbarbarbarbar",
				Outputs:    map[string]ParsedOutput{},
			},
		},
		{
			name: "a bad set",
			input: `
			set foo = bar
			`,
			expected:  ParsedConf{},
			errorText: "unhandled set name: foo",
		},
		{
			name: "double-set variable",
			input: `
			var foo = bar
			var foo = notbar

			set source = $foo
			`,
			expected: ParsedConf{
				SourceURL: "notbar",
				Outputs:   map[string]ParsedOutput{},
			},
		},
		{
			name: "double-set output",
			input: `
			-> mp4 = http://foo.bar/upload
			-> mp4 = http://bar.foo/upload
			`,
			expected:  ParsedConf{},
			errorText: "output already exists: mp4",
		},
		{
			name: "thumbnail with multiple images",
			input: `
			var thumbnails = 12
			var destination = s3://foo:bar@bucket
			-> jpg:300x = $destination/thumbnail_small_#num#.jpg, number=$thumbnails
			`,
			expected: ParsedConf{
				Outputs: map[string]ParsedOutput{
					"jpg:300x": ParsedOutput{
						URL: "s3://foo:bar@bucket/thumbnail_small_#num#.jpg",
						Options: map[string]string{
							"number": "12",
						},
					},
				},
			},
		},
		{
			name: "sample valid config file",
			input: `
			# sample valid creamy-transcode conf file
			var resource_id = 5678

			var s3_key = some-s3-key # key for our S3 bucket
			var s3_secret = some-s3-secret # secret for our S3 bucket
			var s3_bucket = some-s3-bucket

			var base_s3 = s3://$s3_key:$s3_secret@$s3_bucket/videos/$resource_id/transcode

			set source = https://creamy-videos.internal.albinodrought.com/static/videos/$resource_id/video
			set webhook = https://creamy-videos.internal.albinodrought.com/api/video/$resource_id/transcoded, metadata=true

			-> webm = $base_s3/video.webm
			-> mp4 = $base_s3/video.mp4
			-> mp4:720p = $base_s3/video_720p.mp4, if=$source_width >= 1280

			# thumbnails
			-> jpg:300x = $base_s3/thumbnail_small_#num#.jpg, number=6
			-> jpg:160x = $base_s3/thumbnail_sprite.jpg, every=5, sprite=yes, vtt=yes
			`,
			expected: ParsedConf{
				SourceURL:  "https://creamy-videos.internal.albinodrought.com/static/videos/5678/video",
				WebhookURL: "https://creamy-videos.internal.albinodrought.com/api/video/5678/transcoded, metadata=true",
				Outputs: map[string]ParsedOutput{
					"webm": ParsedOutput{
						URL:     "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/video.webm",
						Options: map[string]string{},
					},
					"mp4": ParsedOutput{
						URL:     "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/video.mp4",
						Options: map[string]string{},
					},
					"mp4:720p": ParsedOutput{
						URL: "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/video_720p.mp4",
						Options: map[string]string{
							"if": "$source_width >= 1280",
						},
					},
					"jpg:300x": ParsedOutput{
						URL: "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/thumbnail_small_#num#.jpg",
						Options: map[string]string{
							"number": "6",
						},
					},
					"jpg:160x": ParsedOutput{
						URL: "s3://some-s3-key:some-s3-secret@some-s3-bucket/videos/5678/transcode/thumbnail_sprite.jpg",
						Options: map[string]string{
							"every":  "5",
							"sprite": "yes",
							"vtt":    "yes",
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := ParseString(c.input)

			assert.EqualValues(t, c.expected, actual)
			if c.errorText == "" {
				assert.Nil(t, err)
			} else {
				assert.EqualError(t, err, c.errorText)
			}
		})
	}
}
