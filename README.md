# Singleflight

implementing Singleflight to supress duplicate function calling.


## How to run

run program on directory
```sh
go run main.go
```

## How to test

on 2 different terminal run this curl

```sh
curl -s -w "elapsed: %{time_total}s\n" -i -X GET http://localhost:8080/api/report/download/sales-2024-08
```
