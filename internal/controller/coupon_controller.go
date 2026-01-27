package controller

import (
	"net/http"

	"coupon-system/internal/entity"
	"coupon-system/internal/usecase"

	"github.com/labstack/echo/v4"
)

type CouponController struct {
	useCase *usecase.CouponUseCase
}

func NewCouponController(useCase *usecase.CouponUseCase) *CouponController {
	return &CouponController{useCase: useCase}
}

// Create Coupon - POST /api/coupons
func (c *CouponController) CreateCoupon(e echo.Context) error {
	req := new(entity.CreateCouponRequest)
	if err := e.Bind(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := e.Validate(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	coupon, err := c.useCase.CreateCoupon(req)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusCreated, coupon)
}

// Claim Coupon - POST /api/coupons/claim
func (c *CouponController) ClaimCoupon(e echo.Context) error {
	req := new(entity.ClaimCouponRequest)
	if err := e.Bind(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := e.Validate(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := c.useCase.ClaimCoupon(req)
	if err != nil {
		if err.Error() == "user has already claimed this coupon" {
			return e.JSON(http.StatusConflict, map[string]string{"error": "User has already claimed this coupon"})
		}
		if err.Error() == "coupon not found" || err.Error() == "coupon is not active" || err.Error() == "coupon has no remaining amount" {
			return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		return e.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, map[string]string{"message": "Coupon claimed successfully"})
}

// Get Coupon Details - GET /api/coupons/{name}
func (c *CouponController) GetCouponDetails(e echo.Context) error {
	name := e.Param("name")

	details, err := c.useCase.GetCouponDetails(name)
	if err != nil {
		return e.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
	}

	return e.JSON(http.StatusOK, details)
}
