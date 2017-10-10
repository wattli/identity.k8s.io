http_archive(
    name = "io_bazel_rules_go",
    sha256 = "1f56ddab69248b1df1a1d54f5a5b25d006483beb0a0bed9d9f8a0b9fc2de4733",
    strip_prefix = "rules_go-0cd983e2c6f6665fcae3e93eb303a00e54dc6fc5",
    urls = ["https://github.com/bazelbuild/rules_go/archive/0cd983e2c6f6665fcae3e93eb303a00e54dc6fc5.tar.gz"],
)

http_archive(
    name = "io_kubernetes_build",
    sha256 = "8e49ac066fbaadd475bd63762caa90f81cd1880eba4cc25faa93355ef5fa2739",
    strip_prefix = "repo-infra-e26fc85d14a1d3dc25569831acc06919673c545a",
    urls = ["https://github.com/kubernetes/repo-infra/archive/e26fc85d14a1d3dc25569831acc06919673c545a.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "e86b8764fccc62dddf6e08382ba692b16479a2af478080b1ece4d9add8abbb9a",
    strip_prefix = "rules_docker-28d492bc1dc1275e2c6ff74e51adc864e59ddc76",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/28d492bc1dc1275e2c6ff74e51adc864e59ddc76.tar.gz"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")
load("@io_bazel_rules_docker//docker:docker.bzl", "docker_repositories")

go_rules_dependencies()

go_register_toolchains()

proto_register_toolchains()

docker_repositories()
