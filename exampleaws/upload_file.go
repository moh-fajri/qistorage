package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/moh-fajri/qistorage"
)

func main() {
	// get file
	file, err := os.Open("image.png")
	if err != nil {
		fmt.Println("Error File :", err)
	}
	defer file.Close()
	// convert file to byte
	fileBytes, _ := ioutil.ReadAll(file)
	// set credential aws s3
	storage := qistorage.NewAwsS3(&qistorage.AwsS3{
		AccessKey:  "access key",
		SecretKey:  "secret key",
		BucketName: "bucket name",
		Region:     "region",
	})
	// put file in aws s3
	err = storage.Put(context.Background(), "/example/image.png", fileBytes)
	fmt.Println("Error :", err)
}
