# Chapter 11: Testing

* `go test -v -run="French|Canal"` runs tests that include French or Canal words (regular expression).
* To test a command (a go application with package main), we have to dynamically change the output pipeline from Stdout to whatever we will use (normally buffer).
* `go list -f '{{.GoFiles}}' fmt` Files that are built.
* `go list -f '{{.TestGoFiles}}' fmt` Files that are built for test only.
* `go list -f '{{.XTestGoFiles}}' fmt` Files that are external tests.
* `go test -v -run=Coverage gopl.io/ch7/eval`
* `go test -v -run=Coverage -coverprofile=c.out gopl.io/ch7/eval`
* `go test -v -run=Coverage -coverprofile=c.out -covermode=count gopl.io/ch7/eval` to identify hot zones and cold zones.
* For some reason the above commands does not work as they are reporting no tests where found we can use `go test -cover` instead of `go test -run=Coverage`.
* `go tool cover -html=c.out` After generating a coverage report if we want to display the coverage report as html.
* `go test -bench=.` -bench accepts a regular expression that matches the benchmark functions (. stands for all benchmark functions).
* `go test -cpuprofile=cpu.out`
* `go test -memprofile=mem.out`
* `go test -blockprofile=block.out`
* Long running applications can use `runtime` package to profile the application.
* When profiling is enabled the build does not remove the executable files but saves them under `*.test` for profiling.
* To disable benchmark test cases we add `-run=NONE`
* `-cpuprofile` flag generates a log file which pprof uses with the executable *.test to generate the profiling
* `go tool pprof -text -nodecount=10 ./word3.test cpu.log` for textual profile.
* `go tool pprof -web -nodecount=10 ./word3.test cpu.log` for graphical profile.
