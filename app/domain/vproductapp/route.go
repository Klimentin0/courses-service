package vproductapp

import (
	"net/http"

	"github.com/Klimentin0/courses-service1/app/sdk/auth"
	"github.com/Klimentin0/courses-service1/app/sdk/authclient"
	"github.com/Klimentin0/courses-service1/app/sdk/mid"
	"github.com/Klimentin0/courses-service1/business/domain/userbus"
	"github.com/Klimentin0/courses-service1/business/domain/vproductbus"
	"github.com/Klimentin0/courses-service1/foundation/logger"
	"github.com/Klimentin0/courses-service1/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	UserBus     *userbus.Business
	VProductBus *vproductbus.Business
	AuthClient  *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.AuthClient, auth.RuleAdminOnly)

	api := newApp(cfg.VProductBus)

	app.HandlerFunc(http.MethodGet, version, "/vproducts", api.query, authen, ruleAdmin)
}
