package repositories

import (
	"github.com/go-pg/pg/v10"
	"github.com/hsedjame/gowebapi/src/main/go/models"
	"net/http"
)

type ContratRepository struct {
	DB *pg.DB
}

func NewContratRepository(db *pg.DB) *ContratRepository {
	return &ContratRepository{
		DB: db,
	}
}

func (repo ContratRepository) FindById(id string) (models.Contrat, *models.ApiError) {
	contrat := models.Contrat{NumeroContrat: id}
	if err := repo.DB.Model(&contrat).WherePK().Select(); err != nil {
		return contrat, models.FromError(err, http.StatusInternalServerError)
	}
	return contrat, nil
}

func (repo ContratRepository) FindAll() ([]models.Contrat, *models.ApiError) {
	var contrats []models.Contrat
	if err := repo.DB.Model(&contrats).Select(); err != nil {
		return nil, models.FromError(err, http.StatusInternalServerError)
	}
	return contrats, nil
}

func (repo ContratRepository) Create(contrat *models.Contrat) *models.ApiError {
	if _, err := repo.DB.Model(contrat).Insert(); err != nil {
		return models.FromError(err, http.StatusInternalServerError)
	}
	return nil
}

func (repo ContratRepository) Update(contrat *models.Contrat) *models.ApiError {
	if _, err := repo.DB.Model(contrat).WherePK().Update(); err != nil {
		return models.FromError(err, http.StatusInternalServerError)
	}
	return nil
}

func (repo ContratRepository) Delete(id string) *models.ApiError {
	contrat := models.Contrat{NumeroContrat: id}
	if _, err := repo.DB.Model(&contrat).WherePK().Delete(); err != nil {
		return models.FromError(err, http.StatusInternalServerError)
	}
	return nil
}
