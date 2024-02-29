package health

import (
	"app/internal/bootstrap"
	"app/pkg/errors"
	"app/pkg/utils"
	"net/http"
)

func GetHealth(app *bootstrap.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := utils.Envelope{
			"status": "available",
			"system_info": map[string]string{
				"enviornment": app.Config.Env,
				"version": bootstrap.Version,
			},
		}

		err := utils.WriteJson(w, http.StatusOK, data, nil)
		if err != nil{
			errs.ServerErrorResponse(app.Logger, w, r, err)
		}
	})
}