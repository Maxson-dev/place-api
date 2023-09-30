package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

type Config struct {
	Region   string
	Endpoint string

	AccessKey string
	SecretKey string
}

type client struct {
	cfg  Config
	sess *session.Session
	s3   *s3.S3
}

// New создает клиент s3
func New(cfg Config) (*client, error) {
	s3Session, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint:         aws.String(cfg.Endpoint),
			S3ForcePathStyle: aws.Bool(true),
			Region:           aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(
				cfg.AccessKey,
				cfg.SecretKey,
				"",
			),
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "session.NewSessionWithOptions")
	}
	return &client{
		cfg:  cfg,
		sess: s3Session,
		s3:   s3.New(s3Session, aws.NewConfig()),
	}, nil
}
