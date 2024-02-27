# Full Cycle - Go Expert - Laboratório | Sistema de temperatura por CEP

Este projeto faz parte de um laboratório do curso de pós graduação em Go - Go Expert - da Full Cycle.

## Introdução

Este projeto implementa um servidor web que permite consultar a temperatura de uma localidade em tempo real com base no CEP.

### Como executar

#### Localmente

Para executar este projeto localmente, através do terminal acesse o diretório raiz deste projeto e execute o seguinte comando:

```bash
docker compose up -d
```
Para testar as requisições, é possível utilizar ferramentas visuais como o [Postman](https://www.postman.com/) ou [Insomnia](https://insomnia.rest/). Também é possível usar ferramentas do tipo CLI, tais como [cURL](https://curl.se/) ou [httpie](https://httpie.io/). Se preferir, você também pode user uma navegador.

Abaixo, dois exemplos de requisição usando o **httpie**. A segunda requisição inclui um token no cabeçalho:
```sh
http GET http://localhost:8080?cep=89023-600
http GET http://localhost:8080?cep=askdjajdaslkjd
```

#### Em produção

Para testar o projeto em produção, basta enviar requisição para esta URL: https://api-bt6643nbxa-uc.a.run.app/

### Executando testes de unidade e integração

Para executar os testes de unidade e integração do projeto, através do terminal acesse o diretório raiz deste repositório. Usando a ferramenta **make**, execute:

```sh
make test
```

Se a ferramenta **make** não estiver disponível, execute:

```sh
go test -v ./...
```

Para gerar um relatório de cobertura dos testes, execute:

```sh
make test-cov
```

Se a ferramenta **make** não estiver disponível, execute:
```sh
go test -v ./... -coverprofile=c.out
go tool cover -html=c.out
```