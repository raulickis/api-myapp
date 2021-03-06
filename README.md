# api-myapp

API em **GO v1.14** responsável por cadastrar usuarios e seus endereços.


Para fazer o download das depências e limpar dependências não utilizadas:

``` 
go14 mod tidy
```

## Para executar localmente:
 
1 - Subir o banco de dados
```
docker-compose up postgres
```

2 - Subir o aplicativo localmente pelo GO

```
go run cmd/main.go
```

## Para preparar a versão para implantação:
 
1 - Compilar o programa e gerar a Imagem Docker
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd
docker build -t raulickis/api-myapp .
```

2 - Executar tudo pelo Docker Compose

```
docker-compose up
```

## Como testar a aplicacão:

1 - (Opcional) Para fazer um teste rápido, executar o comando **curl** abaixo:
```
curl http://localhost:9990/cadastro/usuario
```

2 - Importar no Postman o arquivo **myapp.postman_collection.json**

3 - Executar os endpoints para cadastro de usuarios e enderecos


