#
# Base Template Dockerfile for Operator Image.
#

ARG imagebase
FROM ${imagebase} as base

FROM registry.access.redhat.com/ubi9/ubi-minimal


LABEL name="MongoDB Kubernetes Operator" \
      maintainer="support@mongodb.com" \
      vendor="MongoDB" \
      version="1.1.0" \
      release="1" \
      summary="MongoDB Kubernetes Operator Image" \
      description="MongoDB Kubernetes Operator Image"


# Building an UBI-based image: https://red.ht/3n6b9y0
RUN microdnf update \
    --disableplugin=subscription-manager \
    --disablerepo=* --enablerepo=ubi-9-appstream-rpms --enablerepo=ubi-9-baseos-rpms -y \
    && rm -rf /var/cache/yum
RUN microdnf install -y glibc-langpack-en




COPY --from=base /data/mongodb-kubernetes-operator /usr/local/bin/mongodb-kubernetes-operator
COPY --from=base /data/om_version_mapping.json /usr/local/om_version_mapping.json
COPY --from=base /data/licenses /licenses/

USER 2000

ENTRYPOINT exec /usr/local/bin/mongodb-kubernetes-operator


