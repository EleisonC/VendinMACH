package models


type UserModel struct{
	Id string `json:"id"`
	Username string `json:"username" validate:"required,min=3,max=12"`
	Password string	`json:"password" validate:"required,min=5"`
	Deposit int		`json:"deposit"`
	Role string	`json:"role" validate:"required"`
}

type UserLogIn struct {
	Username string `json:"username" validate:"required"`
	Password string	`json:"password" validate:"required"`
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
