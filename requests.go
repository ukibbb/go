package main

import "net/mail"

const minUserNameLength int = 3

type UserRegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *UserRegisterRequest) validate() map[string]string {
	var errors = map[string]string{}
	// enforce some kind of password
	if r.Password != r.PasswordConfirm {
		errors["password"] = "password and password confirm doesn't match"
	}
	if len(r.Username) < minUserNameLength {
		errors["username"] = "username minimum length equals 3"
	}
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		errors["email"] = "invalid email"
	}

	// // if username is not unique
	// errors["username"] = fmt.Sprintf("%s and username is not unique", errors["username"])
	// errors["email"] = fmt.Sprintf("%s and email is not unique", errors["email"])
	// http.StatusConflict

	return errors

}

// return erros or user indetifier email or username
func (r *UserLoginRequest) validate() map[string]string {
	var errors map[string]string
	if len(r.Username) == 0 && len(r.Email) == 0 {
		errors["user"] = "username or passowrd must be provided"
	}
	if len(r.Username) < minUserNameLength {
		errors["username"] = "username minimum length equals 3"
	}
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		errors["email"] = "invalid email"
	}
	return errors
}
