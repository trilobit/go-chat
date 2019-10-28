# Simple Chat Application

Project written with GoLang. Uses PostgreSQL as storage. Registration and simple authorization are included.

## Project setup

Copy `example.config.yml` to `dev.config.yml` for development or `prod.config.yml` for production and fill arguments in it.

### Build binary container

```
docker build -t gochat:latest .
```

### Compile for development

Just run `docker-compose`

```
docker-compose up -d
```

It runs containers with binary and db. And also run migrations if necessary.

## Deployment

It project can be deployed with docker. It depends on another PostgreSQL container.

```
docker run --rm --link [pg-container] -v `pwd`/prod.config.yml:/config.yml -p 9090:9090 -p 9091:9091 gochat:latest
```

To get current version run command:

```
docker run --rm --link [pg-container] -v `pwd`/prod.config.yml:/config.yml -p 9090:9090 -p 9091:9091 gochat:latest /chat --version
```

## Built With

Project created with [GoLang 1.13](https://golang.org/doc/go1.13)

## Authors

* Alexander Savchuk - _main contributor_ - [trilobit](https://github.com/trilobit)
* Pavel Makarenko - _mentor and teacher_ - [m1ome](https://github.com/m1ome)

## License

This project is licensed under the MIT License