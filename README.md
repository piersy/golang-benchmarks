# golang-benchmarks
Indulging my curiosity benchmarking parts of the go language.

# Install

`go get -d github.com/piersy/golang-benchmarks` or clone this repo to
`github.com/piersy/golang-benchmarks` in your GOPATH.

# Usage
`go test -bench .`

## Seeing the assembly behind the code
`go test -cpuprofile cpu.prof -bench=<test_name>`
then
`go tool pprof -weblist . golang-benchmarks.test cpu.prof`
