#### develop

- hotreload のために air を利用する
  - https://github.com/cosmtrek/air

```
go get -u github.com/cosmtrek/air
air
```

## deploy

### Cloud Build を使用してコンテナ イメージをビルド

```
gcloud builds submit --tag gcr.io/cacao-alpha/echo-api
```

### Cloud Run へ Deploy

```
gcloud beta run deploy --image gcr.io/cacao-alpha/cloud-run-go-echo:tag1
gcloud run deploy --image gcr.io/cacao-alpha/echo-api --platform managed

```
