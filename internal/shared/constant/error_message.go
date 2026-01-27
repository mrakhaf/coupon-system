package constant

import "net/http"

const (
	ErrCouponNotFound              = "coupon not found"
	ErrCouponNotActive             = "coupon is not active"
	ErrCouponNoRemainingAmount     = "coupon has no remaining amount"
	ErrUserHasAlreadyClaimedCoupon = "user has already claimed this coupon"
)

var (
	CodeErrorMessage = map[string]int{
		ErrCouponNotFound:              http.StatusBadRequest,
		ErrCouponNotActive:             http.StatusBadRequest,
		ErrCouponNoRemainingAmount:     http.StatusBadRequest,
		ErrUserHasAlreadyClaimedCoupon: http.StatusConflict,
	}
)
