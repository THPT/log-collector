# log-collector

Requirement
--- 
 - Install Golang at https://golang.org/
 - Kafka

Your source code MUST be in folder `$GOPATH/src/log-collector`

Config
---
Copy file `config.go.example` to `config.go` and update it

Build & run
---

- `cd $GOPATH/src/log-collector`
- `go get` to pull all needed libraries
- Just use command `go build` and run binary file `./log-collector`
- Or using command `go run main.go`
