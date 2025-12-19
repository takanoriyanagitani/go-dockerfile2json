package dockerfile2json

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"

	fdp "github.com/moby/buildkit/frontend/dockerfile/parser"
)

type ParseResult struct{ *fdp.Result }

func (r ParseResult) ToJson(enc *json.Encoder) error {
	return enc.Encode(r.Result)
}

func (r ParseResult) ToWriter(wtr io.Writer) error {
	var bw *bufio.Writer = bufio.NewWriter(wtr)
	var enc *json.Encoder = json.NewEncoder(bw)
	return errors.Join(
		r.ToJson(enc),
		bw.Flush(),
	)
}

func (r ParseResult) ToStdout() error { return r.ToWriter(os.Stdout) }

type DockerfileReader struct{ io.Reader }

func (r DockerfileReader) Parse() (ParseResult, error) {
	res, e := fdp.Parse(r.Reader)
	return ParseResult{Result: res}, e
}

func (r DockerfileReader) ToJsonToStdout() error {
	parsed, e := r.Parse()
	if nil != e {
		return e
	}

	return parsed.ToStdout()
}

var DockerfileReaderStdin DockerfileReader = DockerfileReader{
	Reader: bufio.NewReader(os.Stdin),
}

func StdinToDockerfileToParsedToJsonToStdout() error {
	return DockerfileReaderStdin.ToJsonToStdout()
}
