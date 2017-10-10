load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "generated_expansion.go",
        "identity_client.go",
    ],
    importpath = "k8s.io/identity/pkg/client/clientset_generated/clientset/typed/identity/v1alpha1",
    visibility = ["//visibility:public"],
    deps = ["//pkg/client/clientset_generated/clientset/scheme:go_default_library"],
)
