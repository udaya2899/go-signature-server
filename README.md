## Signature Server

### Introduction
Signature Server is a REST API Server written in Go that stores transaction blobs and sings them using ed25519 key

### Setup

* Modify the `config.json` file in the project root as needed. Sample Config:
```json
{
    "port": ":3001",
    "private_key_path": "pem/private_key.pem",
    "public_key_path": "pem/public_key.pem",
    "log_file_name": "server.log"
}
```

* Generate a pem public-private ed25519 key-pair by running the command:
```sh
make keygen
```

It generates two files `public_key.pem` and `private_key.pem` in `pem/` directory

### Tools used
The following tools are used here:
1. REST API Framework - [`gin-gonic/gin`](https://github.com/gin-gonic/gin)
   * High-performance - gin-gonic is a wrapper around `httprouter` and is one of the fastest frameworks
   * Easily configurable - any later custom logger, auth, middleware to be added can be done with ease
   * Simple way to write response back - Easier straightforward way to write API response
  
2. Persistence Key Value store - [`badgerdb`](https://github.com/dgraph-io/badger)
   * Resilient to Crashes - since it uses Write Ahead Log
   * Fast - Writes can be made upto a speed of 160 MB/s \
    Source - https://dgraph.io/badger


### API Documentation

1. `GET PUBLIC KEY`

Returns `public_key` of the generated server keypair

Request:
```shell
curl --request GET \
     --url http://localhost:<port>/public_key \
```

Response:
```json
{
    "public_key": "9d09f5ab-82eb-4fa5-b965-54792ea80131"
}
```

2. `PUT TRANSACTION BLOB`

Takes in a transaction blob encoded as base64 string in the format 
```json
{
    "txn": "\A+B8oD==",
}
```

Returns the unique uuid identifier for the created transaction
```json
{
    "id": "9d09f5ab-82eb-4fa5-b965-54792ea80131"
}
```

The transaction value is persisted in the system



