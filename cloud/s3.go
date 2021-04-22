package cloud

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/javiertlopez/numeral"
)

func (s *storage) PutImage(ctx context.Context, countImage numeral.Image) (numeral.Image, error) {
	reader := bytes.NewReader(countImage.Binary)

	input := &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &countImage.KeyID,
		Body:   reader,
	}

	_, err := s.client.PutObject(ctx, input)
	if err != nil {
		return numeral.Image{}, err
	}

	return countImage, nil
}
