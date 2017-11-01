load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "manager.go",
        "server.go",
    ],
    importpath = "k8s.io/identity/pkg/management",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/apis/idmgr:go_default_library",
        "//pkg/apis/workload:go_default_library",
        "//pkg/client/clientset/typed/identity/v1alpha1:go_default_library",
        "//pkg/uds:go_default_library",
        "//pkg/util:go_default_library",
        "//pkg/workload:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
    ],
)
