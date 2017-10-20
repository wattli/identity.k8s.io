package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_prefix")

go_prefix("k8s.io/identity")

go_library(
    name = "go_default_library",
    srcs = ["interceptors.go"],
    importpath = "k8s.io/identity",
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
