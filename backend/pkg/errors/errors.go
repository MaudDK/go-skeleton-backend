package errs

import (
	"app/pkg/utils"
	"fmt"
	"log"
	"net/http"
)

func LogError (logger *log.Logger, r *http.Request, err error){
	logger.Println(err)
}

func ErrorResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, status int, message any){
	env := utils.Envelope{"error": message}

	err := utils.WriteJson(w, status, env, nil)
	if err != nil{
		LogError(logger, r, err)
		w.WriteHeader(500)
	}
}

//Sends 400 Bad Request and JSON response to client
func BadRequestResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, err error){
	ErrorResponse(logger, w, r, http.StatusBadRequest, err.Error())
}

//Sends 500 Internal Server Error and JSON response (generic error message) to client
func ServerErrorResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, err error){
	LogError(logger, r, err)
	message := "the server encountred a problem and could not process your request"
	ErrorResponse(logger, w, r, http.StatusInternalServerError, message)
}

//Sends 404 Not Found Error and JSON response (generic error message) to client
func NotFoundResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request){
	message := "the requested resource could not be found"
	ErrorResponse(logger, w, r, http.StatusNotFound, message)
}

//Sends 405 Method Not Allowed status code and JSON response to client
func MethodNotAllowedResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request){
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	ErrorResponse(logger, w, r, http.StatusMethodNotAllowed, message)
}

//Sends 442 Unprocessable Entity
func FailedValidationResponse(logger *log.Logger, w http.ResponseWriter, r *http.Request, errors map[string]string){
	ErrorResponse(logger, w, r, http.StatusUnprocessableEntity, errors)
}
