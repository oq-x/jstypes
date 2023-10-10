# JSTypes

JavaScript types (string, array, set, map) in Go

## Example

```js
console.log("HI".toLowerCase()); // "hi"
console.log("HI".includes("H")); // true
console.log(["hi"].includes("hi")); // true
```

```go
package main
import (
  "fmt"

  "github.com/oq-x/jstypes"
)

func point[T any](t T) *T {
  return &t
}

func main() {
  fmt.Println(point(jstypes.String("HI")).ToLowerCase()) // "hi"
  fmt.Println(point(jstypes.String("HI")).Includes("H")) // true
  fmt.Println(point(jstypes.Array[string]([]string{"hi"})).Includes("hi")) // true
}
```
