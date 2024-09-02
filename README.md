# Cross-Site Request Forgery (CSRF)

handling Cross-Site Request Forgery (CSRF)["https://en.wikipedia.org/wiki/Cross-site_request_forgery"]

## How to run program

On terminal type 
```sh
go run .
```
On Browser type
`http://localhost:9000/index`

Test with curl
```
curl -X POST http://localhost:9000/sayhello \
     -H 'Content-Type: application/json' \
     -d '{"name":"myname","gender":"male"}'
```


