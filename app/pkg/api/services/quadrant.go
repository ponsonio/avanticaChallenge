package services

import (
	"fmt"
	"github.com/Kamva/mgm/v3"
	"github.com/ponsonio/avanticaChallenge/pkg/api/models"
	"log"
)

type QuadrantInterface interface {
	Create(*models.Quadrant) (*models.Quadrant, error)
	Get(id string) (*models.Quadrant, error)
	Delete(*models.Quadrant) (err error)
	Update(*models.Quadrant) (*models.Quadrant, error)
}

type QuadrantService struct {
}

func (q *QuadrantService) Create(quadrant *models.Quadrant) (*models.Quadrant, error) {
	err := isValidQuadrant(quadrant)
	if err != nil {
		log.Printf("error updating quadrant %v , error %v\n", quadrant, err)
		return nil, err
	}

	err = mgm.Coll(quadrant).Create(quadrant)
	if err != nil {
		log.Printf("error saving quadrant %v , error %v\n", quadrant, err)
		return nil, err
	}
	return quadrant, nil
}

func (q *QuadrantService) Get(id string) (*models.Quadrant, error) {
	quadrant := &models.Quadrant{}
	err := mgm.Coll(quadrant).FindByID(id, quadrant)
	if err != nil {
		log.Printf("error getting quadrant %v , error %v\n", q, err)
		return nil, err
	}
	return quadrant, nil
}

func (q *QuadrantService) Delete(quadrant *models.Quadrant) (err error) {
	err = mgm.Coll(quadrant).Delete(quadrant)
	if err != nil {
		log.Printf("error deleting quadrant %v , error %v\n", quadrant, err)
	}
	return
}

func (q *QuadrantService) Update(quadrant *models.Quadrant) (*models.Quadrant, error) {

	err := isValidQuadrant(quadrant)
	if err != nil {
		log.Printf("error updating quadrant %v , error %v\n", quadrant, err)
		return nil, err
	}

	err = mgm.Coll(quadrant).Update(quadrant)
	if err != nil {
		log.Printf("error updating quadrant %v , error %v\n", quadrant, err)
		return nil, err
	}
	return quadrant, nil
}

//this can be done with go-validations, but it's implemented as it to show an easy test
func isValidQuadrant(quadrant *models.Quadrant) error {
	if quadrant.Number > 4 || quadrant.Number < 1 {
		return fmt.Errorf("Not a valid quadrant number %v", quadrant.Number)
	}
	return nil
}
