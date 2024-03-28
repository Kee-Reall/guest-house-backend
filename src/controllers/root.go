package controllers

import (
	"encoding/json"
	"example.com/m/v2/src/models"
	"log"
	"net/http"
)

type RootController struct {
	manager *models.ResponseManager
}

func (c *RootController) HelloWorld(writer http.ResponseWriter, request *http.Request) {
	m := c.manager.NewResponse(make([]int, 5))
	if err := json.NewEncoder(writer).Encode(m); err != nil {
		log.Fatal(err)
	}
}

func (c *RootController) PostHandle(writer http.ResponseWriter, request *http.Request) {
	m := c.manager.NewResponse(make([]int, 5))
	if err := json.NewEncoder(writer).Encode(m); err != nil {
		log.Fatal(err)
	}
}

func CreateController() *RootController {
	return &RootController{manager: models.CreateResponseManager()}
}
