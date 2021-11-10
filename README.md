# cgo-cpp-bazel

Are you working with Go in Bazel and need to call C++? This
repo shows how to make that happen.

## The plan

[cgo](https://pkg.go.dev/cmd/cgo) allows for interop between Go and C.
For Go to call C++, you expose a C-compliant surface, and then compile
that into a `cc_library`. This `cc_library` can be consumed as a `cdep`.

This example uses structs to pass data between Go and C.

## Adding dependencies

1. Use **go get** to download the dependency: `go get github.com/stretchr/testify`
2. Update WORKSPACE with the gazelle: `bazel run //:gazelle-update-repos`

## Running benchmarks

`bazel run server:go_server_test -- -test.bench=.`