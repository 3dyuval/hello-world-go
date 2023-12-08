module github.com/3dyuval/go

replace github.com/3dyuval/go/hello => ./hello

replace github.com/3dyuval/go/world => ./world

go 1.21.5

require (
	github.com/3dyuval/go/hello v0.0.0-00010101000000-000000000000
	github.com/3dyuval/go/world v0.0.0-00010101000000-000000000000
)
