package request

type CreateCouponRequest struct {
	Name   string `json:"name" validate:"required"`
	Amount int    `json:"amount" validate:"required"`
}

type ClaimCouponRequest struct {
	CouponName string `json:"name" validate:"required"`
	UserID     string `json:"user_id" validate:"required"`
}
