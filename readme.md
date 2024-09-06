install pprof `go install github.com/google/pprof@latest`

`go run main.go` -> creating file profile.cpu

`pprof profile.cpu` -> for run pprof analyzing