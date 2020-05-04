package main

import (
	"context"
	"fmt"
	"github.com/qasir-id/qistorage"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("image.png")
	if err != nil {
		fmt.Println("Error File :", err)
	}

	defer file.Close()

	fileBytes, _ := ioutil.ReadAll(file)

	storage := qistorage.NewAwsS3(&qistorage.AwsS3{
		AccessKey:  "access key",
		SecretKey:  "secret key",
		BucketName: "bucket name",
		Region:     "region",
	})

	err = storage.Put(context.Background(), "/example/image.png", fileBytes)
	fmt.Println("Error :", err)
}