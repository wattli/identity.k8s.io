#!/bin/bash

# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_PACKAGE=k8s.io/identity
SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo k8s.io/code-generator)}

clientgen="${PWD}/client-gen"
listergen="${PWD}/lister-gen"
informergen="${PWD}/informer-gen"
defaultergen="${PWD}/defaulter-gen"
deepcopygen="${PWD}/deepcopy-gen"
conversiongen="${PWD}/conversion-gen"
# Register function to be called on EXIT to remove generated binary.
cleanup() {
  rm -f "${clientgen:-}"
  rm -f "${listergen:-}"
  rm -f "${informergen:-}"
  rm -f "${defaultergen:-}"
  rm -f "${deepcopygen:-}"
  rm -f "${conversiongen:-}"
}

generate_group() {
  local GROUP_NAME=$1
  local VERSION=$2
  local CLIENT_PKG=${SCRIPT_PACKAGE}/pkg/client
  local LISTERS_PKG=${CLIENT_PKG}/listers_generated
  local INFORMERS_PKG=${CLIENT_PKG}/informers_generated
  local APIS_PKG=${SCRIPT_PACKAGE}/pkg/apis
  local INPUT_APIS="${GROUP_NAME}/,${GROUP_NAME}/${VERSION}"

  echo "Building client-gen"
  go build -o "${clientgen}" ${CODEGEN_PKG}/cmd/client-gen

  echo "generating clientset for group ${GROUP_NAME} and version ${VERSION} at ${SCRIPT_BASE}/${CLIENT_PKG}"
  ${clientgen} \
    --clientset-name clientset \
    --input-base ${APIS_PKG} \
    --input ${INPUT_APIS} \
    --clientset-path ${CLIENT_PKG}
  ${clientgen} \
    --clientset-name clientset \
    --input-base ${APIS_PKG} \
    --input ${INPUT_APIS} \
    --clientset-path ${CLIENT_PKG}

  echo "Building lister-gen"
  go build -o "${listergen}" ${CODEGEN_PKG}/cmd/lister-gen

  echo "generating listers for group ${GROUP_NAME} and version ${VERSION} at ${SCRIPT_BASE}/${LISTERS_PKG}"
  ${listergen} --input-dirs ${APIS_PKG}/${GROUP_NAME} --input-dirs ${APIS_PKG}/${GROUP_NAME}/${VERSION} --output-package ${LISTERS_PKG} --output-base ${SCRIPT_BASE}

  echo "Building informer-gen"
  go build -o "${informergen}" ${CODEGEN_PKG}/cmd/informer-gen

  echo "generating informers for group ${GROUP_NAME} and version ${VERSION} at ${SCRIPT_BASE}/${INFORMERS_PKG}"
  ${informergen} \
    --input-dirs ${APIS_PKG}/${GROUP_NAME} --input-dirs ${APIS_PKG}/${GROUP_NAME}/${VERSION} \
    --versioned-clientset-package ${CLIENT_PKG}/clientset \
    --internal-clientset-package ${CLIENT_PKG}/internalclientset \
    --listers-package ${LISTERS_PKG} \
    --output-package ${INFORMERS_PKG} \
    --output-base ${SCRIPT_BASE}

  echo "Building defaulter-gen"
  go build -o "${defaultergen}" ${CODEGEN_PKG}/cmd/informer-gen

  echo "generating defaults"
  ${defaultergen} \
    --v 1 --logtostderr \
    --input-dirs "${APIS_PKG}/${GROUP_NAME}" \
    --input-dirs "${APIS_PKG}/${GROUP_NAME}/${VERSION}" \
    --output-file-base "zz_generated.defaults"

  echo "Building deepcopy-gen"
  go build -o "${deepcopygen}" ${CODEGEN_PKG}/cmd/deepcopy-gen

  echo "generating deepcopies"
  ${deepcopygen} \
    --v 1 --logtostderr \
    --input-dirs "${APIS_PKG}/${GROUP_NAME}" \
    --input-dirs "${APIS_PKG}/${GROUP_NAME}/${VERSION}" \
    --output-file-base zz_generated.deepcopy

  echo "Building conversion-gen"
  go build -o "${conversiongen}" ${CODEGEN_PKG}/cmd/conversion-gen

  echo "generating conversions"
  ${conversiongen} \
    --v 1 --logtostderr \
    --input-dirs "${APIS_PKG}/${GROUP_NAME}" \
    --input-dirs "${APIS_PKG}/${GROUP_NAME}/${VERSION}" \
    --output-file-base zz_generated.conversion
}

main() {
  trap cleanup EXIT

  generate_group identity v1alpha1
}

main "$@"
