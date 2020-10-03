# learn-go-with-tests
learn-go-with-tests

## basics Test cmds

```
go run xxx.go
go test -v
go test -bench=.
go test -cover
go test -bench=.
```
### Concurrency
fatal error: concurrent map writes

`go test -race`

## settings
go get -u github.com/kisielk/errcheck
errcheck .

## Resources
https://github.com/quii/learn-go-with-tests
