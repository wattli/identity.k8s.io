load("@io_bazel_rules_go//go:def.bzl", "go_library")
load("@io_bazel_rules_go//proto:def.bzl", "go_proto_library")

proto_library(
    name = "leasepb_proto",
    srcs = ["lease.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//etcd/etcdserver/etcdserverpb:etcdserverpb_proto",
        "//gogoproto:gogoproto_proto",
    ],
)

go_proto_library(
    name = "leasepb_go_proto",
    importpath = "github.com/coreos/etcd/lease/leasepb",
    proto = ":leasepb_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//etcd/etcdserver/etcdserverpb:go_default_library",
        "//gogoproto:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = ["lease.pb.go"],
    importpath = "github.com/coreos/etcd/lease/leasepb",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/etcd/etcdserver/etcdserverpb:go_default_library",
        "//vendor/github.com/golang/protobuf/proto:go_default_library",
    ],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["lease.proto"],
    visibility = ["//visibility:public"],
)
