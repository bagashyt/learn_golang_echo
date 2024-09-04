# GRPC and ProtoBuf

implement GRPC with protobuf

## How to setup Protobuf on Ubuntu

On terminal type 

```sh
sudo apt update
sudo apt install protobuf-compiler
```

install protobuf runtime for Go

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

add PATH on ~/.bashrc

`export PATH="$PATH:usr/local/go/bin"`




## How to run program

On service-garage directory terminal type 
```sh
go run main.go
```

On service-user directory terminal type 
```sh
go run main.go
```

On client directory terminal type 
```sh
go run main.go
```



