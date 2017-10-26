load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fake_identity_client.go",
    ],
    importpath = "k8s.io/identity/pkg/client/clientset/typed/identity/v1alpha1/fake",
    visibility = ["//visibility:public"],
    deps = ["//pkg/client/clientset/typed/identity/v1alpha1:go_default_library"],
)
