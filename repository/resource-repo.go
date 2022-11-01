package repository

import "github.com/AugustoEMoreira/viracocha-cspm/entity"

type ResourceRepository interface {
	Create(*entity.Resource) error
	Update(*entity.Resource) (entity.Resource, error)
	Delete(ID string) error
	GetResources(queryStrg string) ([]entity.Resource, error)
}
