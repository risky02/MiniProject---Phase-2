package dto

type User struct {
	ID       int     `json:"id"`
	FullName string  `json:"full_name"`
	Email    string  `json:"email"`
	Username string  `json:"username"`
	Password string  `json:"-"`
	Deposit  float32 `json:"deposit"`
}

type Order struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	EquipmentID int     `json:"equipmentID"`
	RentalDays  int     `json:"rental_days"`
	TotalCost   float32 `json:"total_cost"`
}

type Checkout struct {
	Id          int     `json:"id"`
	UserId      int     `json:"user_id"`
	EquipmentId int     `json:"equipment_id"`
	RentalDate  string  `json:"rental_date"`
	ReturnDate  string  `json:"return_date"`
	RentalDays  string  `json:"rentaldays"`
	TotalCost   float64 `json:"total_cost"`
}

type Payment struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	CheckoutID int     `json:"checkout_id"`
	RentalDays int     `json:"rental_days"`
	TotalCost  float32 `json:"total_cost"`
	Payment    float32 `json:"payment"`
}