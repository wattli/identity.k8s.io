load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "k8s.io/identity/cmd/idcli",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/apis/workload:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/k8s.io/apiserver/pkg/util/logs:go_default_library",
    ],
)

go_binary(
    name = "idcli",
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
    ],
    importpath = "k8s.io/identity/cmd/idcli",
    library = ":go_default_library",
    visibility = ["//cmd/idmgr:__pkg__"],
)
