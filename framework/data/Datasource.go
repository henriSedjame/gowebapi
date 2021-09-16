package data

import (
	"context"
	"errors"
	"fmt"
)

type Datasource interface {
	Start(context.Context) (context.Context, string, error)
	Stop(context.Context)
	CanStart() bool
	LoadEntities(context.Context) error
}

type DatasourceFactory func(*DbProperties) Datasource

func CreateDatasource(properties *DbProperties) (DatasourceFactory, error) {
	if properties != nil{
		switch properties.Type {
		case POSTGRES:
			return func(properties *DbProperties) Datasource {
				return &PgDatasource{
					WithProperties{
						Properties: properties,
					},
				}
			}, nil
		case MONGO:
			return nil, errors.New("NOT IMPLEMENTED YET")
		case ORACLE:
			return func(properties *DbProperties) Datasource {
				return &OracleDatasource{
					WithProperties{
						Properties: properties,
					},
				}
			}, nil
		default:
			return nil, errors.New(fmt.Sprintf("TYPE %s NOT SUPPORTED YET", properties.Type))
		}
	}
	return nil, errors.New("DATASOURCE PROPERTIES ARE MISSING")
}
