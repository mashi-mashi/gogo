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
gcloud run deploy --image gcr.io/cacao-alpha/echo-api --platform managed --region asia-northeast1

```

### 一発でやるやつ

- サービス名も勝手につく

```
gcloud beta run deploy --platform managed --region asia-northeast1 --source .
```
