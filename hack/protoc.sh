#!/usr/bin/bash

source "$(cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd)/utils.sh"

PKGS=(
    github.com/SimonRTC/kubeception/apis/clusters/v1beta1=kubeception.io.apis.clusters.v1beta1
    github.com/SimonRTC/kubeception/apis/nodes/v1beta1=kubeception.io.apis.nodes.v1beta1
    github.com/SimonRTC/kubeception/apis/storage/v1beta1=kubeception.io.apis.storage.v1beta1
)

APIMACHINERY_PKGS=(
    -k8s.io/apimachinery/pkg/util/intstr
    -k8s.io/apimachinery/pkg/api/resource
    -k8s.io/apimachinery/pkg/runtime/schema
    -k8s.io/apimachinery/pkg/runtime
    -k8s.io/apimachinery/pkg/apis/meta/v1
)

[ -d "github.com/SimonRTC" ] || (mkdir -p "github.com/SimonRTC" && ln -s "${PROJECT_DIR}" "github.com/SimonRTC/kubeception")

GO111MODULE=on go-to-protobuf \
    --go-header-file="${SCRIPT_DIR}"/boilerplate.go.txt \
    --packages="$(
        IFS=,
        echo "${PKGS[*]}"
    )" \
    --apimachinery-packages="$(
        IFS=,
        echo "${APIMACHINERY_PKGS[*]}"
    )" \
    --proto-import="${PROJECT_DIR}/vendor" \
    -v 10

[ -d "github.com/SimonRTC" ] && (rm -r "github.com/")