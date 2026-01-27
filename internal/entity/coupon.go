package entity

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name            string    `json:"name" gorm:"uniqueIndex;not null;size:50"`
	Amount          float64   `json:"amount" gorm:"not null"`
	RemainingAmount float64   `json:"remaining_amount" gorm:"not null"`
	MaxUsage        int       `json:"max_usage" gorm:"not null"`
	UsedCount       int       `json:"used_count" gorm:"not null;default:0"`
	IsActive        bool      `json:"is_active" gorm:"not null;default:true"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type CouponClaim struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
	UserID    string    `json:"user_id" gorm:"index;not null;size:100"`
	CouponID  uuid.UUID `json:"coupon_id" gorm:"index;not null"`
	Coupon    Coupon    `json:"-" gorm:"foreignKey:CouponID"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateCouponRequest struct {
	Name   string  `json:"name" validate:"required,min=3,max=50"`
	Amount float64 `json:"amount" validate:"required,gt=0"`
}

type ClaimCouponRequest struct {
	UserID     string `json:"user_id" validate:"required,min=1,max=100"`
	CouponName string `json:"coupon_name" validate:"required,min=3,max=50"`
}

type CouponDetailsResponse struct {
	Name            string   `json:"name"`
	Amount          float64  `json:"amount"`
	RemainingAmount float64  `json:"remaining_amount"`
	ClaimedBy       []string `json:"claimed_by"`
}
