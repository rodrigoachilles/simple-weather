Translations:

* [French](README_fr.md)
* [Portuguese (Brazil)](README_pt_br.md)

---

# 🌤️ Weather System by ZIP Code

![Project Logo](assets/simple_weather-logo.png)

Welcome to the weather system by ZIP code developed in Go! This project allows you to receive a ZIP code, identify the city, and return the current weather in Celsius, Fahrenheit, and Kelvin. The system is deployed on Google Cloud Run.

## 📑&nbsp;Table of Contents

- [📖 Introduction](#introduction)
- [🛠 Prerequisites](#prerequisites)
- [⚙️ Installation](#installation)
- [🚀 Usage](#usage)
- [🔍 Examples](#examples)
- [📤 Deploy on Google Cloud Run](#deploy-on-google-cloud-run)
- [🤝 Contribution](#contribution)
- [📜 License](#license)

## 📖&nbsp;Introduction

This weather system by ZIP code is a Go project that allows you to receive a ZIP code, identify the city, and return the current weather in different temperature units. It uses the viaCEP API to get the location and the WeatherAPI to fetch the temperatures. The system is deployed on Google Cloud Run.

## 🛠&nbsp;Prerequisites

Make sure you have the following items installed before continuing:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- An account on [Google Cloud Platform](https://cloud.google.com/)

## ⚙️&nbsp;Installation

1. Clone this repository:

    ```sh
    git clone git@github.com:rodrigoachilles/simple-weather.git
    cd simple-weather
    ```

2. Create a `.env` file at the root of the project with the following configurations:

    ```env
    KEY_WEATHER_API=your_weather_api_key
    ```

3. Run Docker Compose:

    ```sh
    docker-compose up -d
    ```

## 🚀&nbsp;Usage

After starting Docker Compose, you can configure and use the weather system by ZIP code.

### 🔧&nbsp;Configuration

1. Run the Go server:

    ```sh
    go run main.go
    ```

2. The system will be available on the port configured in the `.env` file (default is 8080).

### 🔧&nbsp;Running Services

1. Navigate to the `api` folder in the `weather-service` directory:

    ```sh
    cd simple-weather/api
    ```

2. Execute the `.http` file using your preferred tool (e.g., VSCode REST Client, Postman):

    ```sh
    # Example for VSCode REST Client
    weather.http
    ```

## 🔍&nbsp;Examples

Here are some examples of using the weather system by ZIP code:

### Request with a valid ZIP code

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

### Request with an invalid ZIP code (incorrect format)

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

### Request with ZIP code not found

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

## 📤&nbsp;Deploy on Google Cloud Run

To deploy on Google Cloud Run, follow the steps below:

1. Authenticate with Google Cloud:

    ```sh
    gcloud auth login
    ```

2. Configure the project:

    ```sh
    gcloud config set project [PROJECT_ID]
    ```

3. Build and push the Docker image:

    ```sh
    gcloud builds submit --tag gcr.io/[PROJECT_ID]/simple-weather
    ```

4. Deploy the image to Cloud Run:

    ```sh
    gcloud run deploy simple-weather --image gcr.io/[PROJECT_ID]/simple-weather --platform managed
    ```

5. Access the provided URL after deployment to use the weather system by ZIP code.

[Simple weather](https://simple-weather-ciw7dykisq-uc.a.run.app/01001000)

## 🤝&nbsp;Contribution

Feel free to open issues or submit pull requests for improvements and bug fixes.

## 📜&nbsp;License

This project is licensed under the MIT License.
