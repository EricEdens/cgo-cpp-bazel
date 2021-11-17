load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")
load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@com_google_protobuf//:protobuf.bzl", "py_proto_library")

go_library(
    name = "solvercgo",
    srcs = ["solvercgo.go"],
    cdeps = [
        ":solver_c_shim",
    ],
    cgo = True,
    importpath = "hpe.one/solver/solvercgo",
    visibility = ["//visibility:public"],
    deps = [
        "//solver:solver_go_proto",
        "@org_golang_google_protobuf//proto",
    ],
)

cc_library(
    name = "solver_c_shim",
    srcs = ["solver_c_shim.cpp"],
    hdrs = [
        "solver_c_shim.h",
    ],
    deps = [
        "//cppgo",
        "//solver:solver_cc_library",
        "//solver:solver_cc_proto",
    ],
)
