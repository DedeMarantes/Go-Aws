# CLI simples para trabalhar com aws
## Criando Instâncias
Esta é uma aplicação CLI desenvolvida em Go para criar instâncias da AWS com a chave de acesso padrão de "aws-key" se não for informada. O usuário pode escolher o tipo de instância entre Ubuntu e Redhat, além de especificar a quantidade de instâncias a serem criadas. **A região desse projeto é us-east-1**.

### Como usar

Para executar o programa, basta rodar o seguinte comando:

```bash
go run main.go create-instance [ubuntu/redhat]
```

Onde `[ubuntu/redhat]` é um argumento obrigatório que deve ser especificado para escolher o tipo de instância.

### Exemplo completo

Abaixo está um exemplo completo de comando para criar 3 instâncias Redhat da AWS com a chave chamada "minhachave":

```bash
go run main.go create-instance -n 3 --key-name="minhachave" redhat
```

## Criando Bucket

É também possível criar uma lista buckets S3 na AWS, através do comando `create-bucket`. E a flag `--public` deixa o bucket livre ao publico. Além de também com o comando `upload-bucket` fazer um upload de um arquivo para uma lista de buckets

Exemplo para criar 2 buckets:

```bash 
go run main.go create-bucket bucket-name45 bucket-otherName
```

Exemplo para fazer upload nos buckets criados:

```bash
go run main.go upload-bucket [arquivo para upload] [lista buckets pelo nome]
```

