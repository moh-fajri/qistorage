package main

import (
	"context"
	"fmt"
	"github.com/qasir-id/qistorage"
)

func main() {
	storage := qistorage.NewAwsS3(&qistorage.AwsS3{
		AccessKey:  "access key",
		SecretKey:  "secret key",
		BucketName: "bucket name",
		Region:     "region",
	})

	res, err := storage.Get(context.Background(), "/example/image.png")
	fmt.Println("Error :", err)
	fmt.Println("Result : ", res)
}
