package token

import "github.com/ams-api/internal/config"

// maker is an interface for managing token
type IMaker interface {
	//CreateToken creates a new token for a specific username and duration
	CreateToken(userId uint) (string, error)

	//VerifyToken checks the token is valid or not
	VerifyToken(token string) (*Payload, error)
}

func NewToken(config config.Config) (IMaker, error) {
	switch config.TokenType {
	case "jwt":
		return NewJWTMaker(config.TokenSymmetricKey)
	default:
		return NewJWTMaker(config.TokenSymmetricKey)
	}
}
