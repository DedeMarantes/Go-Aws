package cmd

import (
	"fmt"
	"os"

	"github.com/DedeMarantes/aws-go/services"
	"github.com/spf13/cobra"
)

func CreateBucket() (*cobra.Command, error) {
	var isPublic bool
	create_bucket := &cobra.Command{
		Use:   "create-bucket [bucket-name]",
		Short: "Cria um bucket s3",
		Run: func(cmd *cobra.Command, args []string) {
			err := services.CreateS3Bucket(ctx, args[0:], isPublic)
			if err != nil {
				fmt.Printf("Erro ao fazer comando %s", err)
				os.Exit(1)
			}
		},
	}
	create_bucket.Flags().BoolVar(&isPublic, "public-access", false, "Acesso publico ao bucket")
	return create_bucket, nil
}
