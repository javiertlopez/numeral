package numeral

import "context"

// Image struct
type Image struct {
	KeyID       string
	ContentType string
	Binary      []byte
}

// Storage interface
type Storage interface {
	PutImage(ctx context.Context, countImage Image) (Image, error)
}
