package dto

type AuthToken struct {
	phoneNumber string
	tokenAuth   string
}

type UserData struct {
	Email  string `gorm:"column:user_email"`
	Expire string `gorm:"column:expire"`
}
