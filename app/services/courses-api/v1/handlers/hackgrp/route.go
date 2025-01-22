package hackgrp

import (
	"net/http"

	"github.com/Klimentin0/courses-service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	app.Handle(http.MethodGet, "/hack", Hack)
}
