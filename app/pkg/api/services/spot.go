package services

import (
	"github.com/Kamva/mgm/v3"
	"github.com/ponsonio/avanticaChallenge/pkg/api/models"
	"log"
)

type SpotServiceInterface interface {
	Create(spot *models.Spot) (*models.Spot, error)
	Get(id string) (*models.Spot, error)
	Delete(*models.Spot) (err error)
	Update(*models.Spot) (*models.Spot, error)
}

type SpotService struct {
}

func (sp *SpotService) Create(spot *models.Spot) (s *models.Spot, err error) {
	err = mgm.Coll(spot).Create(spot)
	if err != nil {
		log.Printf("error saving spot %v , error %v\n", spot, err)
		return nil, err
	}
	return spot, nil
}

func (sp *SpotService) Get(id string) (*models.Spot, error) {
	spot := &models.Spot{}
	err := mgm.Coll(spot).FindByID(id, spot)
	if err != nil {
		log.Printf("error getting spot %v , error %v\n", spot, err)
		return nil, err
	}
	return spot, nil
}

func (sp *SpotService) Delete(spot *models.Spot) (err error) {
	err = mgm.Coll(spot).Delete(spot)
	if err != nil {
		log.Printf("error deleting spot %v , error %v\n", spot, err)
	}
	return
}

func (sp *SpotService) Update(spot *models.Spot) (s *models.Spot, err error) {
	err = mgm.Coll(spot).Update(spot)
	if err != nil {
		log.Printf("error updating spot %v , error %v\n", spot, err)
		return nil, err
	}
	return spot, nil
}
