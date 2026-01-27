-- Create coupons table
CREATE TABLE IF NOT EXISTS coupons (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    remaining_amount DECIMAL(10,2) NOT NULL,
    max_usage INT NOT NULL,
    used_count INT DEFAULT 0,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_name (name),
    INDEX idx_active (is_active),
    INDEX idx_created_at (created_at)
);

-- Create coupon_claims table for claim history
CREATE TABLE IF NOT EXISTS coupon_claims (
    id VARCHAR(36) PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL,
    coupon_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    INDEX idx_user_id (user_id),
    INDEX idx_coupon_id (coupon_id),
    INDEX idx_created_at (created_at),
    
    -- Enforce uniqueness: user can claim each coupon only once
    UNIQUE KEY uk_user_coupon (user_id, coupon_id),
    
    FOREIGN KEY (coupon_id) REFERENCES coupons(id) ON DELETE CASCADE
);

-- Insert sample data
INSERT INTO coupons (id, name, amount, remaining_amount, max_usage, used_count, is_active) VALUES
('12345678-1234-1234-1234-123456789012', 'PROMO_WELCOME', 100.00, 100.00, 1, 0, TRUE),
('23456789-2345-2345-2345-234567890123', 'PROMO_SUMMER', 200.00, 200.00, 1, 0, TRUE),
('34567890-3456-3456-3456-345678901234', 'PROMO_FREESHIP', 50.00, 50.00, 1, 0, TRUE);
