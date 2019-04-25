#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..

if [ -z "$1" ]; then
  echo "Usage: ${SCRIPT_ROOT}/hack/$0 <path-to-liberty-chart>"
  exit 1
fi

CHART_PATH=$1

# build struct-gen binary
go build -o ${SCRIPT_ROOT}/build/_output/bin/struct-gen ${SCRIPT_ROOT}/cmd/struct-gen

# generate appspec based on chart
echo "Generating pkg/apis/liberty/v1alpha1/libertyappspec.go"
${SCRIPT_ROOT}/build/_output/bin/struct-gen v1alpha1 LibertyApp ${CHART_PATH}/values.yaml ${SCRIPT_ROOT}/pkg/apis/liberty/v1alpha1/libertyappspec.go

