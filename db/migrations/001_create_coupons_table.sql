-- Create coupons table
CREATE TABLE IF NOT EXISTS coupons (
    id VARCHAR(36) PRIMARY KEY,
    `name` VARCHAR(50) UNIQUE NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    INDEX idx_name (`name`),
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
