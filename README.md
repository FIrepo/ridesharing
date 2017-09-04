# Ridesharing API

You can found API Documentation in here : http://demo.degananda.com/ridesharing/

# Account

```
driver degananda.ferdian@gmail.com:dega123
passenger ferdian.degananda@gmail.com:dega123
```

password encryoted with bcrypt (base 10).

# Installing

1.Install ridesharin api

```
go get github.com/degananda/ridesharing
```

2.Import sql  ridesharing.sql to your database

3.Setting mysql conneciton & database

model/connection.go

```
func DbConnect() {
	Db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/ridesharing")
}
```

# Running

```
go run *.go
```

## Built With

* [Gin-gonic](https://github.com/gin-gonic/gin) - Framework
* [JWT](github.com/dgrijalva/jwt-go) - Auth
* [Mysql](https://github.com/go-sql-driver/mysql) - Driver
* [Bcrypt](golang.org/x/crypto/bcrypt) - Encyrpt (Base 10)

