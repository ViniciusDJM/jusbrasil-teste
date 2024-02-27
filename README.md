# Tribuna

O Tribuna Process Finder é um projeto Go que consiste em desenvolver uma API para buscar dados de processos em todos os graus dos Tribunais de Justiça de Alagoas (TJAL) e do Ceará (TJCE). O objetivo é coletar informações sobre os processos, como classe, área, assunto, data de distribuição, juiz, valor da ação, partes do processo e lista das movimentações.

## Pré-requisitos

- Go versão 1.20 ou superior instalado no sistema.

## Instalação

Para executar o projeto localmente, siga os passos abaixo:

1. Clone este repositório:

```bash
git clone https://github.com/ViniciusDJM/jusbrasil-teste.git
```

2. Acesse o diretório do projeto:

```bash
cd jusbrasil-teste
```

## Como executar

Você pode escolher entre duas formas de executar comandos relacionados ao projeto: utilizando o `makefile` ou o `taskfile`.

### Utilizando o Makefile

O `makefile` contém os seguintes comandos:

- `all`: Compila o projeto (mesmo que `build`).
- `install-deps`: Instala as dependências do projeto, como o Swag e MockGen.
- `code-gen`: Executa a geração de código e inicializa o Swag para a documentação da API.
- `build`: Compila o projeto com os devidos flags de otimização.
- `clean`: Limpa os binários gerados.

Para utilizar cada comando, execute o seguinte:

```bash
# Instala as dependências
make install-deps

# Executa a geração de código e inicializa o Swag
make code-gen

# Compila o projeto
make build

# Limpa os binários gerados
make clean
```

### Utilizando o Taskfile

Utilizando o Taskfile

> Primeiramente, é necessário instalar o task em seu sistema. Siga as instruções em https://github.com/go-task/task#installation para a instalação adequada.

Com o task instalado, execute o seguinte comando para ver a lista de tarefas disponíveis:

```bash
task -l
```

O `taskfile` contém as seguintes tarefas:

- `install-deps`: Instala as dependências do projeto, como o Swag e MockGen.
- `gen:code`: Executa a geração de código e inicializa o Swag para a documentação da API.
- `build`: Compila o projeto com os devidos flags de otimização.
- `clean`: Limpa os binários gerados.

Para utilizar cada tarefa, execute o seguinte:

```bash
# Instala as dependências
task install-deps

# Executa a geração de código e inicializa o Swag
task gen:code

# Compila o projeto
task build

# Limpa os binários gerados
task clean
```

## Executando o projeto

Para executar o projeto basta executar o comando:
```bash
go run .
```

Caso tenha realizado o build do projeto, basta executar o binario resultante. 

## Documentação da API

A documentação da API está disponível através do Swagger na rota `/api/swagger`. Acesse esta rota em seu navegador para visualizar a documentação detalhada das rotas e parâmetros da API.

A API possui uma rota principal `/api/swagger` que pode ser acessada através de 2 metodos, GET e POST. Onde ambas recebem o número do processo, o qual não se faz necessária a utilização de máscara.

### Requisição GET:

Para fazer uma requisição GET para a rota "/api/search", você precisará fornecer o parâmetro "processNumber" na URL. O "processNumber" é uma string representando o número do processo que você deseja pesquisar.

Exemplo de requisição GET usando o cURL:

```bash
curl -X GET "http://localhost:8000/api/v1/search?processNumber=0710802-55.2018.8.02.0001"
```

### Requisição POST:

Para fazer uma requisição POST para a rota "/api/search", você precisará enviar um corpo (body) no formato JSON contendo o parâmetro "data". O "data" é um objeto que deve ter o campo "number" representando o número do processo que você deseja pesquisar.

Exemplo de requisição POST usando o cURL:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"number": "0710802-55.2018.8.02.0001"}' "http://localhost:8000/api/v1/search"
```

## Imagem Docker

O projeto também está disponível como uma imagem Docker, distribuída através do GitHub Container Registry. Para obter a imagem, utilize o seguinte comando:

```bash
docker pull ghcr.io/viniciusdjm/jusbrasil-teste:latest
```

---
