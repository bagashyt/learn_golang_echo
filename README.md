# Dockerize Golang

using Docker to build Golang app

## How to prepare docker

Create Container on code directory
```sh
docker build -t my-image-hello-world .
```

Start Container my-image-hello-world
```sh
docker run --name my-container-hello-world --rm -it -e PORT=8080 -e INSTANCE_ID="my first instance" -p 8080:8080 my-image-hello-world
```

## How to test

with curl on terminal type:

```sh
curl -X GET http://localhost:8080/
```
with browser type:

`http://localhost:8080/`
