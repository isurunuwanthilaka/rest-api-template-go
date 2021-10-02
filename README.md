# Rest API Microservice Template in Go Lang

This is a sample project for go lang REST microservice.

### What are features included?

1 Three tier pattern (Presentation,Business logics, Persistence)

2 Swagger

3 Dependency management with go mod

4 Documentation with go doc

5 Service discovery [todo]

6 Database integrations (Postgresql/Mysql)

7 Inter-service messaging over rest [todo]

### How to generate swagger docs

Following generates swagger docs in `docs` folder. Then go to `localhost:8080/swagger/index.html`

```shell
swag init -g app/application.go
```

### Dependencies

Once cloned the repo, run `go mod download` to pull all dependencies, run `go mod tidy` to clean dependencies once you add few.

### Formating

1. Use standard `gofmt` tool, add this to your ide better `on save` file.

2. Use standard `goimport` tool.

### Complete documentation

Run `godoc` and check `localhost:6060` or run `go doc <package-name>`
