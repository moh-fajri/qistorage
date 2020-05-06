package qistorage

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"io/ioutil"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
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

// PutBase64 to upload the file with base64 to gcp storage
func (gs *GcpStorage) PutBase64(ctx context.Context, path string, bs64 string) error {
	// convert base64 to file
	data, err := base64.StdEncoding.DecodeString(bs64)
	if err != nil {
		return err
	}
	// Put with byte
	err = gs.Put(ctx, path, data)
	if err != nil {
		return err
	}

	return nil
}

// Put to upload the file with byte to gcp storage
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

// Get to Get the file from gcp storage
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

// Delete to Delete the file from gcp storage
func (gs *GcpStorage) Delete(ctx context.Context, path string) error {
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
