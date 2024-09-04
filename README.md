# JSON Web Token (JWT)

implement JSON Web Token


## How to run program

On terminal type 
```sh
go run .
```

## How to test

with curl on terminal type:

```sh
curl -X POST --user noval:kaliparejaya123 http://localhost:8080/login
```
after got the token you can use it to access /index

```sh
curl -X GET --header "Authorization: Bearer <your Token>" http://localhost:8080/index
```
