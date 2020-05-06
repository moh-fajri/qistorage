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
	// get file from gcp return byte
	byte, err := storage.Get(context.Background(), "example/image.png")
	fmt.Println("Error :", err)
	fmt.Println("Byte :", byte)
}
