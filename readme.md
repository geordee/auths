# Simple Scopes Service

A simple, opinionated user service
to complement https://github.com/geordee/authx

## Test

```bash
source env.local
go run .
curl localhost:9095/users/geordee
```

## Build

```bash
docker build -t auths:latest .
```

## Run

```bash
docker run --env-file env.local -p 9095:9095 auths:latest
```
