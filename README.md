
```bash
go install github.com/golang/mock/mockgen
$GOPATH/bin/mockgen --help
```

```bash
$GOPATH/bin/mockgen --source=router/application.go --destination=router/application_mock.go --package=router
$GOPATH/bin/mockgen --source=router/repository.go --destination=router/repository_mock.go --package=router
```

# Run test

```bash
go test ./...
```
