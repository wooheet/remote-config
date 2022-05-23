package models

// gen:qs
type Users struct {
	//gorm.Model
	ID       uint64 `gorm:"unique" json:"id"`
	Email    string `gorm:"unique" json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
