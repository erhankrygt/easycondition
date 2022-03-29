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

   var item, err = easycondition.First(l, "Name=Dog")
   if err == nil {
	fmt.Println(item.(Foo))
   }
  
   //output
   {Dog 2}


   var biggest, err = easycondition.First(l, "Age>1")
   if err == nil {
	fmt.Println(biggest.(Foo))
   }
   //output
   {Dog 2}
   
   
   var smallest, err = easycondition.First(l, "Age<2")
   if err == nil {
	fmt.Println(smallest.(Foo))
   }
   //output
   {Monkey 1}
   
}
```
# Installation
```
go get -u github.com/erhankrygt/easycondition
```
