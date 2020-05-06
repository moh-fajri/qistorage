package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/moh-fajri/qistorage"
)

func main() {
	//get file
	file, err := os.Open("example.png")
	if err != nil {
		fmt.Println("Error File :", err)
	}
	defer file.Close()
	// convert file to byte
	fileBytes, _ := ioutil.ReadAll(file)
	//set credential gcp
	storage := qistorage.NewGcpStorage(&qistorage.GcpStorage{
		BucketName: "bucket name",
		Credential: "credential base64",
	})
	// put file to gcp
	err = storage.Put(context.Background(), "example/image.png", fileBytes)
	fmt.Println("Error :", err)
}
