package http

import (
	"github.com/Juuhabach/esquenta-go-zero-to-multithreading/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type WebServer struct {
	Products *model.Products
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.Products)
}

func (w WebServer) Serve() {
	e := echo.New()
	e.GET("/product", w.getAll)
	e.POST("/product", w.createProduct)
	e.Logger.Fatal(e.Start(":8585"))
}

func (w WebServer) createProduct(c echo.Context) error {
	product := model.NewProduct()

	if err := c.Bind(product); err != nil {
		return err
	}

	w.Products.Add(*product)
	return c.JSON(http.StatusCreated, product)

}
