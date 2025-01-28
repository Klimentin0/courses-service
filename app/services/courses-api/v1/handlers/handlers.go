package handlers

import (
	"github.com/Klimentin0/courses-service/app/services/courses-api/v1/handlers/checkgrp"
	"github.com/Klimentin0/courses-service/app/services/courses-api/v1/handlers/hackgrp"
	v1 "github.com/Klimentin0/courses-service/business/web/v1"
	"github.com/Klimentin0/courses-service/foundation/web"
)

type Routes struct{}

// Add implements the RouterAdder interface to add all routes.
func (Routes) Add(app *web.App, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(app, hackgrp.Config{
		Auth: apiCfg.Auth,
	})

	checkgrp.Routes(app, checkgrp.Config{
		Build: apiCfg.Build,
		Log:   apiCfg.Log,
	})
}
