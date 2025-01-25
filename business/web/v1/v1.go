package v1

import (
	"os"

	"github.com/Klimentin0/courses-service/business/web/v1/mid"
	"github.com/Klimentin0/courses-service/foundation/logger"
	"github.com/Klimentin0/courses-service/foundation/web"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build    string
	Shutdown chan os.Signal
	Log      *logger.Logger
}

// RouteAdder defines behavoir that sets the routes to bind for an instance of the service
type RouteAdder interface {
	Add(app *web.App, cfg APIMuxConfig)
}

// APIMux constructs an http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, routeAdder RouteAdder) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log))

	routeAdder.Add(app, cfg)

	return app
}
