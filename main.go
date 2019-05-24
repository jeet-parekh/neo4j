// +build oss

package main

import (
	"log"

	private "github.com/jeet-parekh/neo4j/private"
)

func main() {
	log.Println(private.Data)
	log.Println(private.GetData())
}
