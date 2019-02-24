# This is API for main client 

Client lies [here](https://github.com/alzaburetz/restrictor.git)

To get this API

```
mkdir -p $GOROOT/github/alzaburetz
cd $GOROOT/github/alzaburetz
git clone https://github.com/alzaburetz/restrictor-api
```
You also would need sqlite driver and mux router and package from api itself
```
go get github.com/gorilla/mux
go get github.com/mattn/go-sqlite3
go get github.com/alzaburetz/restrictor-api/models
```

## To test this API

Go to API folder and build

```
cd path/to/api
go build
./myrestAPI
```

Clone client

```
git clone https://github.com/alzaburetz/restrictor
cd path/to/client
go run main.go
```

Edit data in database

**Enjoy!**



