package entity

type User struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	FullName string  `json:"full_name" gorm:"not null" validate:"required"`
	Email    string  `json:"email" gorm:"unique;not null" validate:"required"`
	Username string  `json:"username" gorm:"unique;not null" validate:"required"`
	Password string  `json:"-" gorm:"not null" validate:"required"`
	Deposit  float32 `json:"deposit" gorm:"not null;default:0" validate:"required"`
}

type Equipment struct {
	Equipment_id int     `json:"id"`
	Name         string  `json:"product"`
	Price        float32 `json:"price"`
	Description  string  `json:"description"`
}

type Checkout struct {
	Id          int     `json:"id"`
	UserId      int     `json:"user_id"`
	EquipmentId int     `json:"equipment_id"`
	RentalDate  string  `json:"rental_date"`
	ReturnDate  string  `json:"return_date"`
	RentalDays  string  `json:"rentaldays"`
	TotalCost   float32 `json:"total_cost"`
}

type Payment struct {
	ID         int     `json:"id" gorm:"primaryKey"`
	UserID     int     `json:"user_id"`
	CheckoutID int     `json:"checkout_id"`
	RentalDays string  `json:"rental_days" gorm:"not null" validate:"required"`
	TotalCost  float32 `json:"total_cost" gorm:"not null" validate:"required"`
	Payment    float32 `json:"payment"`
}