load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")

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

container_image(
    name = "image",
    base = "@busybox//image",
    entrypoint = ["/identity-apiserver"],
    files = [":identity-apiserver"],
)

container_push(
    name = "push",
    format = "Docker",
    image = ":image",
    registry = "gcr.io",
    repository = "mikedanese-k8s/identity-apiserver",
    tag = "latest",
)
