package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken create a new token for a specific user and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
