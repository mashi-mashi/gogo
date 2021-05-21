### Installation

#### swagger

```
docker pull quay.io/goswagger/swagger
// エイリアス登録
alias swagger="docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
```

#### generate

```
swagger generate server -t server/gen -f swagger/swagger.yaml
```
