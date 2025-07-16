-- Insert essential sample data for testing

-- Sample heritage sites (5 main palaces for testing)
INSERT INTO heritage (name, description, image_url, qr_code, story_text, latitude, longitude) VALUES
('경복궁', '조선 왕조의 정궁으로 1395년에 창건되었습니다.', 'https://example.com/images/gyeongbokgung.jpg', 'QR_GYEONGBOKGUNG_001', '경복궁은 조선 태조 4년(1395년)에 창건된 조선 왕조의 정궁입니다. "경복"은 시경에 나오는 말로 "큰 복을 누리고 왕조의 번영을 기원한다"는 의미를 담고 있습니다.', 37.5796, 126.9770),
('창덕궁', '조선 왕조의 이궁으로 1405년에 건립되었습니다.', 'https://example.com/images/changdeokgung.jpg', 'QR_CHANGDEOKGUNG_002', '창덕궁은 조선 태종 5년(1405년)에 경복궁의 이궁으로 건립되었습니다. 자연과 조화를 이룬 아름다운 궁궐로 유네스코 세계문화유산에 등재되어 있습니다.', 37.5794, 126.9910),
('덕수궁', '조선 후기와 대한제국 시대의 궁궐입니다.', 'https://example.com/images/deoksugung.jpg', 'QR_DEOKSUGUNG_003', '덕수궁은 조선 후기와 대한제국 시대의 궁궐로, 전통 건축과 서구식 건축이 조화를 이루고 있는 특별한 궁궐입니다.', 37.5658, 126.9751),
('창경궁', '조선 성종 때 건립된 이궁입니다.', 'https://example.com/images/changgyeonggung.jpg', 'QR_CHANGGYEONGGUNG_004', '창경궁은 조선 성종 14년(1483년)에 세조의 왕비인 정희왕후, 덕종의 왕비인 소혜왕후, 예종의 왕비인 안순왕후를 위해 지어진 궁궐입니다.', 37.5792, 126.9949),
('경희궁', '조선 광해군 때 건립된 이궁입니다.', 'https://example.com/images/gyeonghuigung.jpg', 'QR_GYEONGHUIGUNG_005', '경희궁은 조선 광해군 8년(1616년)에 건립된 이궁으로, 서궐이라고도 불렸습니다. 일제강점기에 많은 건물이 훼손되었다가 복원되었습니다.', 37.5710, 126.9653);

-- Create quiz entries for each heritage site
INSERT INTO quiz (heritage_id) VALUES (1), (2), (3), (4), (5);

-- Sample questions for 경복궁 (quiz_id = 1)
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(1, '경복궁이 건립된 연도는?', '1392년', '1394년', '1395년', '1396년', 3),
(1, '경복궁의 정전은?', '근정전', '인정전', '중화전', '선정전', 1),
(1, '경복궁의 "경복"의 의미는?', '큰 복을 누린다', '영원한 평화', '왕의 덕을 기린다', '백성의 행복', 1);

-- Sample questions for 창덕궁 (quiz_id = 2)
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(2, '창덕궁이 건립된 연도는?', '1405년', '1406년', '1407년', '1408년', 1),
(2, '창덕궁의 특징은?', '서구식 건축', '자연과의 조화', '화려한 장식', '거대한 규모', 2),
(2, '창덕궁은 유네스코 세계문화유산에 언제 등재되었나?', '1995년', '1996년', '1997년', '1998년', 3);

-- Sample questions for 덕수궁 (quiz_id = 3)
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(3, '덕수궁의 특징은?', '순수 전통 건축', '전통과 서구식 건축의 조화', '순수 서구식 건축', '현대식 건축', 2),
(3, '덕수궁에서 즉위한 황제는?', '고종', '순종', '태조', '세종', 1),
(3, '덕수궁의 서구식 건물은?', '중화전', '석조전', '덕홍전', '정관헌', 2);

-- Sample questions for 창경궁 (quiz_id = 4)
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(4, '창경궁이 건립된 연도는?', '1483년', '1484년', '1485년', '1486년', 1),
(4, '창경궁은 누구를 위해 지어졌나?', '왕의 거처', '세 명의 대비를 위해', '신하들을 위해', '백성들을 위해', 2),
(4, '창경궁을 건립한 왕은?', '태조', '세종', '성종', '중종', 3);

-- Sample questions for 경희궁 (quiz_id = 5)
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(5, '경희궁이 건립된 연도는?', '1615년', '1616년', '1617년', '1618년', 2),
(5, '경희궁의 별칭은?', '동궐', '서궐', '남궐', '북궐', 2),
(5, '경희궁을 건립한 왕은?', '선조', '광해군', '인조', '효종', 2);

-- Create badges for each heritage site
INSERT INTO badges (heritage_id, name, description, image_url) VALUES
(1, '경복궁 마스터', '경복궁 퀴즈를 모두 맞혔습니다!', 'https://example.com/badges/gyeongbokgung.png'),
(2, '창덕궁 마스터', '창덕궁 퀴즈를 모두 맞혔습니다!', 'https://example.com/badges/changdeokgung.png'),
(3, '덕수궁 마스터', '덕수궁 퀴즈를 모두 맞혔습니다!', 'https://example.com/badges/deoksugung.png'),
(4, '창경궁 마스터', '창경궁 퀴즈를 모두 맞혔습니다!', 'https://example.com/badges/changgyeonggung.png'),
(5, '경희궁 마스터', '경희궁 퀴즈를 모두 맞혔습니다!', 'https://example.com/badges/gyeonghuigung.png');