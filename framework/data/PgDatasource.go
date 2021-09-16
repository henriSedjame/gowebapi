package data

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/hsedjame/gowebapi/framework/core"
)

type PgDatasource struct {
	WithProperties
}

func (p PgDatasource) LoadEntities(ctx context.Context)  error {

	db := ctx.Value(core.DBCtxKey).(*pg.DB)

	// Try to create a table for each entities
	for _, mod := range ctx.Value(core.EntitiesCtxKey).([]interface{}) {

		if err := db.Model(mod).CreateTable(&orm.CreateTableOptions{
			Temp:          false,
			IfNotExists:   true,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (p PgDatasource) Start(ctx context.Context) (context.Context, string, error) {

	if p.IsValid() {

		// Create pqsl url
		optStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			p.Properties.Username, p.Properties.Password, p.Properties.Host, p.Properties.Port, p.Properties.Database)

		// Try to parse url
		if opt, err := pg.ParseURL(optStr); err != nil {
			return ctx, "", err
		} else {

			// Connect the database
			db := pg.Connect(opt)

			// Try to query psql version
			// in order to verify if database is working
			var version string
			if _, err := db.QueryOneContext(ctx, pg.Scan(&version), "SELECT version()"); err != nil {
				return ctx, "", err
			} else {

				// Update the application context
				ctx = context.WithValue(ctx, core.DBCtxKey, db)
				return ctx, version, nil
			}
		}
	} else {
		return ctx, "", core.AppError{
			Message: "Database is not set. Please consider to set property { \"db\" : {\"database\" : ****}}",
		}
	}
}

func (p PgDatasource) Stop(ctx context.Context) {
	if db := ctx.Value(core.DBCtxKey).(*pg.DB); db != nil {
		_ = db.Close()
	}
}

func (p PgDatasource) CanStart() bool {
	return p.IsValid()
}

