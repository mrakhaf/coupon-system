package handler

import (
	"net/http"
	"strconv"

	"coupon-system/internal/domain"

	"github.com/labstack/echo/v4"
)

type CouponHandler struct {
	useCase domain.CouponUseCase
}

func NewCouponHandler(useCase domain.CouponUseCase) *CouponHandler {
	return &CouponHandler{useCase: useCase}
}

func (h *CouponHandler) CreateCoupon(c echo.Context) error {
	req := new(domain.CreateCouponRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	coupon, err := h.useCase.CreateCoupon(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, coupon)
}

func (h *CouponHandler) GetCouponByID(c echo.Context) error {
	id := c.Param("id")

	coupon, err := h.useCase.GetCouponByID(id)
	if err != nil {
		if err.Error() == "coupon not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) GetCouponByCode(c echo.Context) error {
	code := c.Param("code")

	coupon, err := h.useCase.GetCouponByCode(code)
	if err != nil {
		if err.Error() == "coupon not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) UpdateCoupon(c echo.Context) error {
	id := c.Param("id")

	req := new(domain.UpdateCouponRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	coupon, err := h.useCase.UpdateCoupon(id, req)
	if err != nil {
		if err.Error() == "coupon not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, coupon)
}

func (h *CouponHandler) DeleteCoupon(c echo.Context) error {
	id := c.Param("id")

	err := h.useCase.DeleteCoupon(id)
	if err != nil {
		if err.Error() == "coupon not found" {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Coupon not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *CouponHandler) ListCoupons(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	coupons, err := h.useCase.ListCoupons(limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, coupons)
}

func (h *CouponHandler) UseCoupon(c echo.Context) error {
	req := new(domain.UseCouponRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	response, err := h.useCase.UseCoupon(req.Code, req.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, response)
}
