package repository

import (
	"backend/app/models"
	"errors"
)

type BeerRepository interface {
	GetBeers() []*models.Beer
	GetBeerById(id string) (*models.Beer, error)
	SaveBeer(beer *models.Beer) error
}

type DBBeerRepository struct {
	beers []*models.Beer
}

func New() *DBBeerRepository {
	return &DBBeerRepository{
		beers: []*models.Beer{
			&models.Beer{"1", "La Trappe Quadrupel Oak Aged", "Ale", "Bierbrouwerij De Koningshoeven"},
			&models.Beer{"2", "Cocoa Psycho", "Porter", "BrewDog"},
			&models.Beer{"3", "American Dream", "Lager", "Mikkeller"},
		},
	}
}

func (r *DBBeerRepository) GetBeers() []*models.Beer {
	return r.beers
}

func (r *DBBeerRepository) GetBeerById(id string) (*models.Beer, error) {
	for _, beer := range r.beers {
		if beer.ID == id {
			return beer, nil
		}
	}
	return nil, errors.New("beer not found")
}

func (r *DBBeerRepository) SaveBeer(beer *models.Beer) error {
	r.beers = append(r.beers, beer)
	return nil
}

var beerRepository *DBBeerRepository

func GetBeerRepository() (r BeerRepository) {
	if beerRepository == nil {
		beerRepository = New()
	}
	return beerRepository
}
