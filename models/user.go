package models

const (
	AdminRole = "admin"
	UserRole  = "user"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Username string `gorm:"not null;unique"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}
