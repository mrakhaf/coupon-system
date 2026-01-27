package usecase

import (
	"fmt"
	"net/http"

	"coupon-system/internal/dto/request"
	"coupon-system/internal/dto/response"
	"coupon-system/internal/entity"
	"coupon-system/internal/repository"
	"coupon-system/internal/shared/constant"
)

type CouponUseCase struct {
	repo *repository.CouponRepository
}

func NewCouponUseCase(repo *repository.CouponRepository) *CouponUseCase {
	return &CouponUseCase{repo: repo}
}

func (uc *CouponUseCase) CreateCoupon(req *request.CreateCouponRequest) (*entity.Coupon, error) {
	coupon := &entity.Coupon{
		Name:     req.Name,
		Amount:   req.Amount,
		IsActive: true,
	}

	if err := uc.repo.Create(coupon); err != nil {
		return nil, fmt.Errorf("failed to create coupon: %w", err)
	}

	return coupon, nil
}

func (uc *CouponUseCase) ClaimCoupon(req *request.ClaimCouponRequest) (code int, err error) {

	err = uc.repo.ClaimCoupon(req.UserID, req.CouponName)

	if err != nil {

		code = constant.CodeErrorMessage[err.Error()]

		if code == 0 {
			code = http.StatusInternalServerError
		}

		return code, fmt.Errorf("failed to claim coupon: %w", err)
	}

	return http.StatusCreated, nil
}

func (uc *CouponUseCase) GetCouponDetails(couponName string) (*response.GetCouponDetailsResponse, int, error) {

	coupon, err := uc.repo.GetByName(couponName)

	if err != nil {

		code := constant.CodeErrorMessage[err.Error()]

		if code == 0 {
			code = http.StatusInternalServerError
		}

		return nil, code, fmt.Errorf("failed to get coupon details: %w", err)
	}

	claimedCoupon, err := uc.repo.GetListClaimedCouponByName(couponName)

	if err != nil {

		code := constant.CodeErrorMessage[err.Error()]

		if code == 0 {
			code = http.StatusInternalServerError
		}

		return nil, code, fmt.Errorf("failed to get coupon details: %w", err)
	}

	claimedBy := []string{}
	for _, c := range claimedCoupon {
		claimedBy = append(claimedBy, c.UserID)
	}

	return &response.GetCouponDetailsResponse{
		Name:            coupon.Name,
		Amount:          coupon.Amount,
		RemainingAmount: coupon.Amount - len(claimedCoupon),
		ClaimedBy:       claimedBy,
	}, http.StatusOK, nil

}
