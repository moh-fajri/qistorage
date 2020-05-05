# qistorage

qistorage is file storage that uses aws s3 and GCP Storage.

## Instalation

When used with Go modules, use the following import path:

```
go get -u github.com/moh-fajri/qistorage
```

## Basic Usage

storage using AWS S3
```go
import "github.com/moh-fajri/qistorage"

storage := qistorage.NewAwsS3(&qistorage.AwsS3{
    AccessKey:  "access key",
    SecretKey:  "seccret key",
    BucketName: "bucket name",
    Region:     "region",
})
err := storage.Put(context.Background(), "path", byte)
```

storage using GCP Storage
```go
import "github.com/moh-fajri/qistorage"

storage := qistorage.NewGcpStorage(&qistorage.GcpStorage{
    BucketName:"bucket name",
    Credential:"credential base64", // convert file credential.json to base64 --> https://www.base64decode.org/
})
err := storage.Put(context.Background(), "path", byte)
```
