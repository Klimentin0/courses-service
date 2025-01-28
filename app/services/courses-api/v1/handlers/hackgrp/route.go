package hackgrp

import (
	"net/http"

	"github.com/Klimentin0/courses-service/business/web/v1/auth"
	"github.com/Klimentin0/courses-service/business/web/v1/mid"
	"github.com/Klimentin0/courses-service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Auth *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	app.Handle(http.MethodGet, version, "/hack", Hack)
	app.Handle(http.MethodGet, version, "/hackauth", Hack, authen, ruleAdmin)
}
