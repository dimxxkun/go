# Go

A clean, idiomatic Go starter repository demonstrating standard project layout, package structure, unit tests, and CLI execution.

## Project Structure

```
go/
├── go.mod
├── README.md
├── .gitignore
├── main.go
└── analyzer/
    ├── analyzer.go
    └── analyzer_test.go
```

## Prerequisites

- Go (version 1.20 or later recommended)

## How to Run & Test

### Run directly:

```bash
go run main.go
```

### Run tests:

```bash
go test -v ./...
```

### Build binary:

```bash
go build -o go-app.exe main.go
./go-app.exe
```

## License

MIT
