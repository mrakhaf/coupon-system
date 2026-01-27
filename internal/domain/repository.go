package domain

import "github.com/google/uuid"

type CouponRepository interface {
	Create(coupon *Coupon) error
	GetByID(id uuid.UUID) (*Coupon, error)
	GetByCode(code string) (*Coupon, error)
	Update(coupon *Coupon) error
	Delete(id uuid.UUID) error
	List(limit, offset int) ([]*Coupon, error)
	IncrementUsage(id uuid.UUID) error
}
