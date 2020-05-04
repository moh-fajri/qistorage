package main

import (
	"fmt"
	"github.com/qasir-id/qistorage"
	"context"
)

func main() {
	storage := qistorage.NewAwsS3(&qistorage.AwsS3{
		AccessKey:  "access key",
		SecretKey:  "secret key",
		BucketName: "bucket name",
		Region:     "region",
	})

	err := storage.Delete(context.Background(), "/example/image.png")
	fmt.Println("Error :", err)
}

