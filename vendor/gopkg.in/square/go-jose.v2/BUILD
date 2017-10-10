load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "asymmetric.go",
        "crypter.go",
        "doc.go",
        "encoding.go",
        "jwe.go",
        "jwk.go",
        "jws.go",
        "shared.go",
        "signing.go",
        "symmetric.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/crypto/ed25519:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2/cipher:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2/json:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "asymmetric_test.go",
        "crypter_test.go",
        "doc_test.go",
        "encoding_test.go",
        "jwe_test.go",
        "jwk_test.go",
        "jws_test.go",
        "signing_test.go",
        "symmetric_test.go",
        "utils_test.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/crypto/ed25519:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2/json:go_default_library",
    ],
)
