package mid

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Klimentin0/courses-service/business/web/v1/metrics"
	"github.com/Klimentin0/courses-service/foundation/web"
)

// Panics recovers from panics and converts the panic to an error so it is
// reported in Metrics and handled in Errors.
func Panics() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {

			//Defer a function to recover from a panic and set the err return
			//variable after the fact.
			//в дефере нету ретёрна, поэтому для этой функции у нас именной ретёрн(err)
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					err = fmt.Errorf("PANIC [%v] TRACE [%s]", rec, string(trace))

					metrics.AddPanics(ctx)
				}
			}()

			return handler(ctx, w, r)
		}
		return h
	}
	return m
}
