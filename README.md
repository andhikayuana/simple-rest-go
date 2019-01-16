# Simple REST API using GO

## Requirements
- [go 1.11 or latest](https://golang.org/doc/install)
- [dep](https://golang.github.io/dep/)

## Dependencies
- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/jinzhu/gorm)
- [sqlite3](https://github.com/mattn/go-sqlite3)

## Install
- `make install-deb`

## Build
- `make build`
thats will be generate 3 binary file for _linux, macos, windows_ in `./build`
or you can specified build using below command :
```
make build-linux
make build-darwin
make build-windows
```

## Run
- `make run` or you can run the binary file in `./build` directory after build success

## Documentation
_todo_