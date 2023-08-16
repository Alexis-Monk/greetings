# Greetings in Go

This package provides a simple way to obtain personalized greetings.

## Installation
Run the following command to install the package:
```bash
go get -u github.com/Alexis-Monk/greetings
```

## Use
Here's an example of how to use the package in your code:

```go
package main

import (
	"fmt"
	"github.com/Alexis-Monk/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // Para no ver fecha y hora

	names := []string{"Alex", "Lake", "Itzel"}
	messages, err := greetings.Hellos(names)

	//message, err := greetings.Hello("Lake")
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(messages, err)
}
```
This example imports the package github/Alexis-Monk/greetings and calls the "hello" function. If an error occurs, an error message is printed.
