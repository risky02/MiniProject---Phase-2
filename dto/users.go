package dto

type User struct {
	ID       int     `json:"id"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Deposit  float64 `json:"deposit"`
}
