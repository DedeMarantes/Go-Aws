package services

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//bucketName = "bucket-teste-go"

func UploadS3Bucket(ctx context.Context, fileName string, bucketNames []string) error {
	file, err := os.Open(fileName)
	client, err := InitS3Client(ctx)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("Erro ao abrir arquivo: %s", err)
	}
	for _, bucket := range bucketNames {
		_, err = client.PutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(fileName),
			Body:   file,
		})
		fmt.Printf("Upload de %s para %s feito com sucesso\n", fileName, bucket)
	}
	if err != nil {
		return fmt.Errorf("Erro ao colocar objeto: %s", err)
	}
	return nil
}
