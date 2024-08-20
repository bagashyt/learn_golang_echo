# HTTP Request Payload with Validation on Echo

A simple HTTP request with validation on Echo Framework

## How to run program

On terminal type 
```sh
go run .
```

## How to test 

with curl

```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"example","email":"example@mail.com","age":"10"}' \
  http://localhost:9000/users
```
with postman you can pass body 

```
{
    "name":"example","email":"example@mail.com","age":"10"
}
```
send to
```
http://localhost:9000/users
```
