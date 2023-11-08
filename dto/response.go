package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponFailed struct {
	Code    uint
	Message string
}

type ResponSuccess struct {
	Code    uint
	Message string
	Data    any
}

type GetToken struct {
	Code    uint
	Message string
	Token   string
}