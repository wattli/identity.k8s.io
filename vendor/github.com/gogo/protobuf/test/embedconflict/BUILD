load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "embedconflict_proto",
    srcs = [
        "eb.proto",
        "ec.proto",
        "ee.proto",
        "em.proto",
        "en.proto",
        "er.proto",
    ],
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:gogoproto_proto"],
)

go_proto_library(
    name = "embedconflict_go_proto",
    importpath = "github.com/gogo/protobuf/test/embedconflict",
    proto = ":embedconflict_proto",
    visibility = ["//visibility:public"],
    deps = ["//github.com/gogo/protobuf/gogoproto:go_default_library"],
)

go_library(
    name = "go_default_library",
    srcs = ["doc.go"],
    importpath = "github.com/gogo/protobuf/test/embedconflict",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["ec_test.go"],
    importpath = "github.com/gogo/protobuf/test/embedconflict",
    library = ":go_default_library",
)

filegroup(
    name = "go_default_library_protos",
    srcs = [
        "eb.proto",
        "ec.proto",
        "ee.proto",
        "em.proto",
        "en.proto",
        "er.proto",
    ],
    visibility = ["//visibility:public"],
)
