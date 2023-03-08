package models


type UserModel struct{
	Id string `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string	`json:"password" validate:"required,min=5"`
	Deposit int		`json:"deposit" validate:"required,oneof=5 10 20 50 100"`
	Role string	`json:"role" validate:"required"`
}

type EditUser struct {
	Username	string `json:"username" validate:"required"`
	Role	string `json:"role" validate:"required"`
}

type Deposit struct {
	Deposit int `json:"deposit" validate:"required,oneof=5 10 20 50 100"`
}

type UserLogIn struct {
	Username string `json:"username" validate:"required"`
	Password string	`json:"password" validate:"required"`
}

type PasswordChange struct {
	OldPassword	string	`json:"oldpassword" validate:"required"`
	NewPassword	string	`json:"newpassword" validate:"required,min=8"`
	ConfirmNewPassword	string	`json:"confirmpassword" validate:"required,eqfield=NewPassword"`
}

type UserModeldb struct{
	Id string `json:"id" db:"id"`
	Username string `json:"username" db:"username" validate:"required,min=3,max=12"`
	Password []byte	`json:"password" db:"password"`
	Deposit int		`json:"deposit" db:"deposit"`
	Role string	`json:"role" db:"role" validate:"required"`
}

type TokenRes struct {
	Message string
	TokenString string
}

type ErrMessageRes struct {
	Message string `json:"message"`
	RawErrorMessage string `json:"raw err message"`
}

type PosMessageRes struct {
	Message string `json:"message"`
}
