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

func main() {
  fmt.Println(jstypes.StringFrom("HI").ToLowerCase()) // "hi"
  fmt.Println(jstypes.StringFrom("HI").Includes("H")) // true
  fmt.Println(jstypes.ArrayFrom([]string{"hi"}).Includes("hi")) // true
}
```
