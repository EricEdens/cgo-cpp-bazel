load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "main",
    srcs = ["main.go"],
    importpath = "hpe.one/cli",
    deps = [
        "//server:go_server",
    ],
)
