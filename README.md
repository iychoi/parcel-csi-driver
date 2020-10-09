## Parcel CSI Driver

Parcel Container Storage Interface (CSI) Driver implements the [CSI Specification](https://github.com/container-storage-interface/spec/blob/master/spec.md) to provide container orchestration engines (like [Kubernetes](https://kubernetes.io/)) public scientific datasets access.

### CSI Specification Compatibility

Parcel CSI Driver only supports CSI Specification Version v1.2.0 or higher.

### Features

Parcel CSI Driver relies on external clients for mounting datasets.
| Client Type | iRODS Client     | Description                     |
|-------------|------------------|---------------------------------|
| irodsfuse   | iRODS FUSE       | For datasets stored in iRODS    |
| webdav      | DavFS2           | For datasets shared via WebDAV  |

### Volume Mount Parameters

Parameters specified in Persistent Volume (PV) and Storage Class (SC) are passed to Parcel CSI Driver to mount a volume.
Depending on driver types, different parameters should be given.

Parameters are given via Persistent Volume (PV).

#### iRODS FUSE Client
| Field | Description | Example |
| --- | --- | --- |
| client | Client type | "irodsfuse" |
| user | iRODS user id | "irods_user" |
| password | iRODS user password | "password" in plane text |
| url | URL | "irods://data.cyverse.org/iplant/home/irods_user" |

Any parameters can be supplied via secrets (using `nodeStageSecretRef`).

#### WebDAV Client
| Field | Description | Example |
| --- | --- | --- |
| client | Client type | "webdav" |
| user | WebDAV user id | "webdav_user" |
| password | WebDAV user password | "password" in plane text |
| url | URL | "https://data.cyverse.org/dav/iplant/home/irods_user" |

Any parameters can be supplied via secrets (using `nodeStageSecretRef`).

### Install & Uninstall

Installation can be done using [Helm Chart](https://github.com/iychoi/parcel-csi-driver/tree/master/helm) or by [manual](https://github.com/cyverse/parcel-csi-driver/tree/master/deploy/kubernetes).

Install using Helm Chart:
```shell script
helm install parcel-csi-driver helm
```

Uninstall using Helm Chart:
```shell script
helm delete parcel-csi-driver
```

### References

The code is based on **iRODS CSI Driver**.
[iRODS-CSI-Driver](https://github.com/cyverse/irods-csi-driver)

### License

This code is licensed under the Apache 2.0 License.
