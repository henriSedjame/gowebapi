package main

import (
	"github.com/go-pg/pg/v10"
	"github.com/hsedjame/gowebapi/framework/core"
	"github.com/hsedjame/gowebapi/framework/std"
	"github.com/hsedjame/gowebapi/src/main/go/controllers"
	"github.com/hsedjame/gowebapi/src/main/go/controllers/contrat"
	"github.com/hsedjame/gowebapi/src/main/go/models"
	"github.com/hsedjame/gowebapi/src/main/go/repositories"
	"log"
	"os"
)

// An api to expose a Crud operations
func main() {

	logger := log.New(os.Stdout, "[CONTRAT-API] ", log.LstdFlags)

	app, err := std.New(logger)

	if err != nil {
		return
	}

	app.
		WithOptions(std.Entities([]interface{}{&models.Contrat{}})).
		WithOptions(std.Rest([]interface{}{

			&controllers.IndexController{},

			&contrat.Controller{
				Handler: &contrat.Handler{
					Repository: &repositories.ContratRepository{
						DB: app.Context().Value(core.DBCtxKey).(*pg.DB),
					},
				},
			},
		})).
		Run()
}
