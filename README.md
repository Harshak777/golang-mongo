# golang-mongo

### Initial project setup
1) Create go.mod file
```
go mod init <module_name>
```
2) Add necessary dependencies - create go.sum for dependencies
```
go get <module_name>
```
Modules required
```
github.com/julienschmidt/httprouter
go.mongodb.org/mongo-driver/mongo
```

### Dependency use case
httprouter - to get routing params
mongo - to connect with Mongo

### Project structure
- controllers
    - user.go
- models
    - user.go
- main.go

### Cleaning
Clean the go mod but using `go mod tidy` to remove unused modules