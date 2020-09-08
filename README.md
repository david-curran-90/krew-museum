# krew-museum

Go based private repository for packaged krew plugins

Self hosted private repository for krew plugins to be used in conjuction with a private index

If you don't have an open way to store and access archived packages.

Krew-museum stores files on the local filesystem.

## Installation


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
curl -X POST -F "file=@pacakge.tar.gz" https://krew-museum/upload
```

Download a pacge in your krew manifest

```yaml
...
    # 'uri' specifies .zip or .tar.gz archive URL of a plugin
    uri: https://krew-museum/download/mypackage.1.0.0.zip
    # 'sha256' is the sha256sum of the package above
    sha256: sha256sumofpackage
...
```

View pacakges

```shell
curl https:/krew-museum/packages
```
