#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
source ./vendor/k8s.io/code-generator/kube_codegen.sh

THIS_PKG="github.com/bborbe/alert"

kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/k8s/apis"

kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/k8s/client" \
    --output-pkg "${THIS_PKG}/k8s/client" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/k8s/apis"

# Generate deepcopy functions
${CODEGEN_PKG}/cmd/deepcopy-gen/deepcopy-gen \
  --input-dirs "my-project/api/v1" \
  --output-file-base zz_generated.deepcopy \
  --bounding-dirs "my-project/api/v1" \
  --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt"