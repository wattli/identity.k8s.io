load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//docker:docker.bzl", "docker_build")

go_library(
    name = "go_default_library",
    srcs = ["true.go"],
    importpath = "k8s.io/identity/cmd/true",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "true",
    importpath = "k8s.io/identity/cmd/true",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

docker_build(
    name = "image",
    entrypoint = ["/true"],
    files = [
        "//cmd/true",
    ],
    visibility = ["//visibility:public"],
)
