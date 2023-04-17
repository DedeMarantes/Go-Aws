package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "root",
	Short: "criar instancias aws",
}

func Execute() error {
	return root.Execute()
}

func init() {
	cobra.OnInitialize()
	instanceCreated, err := CreateInstance()
	if err != nil {
		fmt.Printf("Erro ao criar comando: %s", err)
		os.Exit(1)
	}
	root.AddCommand(instanceCreated)
}
