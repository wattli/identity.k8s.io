load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "generated_expansion.go",
        "identity_client.go",
    ],
    importpath = "k8s.io/identity/pkg/client/clientset_generated/internalclientset/typed/identity/internalversion",
    visibility = ["//visibility:public"],
    deps = ["//pkg/client/clientset_generated/internalclientset/scheme:go_default_library"],
)
