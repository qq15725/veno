package veno

import (
	"github.com/qq15725/veno/container"
	"github.com/qq15725/veno/database/dsn"
	"github.com/qq15725/veno/database/model"
	"net/http"
	"path"
)

type Application struct {
	*container.Container
	Router               *Router
	basePath             string
	loadedConfigurations map[string]bool
}

func (app *Application) registerBindings() {
	app.Bind(newConfig)
}

func (app *Application) bootstrapRouter() {
	app.Router = newRouter()
}

func (app *Application) GetConfigurationPath(name string) string {
	return path.Join(app.basePath, "config", name+".json")
}

func (app *Application) Configure(name string) {
	if app.loadedConfigurations[name] {
		return
	}

	app.loadedConfigurations[name] = true

	cfg := app.Get((*Config)(nil)).(*Config)

	cfg.Load(name, app.GetConfigurationPath(name))
}

func (app *Application) bootstrapDatabase() {
	app.Configure("database")
	cfg := app.Get((*Config)(nil)).(*Config)
	cfgDb := cfg.Get("database").(map[string]interface{})
	cfgConns := cfgDb["connections"].(map[string]interface{})
	connections := make(map[string]dsn.Connector)
	for name, v := range cfgConns {
		conn := v.(map[string]interface{})
		switch conn["driver"] {
		case "mysql":
			connections[name] = &dsn.Mysql{
				Host:     conn["host"].(string),
				Port:     int(conn["port"].(float64)),
				Database: conn["database"].(string),
				Username: conn["username"].(string),
				Password: conn["password"].(string),
				Charset:  conn["charset"].(string),
			}
			break
		}
	}
	model.SetConnections(connections)
	defer model.CloseConnections()
}

func (app *Application) Run(addr string) (err error) {
	return http.ListenAndServe(addr, app.Router)
}

func New(basePath string) *Application {
	app := &Application{
		basePath:             basePath,
		Container:            container.New(),
		loadedConfigurations: make(map[string]bool),
	}

	app.registerBindings()

	app.bootstrapRouter()
	app.bootstrapDatabase()

	return app
}
