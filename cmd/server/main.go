package main

import (
	"encontradev/internal"
	"log"
)

func main() {
	r := internal.SetupRouter()

	log.Println("Servidor rodando em http://localhost:8080")
	r.Run(":8080")
}
