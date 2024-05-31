#!/usr/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_DIR="$(cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd)"
PROJECT_DIR="$(readlink -e "${SCRIPT_DIR}/..")"
CODE_GENERATOR_DIR="$(readlink -e "${PROJECT_DIR}/../code-generator")"

source "${CODE_GENERATOR_DIR}/kube_codegen.sh"

PATH="${PATH}:${PROJECT_DIR}/bin"
GOPATH=$(go env GOPATH)
GOPATH_PROJECT_DIR="${GOPATH}/src/github.com/SimonRTC/kubeception"

[ -e "${GOPATH_PROJECT_DIR}" ] || (mkdir -p "$(dirname "${GOPATH_PROJECT_DIR}")" && ln -s "${PROJECT_DIR}" "${GOPATH_PROJECT_DIR}")