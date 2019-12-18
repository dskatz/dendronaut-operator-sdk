FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/dendronaut-operator-sdk \
    USER_UID=1001 \
    USER_NAME=dendronaut-operator-sdk

# install operator binary
COPY build/_output/bin/dendronaut-operator-sdk ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
