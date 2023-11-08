package entity

type User struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	FullName string  `json:"full_name" gorm:"not null" validate:"required"`
	Email    string  `json:"email" gorm:"unique;not null" validate:"required"`
	Username string  `json:"username" gorm:"unique;not null" validate:"required"`
	Password string  `json:"-" gorm:"not null" validate:"required"`
	Deposit  float64 `json:"deposit" gorm:"not null;default:0" validate:"required"`
}

// type UserActivityLog struct {
// 	ID          int    `json:"id"`
// 	UserID      int    `json:"user_id"`
// 	Description string `json:"description"`
// }