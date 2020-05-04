package qistorage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
)

type GcpStorage struct {
	BucketName string
	Credential string
}

// configuration to set configuration aws s3
func (gs *GcpStorage) configuration(ctx context.Context) (*storage.Client, error) {
	creds, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(gs.Credential)
	if err != nil {
		return nil, err
	}
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(creds))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (gs *GcpStorage) Put(ctx context.Context, path string, fileByte []byte) error {
	client, err := gs.configuration(ctx)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(fileByte)

	wc := client.Bucket(gs.BucketName).Object(path).NewWriter(ctx)
	if _, err = io.Copy(wc, reader); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (gs *GcpStorage) Get(ctx context.Context, path string) ([]byte, error) {
	client, err := gs.configuration(ctx)
	if err != nil {
		return nil, err
	}
	rc, err := client.Bucket(gs.BucketName).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (gs *GcpStorage) Delete(ctx context.Context, path string) error{
	client, err := gs.configuration(ctx)
	if err != nil {
		return err
	}
	o := client.Bucket(gs.BucketName).Object(path)
	if err := o.Delete(ctx); err != nil {
		return err
	}
	return nil
}

// NewGcpStorage create new GcpStorage
func NewGcpStorage(gcpStorage *GcpStorage) *GcpStorage {
	return gcpStorage
}