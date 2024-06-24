Traduções:

* [Inglês](README.md)
* [Francês](README_fr.md)

---

# 🌤️ Sistema de Clima por CEP

![Project Logo](assets/simple_weather-logo.png)

Bem-vindo ao sistema de clima por CEP desenvolvido em Go! Este projeto permite receber um CEP, identificar a cidade e retornar o clima atual em graus Celsius, Fahrenheit e Kelvin. O sistema é implantado no Google Cloud Run.

## 📑&nbsp;Sumário

- [📖 Introdução](#introdução)
- [🛠 Pré-requisitos](#pré-requisitos)
- [⚙️ Instalação](#instalação)
- [🚀 Uso](#uso)
- [🔍 Exemplos](#exemplos)
- [📤 Deploy no Google Cloud Run](#deploy-no-google-cloud-run)
- [🤝 Contribuição](#contribuição)
- [📜 Licença](#licença)

## 📖&nbsp;Introdução

Este sistema de clima por CEP é um projeto desenvolvido em Go que permite receber um CEP, identificar a cidade e retornar o clima atual em diferentes unidades de temperatura. Ele utiliza a API viaCEP para obter a localização e a API WeatherAPI para consultar as temperaturas. O sistema é implantado no Google Cloud Run.

## 🛠&nbsp;Pré-requisitos

Certifique-se de ter os seguintes itens instalados antes de continuar:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- Conta no [Google Cloud Platform](https://cloud.google.com/)

## ⚙️&nbsp;Instalação

1. Clone este repositório:

    ```sh
    git clone git@github.com:rodrigoachilles/simple-weather.git
    cd simple-weather
    ```

2. Crie um arquivo `.env` na raiz do projeto com as seguintes configurações:

    ```env
    KEY_WEATHER_API=sua_chave_api_weather
    ```

3. Execute o Docker Compose:

    ```sh
    docker-compose up -d
    ```

## 🚀&nbsp;Uso

Após iniciar o Docker Compose, você pode configurar e usar o sistema de clima por CEP.

### 🔧&nbsp;Configuração

1. Execute o servidor Go:

    ```sh
    go run main.go
    ```

2. O sistema estará disponível na porta configurada no arquivo `.env` (por padrão, 8080).

### 🔧&nbsp;Executando Serviços

1. Navegue até a pasta `api` no diretório `weather-service`:

    ```sh
    cd simple-weather/api
    ```

2. Execute o arquivo `.http` usando sua ferramenta preferida (por exemplo, VSCode REST Client, Postman):

    ```sh
    # Example for VSCode REST Client
    weather.http
    ```

## 🔍&nbsp;Exemplos

Aqui estão alguns exemplos de uso do sistema de clima por CEP:

### Requisição com CEP válido

**Request:**

```sh
curl -X GET "http://localhost:8080/01001000"
```

**Response:**

```json
{
   "locale": "locale",
   "temp_C": 28.5,
   "temp_F": 83.3,
   "temp_K": 301.5
}
```

### Requisição com CEP inválido (formato incorreto)

**Request:**

```sh
curl -X GET "http://localhost:8080/123"
```

**Response:**

```json
{
   "error": "invalid zipcode"
}
```

### Requisição com CEP não encontrado

**Request:**

```sh
curl -X GET "http://localhost:8080/00000000"
```

**Response:**

```json
{
   "error": "can not find zipcode"
}
```

## 📤&nbsp;Deploy no Google Cloud Run

Para realizar o deploy no Google Cloud Run, siga os passos abaixo:

1. Autentique-se no Google Cloud:

    ```sh
    gcloud auth login
    ```

2. Configure o projeto:

    ```sh
    gcloud config set project [PROJECT_ID]
    ```

3. Construa e envie a imagem Docker:

    ```sh
    gcloud builds submit --tag gcr.io/[PROJECT_ID]/simple-weather
    ```

4. Implante a imagem no Cloud Run:

    ```sh
    gcloud run deploy simple-weather --image gcr.io/[PROJECT_ID]/simple-weather --platform managed
    ```

5. Acesse a URL fornecida após a implantação para usar o sistema de clima por CEP.

[Simple weather](https://simple-weather-ciw7dykisq-uc.a.run.app/01001000)

## 🤝&nbsp;Contribuição

Sinta-se à vontade para abrir issues ou enviar pull requests para melhorias e correções de bugs.

## 📜&nbsp;Licença

Este projeto está licenciado sob a Licença MIT.
