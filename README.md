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

```
port: port to listen on (e.g :3000)  
server: mongodb server (e.g localhost)
database: mongodb database (e.g escherdb)
```

### Run

```
./escher-api
```
