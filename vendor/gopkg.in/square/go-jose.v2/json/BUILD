load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "decode.go",
        "encode.go",
        "indent.go",
        "scanner.go",
        "stream.go",
        "tags.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2/json",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "bench_test.go",
        "decode_test.go",
        "encode_test.go",
        "number_test.go",
        "scanner_test.go",
        "stream_test.go",
        "tagkey_test.go",
        "tags_test.go",
    ],
    data = glob(["testdata/**"]),
    importpath = "gopkg.in/square/go-jose.v2/json",
    library = ":go_default_library",
)
