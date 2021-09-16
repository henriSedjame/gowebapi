package web

import (
	"context"
	"github.com/hsedjame/gowebapi/framework/core/jsonUtils"
	"log"
	"net/http"
)


func PostPutMethodHandler(defaultModel interface{}, modelKey interface{}, errorHandler ErrorHandler) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {
			wr.Header().Add("Content-Type", "application/json")

			/*
			 * Dans le cas d'une methode POST ou PUT
			 * Valider la requête
			 */
			if rq.Method == http.MethodPost || rq.Method == http.MethodPut {

				if err := jsonUtils.FromJson(&defaultModel, rq.Body); err != nil {
					wr.WriteHeader(http.StatusBadRequest)
					_ = errorHandler(err, wr)
					return
				}
				/*else if err := validations.IsValid(defaultModel); err != nil {
					wr.WriteHeader(http.StatusBadRequest)
					_ = errorHandler(err, wr)
					return
				}*/

				ctx := context.WithValue(rq.Context(), modelKey, defaultModel)

				rq := rq.WithContext(ctx)

				next.ServeHTTP(wr, rq)

				return
			}

			next.ServeHTTP(wr, rq)
		})
	}
}

func LoggingMiddleware(logger *log.Logger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {
			logger.Println("---------------------------------")
			logger.Printf(" ----> [ %s %s ]", rq.Method, rq.URL.Path)
			next.ServeHTTP(wr, rq)
			logger.Println("_________________________________")
		})
	}
}