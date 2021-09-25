package auth

import "fmt"

type Auth struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshExpiresIn int    `json:"refresh_expires_in,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
}

func (auth Auth) String() string {
	return fmt.Sprintf("{\nAccessToken: %v, \nRefreshToken: %v, \nExpiresIn: %v, \nRefreshExpiresIn: %v, \nTokenType: %v \n}\n", auth.AccessToken, auth.RefreshToken, auth.ExpiresIn, auth.RefreshExpiresIn, auth.TokenType)
}
