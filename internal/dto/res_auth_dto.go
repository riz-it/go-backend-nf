package dto

type SignUpResponse struct {
	Email       string `json:"email"`
	AccountName string `json:"account_name"`
}

type SignInResponse struct {
	User  CredentialData `json:"user"`
	Token TokenData      `json:"token"`
}

type CredentialData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
