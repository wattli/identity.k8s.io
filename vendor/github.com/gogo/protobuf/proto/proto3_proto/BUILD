load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "proto3_proto_proto",
    srcs = ["proto3.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//testdata:testdata_proto",
        "@com_google_protobuf//:any_proto",
    ],
)

go_proto_library(
    name = "proto3_proto_go_proto",
    importpath = "github.com/gogo/protobuf/proto/proto3_proto",
    proto = ":proto3_proto_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//testdata:go_default_library",
        "@com_github_golang_protobuf//ptypes/any:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["proto3.pb.go"],
    importpath = "github.com/gogo/protobuf/proto/proto3_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto/testdata:go_default_library",
        "//vendor/github.com/gogo/protobuf/types:go_default_library",
    ],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["proto3.proto"],
    visibility = ["//visibility:public"],
)
