package contrat

import (
	"github.com/gorilla/mux"
	"github.com/hsedjame/gowebapi/framework/core/jsonUtils"
	"github.com/hsedjame/gowebapi/src/main/go/models"
	"github.com/hsedjame/gowebapi/src/main/go/repositories"
	"net/http"
)

type Handler struct {
	Repository *repositories.ContratRepository
}

func NewContratHandler(repository repositories.ContratRepository) *Handler {
	return &Handler{
		&repository,
	}
}

func (handler *Handler) FindAll(wr http.ResponseWriter, _ *http.Request)  {
	if contrats, err := handler.Repository.FindAll(); err != nil {
		handleError(err, wr, http.StatusInternalServerError)
		return
	} else if err := jsonUtils.ToJson(contrats, wr); err != nil {
		handleError(err, wr, http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) FindById(wr http.ResponseWriter, rq *http.Request)  {
	if id := mux.Vars(rq)["id"]; id != "" {
		if contrat , err := handler.Repository.FindById(id); err != nil {
			handleError(err, wr, http.StatusInternalServerError)
			return
		} else if err := jsonUtils.ToJson(contrat, wr); err != nil{
			handleError(err, wr, http.StatusInternalServerError)
			return
		}
	}
}

func (handler *Handler) Create(wr http.ResponseWriter, rq *http.Request)  {
	 contrat := rq.Context().Value(ContratKey{}).(models.Contrat)

	 if err := handler.Repository.Create(&contrat); err != nil {
	 	handleError(err, wr, http.StatusInternalServerError)
	 	return
	 } else {
	 	wr.WriteHeader(http.StatusOK)
	 	return
	 }
}

func (handler *Handler) Update(wr http.ResponseWriter, rq *http.Request)  {
	contrat := rq.Context().Value(ContratKey{}).(models.Contrat)
	if err := handler.Repository.Update(&contrat); err != nil {
		handleError(err, wr, http.StatusInternalServerError)
		return
	} else {
		wr.WriteHeader(http.StatusOK)
		return
	}
}

func (handler *Handler) Delete(wr http.ResponseWriter, rq *http.Request)  {
	if id := mux.Vars(rq)["id"]; id != "" {
		if err := handler.Repository.Delete(id); err != nil {
			handleError(err, wr, http.StatusInternalServerError)
			return
		} else {
			wr.WriteHeader(http.StatusOK)
			return
		}
	}
}

func handleError(err error, wr http.ResponseWriter, status int) {
	switch t := err.(type) {
	case *models.ApiError:
		wr.WriteHeader(status)
		_ = jsonUtils.ToJson(t, wr )
	default:
		wr.WriteHeader(status)
		_ = jsonUtils.ToJson( models.ApiError{
			Message: t.Error(),
			Status: status,
		}, wr )
	}

}