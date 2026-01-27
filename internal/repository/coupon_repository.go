package repository

import (
	"fmt"
	"time"

	"coupon-system/internal/entity"
	"coupon-system/internal/shared/constant"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CouponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) *CouponRepository {
	return &CouponRepository{db: db}
}

func (r *CouponRepository) WithTx(tx *gorm.DB) *CouponRepository {
	return &CouponRepository{db: tx}
}

func (r *CouponRepository) Create(coupon *entity.Coupon) error {
	coupon.ID = uuid.New()
	coupon.IsActive = true

	return r.db.Create(coupon).Error
}

func (r *CouponRepository) GetByName(name string) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.First(&coupon, "name = ?", name).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf(constant.ErrCouponNotFound)
		}
		return nil, err
	}
	return &coupon, nil
}

func (r *CouponRepository) ClaimCoupon(userID string, couponName string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		//check if user has already claimed this coupon

		isClaimed, err := r.WithTx(tx).HasUserClaimedCoupon(userID, couponName)
		if err != nil {
			return err
		}

		if isClaimed {
			return fmt.Errorf(constant.ErrUserHasAlreadyClaimedCoupon)
		}

		//check if coupon has remaining amount and is active
		coupon, err := r.WithTx(tx).GetByName(couponName)
		if err != nil {
			return err
		}

		if !coupon.IsActive {
			return fmt.Errorf(constant.ErrCouponNotActive)
		}

		//count coupon claimed with coupon name
		countClaimed, err := r.WithTx(tx).GetCountClaimedCoupon(coupon.ID)
		if err != nil {
			return err
		}

		if int(countClaimed) >= coupon.Amount {
			return fmt.Errorf(constant.ErrCouponNoRemainingAmount)
		}

		//create coupon claim
		couponClaim := &entity.CouponClaim{
			ID:        uuid.New(),
			UserID:    userID,
			CouponID:  coupon.ID,
			CreatedAt: time.Now(),
		}

		if err := tx.Create(couponClaim).Error; err != nil {
			return err
		}

		return nil

	})
}

func (r *CouponRepository) HasUserClaimedCoupon(userID string, couponName string) (bool, error) {
	var count int64

	err := r.db.
		Model(&entity.CouponClaim{}).
		Joins("JOIN coupons ON coupons.id = coupon_claims.coupon_id").
		Where("coupon_claims.user_id = ? AND coupons.name = ?", userID, couponName).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *CouponRepository) GetCountClaimedCoupon(couponID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&entity.CouponClaim{}).Where("coupon_id = ?", couponID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CouponRepository) GetListClaimedCouponByName(couponName string) ([]entity.CouponClaim, error) {
	var couponClaims []entity.CouponClaim
	err := r.db.
		Model(&entity.CouponClaim{}).
		Joins("JOIN coupons ON coupons.id = coupon_claims.coupon_id").
		Where("coupons.name = ?", couponName).
		Find(&couponClaims).Error
	if err != nil {
		return nil, err
	}
	return couponClaims, nil
}
