package main

import (
	"context"
	"fmt"

	"github.com/moh-fajri/qistorage"
)

func main() {
	//set credential gcp
	storage := qistorage.NewGcpStorage(&qistorage.GcpStorage{
		BucketName: "bucket name",
		Credential: "credential base64",
	})
	// delete file in gcp
	err := storage.Delete(context.Background(), "example/image.png")
	fmt.Println("Error :", err)
}
