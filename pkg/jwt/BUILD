load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["signer.go"],
    importpath = "k8s.io/identity/pkg/jwt",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2/jwt:go_default_library",
    ],
)
