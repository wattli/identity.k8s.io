load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "unrecognized_proto",
    srcs = ["unrecognized.proto"],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "unrecognized_go_proto",
    importpath = "github.com/gogo/protobuf/test/unrecognized",
    proto = ":unrecognized_proto",
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["unrecognized.pb.go"],
    importpath = "github.com/gogo/protobuf/test/unrecognized",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/descriptor:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "oldnew_test.go",
        "unrecognizedpb_test.go",
    ],
    importpath = "github.com/gogo/protobuf/test/unrecognized",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/jsonpb:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
    ],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["unrecognized.proto"],
    visibility = ["//visibility:public"],
)
