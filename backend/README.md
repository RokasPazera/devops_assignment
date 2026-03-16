# NSE Thermometer Backend
This repository contains the Go backend for managing stretch goals and progress tracking.

The backend provides a REST API with the following functionality:

- `/api/goals`: Manage stretch goals (view all goals, add new goals)
- `/api/percentage`: Track current progress percentage (view and update)

The backend is built using:

- [Go](https://go.dev/) for high-performance server logic
- [Gin](https://gin-gonic.com/) as the HTTP web framework
- [GORM](https://gorm.io/) as the ORM for database access
- [SQLite](https://sqlite.org/) by default, with optional support for [MySQL](https://www.mysql.com/), [MariaDB](https://mariadb.org/) or [MSSQL](https://www.microsoft.com/sql-server).

## Environment variables
The backend uses the following environment variables to configure database and server behavior:

| Variable        | Description                                                            | Default  |
| --------------- | ---------------------------------------------------------------------- | -------- |
| `DB_TYPE`       | Type of database to use. Supported values: `sqlite`, `mysql`, `mssql`. | `sqlite` |
| `DB_CONNECTION` | Database connection string (required for MySQL or MSSQL).              | —        |
| `SERVER_PORT`   | Port on which the backend server will listen.                          | `8000`   |
| `GIN_MODE`      | Gin framework mode. Can be `debug` or `release`.                       | `debug`  |

## Building and running the project

1. Build the Go backend binary:

```bash
go build -o backend
```

2. Run the backend:

```bash
./backend
```

- The server will start on the port specified in `SERVER_PORT`.
- Make sure your database is reachable and the connection string is correct.
- When using SQLite, the database is stored in the file `app.db`.

## Quality control

### Testing

All tests are written using the Go `testing` package, Gin for HTTP handler testing, and `testify/assert` for assertions.  

Tests are located in the `tests/` folder and cover both the `/api/goals` and `/api/percentage` endpoints.

To run all tests and see the output:

```bash
go test ./... -v
```

This will execute all test files and print detailed results in the terminal.

Each test runs in an isolated in-memory database, so you do **NOT** need to start the backend server to run the tests.

### Code coverage report
To generate a code coverage report:

```bash
go test ./... -coverpkg=./... -coverprofile=coverage.out
```

This will produce a `coverage.out` file containing coverage data for all packages.

To view the coverage in a browser-friendly HTML format:

```bash
go tool cover -html=coverage.out -o coverage.html
```

### Linting
We use [Golangci-lint](https://golangci-lint.run/) to enforce code style, formatting, and common Go best practices. Please follow the installation instructions on the website.

To run the linter:

```bash
golangci-lint run ./...
```

It helps maintain clean, readable, and maintainable code. This will scan all packages and report issues such as unused variables or imports, potential bugs (e.g. nil pointer dereferences) and formatting/style inconsistencies.

### Security
We use the [GoSec](https://github.com/securego/gosec) to perform static security analysis on the codebase. GoSec inspects your Go code for common vulnerabilities such as SQL injection risks, hardcoded credentials and unsafe function usage. Please follow the installation instructions on the website.

To run a security scan:

```bash
gosec ./...
```