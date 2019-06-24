# Chapter 7: Interfaces

* To implement an interface do not use a pointer receiver just use a concrete one
  ```go
    type SpecialWriter struct {

    }

    func (s SpecialWriter) Write(p []byte) (int, error) {
      // Implementation details
    }
  ```
  If we tried to implement an interface using a pointer receiver the compiler will give us this error `does not implement io.Writer (Write method has pointer receiver)`.
* The sort package provide `Strings` method for `StringSlice` type.
* text/tabwriter package produces a table whose columns are neatly aligned.
* `http.Error()` function can be used in http errors.
* When receiving data from client in the http handler func we can use `r.ParseForm()` followed by `r.Form.Get("something")`