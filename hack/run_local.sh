#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..

# build
go build -o ${SCRIPT_ROOT}/bin/websphere-liberty-operator ${SCRIPT_ROOT}/cmd/manager

# run locally
WATCH_NAMESPACE=plantsbywebsphere
${SCRIPT_ROOT}/bin/websphere-liberty-operator

