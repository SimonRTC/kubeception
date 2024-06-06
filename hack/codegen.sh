#!/usr/bin/bash

source "$(cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd)/utils.sh"

kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_DIR}/boilerplate.go.txt" \
    "${PROJECT_DIR}"

kube::codegen::gen_client \
    --with-watch \
    --with-applyconfig \
    --boilerplate "${SCRIPT_DIR}/boilerplate.go.txt" \
    --output-dir "${PROJECT_DIR}/pkg/generated" \
    --output-pkg "github.com/SimonRTC/kubeception/pkg/generated" \
    "${PROJECT_DIR}/apis"

kube::codegen::gen_openapi \
    --update-report \
    --report-filename "${SCRIPT_DIR}/codegen_violation_exceptions.list" \
    --boilerplate "${SCRIPT_DIR}/boilerplate.go.txt" \
    --output-dir "${PROJECT_DIR}/pkg/generated/openapi" \
    --output-pkg "github.com/SimonRTC/kubeception/pkg/generated/openapi" \
    "${PROJECT_DIR}/apis"