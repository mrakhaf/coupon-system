package usecase

import (
	"fmt"

	"coupon-system/internal/entity"
	"coupon-system/internal/repository"
)

type CouponUseCase struct {
	repo *repository.CouponRepository
}

func NewCouponUseCase(repo *repository.CouponRepository) *CouponUseCase {
	return &CouponUseCase{repo: repo}
}

func (uc *CouponUseCase) CreateCoupon(req *entity.CreateCouponRequest) (*entity.Coupon, error) {
	coupon := &entity.Coupon{
		Name:            req.Name,
		Amount:          req.Amount,
		RemainingAmount: req.Amount,
		MaxUsage:        1,
		UsedCount:       0,
		IsActive:        true,
	}

	if err := uc.repo.Create(coupon); err != nil {
		return nil, fmt.Errorf("failed to create coupon: %w", err)
	}

	return coupon, nil
}

func (uc *CouponUseCase) ClaimCoupon(req *entity.ClaimCouponRequest) error {
	// Check if user has already claimed this coupon
	alreadyClaimed, err := uc.repo.HasUserClaimedCoupon(req.UserID, req.CouponName)
	if err != nil {
		return fmt.Errorf("failed to check claim status: %w", err)
	}

	if alreadyClaimed {
		return fmt.Errorf("user has already claimed this coupon")
	}

	// Check if coupon has remaining amount and is active
	coupon, err := uc.repo.GetByName(req.CouponName)
	if err != nil {
		return fmt.Errorf("coupon not found: %w", err)
	}

	if !coupon.IsActive {
		return fmt.Errorf("coupon is not active")
	}

	if coupon.RemainingAmount <= 0 {
		return fmt.Errorf("coupon has no remaining amount")
	}

	// Use transaction to ensure atomicity
	return uc.repo.ClaimCouponTransaction(req.UserID, coupon.ID, coupon.Amount)
}

func (uc *CouponUseCase) GetCouponDetails(name string) (*entity.CouponDetailsResponse, error) {
	coupon, err := uc.repo.GetByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to get coupon: %w", err)
	}

	claimedBy, err := uc.repo.GetClaimedByUsers(name)
	if err != nil {
		return nil, fmt.Errorf("failed to get claimed users: %w", err)
	}

	return &entity.CouponDetailsResponse{
		Name:            coupon.Name,
		Amount:          coupon.Amount,
		RemainingAmount: coupon.RemainingAmount,
		ClaimedBy:       claimedBy,
	}, nil
}
