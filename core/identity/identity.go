package identity

import (
	"errors"

	"github.com/Jaime2003z/Agora/core/commons"
)

type IdentityType string

const (
	User     IdentityType = "user"
	Executor IdentityType = "executor"
)

type Identity struct {
	PublicKey  string
	Location   commons.LocalityID
	Reputation int
	UserType   IdentityType
}

// constructor
func NewIdentity(publicKey string, location *commons.LocalityID) (*Identity, error) {
	if publicKey == "" {
		return nil, errors.New("public key is required")
	}

	if location == nil {
		return nil, errors.New("location is required")
	}

	return &Identity{
		PublicKey:  publicKey,
		Location:   *location,
		Reputation: 10,
		UserType:   User,
	}, nil
}
