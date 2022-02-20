package user

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID           int64       `json:"id" db:"user_id"`
	Name         string      `json:"name" db:"name"`
	Status       int         `json:"status" db:"status"`
	Phone        string      `json:"phone" db:"phone"`
	PhoneFilter  string      `json:"phone_filter" db:"phone_filter"`
	OTP          null.String `json:"otp" db:"otp"`
	OTPCreatedAt *time.Time  `json:"otp_created_at" db:"otp_created_at"`
	CreatedAt    time.Time   `json:"-" db:"created_at"`
	UpdatedAt    *time.Time  `json:"-" db:"updated_at"`
	UpdatedBy    *int64      `json:"-" db:"updated_by"`
}

type CreateUser struct {
	Name         string
	Status       int
	Phone        string
	PhoneFilter  string
	Password     string
	OTP          null.String
	OTPCreatedAt null.Time
}

type VerifyUser struct {
	PhoneFilter string
	OTP         string
}

type CreateUserRequest struct {
	Name  string `json:"name" validate:"empty=false"`
	Phone string `json:"phone" validate:"empty=false"`
}

type VerifyOTPRequest struct {
	Phone       string `json:"phone" validate:"empty=false"`
	OTP         string `json:"otp" validate:"empty=false"`
	IsFirstOTP  bool   `json:"is_first_otp"`
	PhoneFilter string
}

type UserLoginRequest struct {
	Phone string `json:"phone" validate:"empty=false"`
}

type UserDetailResponse struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Phone  string `json:"phone"`
}
