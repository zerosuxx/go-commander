# go-http-commander

## build
```
make build
```

## running the application
```
make start # listening on localhost:1234
```

## available routes
```
POST /cmd # Request body: '["command", "arg1", ...]'
GET /healthcheck
```