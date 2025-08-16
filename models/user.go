package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// TableName allows you to specify a custom table name for the User model
func (u *User) TableName() string {
	return "users"
}
