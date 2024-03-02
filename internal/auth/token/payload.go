package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Payload contains the payload data of token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserId    uint      `json:"user_id"`
	AgentId   uint      `json:"agent_id"`
	PlayerId  uint      `json:"player_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// NewPayload creates  a new token payload with a specific username and duration
func NewPayload(userId uint) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		UserId:    userId,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(168 * time.Hour),
	}
	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
