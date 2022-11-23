package main

import (
	"log"

	"go.uber.org/fx"
)

func main() {
	log.Println("running server")

	fx.New()
}
