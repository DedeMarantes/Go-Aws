package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var ctx = context.Background()

var root = &cobra.Command{
	Use:   "root",
	Short: "criar instancias aws",
}

func Execute() error {
	return root.Execute()
}

func init() {
	cobra.OnInitialize()
	instanceCommand, err := CreateInstance()
	checkError(err)
	bucketCommand, err := CreateBucket()
	checkError(err)
	uploadCommand, err := UploadToBucket()
	checkError(err)
	root.AddCommand(instanceCommand)
	root.AddCommand(bucketCommand)
	root.AddCommand(uploadCommand)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("Erro ao criar comando: %s", err)
		os.Exit(1)
	}
}
