package video

import (
	"bytes"
	"io"
)

// TODO
type Video struct {
	input  io.Reader
	output io.Writer
}

func New(raw []byte) *Video {
	return &Video{
		input: bytes.NewReader(raw),
	}
}
