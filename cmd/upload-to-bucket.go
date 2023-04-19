package cmd

import (
	"fmt"
	"os"

	"github.com/DedeMarantes/aws-go/services"
	"github.com/spf13/cobra"
)

func UploadToBucket() (*cobra.Command, error) {
	upload_bucket := &cobra.Command{
		Use:   "upload-bucket [file_name] [bucket_name]",
		Short: "Upload de um arquivo para o bucket nomeado",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := services.UploadS3Bucket(ctx, args[0], args[1:])
			if err != nil {
				fmt.Printf("Erro no comando: %s", err)
				os.Exit(1)
			}
		},
	}
	return upload_bucket, nil
}
