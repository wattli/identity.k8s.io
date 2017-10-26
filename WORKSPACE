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
    sha256 = "2bad6a66735a06756a7abc8974f29f762e314a8c8febfd56f3ceb0ebd1c0d356",
    strip_prefix = "rules_docker-58d022892232e5d59daba7760289976d5f6e7433",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/58d022892232e5d59daba7760289976d5f6e7433.tar.gz"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@io_bazel_rules_go//proto:def.bzl", "proto_register_toolchains")
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
    container_repositories = "repositories",
)

go_rules_dependencies()

go_register_toolchains()

proto_register_toolchains()

container_repositories()

container_pull(
    name = "busybox",
    digest = "sha256:be3c11fdba7cfe299214e46edc642e09514dbb9bbefcd0d3836c05a1e0cd0642",
    registry = "index.docker.io",
    repository = "library/busybox",
    tag = "latest",  # ignored, but kept here for documentation
)
