package main

import (
	"context"
	"fmt"
	"github.com/moh-fajri/qistorage"
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
