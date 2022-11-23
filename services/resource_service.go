package services

import (
	"github.com/AugustoEMoreira/viracocha-cspm/entity"
	"github.com/AugustoEMoreira/viracocha-cspm/lib"
	"github.com/AugustoEMoreira/viracocha-cspm/repository"
)

type ResourceService struct {
	logger     lib.Logger
	repository repository.ResourceRepository
}

func NewResourceService(logger lib.Logger, repository repository.ResourceRepository) ResourceService {
	return ResourceService{
		logger:     logger,
		repository: repository,
	}
}

func (s ResourceService) GetResourceById(id string) (resource entity.Resource, err error) {
	res, err := s.repository.FindById(id)
	return res, err
}

func (s ResourceService) WriteResource(resource entity.Resource) error {
	err := s.repository.WriteResource(resource)
	return err
}
