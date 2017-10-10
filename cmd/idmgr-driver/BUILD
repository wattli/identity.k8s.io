load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//docker:docker.bzl", "docker_build", "docker_push")

go_library(
    name = "go_default_library",
    srcs = ["volume.go"],
    importpath = "k8s.io/identity/cmd/idmgr-driver",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/apis/idmgr:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/logs:go_default_library",
    ],
)

go_binary(
    name = "idmgr-driver",
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
    ],
    importpath = "k8s.io/identity/cmd/idmgr-driver",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

docker_build(
    name = "image",
    base = "//cmd/true:image",
    directory = "/flexvolume/k8s~idmgr/",
    files = [
        ":idmgr-driver",
    ],
)

docker_push(
    name = "push",
    image = ":image",
    registry = "gcr.io",
    repository = "mikedanese-k8s/idmgr-driver",
    tag = "latest",
)
