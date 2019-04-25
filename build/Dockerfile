FROM registry.access.redhat.com/ubi7-dev-preview/ubi-minimal:7.6

ENV OPERATOR=/usr/local/bin/websphere-liberty-operator \
    USER_UID=1001 \
    USER_NAME=websphere-liberty-operator

# install operator binary
COPY build/_output/bin/websphere-liberty-operator ${OPERATOR}

RUN mkdir -p /opt/helm/charts
COPY build/_output/chart/stable/ibm-websphere-liberty/* /opt/helm/charts/

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
