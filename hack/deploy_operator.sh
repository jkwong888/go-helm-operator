#!/bin/bash

SCRIPT_ROOT=$(dirname "${BASH_SOURCE}")/..
SCRIPT_BASE=${SCRIPT_ROOT}/../..

# build binaries in an x86_64 go container
echo "Compile for x86_64 using docker ..."
docker run \
  --rm \
  -v ${GOPATH}:${GOPATH} \
  -e GOPATH=${GOPATH} \
  -w `pwd` \
  golang:1.11 \
  go build -v -o build/_output/bin/websphere-liberty-operator ./cmd/manager

# build the docker image
CHART_VERSION=`cat build/_output/chart/stable/ibm-websphere-liberty/Chart.yaml | grep '^version:' | cut -d: -f2 | sed -e 's/ *//g'`
IMAGE=harbor.jkwong.cloudns.cx/jkwong/websphere-liberty-operator:${CHART_VERSION}

echo "Build image ${IMAGE} ..."
docker build -t ${IMAGE} -f build/Dockerfile .

# push the image
echo "Push image ${IMAGE} ..."
docker push ${IMAGE}

sed -i'' -e 's|REPLACE_IMAGE|'${IMAGE}'|g' ${SCRIPT_ROOT}/deploy/operator.yaml

# create the CRD on the target cluster first
echo "Create CRD ..."
kubectl apply -f ${SCRIPT_ROOT}/deploy/crds/liberty_v1alpha1_libertyapp_crd.yaml

# deploy the serviceaccount, role, and rolebinding
echo "Create ServiceAccount ..."
kubectl apply -f ${SCRIPT_ROOT}/deploy/role.yaml
kubectl apply -f ${SCRIPT_ROOT}/deploy/service_account.yaml
kubectl apply -f ${SCRIPT_ROOT}/deploy/role_binding.yaml

# now deploy the operator
echo "Create Operator Deployment ..."
kubectl apply -f ${SCRIPT_ROOT}/deploy/operator.yaml