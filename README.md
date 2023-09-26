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

func main() {
  fmt.Println(jstypes.String("HI").ToLowerCase()) // "hi"
  fmt.Println(jstypes.String("HI").Includes("h")) // true
  fmt.Println(jstypes.Array[string]([]string{"hi"}).Includes("hi")) // true
}
```
