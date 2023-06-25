# SBOM Prototype

## Setup

> prerequisite: docker, go, taskfile

```
task bootstrap
task pb
task tls
```

## Generate migrate files

```
cd <MODULE>
migrate create -ext sql -dir storage/migrations -seq <create_XYZ_table>
```


