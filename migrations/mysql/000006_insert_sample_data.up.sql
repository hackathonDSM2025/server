-- Insert sample heritage data
INSERT INTO heritage (name, description, image_url, qr_code, story_text, latitude, longitude) VALUES
('경복궁', '조선 왕조의 정궁으로 1395년에 창건되었습니다.', 'https://example.com/images/gyeongbokgung.jpg', 'QR_GYEONGBOKGUNG_001', '경복궁은 조선 태조 4년(1395년)에 창건된 조선 왕조의 정궁입니다. "경복"은 시경에 나오는 말로 "큰 복을 누리고 왕조의 번영을 기원한다"는 의미를 담고 있습니다.', 37.5796, 126.9770),
('창덕궁', '조선 왕조의 이궁으로 1405년에 건립되었습니다.', 'https://example.com/images/changdeokgung.jpg', 'QR_CHANGDEOKGUNG_002', '창덕궁은 조선 태종 5년(1405년)에 경복궁의 이궁으로 건립되었습니다. 자연과 조화를 이룬 아름다운 궁궐로 유네스코 세계문화유산에 등재되어 있습니다.', 37.5794, 126.9910),
('덕수궁', '조선 후기와 대한제국 시대의 궁궐입니다.', 'https://example.com/images/deoksugung.jpg', 'QR_DEOKSUGUNG_003', '덕수궁은 조선 후기와 대한제국 시대의 궁궐로, 전통 건축과 서구식 건축이 조화를 이루고 있는 특별한 궁궐입니다.', 37.5658, 126.9751),
('종묘', '조선 왕조 역대 왕과 왕비의 신주를 모신 사당입니다.', 'https://example.com/images/jongmyo.jpg', 'QR_JONGMYO_004', '종묘는 조선 왕조 역대 왕과 왕비의 신주를 모신 유교 사당으로, 조선 태조 때 건립되어 500년 이상 제례를 지내고 있습니다.', 37.5743, 126.9936),
('남대문', '서울의 정문 역할을 했던 조선시대 성문입니다.', 'https://example.com/images/namdaemun.jpg', 'QR_NAMDAEMUN_005', '남대문(숭례문)은 조선 태조 7년(1398년)에 축조된 한양도성의 정문으로, 국보 제1호로 지정되어 있습니다.', 37.5601, 126.9754);

-- Insert quiz data
INSERT INTO quiz (heritage_id) VALUES
(1), (2), (3), (4), (5);

-- Insert questions for 경복궁
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(1, '경복궁이 건립된 연도는?', '1392년', '1394년', '1395년', '1396년', 3),
(1, '경복궁의 정전은?', '근정전', '인정전', '중화전', '선정전', 1),
(1, '경복궁의 "경복"의 의미는?', '큰 복을 누린다', '영원한 평화', '왕의 덕을 기린다', '백성의 행복', 1);

-- Insert questions for 창덕궁
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(2, '창덕궁이 건립된 연도는?', '1405년', '1406년', '1407년', '1408년', 1),
(2, '창덕궁의 특징은?', '서구식 건축', '자연과의 조화', '화려한 장식', '거대한 규모', 2),
(2, '창덕궁은 유네스코 세계문화유산에 언제 등재되었나?', '1995년', '1996년', '1997년', '1998년', 3);

-- Insert questions for 덕수궁
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(3, '덕수궁의 특징은?', '순수 전통 건축', '전통과 서구식 건축의 조화', '순수 서구식 건축', '현대식 건축', 2),
(3, '덕수궁에서 즉위한 황제는?', '고종', '순종', '태조', '세종', 1),
(3, '덕수궁의 서구식 건물은?', '중화전', '석조전', '덕홍전', '정관헌', 2);

-- Insert questions for 종묘
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(4, '종묘의 용도는?', '왕의 거처', '신하들의 회의장', '왕과 왕비의 신주를 모신 사당', '군사 훈련장', 3),
(4, '종묘제례는 언제 열리나?', '매월', '매년 5월', '3년마다', '5년마다', 2),
(4, '종묘 제례악은 무엇인가?', '궁중음악', '민속음악', '서구음악', '현대음악', 1);

-- Insert questions for 남대문
INSERT INTO questions (quiz_id, question_text, option1, option2, option3, option4, correct_answer) VALUES
(5, '남대문의 정식 명칭은?', '광화문', '숭례문', '흥인지문', '돈의문', 2),
(5, '남대문이 건립된 연도는?', '1396년', '1397년', '1398년', '1399년', 3),
(5, '남대문은 국보 몇 호인가?', '제1호', '제2호', '제3호', '제4호', 1);

-- Insert badge data
INSERT INTO badges (name, description, image_url, condition_type) VALUES
('퀴즈 마스터', '퀴즈를 성공적으로 완료했습니다!', 'https://example.com/badges/quiz_master.png', 'quiz_completion'),
('문화재 탐험가', '첫 번째 문화재를 방문했습니다!', 'https://example.com/badges/explorer.png', 'first_visit'),
('궁궐 전문가', '모든 궁궐을 방문했습니다!', 'https://example.com/badges/palace_expert.png', 'palace_completion'),
('서울 문화재 지킴이', '서울의 주요 문화재를 모두 탐방했습니다!', 'https://example.com/badges/seoul_guardian.png', 'seoul_completion');