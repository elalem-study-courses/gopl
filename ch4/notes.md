# Chapter 4: Composite Types

| Programs | Description|
| --- | --- |
| append | append with table doubling technique |
| currency | Array with constants as indices |
| nonempty | similar to compact in ruby |
| rev | reversing |
| sha256 | array |
| remove| remove an element from a slice |
| graph | complex map |
| charcount | counting chars with map |
| dedup | removing duplicate lines |
| embed | embedding struct |
| movie | json |


* The array length in the next snippet is determined by the initializer size
  ```golang
  q := [...]int{1, 2, 3}
  ```
* `[3]int` and `[4]int` are different types
  ```golang
  q := [3]int{1, 2, 3}
  q = [4]int{1, 2, 3, 4} // compile error cannot assign [4]int to [3]int
  ```
* ```golang
  r := [...]int{99: -1}
  ```
  defines an array r with 100 elements, all zero except the last, which has value -1.
* If the array element type is comparable then the array type is comparable also.
  For example, the `==` operator reports whether all corrosponding elements are equal.
* Array types are passed by values.
* Slicing is done by the operator [i:j].
* Slicing has three components (pointer, length, capacity).
* If slicing is done beyond the length then the slice will extend.
* If the slicing is done beyond the capacity it panics.
* The nil value of the slice can be written using a conversion expression `[]int(nil)`.
* Struct is comparable if all of its fields are comparable.
* Comparable struct can be used as a key for map.
* Field-tags for structs are usually interpreted as key:value pairs.
* Unmarshaling JSON is case-insensitive.
* Unmarshal vs Decode
  * Use unmarshal if you already have the json object loaded in memory.
  * Use decode if you expect a stream (a reader).