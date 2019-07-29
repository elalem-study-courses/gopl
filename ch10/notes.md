# Chapter 10 Packages and the go tool

* `go build <absolute path or go import path>` to build a library/command that is not in the current directory
* `go install` saves the compiled code for each package and command instead of throwing it away as in `go build` command. (Located in $GOPATH/pkg)
* `go build -i` installs the packages that are dependencies of the build target.
* Build tags are special comments in the source file that instructs the go tool to do certain actions while compiling.
  For example
  ```go
    // +build linux darwin
  ```

  ```go
    // +build ignore
  ```
* go doc examples are as follows
  ```bash
  go doc time                   # package documentation
  go doc time.Since             # Package member documentation
  go doc time.Duration.Seconds  # method documentation
  go doc json.Decode            # Full identifier documentation
  go doc json.decode            # Small case identifier documentation
  ```
* go doc is case insensitive 
* `godoc -http :8000 -analysis=type`
* `godoc -http :8000 -analysis=pointer`
* In case of internal packages: *net/http/internal/chunked* can be imported from *net/http/httputil* or *net/http* but not from *net/url*
* ```bash
  go list ... # all
  go list github.com/... # All under github.com
  go list ...xml... # All that contains word "xml"
  go list -json hash
  go list -f '{{join .Deps " "}}' strconv
  ```
