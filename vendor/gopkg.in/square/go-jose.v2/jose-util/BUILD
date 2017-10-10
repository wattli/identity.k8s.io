load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "utils.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2/jose-util",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/gopkg.in/alecthomas/kingpin.v2:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
    ],
)

go_binary(
    name = "jose-util",
    importpath = "gopkg.in/square/go-jose.v2/jose-util",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
