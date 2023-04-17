package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	var (
		instanceId string
		err        error
	)
	ctx := context.Background()
	if instanceId, err = createEc2(ctx, "us-east-1"); err != nil {
		fmt.Printf("Erro é %s", err)
		os.Exit(1)
	}
	fmt.Printf("a instancia é %s\n", instanceId)

}

//criar instancia ec2
func createEc2(ctx context.Context, region string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}
	ec2Client := ec2.NewFromConfig(cfg)
	keyPair, err := ec2Client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		KeyNames: []string{"aws-key"},
	})
	if err != nil && !strings.Contains(err.Error(), "InvalidKeyPair.NotFound") {
		return "", fmt.Errorf("describesKeyPair error: %s", err)
	}
	//Se chave não existir, é para criar
	if keyPair == nil || len(keyPair.KeyPairs) == 0 {
		//Criar par de chaves
		CreatedKeyPair, err := ec2Client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
			KeyName: aws.String("aws-key"),
		})
		if err != nil {
			return "", fmt.Errorf("CreateKey error: %s", err)
		}
		//Escrever o arquivo da chave para ter acesso a chave
		err = os.WriteFile("aws-key.pem", []byte(*CreatedKeyPair.KeyMaterial), 0400)
		if err != nil {
			return "", fmt.Errorf("WriteFile error: %s", err)
		}
	}
	if err != nil {
		return "", fmt.Errorf("key error: %s", err)
	}

	//Descrever imagem ubuntu
	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{"ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20230325"},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{"099720109477"},
	})
	if err != nil {
		return "", fmt.Errorf("image error: %s", err)
	}
	if len(imageOutput.Images) == 0 {
		return "", fmt.Errorf("image length is empty")
	}
	//Saida da instancia, pegar o id da imagem
	instanceOutput, err := ec2Client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      imageOutput.Images[0].ImageId,
		KeyName:      aws.String("aws-key"),
		InstanceType: types.InstanceTypeT2Micro,
		MinCount:     aws.Int32(1),
		MaxCount:     aws.Int32(1),
	})
	if err != nil {
		return "", fmt.Errorf("instance error: %s", err)
	}
	if len(instanceOutput.Instances) == 0 {
		return "", fmt.Errorf("image length is empty")
	}
	return *instanceOutput.Instances[0].InstanceId, nil //retorna o id da instancia
}
