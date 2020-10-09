# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

PKG=github.com/iychoi/parcel-csi-driver
CSI_DRIVER_BUILD_IMAGE=parcel_csi_driver_build
CSI_DRIVER_BUILD_DOCKERFILE=deploy/image/parcel_csi_driver_build.dockerfile
IRODS_FUSE_CLIENT_BUILD_IMAGE=irods_fuse_client_build
IRODS_FUSE_CLIENT_BUILD_DOCKERFILE=deploy/image/irods_fuse_build.dockerfile
CSI_DRIVER_IMAGE?=iychoi/parcel-csi-driver
CSI_DRIVER_DOCKERFILE=deploy/image/parcel_csi_driver_image.dockerfile
VERSION=v0.1.0
GIT_COMMIT?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS?="-X ${PKG}/pkg/driver.driverVersion=${VERSION} -X ${PKG}/pkg/driver.gitCommit=${GIT_COMMIT} -X ${PKG}/pkg/driver.buildDate=${BUILD_DATE}"
GO111MODULE=on
GOPROXY=direct
GOPATH=$(shell go env GOPATH)

.EXPORT_ALL_VARIABLES:

.PHONY: parcel-csi-driver
parcel-csi-driver:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -ldflags ${LDFLAGS} -o bin/parcel-csi-driver ./cmd/

.PHONY: irods_fuse_build
irods_fuse_build:
	docker build -t $(IRODS_FUSE_CLIENT_BUILD_IMAGE):latest -f $(IRODS_FUSE_CLIENT_BUILD_DOCKERFILE) .

.PHONY: driver_build
driver_build:
	docker build -t $(CSI_DRIVER_BUILD_IMAGE):latest -f $(CSI_DRIVER_BUILD_DOCKERFILE) .

.PHONY: image
image: irods_fuse_build driver_build
	docker build -t $(CSI_DRIVER_IMAGE):latest -f $(CSI_DRIVER_DOCKERFILE) .

.PHONY: push
push: image
	docker push $(CSI_DRIVER_IMAGE):latest

.PHONY: image-release
image-release:
	docker build -t $(CSI_DRIVER_IMAGE):$(VERSION) -f $(CSI_DRIVER_DOCKERFILE) .

.PHONY: push-release
push-release:
	docker push $(CSI_DRIVER_IMAGE):$(VERSION)
