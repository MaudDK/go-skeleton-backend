package models

import (
	"app/pkg/validator"
	"database/sql"
	"errors"
	"time"
)


type Account struct{
	ID 				int64 		`json:"id"`				//ID for account
	Email 			string 		`json:"email"`								//Email for account
	Password 		string 		`json:"-"`						//Password for account
	Verified 		bool		`json:"activated"`				//account verification status (false, true)
	CreatedAt 		time.Time 	`json:"created_at"`				//Timestamp for when account is added to database
	UpdatedAt 		time.Time 	`json:"updated_at,omitempty"`	//Timestamp for when account is added to database
}

var(
	ErrDuplicateEmail = errors.New("duplicate email")
)

//CRUD
type AccountModel struct {
	DB *sql.DB
}

//Create
func (m AccountModel) Create(account *Account) error {
	query := `
		INSERT INTO accounts (email, password, verified)
		VALUES($1, $2, $3)
		RETURNING id, created_at`

	args := []any{
		account.Email, 
		account.Password, 
		account.Verified}
	
	err := m.DB.QueryRow(query, args...).Scan(&account.ID, &account.CreatedAt)
	if err != nil {
		switch{
		case err.Error() == `duplicate key value violates unique constraint "accounts_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

//Read
func (m AccountModel) Get(id int64) (*Account, error) {
	if id < 1{
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, email, verified, created_at, updated_at
		FROM accounts
		where id = $1
	`

	var account Account

	err := m.DB.QueryRow(query, id).Scan(
		&account.ID,
		&account.Email,
		&account.Verified,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil{
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &account, nil
}

//Update
func (m AccountModel) Update(account *Account) error {
	return nil
}

//Delete
func (m AccountModel) Delete(id int64) error {
	return nil
}

func ValidateAccount(v *validator.Validator, account *Account){
		//Email Validation
		v.Check(account.Email != "", "email", "must be provided")
		v.Check(validator.Matches(account.Email, validator.EmailRX), "email", "must be a valid email address")
		
		//Password Validation
		v.Check(len(account.Password) > 7, "password", "password must be atleast 8 characters long")
}