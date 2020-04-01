#单元测试
```go
go test -v
go test -v -run=byAdd
```


#基准测试
```go
go test -run=none -bench=BenchmarkByAdd -count=10
go test -run=none -bench=BenchmarkBySprintf -count=10
```