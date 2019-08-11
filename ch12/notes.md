# Chapter 12: Reflection

* `reflect.TypeOf` returns the dynamic type of an interface for example
  ```go
  var w io.Writer = os.Stdout
  fmt.Println(reflect.TypeOf(w)) // *os.File
  ```
* important methods for reflections
  ```go
  t := reflect.TypeOf(3) // returns reflect.Type that holds type "int"
  v := reflect.ValueOf(3) // returns reflect.Value that holds the value of 3 with type of "int"
  t = v.Type() // returns the reflect.Type of v same as reflect.TypeOf(3)
  x := v.Interface() // returns an interface type of v now we can use type conversion
  i := x.(int) // returns an integer of x
  ```
* `reflect.ValueOf()` always returns the concrete type.
