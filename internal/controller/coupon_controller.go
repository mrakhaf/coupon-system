package controller

import (
	"net/http"

	"coupon-system/internal/dto/request"
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
	req := new(request.CreateCouponRequest)
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
	req := new(request.ClaimCouponRequest)
	if err := e.Bind(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := e.Validate(req); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	statusCode, err := c.useCase.ClaimCoupon(req)
	if err != nil {
		return e.JSON(statusCode, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, map[string]string{"message": "Coupon claimed successfully"})
}

// Get Coupon Details - GET /api/coupons/{name}
func (c *CouponController) GetCouponDetails(e echo.Context) error {
	name := e.Param("name")

	details, code, err := c.useCase.GetCouponDetails(name)
	if err != nil {
		return e.JSON(code, map[string]string{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, details)
}
