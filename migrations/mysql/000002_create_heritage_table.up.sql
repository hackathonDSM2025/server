CREATE TABLE heritage (
    heritage_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(500),
    qr_code VARCHAR(255) UNIQUE NOT NULL,
    story_text TEXT,
    latitude DECIMAL(10, 8),
    longitude DECIMAL(11, 8),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_qr_code (qr_code),
    INDEX idx_name (name)
);