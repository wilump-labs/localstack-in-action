package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

var client *s3.S3

func initClient(sess ...*session.Session) error {
	if len(sess) > 0 {
		inputSession := sess[0]
		client = s3.New(inputSession)
		return nil
	}

	cfg := &aws.Config{
		Credentials: credentials.NewStaticCredentials("not", "empty", ""),
		DisableSSL:  aws.Bool(true),
		Region:      aws.String(endpoints.UsWest1RegionID),
		Endpoint:    aws.String("http://localhost:4566"),
	}

	defaultSession, err := session.NewSession(cfg)
	if err != nil {
		return err
	}

	client = s3.New(defaultSession)
	return nil
}

func Upload(file *os.File) error {
	// TODO implement
	return nil
}
