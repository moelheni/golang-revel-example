package controllers

import (
	"backend/app/models"
	"backend/app/repository"

	"github.com/revel/revel"
)

type BeersController struct {
	*revel.Controller
}

func (c BeersController) List() revel.Result {
	beers := repository.GetBeerRepository().GetBeers()

	response := JsonResponse{}
	response.Success = true
	response.Data = beers

	return c.RenderJSON(response)
}

func (c BeersController) Show(beerID string) revel.Result {
	beer, err := repository.GetBeerRepository().GetBeerById(beerID)

	response := JsonResponse{}
	response.Success = err == nil
	response.Data = beer
	if err != nil {
		response.Error = err.Error()
	}

	return c.RenderJSON(response)
}

func (c BeersController) Save() revel.Result {
	beer := models.Beer{}
	c.Params.BindJSON(&beer)

	err := repository.GetBeerRepository().SaveBeer(&beer)

	response := JsonResponse{}
	response.Success = err == nil
	response.Data = beer
	if err != nil {
		response.Error = err.Error()
	}
	return c.RenderJSON(response)
}
