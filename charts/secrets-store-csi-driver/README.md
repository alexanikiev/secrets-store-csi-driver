# secrets-store-csi-driver

## Installation

Quick start instructions for the setup and configuration of secrets-store-csi-driver using Helm.

### Prerequisites

- [Helm v3.0+](https://helm.sh/docs/intro/quickstart/#install-helm)

### Installing the chart

```bash
$ helm repo add secrets-store-csi-driver https://raw.githubusercontent.com/kubernetes-sigs/secrets-store-csi-driver/master/charts
$ helm install csi-secrets-store secrets-store-csi-driver/secrets-store-csi-driver
```

### Configuration

The following table lists the configurable parameters of the csi-secrets-store-provider-azure chart and their default values.

| Parameter                               | Description                                                                                                                       | Default                                             |
| --------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------- |
| `nameOverride`                          | String to partially override secrets-store-csi-driver.fullname template with a string (will prepend the release name)             | `""`                                                |
| `fullnameOverride`                      | String to fully override secrets-store-csi-driver.fullname template with a string                                                 | `""`                                                |
| `linux.image.repository`                | Linux image repository                                                                                                            | `k8s.gcr.io/csi-secrets-store/driver`               |
| `linux.image.pullPolicy`                | Linux image pull policy                                                                                                           | `Always`                                            |
| `linux.image.tag`                       | Linux image tag                                                                                                                   | `v0.0.21`                                           |
| `linux.driver.resources`                | The resource request/limits for the linux secrets-store container image                                                           | `limits: 200m CPU, 200Mi; requests: 50m CPU, 100Mi` |
| `linux.enabled`                         | Install secrets store csi driver on linux nodes                                                                                   | true                                                |
| `linux.kubeletRootDir`                  | Configure the kubelet root dir                                                                                                    | `/var/lib/kubelet`                                  |
| `linux.providersDir`                    | Configure the providers root dir                                                                                                  | `/etc/kubernetes/secrets-store-csi-providers`       |
| `linux.nodeSelector`                    | Node Selector for the daemonset on linux nodes                                                                                    | `{}`                                                |
| `linux.tolerations`                     | Tolerations for the daemonset on linux nodes                                                                                      | `[]`                                                |
| `linux.metricsAddr`                     | The address the metric endpoint binds to                                                                                          | `:8095`                                             |
| `linux.registrarImage.repository`       | Linux node-driver-registrar image repository                                                                                      | `k8s.gcr.io/sig-storage/csi-node-driver-registrar`  |
| `linux.registrarImage.pullPolicy`       | Linux node-driver-registrar image pull policy                                                                                     | `Always`                                            |
| `linux.registrarImage.tag`              | Linux node-driver-registrar image tag                                                                                             | `v2.1.0`                                            |
| `linux.registrar.resources`             | The resource request/limits for the linux node-driver-registrar container image                                                   | `limits: 100m CPU, 100Mi; requests: 10m CPU, 20Mi`  |
| `linux.registrar.logVerbosity`          | Log level for node-driver-registrar. Uses V logs (klog)                                                                           | `5`                                                 |
| `linux.livenessProbeImage.repository`   | Linux liveness-probe image repository                                                                                             | `k8s.gcr.io/sig-storage/livenessprobe`              |
| `linux.livenessProbeImage.pullPolicy`   | Linux liveness-probe image pull policy                                                                                            | `Always`                                            |
| `linux.livenessProbeImage.tag`          | Linux liveness-probe image tag                                                                                                    | `v2.2.0`                                            |
| `linux.livenessProbe.resources`         | The resource request/limits for the linux liveness-probe container image                                                          | `limits: 100m CPU, 100Mi; requests: 10m CPU, 20Mi`  |
| `linux.env`                             | Environment variables to be passed for the daemonset on linux nodes                                                               | `[]`                                                |
| `linux.priorityClassName`               | Indicates the importance of a Pod relative to other Pods.                                                                         | `""`                                                |
| `linux.daemonsetAnnotations`            | Linux *DaemonSet* annotations                                                                                                     | `{}`                                                |
| `linux.podAnnotations`                  | Linux *Pod* annotations                                                                                                           | `{}`                                                |
| `linux.podLabels`                       | Linux *Pod* labels                                                                                                                | `{}`                                                |
| `linux.updateStrategy`                  | Configure a custom update strategy for the daemonset on linux nodes                                                               | `RollingUpdate with 1 maxUnavailable`               |
| `windows.image.repository`              | Windows image repository                                                                                                          | `k8s.gcr.io/csi-secrets-store/driver`               |
| `windows.image.pullPolicy`              | Windows image pull policy                                                                                                         | `IfNotPresent`                                      |
| `windows.image.tag`                     | Windows image tag                                                                                                                 | `v0.0.21`                                           |
| `windows.driver.resources`              | The resource request/limits for the windows secrets-store container image                                                         | `limits: 400m CPU, 400Mi; requests: 50m CPU, 100Mi` |
| `windows.enabled`                       | Install secrets store csi driver on windows nodes                                                                                 | false                                               |
| `windows.kubeletRootDir`                | Configure the kubelet root dir                                                                                                    | `C:\var\lib\kubelet`                                |
| `windows.providersDir`                  | Configure the providers root dir                                                                                                  | `C:\k\secrets-store-csi-providers`                  |
| `windows.nodeSelector`                  | Node Selector for the daemonset on windows nodes                                                                                  | `{}`                                                |
| `windows.tolerations`                   | Tolerations for the daemonset on windows nodes                                                                                    | `[]`                                                |
| `windows.metricsAddr`                   | The address the metric endpoint binds to                                                                                          | `:8095`                                             |
| `windows.registrarImage.repository`     | Windows node-driver-registrar image repository                                                                                    | `k8s.gcr.io/sig-storage/csi-node-driver-registrar`  |
| `windows.registrarImage.pullPolicy`     | Windows node-driver-registrar image pull policy                                                                                   | `Always`                                            |
| `windows.registrarImage.tag`            | Windows node-driver-registrar image tag                                                                                           | `v2.1.0`                                            |
| `windows.registrar.resources`           | The resource request/limits for the windows node-driver-registrar container image                                                 | `limits: 200m CPU, 200Mi; requests: 10m CPU, 20Mi`  |
| `windows.registrar.logVerbosity`        | Log level for node-driver-registrar. Uses V logs (klog)                                                                           | `5`                                                 |
| `windows.livenessProbeImage.repository` | Windows liveness-probe image repository                                                                                           | `k8s.gcr.io/sig-storage/livenessprobe`              |
| `windows.livenessProbeImage.pullPolicy` | Windows liveness-probe image pull policy                                                                                          | `Always`                                            |
| `windows.livenessProbeImage.tag`        | Windows liveness-probe image tag                                                                                                  | `v2.2.0`                                            |
| `windows.livenessProbe.resources`       | The resource request/limits for the windows liveness-probe container image                                                        | `limits: 200m CPU, 200Mi; requests: 10m CPU, 20Mi`  |
| `windows.env`                           | Environment variables to be passed for the daemonset on windows nodes                                                             | `[]`                                                |
| `windows.priorityClassName`             | Indicates the importance of a Pod relative to other Pods.                                                                         | `""`                                                |
| `windows.daemonsetAnnotations`          | Windows *DaemonSet* annotations                                                                                                   | `{}`                                                |
| `windows.podAnnotations`                | Windows *Pod* annotations                                                                                                         | `{}`                                                |
| `windows.podLabels`                     | Windows *Pod* labels                                                                                                              | `{}`                                                |
| `windows.updateStrategy`                | Configure a custom update strategy for the daemonset on windows nodes                                                             | `RollingUpdate with 1 maxUnavailable`               |
| `logVerbosity`                          | Log level. Uses V logs (klog)                                                                                                     | `0`                                                 |
| `logFormatJSON`                         | Use JSON logging format                                                                                                           | `false`                                             |
| `livenessProbe.port`                    | Liveness probe port                                                                                                               | `9808`                                              |
| `livenessProbe.logLevel`                | Liveness probe container logging verbosity level                                                                                  | `2`                                                 |
| `rbac.install`                          | Install default rbac roles and bindings                                                                                           | true                                                |
| `syncSecret.enabled`                    | Enable rbac roles and bindings required for syncing to Kubernetes native secrets (the default will change to false after v0.0.14) | true                                                |
| `minimumProviderVersions`               | [**DEPRECATED**] A comma delimited list of key-value pairs of minimum provider versions with driver                               | `""`                                                |
| `enableSecretRotation`                  | Enable secret rotation feature [alpha]                                                                                            | `false`                                             |
| `rotationPollInterval`                  | Secret rotation poll interval duration                                                                                            | `"120s"`                                            |
| `filteredWatchSecret`                   | Enable filtered watch for NodePublishSecretRef secrets with label `secrets-store.csi.k8s.io/used=true`                            | `false`                                             |
