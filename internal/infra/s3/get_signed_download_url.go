package s3

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

func (c *client) GetSignedDownloadURL(bucket, key, name, contentType string, expire time.Duration) (string, error) {
	req, _ := c.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket:                     aws.String(bucket),
		Key:                        aws.String(key),
		ResponseContentType:        aws.String(contentType),
		ResponseContentDisposition: aws.String(fmt.Sprintf(`attachment; filename="%s"`, name)),
	})
	signURL, err := req.Presign(expire)
	if err != nil {
		return "", errors.Wrap(err, "req.Presign")
	}
	return signURL, nil
}
