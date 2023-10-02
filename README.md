# place-api

### Requirements

- Golang version 1.19 or higher

### Installation

1. Clone the repository:

```
git clone <repository_url>
```

2. Install the required dependencies:

```
make download-deps
```

This command will install the following dependencies:

- [golangci-lint](https://github.com/golangci/golangci-lint): Linter for Go programming language.
- [goose](https://github.com/pressly/goose): Database migration tool for Go applications.
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports): Tool for organizing imports in Go code.
- [swag](https://github.com/swaggo/swag): Toolkit for Go Swagger.

### Usage

The following commands are available for use:

- `make install-lint`: Installs the golangci-lint binary.
- `make download-deps`: Downloads project-specific dependencies.
- `make go-generate`: Runs go-generate.
- `make swag-v1`: Initializes Swagger documentation.
- `make generate`: Runs all code generations.
- `make lint`: Runs the linter for Go code.
- `make go-fmt`: Runs 'go fmt' to format Go code.
- `make new-migration`: Creates a new migration file.
- `make up`: Runs the application using Docker Compose.
- `make down`: Stops the running application using Docker Compose.
- `make build-up`: Builds and runs the application using Docker Compose.
- `make test`: Runs the tests for the application.

To execute any of the above commands, use the following syntax:

```
make <command>
```

For example, to install the golangci-lint binary, run:

```
make install-lint
```

### Contributing

If you would like to contribute to this project, please fork the repository and create a pull request.

### License

This project is licensed under the [MIT License](LICENSE).