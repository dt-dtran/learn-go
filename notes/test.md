# Testing in Go

Create go.mod file: go mod init projectDirectoryName

Test naming convention include: \_test.go

Run test: go test

````go
func TestFeature(t *testing.T) {
    ...
    t.Errorf("Expected x, but got %v", value)
}
```s
````
