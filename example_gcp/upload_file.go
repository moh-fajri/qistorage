package main

import (
	"context"
	"fmt"
	"github.com/qasir-id/qistorage"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("exmaple.png")
	if err != nil {
		fmt.Println("Error File :", err)
	}
	defer file.Close()

	fileBytes, _ := ioutil.ReadAll(file)

	storage := qistorage.NewGcpStorage(&qistorage.GcpStorage{
		BucketName: "bucket name",
		Credential: "credential base64",
	})

	err = storage.Put(context.Background(), "example/image.png", fileBytes)
	fmt.Println("Error :", err)
}