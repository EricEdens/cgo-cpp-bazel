load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

cc_library(
    name = "server",
    srcs = [
        "server.cpp",
    ],
    hdrs = ["server.h"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_server",
    srcs = ["server.go"],
    cdeps = [":server"],
    cgo = True,
    importpath = "hpe.one/server",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_server_test",
    srcs = ["server_test.go"],
    deps = [
        ":go_server",
        "@com_github_stretchr_testify//assert",
    ],
)
