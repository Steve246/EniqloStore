package dto

type RequestRegistBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

type RequestLoginBody struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
