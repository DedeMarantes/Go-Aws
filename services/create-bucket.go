package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// const bucketName = "bucket-teste-go"
const region = "us-east-1"

func CreateS3Bucket(ctx context.Context, bucketNames []string, isPublic bool) error {
	s3Client, err := InitS3Client(ctx)
	if err != nil {
		return fmt.Errorf("Erro no s3Client: %s\n", err)
	}
	for _, bucketName := range bucketNames {
		input := &s3.CreateBucketInput{
			Bucket: aws.String(bucketName),
			//Regiao onde o bucket será criado
			/* CreateBucketConfiguration: &types.CreateBucketConfiguration{
				LocationConstraint: types.BucketLocationConstraintAfSouth1,
			}, */
		}
		if isPublic {
			input.ACL = types.BucketCannedACLPublicRead
		}
		//criar bucket através do cliente s3
		found, err := checkS3Bucket(ctx, s3Client, bucketName)
		if err != nil {
			return fmt.Errorf("Erro ao checar bucket %s", err)
		}
		//Se bucket nao for encontrado criar um novo
		if !found {
			_, err := s3Client.CreateBucket(ctx, input)
			if err != nil {
				return fmt.Errorf("Erro ao criar bucket: %s", err)
			}
		}
		fmt.Printf("bucket %s criado\n", bucketName)
	}
	return nil
}

func InitS3Client(ctx context.Context) (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("Erro ao configurar: %s", err)
	}
	return s3.NewFromConfig(cfg), nil
}

func checkS3Bucket(ctx context.Context, client *s3.Client, bucketName string) (bool, error) {
	allBuckets, err := client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		return false, fmt.Errorf("List buckets erro: %s", err)
	}
	found := false
	for _, bucket := range allBuckets.Buckets {
		if *bucket.Name == bucketName {
			found = true
			fmt.Printf("Bucket já existe\n")
		}
	}
	return found, nil
}
