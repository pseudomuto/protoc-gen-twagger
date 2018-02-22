package main

import (
	"github.com/pseudomuto/protokit"

	"log"

	"github.com/pseudomuto/protoc-gen-twagger/internal"
)

func main() {
	if err := protokit.RunPlugin(new(internal.Plugin)); err != nil {
		log.Fatal(err)
	}
}
