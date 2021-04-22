package cloud

import (
	"github.com/javiertlopez/numeral"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// repository struct holds the MongoDB client
type storage struct {
	bucket string
	client *s3.Client
}

// New returns a storage instance
func New(
	b string,
	c *s3.Client,
) numeral.Storage {
	return &storage{
		bucket: b,
		client: c,
	}
}
