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
	//base64 string without data:image/png;base64
	bs64 := "iVBORw0KGgoAAAANSUhEUgAACy4A"
	// put file in aws s3
	err := storage.PutBase64(context.Background(), "/example/image.png", bs64)
	fmt.Println("Error :", err)
}
