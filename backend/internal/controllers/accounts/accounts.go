package accounts

import (
	"app/internal/bootstrap"
	"app/internal/database/models"
	"app/pkg/errors"
	"app/pkg/utils"
	"app/pkg/validator"
	"errors"
	"fmt"
	"net/http"
)

//Create a new user
func CreateAccount(app *bootstrap.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input struct{
			Email string `json:"email"`
			Password string `json:"password"`
		}

		err := utils.ReadJSON(w, r, &input)

		if err != nil{
			errs.BadRequestResponse(app.Logger, w, r, err)
			return
		}

		account := &models.Account{
			Email: input.Email,
			Password: input.Password,
		}
		
		v:= validator.New();

		if models.ValidateAccount(v, account); !v.Valid(){
			errs.FailedValidationResponse(app.Logger, w, r, v.Errors)
			return
		}
		err = app.Models.Accounts.Create(account)

		if err != nil{
			errs.ServerErrorResponse(app.Logger, w, r, err)
			return
		}

		headers := make(http.Header)
		headers.Set("Location", fmt.Sprintf("/v1/accounts//%d", account.ID))

		err = utils.WriteJson(w, http.StatusCreated, utils.Envelope{"account":account}, headers)
		if err != nil{
			errs.ServerErrorResponse(app.Logger, w, r, err)
		}
	})
}


//Gets a user
func GetAccount(app *bootstrap.Application) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ReadParam("id", r)
		if err != nil {
			errs.NotFoundResponse(app.Logger, w, r)
			return
		}

		account, err := app.Models.Accounts.Get(id)

		if err != nil{
			switch{
			case errors.Is(err, models.ErrRecordNotFound):
				errs.NotFoundResponse(app.Logger, w, r)
			default:
				errs.ServerErrorResponse(app.Logger, w, r, err)
			}
			return
		}

		err = utils.WriteJson(w, http.StatusOK, utils.Envelope{"account":account}, nil)
		if err != nil {
			errs.ServerErrorResponse(app.Logger, w, r, err)
		}
	})
}