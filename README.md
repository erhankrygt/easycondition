# easycondition

easycondition allows to reach the target in a single line without searching for conditions in lists.

# Example
```go
package main

import (
	"fmt"
	auto "github.com/erhankrygt/easycondition"
)

type Foo struct {
	Name string
	Age  int
}

// Main function
func main() {
  l := []Foo{
		{
			Name: "Monkey",
			Age:  1,
		},
		{
			Name: "Dog",
			Age:  2,
		},
	}

	var item, err = easycondition.FirstDefault(l, "Age=2")
	if err == nil {
		fmt.Println(item.(Foo))
	}
  
  //output
  {Dog 2}

}
```
# Installation
```
go get -u github.com/erhankrygt/easycondition
```
