load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@io_bazel_rules_docker//container:container.bzl", "container_image", "container_push")
load("@bazel_tools//tools/build_defs/pkg:pkg.bzl", "pkg_tar")

go_library(
    name = "go_default_library",
    srcs = ["idmgr.go"],
    importpath = "k8s.io/identity/cmd/idmgr",
    visibility = ["//visibility:private"],
    deps = [
        "//pkg/apis/idmgr:go_default_library",
        "//pkg/client/clientset/typed/identity/v1alpha1:go_default_library",
        "//pkg/management:go_default_library",
        "//pkg/uds:go_default_library",
        "//pkg/util:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/k8s.io/client-go/rest:go_default_library",
        "//vendor/k8s.io/client-go/tools/clientcmd:go_default_library",
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

pkg_tar(
    name = "idmgr-pkg",
    srcs = [
        ":idmgr",
        "//cmd/idcli",
        "//cmd/idmgr-driver",
    ],
    package_dir = "usr/local/bin",
    visibility = ["//visibility:private"],
)

container_image(
    name = "image",
    base = "@busybox//image",
    entrypoint = ["/usr/local/bin/idmgr"],
    tars = [":idmgr-pkg"],
)

container_push(
    name = "push",
    format = "Docker",
    image = ":image",
    registry = "gcr.io",
    repository = "mikedanese-k8s/idmgr",
    tag = "latest",
)
