load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "builder.go",
        "claims.go",
        "doc.go",
        "errors.go",
        "jwt.go",
        "validation.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2/jwt",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2/json:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "builder_test.go",
        "claims_test.go",
        "jwt_test.go",
        "validation_test.go",
    ],
    importpath = "gopkg.in/square/go-jose.v2/jwt",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2/json:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = ["example_test.go"],
    importpath = "gopkg.in/square/go-jose.v2/jwt_test",
    deps = [
        ":go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
    ],
)
