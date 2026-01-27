package domain

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID          uuid.UUID `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Discount    float64   `json:"discount"`
	MinAmount   float64   `json:"min_amount"`
	MaxUsage    int       `json:"max_usage"`
	UsedCount   int       `json:"used_count"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateCouponRequest struct {
	Code        string  `json:"code" validate:"required,min=3,max=20"`
	Description string  `json:"description" validate:"required,max=200"`
	Discount    float64 `json:"discount" validate:"required,gt=0,lt=100"`
	MinAmount   float64 `json:"min_amount" validate:"gte=0"`
	MaxUsage    int     `json:"max_usage" validate:"gte=1"`
}

type UpdateCouponRequest struct {
	Description string  `json:"description" validate:"required,max=200"`
	Discount    float64 `json:"discount" validate:"required,gt=0,lt=100"`
	MinAmount   float64 `json:"min_amount" validate:"gte=0"`
	MaxUsage    int     `json:"max_usage" validate:"gte=1"`
	IsActive    bool    `json:"is_active"`
}

type UseCouponRequest struct {
	Code string  `json:"code" validate:"required"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type UseCouponResponse struct {
	Success    bool    `json:"success"`
	Message    string  `json:"message"`
	Discount   float64 `json:"discount,omitempty"`
	FinalAmount float64 `json:"final_amount,omitempty"`
}