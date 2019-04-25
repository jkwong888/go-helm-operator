#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..

# build
go build -o ${SCRIPT_ROOT}/build/_output/websphere-liberty-operator ${SCRIPT_ROOT}/cmd/manager

# run locally
WATCH_NAMESPACE=plantsbywebsphere
OPERATOR_NAME=websphere-liberty-operator
${SCRIPT_ROOT}/build/_output/websphere-liberty-operator

