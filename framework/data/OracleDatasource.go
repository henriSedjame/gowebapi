package data

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hsedjame/gowebapi/framework/core"
	go_ora "github.com/sijms/go-ora/v2"
)

type OracleDatasource struct {
	WithProperties
}

func (d *OracleDatasource) LoadEntities(ctx context.Context) error {
	db := ctx.Value(core.DBCtxKey).(*sql.DB)
	if driver, ok := db.Driver().(*go_ora.OracleDriver); ok {
		return  core.AppError{Message: ""}
	} else {
		for _, entity := range ctx.Value(core.EntitiesCtxKey).([]Entity) {
			if err := driver.Conn.RegisterType2(entity.TableName(), entity); err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *OracleDatasource) Start(ctx context.Context) (context.Context, string, error) {

	if d.IsValid() {
		// Database Url
		url := fmt.Sprintf("oracle://%s:%s@%s:%d/%s",
			d.Properties.Username, d.Properties.Password, d.Properties.Host, d.Properties.Port, d.Properties.Database)

		// Connect DB
		if db, err := sql.Open("oracle", url); err != nil {
			return ctx, "", err
		} else {

			// Test connection
			if err := db.Ping(); err != nil {
				return ctx, "", err
			}
			
			// Retrieve db version description
			dbDesc := ""
			if rows, err := db.Query("SELECT BANNER FROM v$version\nWHERE banner LIKE ?", "Oracle%"); err != nil {
				return ctx, "", err
			} else {
				for rows.Next() {
					if err:= rows.Scan(dbDesc); err != nil {
						return ctx, "", err
					}
				}
			}

			// Update the application context
			ctx = context.WithValue(ctx, core.DBCtxKey, db)
			
			return ctx, dbDesc, nil
		}
		
	} else {
		return ctx, "", core.AppError{Message: "Database is not set. Please consider to set property datasource to your application.yml file."}
	}

}

func (d *OracleDatasource) Stop(ctx context.Context) {
	if db := ctx.Value(core.DBCtxKey).(*sql.DB); db != nil {
		_ = db.Close()
	}
}

func (d *OracleDatasource) CanStart() bool {
	return d.IsValid()
}

