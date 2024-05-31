#!/usr/bin/bash

source "$(cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd)/utils.sh"

kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_DIR}/boilerplate.go.txt" \
    "${PROJECT_DIR}"

if [[ -n "${API_KNOWN_VIOLATIONS_DIR:-}" ]]; then
    report_filename="${API_KNOWN_VIOLATIONS_DIR}/codegen_violation_exceptions.list"
    if [[ "${UPDATE_API_KNOWN_VIOLATIONS:-}" == "true" ]]; then
        update_report="--update-report"
    fi
fi

kube::codegen::gen_client \
    --with-watch \
    --with-applyconfig \
    --output-dir "${PROJECT_DIR}/client" \
    --output-pkg "github.com/SimonRTC/kubeception/apis" \
    --boilerplate "${SCRIPT_DIR}/boilerplate.go.txt" \
    "${PROJECT_DIR}/apis"