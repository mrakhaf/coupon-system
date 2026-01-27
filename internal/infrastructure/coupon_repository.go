package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"coupon-system/internal/domain"

	"github.com/google/uuid"
)

type couponRepository struct {
	db *sql.DB
}

func NewCouponRepository(db *sql.DB) domain.CouponRepository {
	return &couponRepository{db: db}
}

func (r *couponRepository) Create(coupon *domain.Coupon) error {
	query := `
		INSERT INTO coupons (id, code, description, discount, min_amount, max_usage, used_count, is_active, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	coupon.ID = uuid.New()
	coupon.CreatedAt = time.Now()
	coupon.UpdatedAt = time.Now()
	coupon.UsedCount = 0
	coupon.IsActive = true

	_, err := r.db.Exec(query, coupon.ID, coupon.Code, coupon.Description, coupon.Discount,
		coupon.MinAmount, coupon.MaxUsage, coupon.UsedCount, coupon.IsActive, coupon.CreatedAt, coupon.UpdatedAt)

	return err
}

func (r *couponRepository) GetByID(id uuid.UUID) (*domain.Coupon, error) {
	query := `SELECT id, code, description, discount, min_amount, max_usage, used_count, is_active, created_at, updated_at FROM coupons WHERE id = ?`

	var coupon domain.Coupon
	err := r.db.QueryRow(query, id).Scan(
		&coupon.ID, &coupon.Code, &coupon.Description, &coupon.Discount,
		&coupon.MinAmount, &coupon.MaxUsage, &coupon.UsedCount, &coupon.IsActive,
		&coupon.CreatedAt, &coupon.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("coupon not found")
		}
		return nil, err
	}

	return &coupon, nil
}

func (r *couponRepository) GetByCode(code string) (*domain.Coupon, error) {
	query := `SELECT id, code, description, discount, min_amount, max_usage, used_count, is_active, created_at, updated_at FROM coupons WHERE code = ?`

	var coupon domain.Coupon
	err := r.db.QueryRow(query, code).Scan(
		&coupon.ID, &coupon.Code, &coupon.Description, &coupon.Discount,
		&coupon.MinAmount, &coupon.MaxUsage, &coupon.UsedCount, &coupon.IsActive,
		&coupon.CreatedAt, &coupon.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("coupon not found")
		}
		return nil, err
	}

	return &coupon, nil
}

func (r *couponRepository) Update(coupon *domain.Coupon) error {
	query := `
		UPDATE coupons SET description = ?, discount = ?, min_amount = ?, max_usage = ?, is_active = ?, updated_at = ?
		WHERE id = ?
	`

	coupon.UpdatedAt = time.Now()

	_, err := r.db.Exec(query, coupon.Description, coupon.Discount, coupon.MinAmount,
		coupon.MaxUsage, coupon.IsActive, coupon.UpdatedAt, coupon.ID)

	return err
}

func (r *couponRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM coupons WHERE id = ?`

	_, err := r.db.Exec(query, id)
	return err
}

func (r *couponRepository) List(limit, offset int) ([]*domain.Coupon, error) {
	query := `
		SELECT id, code, description, discount, min_amount, max_usage, used_count, is_active, created_at, updated_at 
		FROM coupons 
		ORDER BY created_at DESC 
		LIMIT ? OFFSET ?
	`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coupons []*domain.Coupon
	for rows.Next() {
		var coupon domain.Coupon
		err := rows.Scan(
			&coupon.ID, &coupon.Code, &coupon.Description, &coupon.Discount,
			&coupon.MinAmount, &coupon.MaxUsage, &coupon.UsedCount, &coupon.IsActive,
			&coupon.CreatedAt, &coupon.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		coupons = append(coupons, &coupon)
	}

	return coupons, nil
}

func (r *couponRepository) IncrementUsage(id uuid.UUID) error {
	query := `
		UPDATE coupons 
		SET used_count = used_count + 1, updated_at = ?
		WHERE id = ? AND used_count < max_usage
	`

	result, err := r.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("coupon usage limit reached or coupon not found")
	}

	return nil
}
