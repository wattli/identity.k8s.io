load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//docker:docker.bzl", "docker_build", "docker_push")

go_library(
    name = "go_default_library",
    srcs = [
        "idmgr.go",
        "server.go",
    ],
    importpath = "k8s.io/identity/cmd/idmgr",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/apis/idmgr:go_default_library",
        "//pkg/volumemgr:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/health:go_default_library",
    ],
)

go_binary(
    name = "idmgr",
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
    ],
    importpath = "k8s.io/identity/cmd/idmgr",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

docker_build(
    name = "image",
    entrypoint = ["/idmgr"],
    files = [":idmgr"],
)

docker_push(
    name = "push",
    image = ":image",
    registry = "gcr.io",
    repository = "mikedanese-k8s/idmgr",
    tag = "latest",
)
