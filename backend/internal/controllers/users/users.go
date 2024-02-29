package users

import (
	"app/internal/database/fields"
	"app/internal/database/models"
	"app/pkg/errors"
	"app/pkg/utils"
	"app/pkg/validator"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Create a new user
func CreateUser(logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var newUser struct{
			Username string `json:"username"`
			Email string `json:"email"`
			Password string `json:"password"`
			Version fields.Runtime `json:"runtime"`

		}

		err := utils.ReadJSON(w, r, &newUser)

		if err != nil{
			errs.BadRequestResponse(logger, w, r, err)
			return
		}

		user := &models.User{
			Username: newUser.Username,
			Email: newUser.Email,
			Password: newUser.Password,
			Version: newUser.Version,
		}
		
		v:= validator.New();

		if models.ValidateUser(v, user); !v.Valid(){
			errs.FailedValidationResponse(logger, w, r, v.Errors)
			return
		}

		fmt.Fprintf(w, "%+v\n", newUser)
	})
}


//Gets a user
func GetUser(logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.ReadParam("id", r)
		if err != nil {
			errs.NotFoundResponse(logger, w, r)
			return
		}

		user := models.User{
			UUID: id,
			CreatedAt: time.Now(),
			Username: "ModD",
			Email: "modd@gmail.com",
			Password: "moddgoat",
			Verified: true,
			Premium: true,
			PremiumExpiry: time.Now().AddDate(1000, 0, 0),
			Version: 1,
		}

		err = utils.WriteJson(w, http.StatusOK, utils.Envelope{"user":user}, nil)
		if err != nil {
			errs.ServerErrorResponse(logger, w, r, err)
		}
	})
}