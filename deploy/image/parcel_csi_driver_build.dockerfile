# Parcel-CSI-Driver-Build
#
# VERSION	1.0


##############################################
# Build parcel-csi-driver
##############################################
FROM golang:1.14.4-stretch
LABEL maintainer="Illyoung Choi <iychoi@email.arizona.edu>"
LABEL version="0.1"
LABEL description="Parcel CSI Driver Build Image"

ARG SRC_DIR="/go/src/github.com/iychoi/parcel-csi-driver/"

WORKDIR ${SRC_DIR}

# Cache go modules
ENV GOPROXY=direct
COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . .
RUN make parcel-csi-driver

