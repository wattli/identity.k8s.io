load("@io_bazel_rules_go//proto:def.bzl", "go_grpc_library")

package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "conversion.go",
        "defaults.go",
        "doc.go",
        "generated.pb.go",
        "register.go",
        "types.go",
        "types_swagger_doc_generated.go",
        "zz_generated.conversion.go",
        "zz_generated.deepcopy.go",
        "zz_generated.defaults.go",
    ],
    importpath = "k8s.io/apimachinery/pkg/apis/testapigroup/v1",
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/sortkeys:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/apis/testapigroup:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/conversion:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["generated.proto"],
    visibility = ["//visibility:public"],
)

proto_library(
    name = "v1_proto",
    srcs = ["generated.proto"],
    visibility = ["//visibility:public"],
    deps = [
        "//k8s.io/apimachinery/pkg/apis/meta/v1:v1_proto",
        "//k8s.io/apimachinery/pkg/runtime:runtime_proto",
        "//k8s.io/apimachinery/pkg/runtime/schema:schema_proto",
        "//k8s.io/apimachinery/pkg/util/intstr:intstr_proto",
    ],
)

go_grpc_library(
    name = "v1_go_proto",
    importpath = "k8s.io/apimachinery/pkg/apis/testapigroup/v1",
    proto = ":v1_proto",
    visibility = ["//visibility:public"],
    deps = [
        "//k8s.io/apimachinery/pkg/apis/meta/v1:go_default_library",
        "//k8s.io/apimachinery/pkg/runtime:go_default_library",
        "//k8s.io/apimachinery/pkg/runtime/schema:go_default_library",
        "//k8s.io/apimachinery/pkg/util/intstr:go_default_library",
    ],
)
