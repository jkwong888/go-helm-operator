#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..

if [ -d ${SCRIPT_ROOT}/build/_output/chart ]; then
  echo "Removing ${SCRIPT_ROOT}/build/_output/chart ..."
  rm -rf ${SCRIPT_ROOT}/build/_output/chart
fi

git clone --depth 1 --single-branch https://github.com/IBM/charts ${SCRIPT_ROOT}/build/_output/chart 

${SCRIPT_ROOT}/hack/libertyappspec_gen.sh ${SCRIPT_ROOT}/build/_output/chart/stable/ibm-websphere-liberty
${SCRIPT_ROOT}/hack/deepcopy_gen.sh 