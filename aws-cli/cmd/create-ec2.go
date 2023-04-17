package cmd

import (
	"context"
	"fmt"

	"github.com/DedeMarantes/aws-go/instances"
	"github.com/spf13/cobra"
)

var ctx = context.Background()

//Criando comando para criar instancia
func CreateInstance() (*cobra.Command, error) {
	var (
		count   int
		keyname string
	)
	create_instance := &cobra.Command{
		Use:       "create-instance [ubuntu/redhat]",
		Short:     "Criar uma instancia Ec2 redhat ou ubuntu",
		ValidArgs: []string{"redhat", "ubuntu"},
		Args:      cobra.MatchAll(cobra.ExactArgs(1)),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < count; i++ {
				switch args[0] {
				case "ubuntu":
					_, err := instances.CreateUbuntuInstance(ctx, "us-east-1", keyname)
					if err != nil {
						fmt.Printf("Erro ao criar instancia: %s", err)
					}
				case "redhat":
					_, err := instances.CreateRedHatInstance(ctx, "us-east-1", keyname)
					if err != nil {
						fmt.Printf("Erro ao criar instancia: %s", err)
					}
				}
			}
		},
	}
	//flags para quantidade de instancias e nome da chave
	create_instance.Flags().IntVarP(&count, "number", "n", 1, "numero de instancias a ser criada")
	create_instance.Flags().StringVar(&keyname, "keyname", "aws-key", "nome da chave a ser usada, se não existir ela é criada")
	return create_instance, nil
}
