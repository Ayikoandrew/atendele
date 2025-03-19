package core

import "io"

type Decoder[T any] interface {
	Decode(r io.Reader, value T) (T, error)
}

type Encoder[T any] interface {
	Encode(w io.Writer, value T) error
}
