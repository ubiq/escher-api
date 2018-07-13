# Escher-API

API for escher-interface written in golang. Serves data from mongodb that has been populated by escher-crawler.

### Build

clone to ```$GOPATH/src/github.com/ubiq/escher-api```

```
cd $GOPATH/src/github.com/ubiq/escher-api
go get
go build
```

### Configure

Configure via config.toml

port: port to listen on
server: mongodb server
database: mongodb database

### Run

```
./escher-api
```
