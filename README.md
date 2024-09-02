# Manage Session with Gorilla Sessions

Using [Gorilla Sessions]("https://github.com/gorilla/sessions") for Manage Session

## How to run program

On terminal type 
```sh
go run .
```

## How to test 

with curl

```sh
curl -X GET http://localhost:9000/get
```
with postman send GET


```
http://localhost:9000/get
```

## How to setup postgresql on Ubuntu

open postgres cli

```
sudo -u postgres psql
```

create user and database with psql

```
CREATE DATABASE mydb;
CREATE USER myuser WITH ENCRYPTED PASSWORD 'yourpass';
GRANT ALL PRIVILEGES ON DATABASE mydb TO myuser;
```

if getting error 'FATAL: password authentication failed for user myuser' set this on psql

'''
\password myuser
'''

then type the password

