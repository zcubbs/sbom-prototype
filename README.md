# SBOM Prototype

## Setup

```
task bootstrap
```

## Generate migrate files

```
cd <MODULE>
migrate create -ext sql -dir storage/migrations -seq <create_XYZ_table>
```

## Setup proto linter 

```
go install github.com/googleapis/api-linter/cmd/api-linter@latest
```


