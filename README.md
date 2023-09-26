# JSTypes
JavaScript types (string, array, set, map) in Go

## Example
```js
  console.log("HI".toLowerCase()) // "hi"
  console.log("HI".includes("h")) // true
  console.log(["hi"].includes("hi")) // true
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
  fmt.Println(point(jstypes.String("HI")).Includes("h")) // true
  fmt.Println(point(jstypes.Array[string]([]string{"hi"})).Includes("hi")) // true
}
```
