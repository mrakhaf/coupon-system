package repository

import (
	"fmt"

	"coupon-system/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CouponRepository struct {
	db *gorm.DB
}

func NewCouponRepository(db *gorm.DB) *CouponRepository {
	return &CouponRepository{db: db}
}

func (r *CouponRepository) Create(coupon *entity.Coupon) error {
	coupon.ID = uuid.New()
	coupon.UsedCount = 0
	coupon.IsActive = true

	return r.db.Create(coupon).Error
}

func (r *CouponRepository) GetByID(id uuid.UUID) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.First(&coupon, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("coupon not found")
		}
		return nil, err
	}
	return &coupon, nil
}

func (r *CouponRepository) GetByCode(code string) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.First(&coupon, "code = ?", code).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("coupon not found")
		}
		return nil, err
	}
	return &coupon, nil
}

func (r *CouponRepository) GetByName(name string) (*entity.Coupon, error) {
	var coupon entity.Coupon
	err := r.db.First(&coupon, "name = ?", name).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("coupon not found")
		}
		return nil, err
	}
	return &coupon, nil
}

func (r *CouponRepository) HasUserClaimedCoupon(userID, couponName string) (bool, error) {
	var count int64
	err := r.db.Model(&entity.CouponClaim{}).
		Joins("JOIN coupons ON coupon_claims.coupon_id = coupons.id").
		Where("coupon_claims.user_id = ? AND coupons.name = ?", userID, couponName).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *CouponRepository) ClaimCouponTransaction(userID string, couponID uuid.UUID, amount float64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Lock the coupon row for update
		var coupon entity.Coupon
		err := tx.Set("gorm:query_option", "FOR UPDATE").First(&coupon, couponID).Error
		if err != nil {
			return err
		}

		// Check remaining amount
		if coupon.RemainingAmount <= 0 {
			return fmt.Errorf("coupon has no remaining amount")
		}

		// Create claim
		claim := &entity.CouponClaim{
			ID:       uuid.New(),
			UserID:   userID,
			CouponID: couponID,
		}

		if err := tx.Create(claim).Error; err != nil {
			return err
		}

		// Update coupon remaining amount
		coupon.RemainingAmount -= amount
		coupon.UsedCount++

		return tx.Save(&coupon).Error
	})
}

func (r *CouponRepository) GetClaimedByUsers(couponName string) ([]string, error) {
	var users []string
	err := r.db.Model(&entity.CouponClaim{}).
		Joins("JOIN coupons ON coupon_claims.coupon_id = coupons.id").
		Where("coupons.name = ?", couponName).
		Pluck("coupon_claims.user_id", &users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *CouponRepository) Update(coupon *entity.Coupon) error {
	return r.db.Save(coupon).Error
}

func (r *CouponRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entity.Coupon{}, "id = ?", id).Error
}

func (r *CouponRepository) List(limit, offset int) ([]*entity.Coupon, error) {
	var coupons []*entity.Coupon
	err := r.db.Order("created_at DESC").Limit(limit).Offset(offset).Find(&coupons).Error
	if err != nil {
		return nil, err
	}
	return coupons, nil
}

func (r *CouponRepository) IncrementUsage(id uuid.UUID) error {
	result := r.db.Model(&entity.Coupon{}).
		Where("id = ? AND used_count < max_usage", id).
		Update("used_count", gorm.Expr("used_count + 1"))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("coupon usage limit reached or coupon not found")
	}

	return nil
}
