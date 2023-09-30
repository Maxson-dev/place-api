package s3

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

func (c *client) UploadFile(bucket, key, contentType string, body io.ReadSeeker) (string, error) {
	_, err := c.s3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        body,
	})
	if err != nil {
		return "", errors.Wrap(err, "c.s3.PutObject")
	}

	return fmt.Sprintf("%s/%s/%s", c.cfg.Endpoint, bucket, key), nil
}
