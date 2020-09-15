# krew-museum
![Release Charts](https://github.com/schizoid90/krew-museum/workflows/Release%20Charts/badge.svg)
![Docker](https://github.com/schizoid90/krew-museum/workflows/Docker/badge.svg)

Go based private repository for packaged krew plugins

Self hosted private repository for krew plugins to be used in conjuction with a private index

If you don't have an open way to store and access archived packages.

Krew-museum stores files on the local filesystem.

## Installation

The application defaults to listening on `127.0.0.1:8090`, this can be changed by setting the `BIND_SERVER` and `BIND_PORT` environment variables.

```shell
export BIND_SERVER=0.0.0.0
export BIND_PORT=80
```

or in helm

```yaml
bindserver: 0.0.0.0
bindport: 80
```

**N.B.: If changing the port in a values.yaml file you will need to edit the readiness and livenes probes**

### local installation

Download the release binary and run

```shell
./krew-museum
```

### Docker

```shell
docker run schizoid90/krew-museum
```

### Kubernetes

```shell
helm install krew-museum krew-museum 
```

## Usage

The REST API is quite simple with just a few endpoints

| Endpoint  | Description                           |
|-----------|---------------------------------------| 
| /status   | View information about the repository |
| /upload   | Upload a package                      |
| /download | Download a pacakge                    |
| /packages | View package information              |

Upload a package

```shell
curl -X POST -F "file=@pacakge.1.0.0.zip" https://krew-museum/upload/package
```

Download a pacge in your krew manifest

```yaml
...
    # 'uri' specifies .zip or .tar.gz archive URL of a plugin
    uri: https://krew-museum/download/package/package.1.0.0.zip
    # 'sha256' is the sha256sum of the package above
    sha256: sha256sumofpackage
...
```

View pacakges

```shell
curl https:/krew-museum/packages
```

## Values

| Value                     | Description                                                   | Default                                    |
|---------------------------|---------------------------------------------------------------|--------------------------------------------|
| bindport                  | Port the server binds to                                      | `80`                                       |
| bindserver                | IP address the server listens on                              | `0.0.0.0`                                  |
| env                       | Set environment variables                                     | `[]`                                       |
| gracePeriod               | Time to wait for pod to end gracefully (seconds)              | `10`                                       |
| image.name                | URI of container image                                        | `ghcr.io/schizoid90/krew-museum:{version}` |
| image.pullPolicy          | PullPolicy for the image                                      | `IfNotPresent`                             |
| image.pullSecrets         | Name of secret containing container registry details          | `""`                                       |
| ingress.enabled           | Enable Kubernetes ingress for the application                 | `false`                                    |
| ingress.hostName          | Ingress hostname                                              | `""`                                       |
| ingress.path              | Path to route in the ingress                                  | `"/"`                                      |
| ingress.tls.enabled       | Enable TLS                                                    | `false`                                    |
| ingress.tls.hosts         | TLS hosts                                                     | `[]`                                       |
| ingress.tls.secretName    | Secret containing TLS cert and key                            | `""`                                       |
| livenessProbe             | Configure LivenessProbe                                       | see values.yaml                            |
| nameOverride              | Set the app and container name                                | `""`                                       |
| namespaceOverride         | Set the app namespace                                         | `""`                                       |
| persistence.accessMode    | Set the access mode on the storage                            | `"ReadWriteOnce"`                          |
| persistence.enabled       | Enable/Disable persistent storage (required for statefulsets) | `true`                                     |
| persistence.mounts        | Confgiure container mounts                                    | see values.yaml                            |
| persistence.size          | Configure size of storage                                     | `2G`                                       |
| persistence.storageClass  | Set the StorageClass                                          | `""`                                       |
| persistence.type          | Set the type of persistence ("pvc, statefulset")              | `"statefulset"`                            |
| persistence.volumes       | Configure storage volumes                                     | see values.yaml                            |
| readinessProbe            | Configure ReadinessProbs                                      | see values.yaml                            |
| replicas                  | Number of replicas to run                                     | `1`                                        |
| resources                 | Set resource limits/requests                                  | `[]`                                       |
| service.enabled           | Enable Kubernetes service for the application                 | `false`                                    |
| service.type              | Type of service ("LoadBalancer, NodePort, ClusterIP")         | `NodePort`                                 |
| service.nodePort          | Specify the NodePort (requires service.type: NodePort)        | `""`                                       |
| service.sourceRanges      | Specify source ranges (requires service.type: LoadBalancer)   | `{}`                                       |
| service.clusterIP         | Specify cluster IP (requires service.type: ClusterIP)         | `""`                                       |
| serivce.externalPort      | Set the port for the service to expose on                     | `8090`                                     |
| service.targetPort        | Set the target port of the pod                                | `8090`                                     |
| service.portProtocol      | Set port protcol                                              | `tcp`                                      |
| service.portName          | Set the name of the service port                              | `"krew-museum"`                            |
| tolerations               | Tolerate node taints                                          | `[]`                                       |