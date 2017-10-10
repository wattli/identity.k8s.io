load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//docker:docker.bzl", "docker_build", "docker_push")

go_library(
    name = "go_default_library",
    srcs = ["api.go"],
    importpath = "k8s.io/identity/cmd/identity-apiserver",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/server:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/server:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/logs:go_default_library",
    ],
)

go_binary(
    name = "identity-apiserver",
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
    ],
    importpath = "k8s.io/identity/cmd/identity-apiserver",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

docker_build(
    name = "image",
    entrypoint = ["/identity-apiserver"],
    files = [":identity-apiserver"],
)

docker_push(
    name = "push",
    image = ":image",
    registry = "gcr.io",
    repository = "mikedanese-k8s/identity-apiserver",
    tag = "latest",
)
