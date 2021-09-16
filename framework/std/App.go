package std

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hsedjame/gowebapi/framework/core"
	"github.com/hsedjame/gowebapi/framework/data"
	"github.com/hsedjame/gowebapi/framework/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

type App struct {
	Logger *log.Logger
	server *http.Server
	ctx    core.AppCtx
	classpath string
	properties *AppProperties
	datasource data.Datasource
}

// New : CreateDatasource new app
func New(logger *log.Logger, options ...AppOptions) (*App, error) {

	app := &App{
		Logger: logger,
		ctx: context.Background(),
	}

	if err := app.setClasspath(); err != nil {
		return nil, err
	}

	return app.
		addToContext(core.EntitiesCtxKey, []interface{}{}).
		addToContext(core.ControllersCtxKey, []interface{}{}).
		loadProperties().
		WithOptions(options...), nil
}

// WithOptions : Add options to application
func (app *App) WithOptions(options ...AppOptions) *App {
	for _, option := range options {
		option(app)
	}
	return app
}

// Context : Get application context
func (app App) Context() core.AppCtx {
	return app.ctx
}

// WithContext : set an application context
func (app *App) WithContext(ctx core.AppCtx) {
	app.ctx = ctx
}

// Add new key to application context
func (app *App) addToContext(key core.CtxKey, value interface{}) *App {
	ctx := context.WithValue(app.Context(), key, value)
	app.WithContext(ctx)
	return app
}

// Set application classpath
func (app *App) setClasspath() error {
	if rootDir, err := os.Getwd(); err != nil {
		return err
	} else {
		app.classpath = fmt.Sprintf("%s/%s", rootDir, core.ResourcesLocation)
		//app.Logger.Printf("Application classpath configured to : %s \n", app.classpath)
		return nil
	}
}

func (app *App) configureWebServer() {

	router := mux.NewRouter()

	for _, controller := range app.Context().Value(core.ControllersCtxKey).([]interface{}) {
		ctrl := (controller).(web.RestController)
		subRouter := router.PathPrefix(ctrl.Path()).Subrouter()

		if ctrl.MiddleWare != nil {
			subRouter.Use(ctrl.MiddleWare)
		}
		subRouter.Use(web.LoggingMiddleware(app.Logger))

		/*if ctrl.DefaultModel() != nil && ctrl.ModelKey() != nil {
			eh := ctrl.ErrorHandler()
			if eh == nil {
				eh = func(err error, writer http.ResponseWriter) error {
					return jsonUtils.ToJson(core.AppError{ Message: err.Error()}, writer)
				}
			}
			subRouter.Use(web.PostPutMethodHandler(ctrl.DefaultModel(), ctrl.ModelKey(), eh))
		}*/

		for _, endpoint := range ctrl.Endpoints() {
			subRouter.HandleFunc(endpoint.Path(), endpoint.Handler()).Methods(endpoint.HttpMethod())
		}
	}

	/* Configure CORS */
	var opts []handlers.CORSOption
	cors := app.properties.Cors
	origins := cors.AllowedOrigins
	headers := cors.AllowedHeaders
	methods := cors.AllowedMethods

	if origins != "" {
		opts = append(opts, handlers.AllowedOrigins(strings.Split(origins, ",")))
	}
	if headers != "" {
		opts = append(opts, handlers.AllowedHeaders(strings.Split(headers, ",")))
	}
	if methods != "" {
		opts = append(opts, handlers.AllowedHeaders(strings.Split(methods, ",")))
	}

	corsHandlers := handlers.CORS(opts...)

	/* Configure server */
	port := app.properties.Server.Port
	if port == 0 {
		port = 8080
	}

	app.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      corsHandlers(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

}

func (app App) Run()  {


	// Start web server
	app.configureWebServer()

	// Launch the application
	go func() {
		app.Logger.Fatal(app.server.ListenAndServe())
	}()

	// Shutdown the application
	app.shutDownGracefully()
}

func (app *App) loadProperties() *App {
	// Load properties
	properties := AppDefaultProperties()
	if err := properties.Load(app.classpath); err != nil {
		app.Logger.Fatal(properties.Load(app.classpath))
	}
	app.properties = properties

	// Start database
	var datasource data.Datasource

	dbProperties := properties.Datasource
	if source, err := data.CreateDatasource(dbProperties); err != nil {
		app.Logger.Fatal(err)
	} else {
		datasource = source(dbProperties)
		app.datasource = datasource

		if datasource.CanStart() {
			if ctx, desc, err := datasource.Start(app.Context()); err != nil {
				app.Logger.Fatal(err)
			} else {
				app.WithContext(ctx)
				app.Logger.Printf("Database connected successfully with description %s", desc)
			}
		}
	}
	return app
}

func (app App) shutDownGracefully() {

	// Create a channel to listen OS signals
	osSignalsChannel := make(chan os.Signal)

	// Send a message to the channel when
	//  - interruption occurs
	//  - os is killed
	signal.Notify(osSignalsChannel, os.Interrupt)
	signal.Notify(osSignalsChannel, os.Kill)

	// Wait for new signal
	_  = <- osSignalsChannel

	app.Logger.Println(" ### Arrêt du serveur ....")

	deadline, cancel := context.WithTimeout(app.Context(), 30 * time.Second)

	defer cancel()

	// stop datasource
	if app.datasource != nil {
		app.datasource.Stop(app.Context())
	}

	// Shutdown the server
	app.Logger.Fatal(app.server.Shutdown(deadline))

}

func (app App) DB() (interface{}, error) {

	datasource := app.properties.Datasource

	if datasource != nil {
		db := app.Context().Value(core.DBCtxKey)

		switch datasource.Type {
		case data.POSTGRES:
			return db.(*pg.DB), nil
		case data.MONGO:
			return nil, errors.New("NOT IMPLEMENTED YET")
		case data.ORACLE:
			return db.(*sql.DB), nil
		default:
			return nil, errors.New("DATASOURCE TYPE NOT SUPPORTED")
		}
	}
	return nil, errors.New("DATASOURCE NOT PROVIDED")
}