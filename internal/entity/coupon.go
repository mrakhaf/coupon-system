package entity

import (
	"time"

	"github.com/google/uuid"
)

type Coupon struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Name      string    `json:"name" gorm:"uniqueIndex;not null;size:50"`
	Amount    int       `json:"amount" gorm:"not null"`
	IsActive  bool      `json:"is_active" gorm:"not null;default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CouponClaim struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:varchar(36)"`
	UserID    string    `json:"user_id" gorm:"index;not null;size:100"`
	CouponID  uuid.UUID `json:"coupon_id" gorm:"index;not null"`
	Coupon    Coupon    `json:"-" gorm:"foreignKey:CouponID"`
	CreatedAt time.Time `json:"created_at"`
}
