# Go Strings
This is a project used to enhance string in golang, reduce duplicate codes for handling string in different scenario.

# Usage
```
go get github.com/x-punch/go-strings
```
```go
package main

import (
    "fmt"

    "github.com/x-punch/go-strings"
)

func main() {
    fmt.Println(strings.GetRandomString(10))
    fmt.Println(strings.GetRandomNumber(4))
    fmt.Println(strings.SubString("abcdefg", 2, 10))
    fmt.Println(strings.ContainString("abc,def,g", ",", "def"))
}
```