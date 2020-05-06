package qistorage

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AwsS3 object
type AwsS3 struct {
	AccessKey  string
	SecretKey  string
	BucketName string
	Region     string
	Token      string
}

// configuration to set configuration aws s3
func (as *AwsS3) configuration() (*s3.S3, error) {
	creds := credentials.NewStaticCredentials(as.AccessKey, as.SecretKey, as.Token)
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(as.Region).WithCredentials(creds)

	return s3.New(session.New(), cfg), nil
}

// PutBase64 to upload the file with base64 to aws s3
func (as *AwsS3) PutBase64(ctx context.Context, path string, bs64 string) error {
	// convert base64 to file
	data, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		return err
	}
	// Put with byte
	err = as.Put(ctx, path, data)
	if err != nil {
		return err
	}

	return nil
}

// Put to upload the file to aws s3
func (as *AwsS3) Put(ctx context.Context, path string, byte []byte) error {
	// set configuration aws s3
	svc, err := as.configuration()
	if err != nil {
		return err
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(as.BucketName),
		Key:    aws.String(path),
		Body:   bytes.NewReader(byte),
		ACL:    aws.String("public-read"),
	}

	// Upload the file to S3.
	_, err = svc.PutObjectWithContext(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

// get to read file from aws s3
func (as *AwsS3) Get(ctx context.Context, path string) ([]byte, error) {
	// set configuration aws s3
	svc, err := as.configuration()
	if err != nil {
		return nil, err
	}
	input := &s3.GetObjectInput{
		Bucket: aws.String(as.BucketName),
		Key:    aws.String(path),
	}

	// Get the file from S3.
	res, err := svc.GetObjectWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, res.Body); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Delete to delete file from aws
func (as *AwsS3) Delete(ctx context.Context, path string) error {
	// set configuration aws s3
	svc, err := as.configuration()
	if err != nil {
		return err
	}
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(as.BucketName),
		Key:    aws.String(path),
	}

	// delete the file to S3.
	_, err = svc.DeleteObjectWithContext(ctx, input)
	if err != nil {
		return err
	}
	return nil
}

// NewAwsS3 create new AwsS3
func NewAwsS3(awsS3 *AwsS3) *AwsS3 {
	return awsS3
}
