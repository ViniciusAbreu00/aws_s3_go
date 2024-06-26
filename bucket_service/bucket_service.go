package bucket_service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3client *s3.Client

func inits3Client() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("ViniciusAbreu"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	s3client = s3.NewFromConfig(cfg)

}

func GetBucketContent() ([]string, error) {
	if s3client == nil {
		inits3Client()
	}

	output, err := s3client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{Bucket: aws.String("seiglu-solutions")})
	if err != nil {
		return nil, err
	}

	var data []string
	for _, object := range output.Contents {
		data = append(data, fmt.Sprintf("key=%s size=%d", aws.ToString(object.Key), object.Size))
	}

	return data, nil
}

func UploadFile(pathname string) (string, error) {
	if s3client == nil {
		inits3Client()
	}

	file, err := os.Open(pathname)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = s3client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("seiglu-solutions"),
		Key:    aws.String(pathname),
		Body:   file,
	})
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("File uploaded successfully: %s", pathname)

	return message, nil
}
