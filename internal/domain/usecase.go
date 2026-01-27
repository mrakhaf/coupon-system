package domain

type CouponUseCase interface {
	CreateCoupon(req *CreateCouponRequest) (*Coupon, error)
	GetCouponByID(id string) (*Coupon, error)
	GetCouponByCode(code string) (*Coupon, error)
	UpdateCoupon(id string, req *UpdateCouponRequest) (*Coupon, error)
	DeleteCoupon(id string) error
	ListCoupons(limit, offset int) ([]*Coupon, error)
	UseCoupon(code string, amount float64) (*UseCouponResponse, error)
}
