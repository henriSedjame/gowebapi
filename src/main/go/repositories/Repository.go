package repositories

import "github.com/hsedjame/gowebapi/src/main/go/models"

type Repository interface {
	FindById(interface{}) (interface{}, *models.ApiError)
	FindAll()([]interface{}, *models.ApiError)
	Create(interface{}) *models.ApiError
	Update(interface{}) *models.ApiError
	Delete(interface{}) *models.ApiError
}
