load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "generated_expansion.go",
        "identity_client.go",
        "identitydocument.go",
        "identitydocument_expansion.go",
    ],
    importpath = "k8s.io/identity/pkg/client/clientset/typed/identity/v1alpha1",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/apis/identity/v1alpha1:go_default_library",
        "//pkg/client/clientset/scheme:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/serializer:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/client-go/transport:go_default_library",
    ],
)
