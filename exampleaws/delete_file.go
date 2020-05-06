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
	// delete file in aws s3
	err := storage.Delete(context.Background(), "/example/image.png")
	fmt.Println("Error :", err)
}
