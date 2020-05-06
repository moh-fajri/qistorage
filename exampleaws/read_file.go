package main

import (
	"context"
	"fmt"

	"github.com/moh-fajri/qistorage"
)

func main() {
	// set credential aws s3
	storage := qistorage.NewAwsS3(&qistorage.AwsS3{
		AccessKey:  "access key",
		SecretKey:  "secret key",
		BucketName: "bucket name",
		Region:     "region",
	})
	// get file from aws s3 return byte
	res, err := storage.Get(context.Background(), "/example/image.png")
	fmt.Println("Error :", err)
	fmt.Println("Result : ", res)
}
