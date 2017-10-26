load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    importpath = "k8s.io/identity/pkg/workload",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/apis/idmgr:go_default_library",
        "//pkg/apis/workload:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
    ],
)
