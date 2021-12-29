package std

import (
	"github.com/hsedjame/gowebapi/framework/core"
	"github.com/hsedjame/gowebapi/framework/security"
)

// AppOptions : Application options
type AppOptions func(app *App)

// Rest : Declare application restControllers
func Rest(controllers []interface{}) AppOptions {
	return func(app *App) {
		restControllers := app.Context().Value(core.ControllersCtxKey).([]interface{})
		if restControllers == nil {
			restControllers = []interface{}{}
		}
		restControllers = append(restControllers, controllers...)
		app.addToContext(core.ControllersCtxKey, restControllers)
	}
}

// Entities : Declare application entities
func Entities(datas []interface{}) AppOptions {
	return func(app *App) {
		entities := app.Context().Value(core.EntitiesCtxKey).([]interface{})
		if entities == nil {
			entities = []interface{}{}
		}
		entities = append(entities, datas...)

		app.addToContext(core.EntitiesCtxKey, entities)

		if err := app.datasource.LoadEntities(app.Context()); err != nil {
			app.Logger.Fatal(err)
		}
	}
}

// Security : configure security
func Security(configuration *security.Configuration) AppOptions {
	return func(app *App) {
		app.security = configuration
	}
}
