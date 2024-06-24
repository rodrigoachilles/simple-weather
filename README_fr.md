Traductions:

* [Anglais](README.md)
* [Portugais (Brésil)](README_pt_br.md)

---

# 🌤️ Système Météo par Code Postal

![Project Logo](assets/simple_weather-logo.png)

Bienvenue dans le système météo par code postal développé en Go ! Ce projet vous permet de recevoir un code postal, d'identifier la ville et de retourner le climat actuel en Celsius, Fahrenheit et Kelvin. Le système est déployé sur Google Cloud Run.

## 📑&nbsp;Table des Matières

- [📖 Introduction](#introduction)
- [🛠 Prérequis](#prérequis)
- [⚙️ Installation](#installation)
- [🚀 Utilisation](#utilisation)
- [🔍 Exemples](#exemples)
- [🤝 Contribution](#contribution)
- [📜 Licence](#licence)

## 📖&nbsp;Introduction

Ce système météo par code postal est un projet développé en Go qui permet de recevoir un code postal, d'identifier la ville et de retourner le climat actuel en différentes unités de température. Il utilise l'API viaCEP pour obtenir la localisation et l'API WeatherAPI pour consulter les températures. Le système est déployé sur Google Cloud Run.

## 🛠&nbsp;Prérequis

Assurez-vous d'avoir les éléments suivants installés avant de continuer :

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- Un compte sur [Google Cloud Platform](https://cloud.google.com/)

## ⚙️&nbsp;Installation

1. Clonez ce dépôt :

    ```sh
    git clone git@github.com:rodrigoachilles/simple-weather.git
    cd simple-weather
    ```

2. Créez un fichier `.env` à la racine du projet avec les configurations suivantes :

    ```env
    KEY_WEATHER_API=votre_clé_api_weather
    ```

3. Exécutez Docker Compose :

    ```sh
    docker-compose up -d
    ```

## 🚀&nbsp;Utilisation

Après avoir démarré Docker Compose, vous pouvez configurer et utiliser le système météo par code postal.

### 🔧&nbsp;Configuration

1. Exécutez le serveur Go :

    ```sh
    go run main.go
    ```

2. Le système sera disponible sur le port configuré dans le fichier `.env` (par défaut, 8080).

### 🔧 Exécution des Services

1. Accédez au dossier `api` dans le répertoire `weather-service` :

    ```sh
    cd simple-weather/api
    ```

2. Exécutez le fichier `.http` en utilisant votre outil préféré (par exemple, VSCode REST Client, Postman) :

    ```sh
    # Example for VSCode REST Client
    weather.http
    ```

## 🔍&nbsp;Exemples

Voici quelques exemples d'utilisation du système météo par code postal :

### Requête avec un code postal valide

**Requête :**

```sh
curl -X GET "http://localhost:8080/01001000"
```

**Réponse :**

```json
{
  "locale": "locale",
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

### Requête avec un code postal invalide (format incorrect)

**Requête :**

```sh
curl -X GET "http://localhost:8080/123"
```

**Réponse :**

```json
{
  "error": "invalid zipcode"
}
```

### Requête avec un code postal non trouvé

**Requête :**

```sh
curl -X GET "http://localhost:8080/00000000"
```

**Réponse :**

```json
{
  "error": "can not find zipcode"
}
```

## 🤝&nbsp;Contribution

N'hésitez pas à ouvrir des issues ou à soumettre des pull requests pour des améliorations et des corrections de bugs.

## 📜&nbsp;Licence

Ce projet est sous licence MIT.

---

## 📤&nbsp;Déploiement sur Google Cloud Run

Pour déployer sur Google Cloud Run, suivez les étapes ci-dessous :

1. Authentifiez-vous auprès de Google Cloud :

    ```sh
    gcloud auth login
    ```

2. Configurez le projet :

    ```sh
    gcloud config set project [PROJECT_ID]
    ```

3. Construisez et envoyez l'image Docker :

    ```sh
    gcloud builds submit --tag gcr.io/[PROJECT_ID]/simple-weather
    ```

4. Déployez l'image sur Cloud Run :

    ```sh
    gcloud run deploy simple-weather --image gcr.io/[PROJECT_ID]/simple-weather --platform managed
    ```

5. Accédez à l'URL fournie après le déploiement pour utiliser le système météo par code postal.
