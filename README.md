# Rest API Microservice Template in Go Lang

This is a sample project for go lang REST microservice.

### What are features included?

1   Three tier pattern (Presentation,Business logics, Persistence)

2   Swagger

3   Dependency management with go mod

4   Documentation with go doc

5   Service discovery [todo]

6   Database integrations (Postgresql/Mysql)

7   Inter-service messaging over rest [todo]

### How to generate swagger docs

Following generates swagger docs in `docs` folder. Then go to `localhost:8080/swagger/index.html`

```shell
swag init -g app/application.go
```

### Complete documentation

Run `godoc` and check `localhost:6060` or run `go doc <package-name>`