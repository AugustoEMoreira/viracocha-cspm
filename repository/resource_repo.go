package repository

import (
	"errors"

	"github.com/AugustoEMoreira/viracocha-cspm/entity"
	"github.com/AugustoEMoreira/viracocha-cspm/lib"
)

type ResourceRepository struct {
	lib.Database
	logger lib.Logger
}

func NewResourceRepository(db lib.Database, logger lib.Logger) ResourceRepository {
	return ResourceRepository{
		Database: db,
		logger:   logger,
	}
}

func (r ResourceRepository) FindById(id string) (entity.Resource, error) {
	err := errors.New("to be implemented")
	return entity.Resource{}, err
}
func (r ResourceRepository) WriteResource(resource entity.Resource) error {
	err := errors.New("to be implemented")
	return err
}
