load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

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
    visibility = ["//cmd/idmgr:__pkg__"],
)
