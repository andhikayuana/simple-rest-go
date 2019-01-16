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

* `BASE_URL=http://localhost:3000`
* `GET` `/products` _get all persons_

response example

```json
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "FirstName": "",
            "LastName": "",
            "Avatar": "",
            "Age": 0
        },
        {
            "id": 2,
            "FirstName": "",
            "LastName": "",
            "Avatar": "https://api.adorable.io/avatars/120/jarjitsingh.png",
            "Age": 10
        },
        {
            "id": 3,
            "FirstName": "jarjit",
            "LastName": "singh",
            "Avatar": "https://api.adorable.io/avatars/120/jarjitsingh.png",
            "Age": 10
        }
    ],
    "msg": "success"
}
``` 

* `GET` `/persons/{id}` _get person by id_

response example

```json
{
    "code": 200,
    "msg": "Success",
    "data": {
        "id": 1,
        "name": "Sandal Mahal",
        "price": 5000000,
        "image": "https://anu.com/sandal-mahal.jpg"
    }
}
```


* `POST` `/persons` _insert person data_

request example

```json
{
	"firstname":"jarjit",
	"lastname":"singh",
	"avatar":"https://api.adorable.io/avatars/120/jarjitsingh.png",
	"age":10
}
```

response example

```json
{
    "code": 200,
    "data": {
        "id": 4,
        "FirstName": "jarjit",
        "LastName": "singh",
        "Avatar": "https://api.adorable.io/avatars/120/jarjitsingh.png",
        "Age": 10
    },
    "msg": "success"
}
```