package core

import "context"

type CtxKey string

const (
	EntitiesCtxKey    CtxKey = "entities-ctx-key"
	ControllersCtxKey CtxKey = "controllers-ctx-key"
	DBCtxKey          CtxKey = "db-ctx-key"
	SecurityCtxKey    CtxKey = "security-context-key"
)

type AppCtx = context.Context
