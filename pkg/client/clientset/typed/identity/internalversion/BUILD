load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "generated_expansion.go",
        "identity_client.go",
        "identitydocument.go",
    ],
    importpath = "k8s.io/identity/pkg/client/clientset/typed/identity/internalversion",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/client/clientset/scheme:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
    ],
)
