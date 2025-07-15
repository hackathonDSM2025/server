CREATE TABLE user_visits (
    visit_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    heritage_id INT NOT NULL,
    quiz_completed BOOLEAN DEFAULT FALSE,
    quiz_score INT DEFAULT 0,
    visited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (heritage_id) REFERENCES heritage(heritage_id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_heritage_id (heritage_id),
    UNIQUE KEY unique_user_heritage (user_id, heritage_id)
);