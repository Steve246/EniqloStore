package dto

type SuccessLoginBody struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type RequestLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
