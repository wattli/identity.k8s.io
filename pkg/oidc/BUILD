load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["oidc.go"],
    importpath = "k8s.io/identity/pkg/oidc",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/authorization/authorizer:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["oidc_test.go"],
    importpath = "k8s.io/identity/pkg/oidc",
    library = ":go_default_library",
)
