package services

import (
	"github.com/Kamva/mgm/v3"
	"github.com/ponsonio/avanticaChallenge/pkg/api/models"
	"log"
)

type PathServiceInterface interface {
	Create(*models.Path) (*models.Path, error)
	Get(id string) (*models.Path, error)
	Delete(*models.Path) (err error)
	Update(*models.Path) (*models.Path, error)
}

type PathService struct {
}

func (p *PathService) Create(path *models.Path) (*models.Path, error) {
	err := mgm.Coll(path).Create(path)
	if err != nil {
		log.Printf("error saving path %v , error %v\n", path, err)
		return nil, err
	}
	return path, nil
}

func (p *PathService) Get(id string) (*models.Path, error) {
	path := &models.Path{}
	err := mgm.Coll(path).FindByID(id, path)
	if err != nil {
		log.Printf("error getting path %v , error %v\n", path, err)
		return nil, err
	}
	return path, nil
}

func (p *PathService) Delete(path *models.Path) (err error) {
	err = mgm.Coll(path).Delete(path)
	if err != nil {
		log.Printf("error deleting path %v , error %v\n", path, err)
	}
	return
}

func (p *PathService) Update(path *models.Path) (*models.Path, error) {
	err := mgm.Coll(path).Update(path)
	if err != nil {
		log.Printf("error updating path %v , error %v\n", path, err)
		return nil, err
	}
	return path, nil
}
