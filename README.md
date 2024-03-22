# GoDI

[![Make](https://github.com/kukymbr/godi/actions/workflows/test.yml/badge.svg)](https://github.com/kukymbr/godi/actions/workflows/test.yml)
[![GoDoc](https://godoc.org/github.com/kukymbr/godi?status.svg)](https://godoc.org/github.com/kukymbr/godi)
[![GoReport](https://goreportcard.com/badge/github.com/kukymbr/godi)](https://goreportcard.com/report/github.com/kukymbr/godi)

The [Golang](https://go.dev) Dependency Injector with no-magic initializations.

## Usage

### Installation

The go modules is the only supported way to use this package:

```shell
go get github.com/kukymbr/godi
```

### Building the container

Create the `godi.Builder` instance and add dependency definitions into it:

```go
builder := &godi.Builder{}

err := builder.Add(godi.Def{
    Name: "db",
    Build: func(ctn *godi.Container) (any, err) {
        return database.NewDB()	
    }
})
if err != nil {
    panic(err)
}

ctn, err := builder.Build()
if err != nil {
    panic(err)
}
```

### Accessing the dependencies

With panic on error:

```go
db := ctn.Get("db").(database.DB)
```

With error handling:

```go
dbAny, err := ctn.SafeGet("db")
if err != nil {
    // ...
}

db, ok := dbAny.(database.DB)
if !ok {
    // ...
}
```

Checking if dependency is registered in the container:

```go
if ctn.Has("db") {
    // ...
}
```

## License

MIT.