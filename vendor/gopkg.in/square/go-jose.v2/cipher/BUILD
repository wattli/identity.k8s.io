load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "cbc_hmac.go",
        "concat_kdf.go",
        "ecdh_es.go",
        "key_wrap.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2/cipher",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "cbc_hmac_test.go",
        "concat_kdf_test.go",
        "ecdh_es_test.go",
        "key_wrap_test.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2/cipher",
    library = ":go_default_library",
)
