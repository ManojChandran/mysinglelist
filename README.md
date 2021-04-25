# Single Linked List
This simple go package will implement help to implement single linked list. 

# Get package 

# Main functions

# Sample Implemenatation code

```go
package main

import (
	"fmt"

	llist "github.com/ManojChandran/mysinglelist"
)

func main() {
	mylist := llist.Getlist()
	mylist.Add("Manoj")
	mylist.ListAll()
	fmt.Println(mylist.Delete("Chandran"))
	mylist.Add("Archana")
	fmt.Println(mylist.Insert("Chandran", "Manoj"))
	mylist.ListAll()
	fmt.Println(mylist.Delete("Archana"))
	mylist.ListAll()
}

```