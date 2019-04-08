#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..

function create_doc_go {
    echo "Generating doc.go in ${1} ..."

    package_name=`basename ${1}`

    cat > ${1}/doc.go <<EOF
// Package v1alpha1 contains API Schema definitions for the liberty v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=liberty.ibm.com
package ${package_name}
EOF
}

# build deepcopy-gen binary
go build -o ${SCRIPT_ROOT}/bin/deepcopy-gen ${SCRIPT_ROOT}/vendor/k8s.io/code-generator/cmd/deepcopy-gen

# add doc.go to the following packages:

dirs="\
k8s.io/helm/pkg/proto/hapi/release \
k8s.io/helm/pkg/proto/hapi/chart \
github.com/golang/protobuf/ptypes/timestamp \
github.com/golang/protobuf/ptypes/any"

for d in ${dirs}; do
    create_doc_go ${SCRIPT_ROOT}/vendor/${d}
done

echo "Generating zz_generated.deepcopy.go ..."
${SCRIPT_ROOT}/bin/deepcopy-gen \
    --input-dirs `echo ${dirs} | sed -e 's/ /,/g'` \
    --output-package github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1 \
    -O zz_generated.deepcopy

${SCRIPT_ROOT}/bin/deepcopy-gen \
    --input-dirs github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1 \
    --output-package github.com/jkwong888/websphere-liberty-operator/pkg/apis/liberty/v1alpha1 \
    -O zz_generated.deepcopy