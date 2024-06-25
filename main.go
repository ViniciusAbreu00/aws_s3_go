package main

import (
	"aws_s3_go/bucket_service"
	"fmt"
	"log"
)

func searchAndUploadFile(bucket, key, filename string, mode string) string {
	switch mode {
	case "upload":
		mess, err := bucket_service.UploadFile(filename)
		if err != nil {
			log.Fatalf("error uploading file: %f", err)
		}
		return mess
	case "search":
		return searchFile(bucket, key, filename)
	default:
		return fmt.Errorf("invalid mode, must be 'upload' or 'search'")
	}
}
