load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_grpc_library")

proto_library(
    name = "grpc_lb_v1_proto",
    srcs = ["grpclb.proto"],
    visibility = ["//visibility:public"],
)

go_grpc_library(
    name = "grpc_lb_v1_go_proto",
    importpath = "google.golang.org/grpc/grpclb/grpc_lb_v1",
    proto = ":grpc_lb_v1_proto",
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["grpclb.pb.go"],
    importpath = "google.golang.org/grpc/grpclb/grpc_lb_v1",
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/golang/protobuf/proto:go_default_library"],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["grpclb.proto"],
    visibility = ["//visibility:public"],
)
