package web

import (
	"bytes"
	"context"
	"github.com/hsedjame/gowebapi/framework/core"
	"github.com/hsedjame/gowebapi/framework/core/jsonUtils"
	"github.com/hsedjame/gowebapi/framework/security"
	"log"
	"net/http"
	"runtime"
	"strconv"
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
			id := GetGID()
			logger.Printf(" Goroutine #%v ---------------------------------\n", id)
			logger.Printf(" ----> [ %s %s ]", rq.Method, rq.URL.Path)
			next.ServeHTTP(wr, rq)
			logger.Println("_________________________________")
		})
	}
}

func SecurityMiddleware(config *security.Configuration, ctx context.Context, errorHandler ErrorHandler) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {

			if securityContext, b := config.Apply(rq); b {
				ctx = context.WithValue(ctx, core.SecurityCtxKey, &securityContext)

				rq := rq.WithContext(ctx)

				next.ServeHTTP(wr, rq)

			} else {
				wr.WriteHeader(http.StatusUnauthorized)
				_ = errorHandler(core.AppError{Message: " Unauthorized"}, wr)
				return
			}

		})
	}
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
