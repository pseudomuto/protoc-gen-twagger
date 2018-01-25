package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Could not read contents from stdin")
	}

	ioutil.WriteFile("codegen.req", data, 0666)
}
