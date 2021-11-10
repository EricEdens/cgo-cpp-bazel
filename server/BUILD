load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@rules_proto//proto:defs.bzl", "proto_library")

proto_library(
    name = "response_proto",
    srcs = ["response.proto"],
)

cc_proto_library(
    name = "response_cc_proto",
    deps = [":response_proto"],
)

go_proto_library(
    name = "response_go_proto",
    importpath = "hpe.one/serverpb",
    protos = [":response_proto"],
)

cc_library(
    name = "server",
    srcs = [
        "server.cpp",
    ],
    hdrs = ["server.h"],
    visibility = ["//visibility:public"],
    deps = ["response_cc_proto"],
)

go_library(
    name = "go_server",
    srcs = ["server.go"],
    cdeps = [":server"],
    cgo = True,
    importpath = "hpe.one/server",
    visibility = ["//visibility:public"],
    deps = [
        ":response_go_proto",
        "@org_golang_google_protobuf//proto",
    ],
)

go_test(
    name = "go_server_test",
    srcs = ["server_test.go"],
    importpath = "hpe.one/server_test",
    deps = [
        ":go_server",
        ":response_go_proto",
        "@com_github_stretchr_testify//assert",
    ],
)
