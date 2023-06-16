package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/elgohr/go-localstack"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func getSession(service localstack.Service) *session.Session {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	l, err := localstack.NewInstance()
	if err != nil {
		log.Fatalf("Failed to connect to Docker %v", err)
	}
	if err := l.StartWithContext(ctx); err != nil {
		log.Fatalf("Failed to start localstack %v", err)
	}

	cfg := &aws.Config{
		Credentials: credentials.NewStaticCredentials("not", "empty", ""),
		DisableSSL:  aws.Bool(true),
		Region:      aws.String(endpoints.UsWest1RegionID),
		Endpoint:    aws.String(l.Endpoint(service)),
	}

	sess, err := session.NewSession(cfg)
	if err != nil {
		log.Fatalf("Failed to create session %v", err)
	}

	return sess
}

func makeDummyTextFile(text string) *os.File {
	b := []byte(text)
	f, err := os.Create("dummy.txt")
	if err != nil {
		log.Fatalf("Failed to create file %v", err)
	}
	defer f.Close()

	fmt.Fprintf(f, string(b))
	return f
}

func TestUpload(t *testing.T) {
	if err := initClient(getSession(localstack.S3)); err != nil {
		t.Fatalf("Failed to init client %v", err)
	}

	text := "hello world!"
	assert.Error(t, nil, Upload(makeDummyTextFile(text)))
}
