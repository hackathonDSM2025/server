CREATE TABLE heritage_reviews (
    review_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    heritage_id INT NOT NULL,
    rating INT NOT NULL CHECK (rating >= 1 AND rating <= 5),
    review_text TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (heritage_id) REFERENCES heritage(heritage_id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_heritage_id (heritage_id),
    INDEX idx_created_at (created_at),
    UNIQUE KEY unique_user_heritage_review (user_id, heritage_id)
);