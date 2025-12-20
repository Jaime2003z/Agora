package identity

import "errors"

type IdentityType string

const (
	User     IdentityType = "user"
	Executor IdentityType = "executor"
)

type Identity struct {
	PublicKey  string
	Location   string
	Reputation int
	UserType   IdentityType
}

// constructor
func NewIdentity(publicKey string, location string) (*Identity, error) {
	if publicKey == "" {
		return nil, errors.New("public key is required")
	}

	if location == "" {
		return nil, errors.New("location is required")
	}

	return &Identity{
		PublicKey:  publicKey,
		Location:   location,
		Reputation: 0,
		UserType:   User,
	}, nil
}
