package services

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

// criar instancia ec2
func CreateEc2(ctx context.Context, region, name, owner, keyname string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return "", fmt.Errorf("unable to load SDK config, %v", err)
	}
	ec2Client := ec2.NewFromConfig(cfg)
	//Checar se chave existe, se não existe criar uma
	err = checkKey(ctx, ec2Client, keyname)
	if err != nil {
		return "", fmt.Errorf("Erro ao checar chave: %s", err)
	}
	//Descrever imagem ubuntu
	imageOutput, err := ec2Client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: []types.Filter{
			{
				Name:   aws.String("name"),
				Values: []string{name},
			},
			{
				Name:   aws.String("virtualization-type"),
				Values: []string{"hvm"},
			},
		},
		Owners: []string{owner},
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
		KeyName:      aws.String(keyname),
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
	fmt.Printf("a instancia é %s\n", *instanceOutput.Instances[0].InstanceId)
	return *instanceOutput.Instances[0].InstanceId, nil //retorna o id da instancia
}

func CreateUbuntuInstance(ctx context.Context, region, keyname string) (string, error) {
	owner := "099720109477"
	imageName := "ubuntu/images/hvm-ssd/ubuntu-jammy-22.04-amd64-server-20230325"
	return CreateEc2(ctx, region, imageName, owner, keyname)
}

func CreateRedHatInstance(ctx context.Context, region, keyname string) (string, error) {
	owner := "309956199498"
	imageName := "RHEL-9.0.0_HVM-20220513-x86_64-0-Hourly2-GP2"
	return CreateEc2(ctx, region, imageName, owner, keyname)
}

func checkKey(ctx context.Context, client *ec2.Client, keyname string) error {
	keyPair, err := client.DescribeKeyPairs(ctx, &ec2.DescribeKeyPairsInput{
		KeyNames: []string{keyname},
	})
	if err != nil && !strings.Contains(err.Error(), "InvalidKeyPair.NotFound") {
		return fmt.Errorf("describesKeyPair error: %s", err)
	}
	//Se chave não existir, é para criar
	if keyPair == nil || len(keyPair.KeyPairs) == 0 {
		//Criar par de chaves
		CreatedKeyPair, err := client.CreateKeyPair(ctx, &ec2.CreateKeyPairInput{
			KeyName: aws.String(keyname),
		})
		if err != nil {
			return fmt.Errorf("CreateKey error: %s", err)
		}
		//Escrever o arquivo da chave para ter acesso a chave
		err = os.WriteFile(keyname+".pem", []byte(*CreatedKeyPair.KeyMaterial), 0400)
		if err != nil {
			return fmt.Errorf("WriteFile error: %s", err)
		}
	}
	return nil
}
