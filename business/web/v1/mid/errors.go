package mid

import (
	"context"
	"net/http"

	"github.com/Klimentin0/courses-service/business/web/v1/auth"
	"github.com/Klimentin0/courses-service/business/web/v1/response"
	"github.com/Klimentin0/courses-service/foundation/logger"
	"github.com/Klimentin0/courses-service/foundation/validate"
	"github.com/Klimentin0/courses-service/foundation/web"
)

// Errors handles error coming out of the call chain.
// It detects normal application errors which are used to respond
// to the client in a uniform way. Unexpected error (status >= 500) are logged.
func Errors(log *logger.Logger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := handler(ctx, w, r); err != nil {
				log.Error(ctx, "message", "msg", err)

				var er response.ErrorDocument
				var status int

				switch {
				//Trused error
				case response.IsError(err):
					reqErr := response.GetError(err)

					if validate.IsFieldErrors(reqErr.Err) {
						fieldErrors := validate.GetFieldErrors(reqErr.Err)
						er = response.ErrorDocument{
							Error:  "data validation error",
							Fields: fieldErrors.Fields(),
						}
						status = reqErr.Status
						break
					}
					er = response.ErrorDocument{
						Error: reqErr.Error(),
					}
					status = reqErr.Status

				case auth.IsAuthError(err):
					er = response.ErrorDocument{
						Error: http.StatusText(http.StatusUnauthorized),
					}
					status = http.StatusUnauthorized
				//NON-Trusted error
				default:
					er = response.ErrorDocument{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				if err := web.Respond(ctx, w, er, status); err != nil {
					return nil
				}

				//If we receive the shutdown err we need to return it
				// back to the base handler to shut down the service.
				if web.IsShutdown(err) {
					return err
				}
			}
			return nil
		}
		return h
	}
	return m
}
