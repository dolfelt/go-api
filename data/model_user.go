package data

// User model
type User struct {
	BaseModel
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone"`
	Password    string `json:"-"`
	HasPassword bool   `gorm:"-"`
}
