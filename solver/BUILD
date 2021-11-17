load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_binary")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@com_google_protobuf//:protobuf.bzl", "py_proto_library")

proto_library(
    name = "solver_proto",
    srcs = ["solver.proto"],
)

go_binary(
    name = "solver",
    srcs = [
        "main.go",
    ],
    deps = [
        ":solver_go_proto",
        "//solver/solvercgo",
    ],
)

cc_library(
    name = "solver_cc_library",
    srcs = ["solver.cpp"],
    hdrs = [
        "solver.hpp",
    ],
    visibility = ["//visibility:public"],
    deps = [
        ":solver_cc_proto",
    ],
)

cc_proto_library(
    name = "solver_cc_proto",
    visibility = ["//visibility:public"],
    deps = [":solver_proto"],
)

go_proto_library(
    name = "solver_go_proto",
    importpath = "hpe.one/solver/solverpb",
    protos = [":solver_proto"],
    visibility = ["//visibility:public"],
)
