package dto

type AuthToken struct {
	phoneNumber string
	tokenAuth   string
}

type UserData struct {
	UserUniqueId string `gorm:"column:user_unique_id"`
	Expire       string `gorm:"column:expire"`
}
