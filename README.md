# GQLGEN Todo Example

This service is to demonstrate how we can develop GraphQL service in Golang. 
This example emphasize on how we can structure our Golang code with GraphQL, use dependency injection pattern, and follow Effective Go approach.

## Running The Service

To run the service, you can run it directly with `go run cmd/main.go`.

Don't worry about not be able to resolve dependencies, you don't have to download dependencies. It will always be there (committed).

If you wish to regenerate the generated file, you can use `go run github.com/99designs/gqlgen generate`.

## Code Structure

```text
.
├── README.md
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── model
│   ├── schema.graphqls
│   └── z_graph.go
├── internal
│   ├── mutation
│   ├── query
│   ├── resolver
│   └── storage
├── tools.go
└── vendor
    ├── github.com
    ├── golang.org
    ├── gopkg.in
    └── modules.txt

12 directories, 9 files
```

`internal` folder is created and inside is declaration of GQL handler for `mutation`, `query`, and `resolver`.

The `storage` package used for connecting to database / cache / file storage.

The initialization happened in `cmd/main.go`. In `main.go`, we define how we can inject `TODO Storage` to GraphQL resolver. Later, `TODO Storage` used by `mutation` and `query` resolver.

## Unit Testing

We use `testify` and `mockery` for unit testing approach. This may not be the right taste for some people. For me, this is the robust and easiest way to do.

## Credit

Made in Jakarta with love (c) 2023.
