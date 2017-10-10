load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["rest.go"],
    importpath = "k8s.io/identity/pkg/registry/identity/identitydocument",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/apis/identity:go_default_library",
        "//pkg/jwt:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/endpoints/request:go_default_library",
    ],
)
