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
	//base64 string without data:image/png;base64
	bs64 := "iVBORw0KGgoAAAANSUhEUgAACy4A"
	// put file to gcp
	err := storage.PutBase64(context.Background(), "example/image.png", bs64)
	fmt.Println("Error :", err)
}
