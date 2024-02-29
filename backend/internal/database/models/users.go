package models

import (
	"app/internal/database/fields"
	"app/pkg/validator"
	"database/sql"
	"time"
)


type User struct{
	UUID 			int64 		`json:"id,omitempty,string"`	//UUID for user
	CreatedAt 		time.Time 	`json:"-"`						//Timestamp for when user is added to database
	Username 		string 										//Username for user
	Email 			string 										//Email for user
	Password 		string 		`json:"-"`						//Password for user
	Verified 		bool										//User verification status (false, true)
	Premium 		bool										//User premium status (false, true)
	PremiumExpiry 	time.Time 									//Premium Expiry Date
	Version 		fields.Runtime 	`json:"runtime,omitempty"`
	
}

//CRUD
type UserModel struct {
	DB *sql.DB
}

//Create
func (m UserModel) Create(user *User) error {
	return nil
}

//Read
func (m UserModel) Get(id int64) (*User, error) {
	return nil, nil
}

//Update
func (m UserModel) Update(user *User) error {
	return nil
}

//Delete
func (m UserModel) Delete(id int64) error {
	return nil
}

func ValidateUser(v *validator.Validator, user *User){
	
		//Username Validation
		v.Check(user.Username != "", "username", "must be provided")
		v.Check(len(user.Username) > 2, "username", "username is too short, must be atleast 3 characters")
		v.Check(len(user.Username) < 16, "username", "username is too long, must be max 15 characters")
		
		//Email Validation
		v.Check(user.Email != "", "email", "must be provided")
		v.Check(validator.Matches(user.Email, validator.EmailRX), "email", "must be a valid email address")
		
		//Password Validation
		v.Check(len(user.Password) > 7, "password", "password must be atleast 8 characters long")

		//Runtime Validation
		v.Check(user.Version != 0, "version", "must be provided")
		v.Check(user.Version > 0, "version", "must be a positive integer")
}

//Aliasing Approach (Note: field order not maintained)
// func (u User) MarshalJSON() ([]byte, error){
// 	var runtime string

// 	if u.Runtime != 0{
// 		runtime = fmt.Sprintf("%d mins", u.Runtime)
// 	}

// 	//Define an alias type which has underlying type user.
// 	//The alias type will contain all fields that user has but not the methods.
// 	type UserAlias User


// 	aux := struct{
// 		UserAlias
// 		Runtime string `json:"runtime,omitempty"`
// 	}{
// 		UserAlias: UserAlias(u),
// 		Runtime: runtime,
// 	}

// 	return json.Marshal(aux)
// }