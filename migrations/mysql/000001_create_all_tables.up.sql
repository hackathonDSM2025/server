-- Heritage Tour Database Schema
-- Create all tables in correct order to handle foreign key constraints

-- 1. Users table
CREATE TABLE users (
    user_id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_username (username)
);

-- 2. Heritage table
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

-- 3. User visits table
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

-- 4. Heritage reviews table
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
    UNIQUE KEY unique_user_heritage_review (user_id, heritage_id)
);

-- 5. Quiz table
CREATE TABLE quiz (
    quiz_id INT AUTO_INCREMENT PRIMARY KEY,
    heritage_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (heritage_id) REFERENCES heritage(heritage_id) ON DELETE CASCADE,
    INDEX idx_heritage_id (heritage_id)
);

-- 6. Questions table
CREATE TABLE questions (
    question_id INT AUTO_INCREMENT PRIMARY KEY,
    quiz_id INT NOT NULL,
    question_text TEXT NOT NULL,
    option1 VARCHAR(255) NOT NULL,
    option2 VARCHAR(255) NOT NULL,
    option3 VARCHAR(255) NOT NULL,
    option4 VARCHAR(255) NOT NULL,
    correct_answer INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (quiz_id) REFERENCES quiz(quiz_id) ON DELETE CASCADE,
    INDEX idx_quiz_id (quiz_id)
);

-- 7. Badges table
CREATE TABLE badges (
    badge_id INT AUTO_INCREMENT PRIMARY KEY,
    heritage_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (heritage_id) REFERENCES heritage(heritage_id) ON DELETE CASCADE,
    INDEX idx_heritage_id (heritage_id)
);

-- 8. User badges table
CREATE TABLE user_badges (
    user_badge_id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    badge_id INT NOT NULL,
    earned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY (badge_id) REFERENCES badges(badge_id) ON DELETE CASCADE,
    INDEX idx_user_id (user_id),
    INDEX idx_badge_id (badge_id),
    UNIQUE KEY unique_user_badge (user_id, badge_id)
);