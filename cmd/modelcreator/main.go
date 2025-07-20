// cmd/modelcreator/main.go
package main

import (
	"flag"
	"log"
	"os"

	"modelcreator/internal/modelcreator"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := modelcreator.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
