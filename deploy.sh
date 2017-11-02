#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset

main() {
  local k="${HOME}/go/src/k8s.io/kubernetes/bazel-bin/cmd/kubectl/kubectl"
  bazel run //cmd/identity-apiserver:push
  bazel run //cmd/idmgr:push
  "${k}" apply -f  example/identity-apiserver.yaml
  "${k}" apply -f  example/identity-apiserver-deployment.yaml
  "${k}" delete -f example/idmgr.yaml || true
  "${k}" apply -f  example/idmgr.yaml
}

main "${@}"
