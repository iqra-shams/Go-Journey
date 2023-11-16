// main.go
package main

import (
	"log"
	"github.com/iqra-shams/cobra-example/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
