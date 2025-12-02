# simple-go-app

A simple Go app to serve a static site

## Run the app

```bash
    go run main.go
```

## To deploy

Install and setup your Google Cloud SDK.

Run:

```bash
    gcloud app deploy .
```

## Build docker image

```bash
docker build -t simple-go-app .
```
