#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..

# build struct-gen binary
go build -o ${SCRIPT_ROOT}/bin/struct-gen ${SCRIPT_ROOT}/cmd/struct-gen

# generate appspec based on chart
echo "Generating pkg/apis/liberty/v1alpha1/libertyappspec.go"
${SCRIPT_ROOT}/bin/struct-gen v1alpha1 LibertyApp ibm-websphere-liberty-rhel/values.yaml ${SCRIPT_ROOT}/pkg/apis/liberty/v1alpha1/libertyappspec.go
