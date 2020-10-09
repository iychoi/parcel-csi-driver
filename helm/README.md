## Parcel CSI Driver Helm Chart
This script enables easy installation of Parcel CSI Driver using Helm Chart.

### Compatibility
- Helm 3+
- Kubernetes > 1.17.x, can be deployed to any namespace.
- Kubernetes < 1.17.x, namespace **must** be `kube-system`, as `system-cluster-critical` hard coded to this namespace.

### Install

Kubernetes > 1.17.x
```shell script
helm install parcel-csi-driver .
```

Kubernetes < 1.17.x
```shell script
helm install parcel-csi-driver --namespace kube-system .
```

### Uninstall
```shell script
helm delete parcel-csi-driver
```

