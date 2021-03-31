package main

import (
	"github.com/Juuhabach/esquenta-go-zero-to-multithreading/http"
	"github.com/Juuhabach/esquenta-go-zero-to-multithreading/model"
	uuid "github.com/satori/go.uuid"
)

func main() {
	produto1 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Carrinho",
	}

	produto2 := model.Product{
		ID:   uuid.NewV4().String(),
		Name: "Boneca",
	}

	products := model.Products{}
	products.Add(produto1)
	products.Add(produto2)

	server := http.NewWebServer()
	server.Products = &products
	server.Serve()

}
