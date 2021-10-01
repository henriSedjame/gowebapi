package contrat

import (
	"context"
	"github.com/hsedjame/gowebapi/framework/core/jsonUtils"
	"github.com/hsedjame/gowebapi/framework/core/validations"
	"github.com/hsedjame/gowebapi/framework/web"
	"github.com/hsedjame/gowebapi/src/main/go/controllers/contrat/endpoints"
	"github.com/hsedjame/gowebapi/src/main/go/models"
	"net/http"
)

type Controller struct {
	Handler *Handler
}

func (c Controller) Endpoints() []web.Endpoint {
	return []web.Endpoint{
		endpoints.CreateEndpoint{RHandler: c.Handler.Create},
		endpoints.FindAllEndpoint{RHandler: c.Handler.FindAll},
		endpoints.FindByIdEndpoint{RHandler: c.Handler.FindById},
		endpoints.UpdateEndpoint{RHandler: c.Handler.Update},
		endpoints.DeleteEndpoint{RHandler: c.Handler.Delete},
	}
}

func (c Controller) ErrorHandler() web.ErrorHandler {
	return func(err error, writer http.ResponseWriter) error {
		return jsonUtils.ToJson(err, writer)
	}
}

func (c Controller) Path() string {
	return "/contrats"
}

func (c Controller) MiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {
		wr.Header().Add("Content-Type", "application/json")

		/*
		 * Dans le cas d'une methode POST ou PUT
		 * Valider la requête
		 */
		if rq.Method == http.MethodPost || rq.Method == http.MethodPut {

			var contrat models.Contrat
			if err := jsonUtils.FromJson(&contrat, rq.Body); err != nil {
				wr.WriteHeader(http.StatusBadRequest)
				_ = c.ErrorHandler()(err, wr)
				return
			} else if err := validations.IsValid(contrat); err != nil {
				wr.WriteHeader(http.StatusBadRequest)
				_ = c.ErrorHandler()(err, wr)
				return
			}

			ctx := context.WithValue(rq.Context(), ContratKey{}, contrat)

			rq := rq.WithContext(ctx)

			next.ServeHTTP(wr, rq)

			return
		}

		next.ServeHTTP(wr, rq)
	})
}
