# Chapter 3 Basic Datetypes

* fmt trick
* runes
* surface program (rendering of an image using floating point computations)
* mandelbrot generator
* | Character | Description |
  | ---       | ---         |
  | \n | "alert" or bell |
  | \b | backspace |
  | \f | form feed |
  | \n | newline |
  | \r | carriage return |
  | \t | tab |
  | \v | vertical tab |
* \u for $16-bit$ runes. Where \U is for $32-bit" runes. (e.g. \u4e16 and \U00004e16)
* Important packages for manipulating strings
  * **strings** provide functions for (searching, replacing, comparing, trimming, splitting, joining strings)
  * **bytes** is similar to **strings**. But because strings are immutable, building up strings incrementally can invoke a lot of allocation and copying. In such cases it's more efficient to use the *bytes.Buffer* type.
  * **strconv** provides functions for converting boolean, integer, floating-point values to and from their string representations, and functions for quoting  and unquoting strings.
  * **unicode** package provides functions like *IsDigit*, *IsLetter*, *IsUpper*, *IsLower* for classifying rune. Each function takes a single rune argument and returns a boolean.