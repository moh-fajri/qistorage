package main

import (
	"context"
	"fmt"
	"github.com/qasir-id/qistorage"
)

func main() {
	storage := qistorage.NewGcpStorage(&qistorage.GcpStorage{
		BucketName: "bucket name",
		Credential: "credential base64",
	})

	byte , err := storage.Get(context.Background(), "example/image.png")
	fmt.Println("Error :", err)
	fmt.Println("Byte :", byte)
}

