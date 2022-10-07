package data

import "time"

type User struct {
	ID          string    `json:"id" sql:"id"`
	Email       string    `json:"email" sql:"email" validate:"required"`
	Password    string    `json:"password" sql:"password" validate:"required"`
	Username    string    `json:"username" sql:"username"`
	TokenHash   string    `json:"tokenhash" sql:"tokenhash"`
	IsCertified bool      `json:"iscertified" sql:"iscertified"`
	CreatedAt   time.Time `json:"createdat" sql:"createdat"`
	UpdatedAt   time.Time `json:"updatedat" sql:"updatedat"`
}

type VerificationDataType int

const (
	MailConfirmation VerificationDataType = iota + 1
	PassReset
)

// VerificationData represents the type for the data stored for verification.
type VerificationData struct {
	Email     string               `json:"email" sql:"email" validate:"required"`
	Code      string               `json:"code" sql:"code" validate:"required"`
	ExpiresAt time.Time            `json:"expiresat" sql:"expiresat"`
	Type      VerificationDataType `json:"type" sql:"type"`
}
