package goproxy

import "io"

type KnownLengthReadCloser struct {
	io.ReadCloser
	l int64
}

func (r *KnownLengthReadCloser) Length() int64 {
	return r.l
}

func NewKnownLengthReadCloser(r io.ReadCloser, l int64) KnownLengthReadCloser {
	return KnownLengthReadCloser{r, l}
}
