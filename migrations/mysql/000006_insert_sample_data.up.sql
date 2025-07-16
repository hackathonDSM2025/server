-- Insert sample heritage data
INSERT INTO
    heritage (
        name,
        description,
        image_url,
        qr_code,
        story_text,
        latitude,
        longitude
    )
VALUES
    -- 궁궐
    (
        '경복궁',
        '조선 왕조의 정궁으로 1395년에 창건되었습니다.',
        'https://example.com/images/gyeongbokgung.jpg',
        'QR_GYEONGBOKGUNG_001',
        '경복궁은 조선 태조 4년(1395년)에 창건된 조선 왕조의 정궁입니다. "경복"은 시경에 나오는 말로 "큰 복을 누리고 왕조의 번영을 기원한다"는 의미를 담고 있습니다.',
        37.5796,
        126.9770
    ),
    (
        '창덕궁',
        '조선 왕조의 이궁으로 1405년에 건립되었습니다.',
        'https://example.com/images/changdeokgung.jpg',
        'QR_CHANGDEOKGUNG_002',
        '창덕궁은 조선 태종 5년(1405년)에 경복궁의 이궁으로 건립되었습니다. 자연과 조화를 이룬 아름다운 궁궐로 유네스코 세계문화유산에 등재되어 있습니다.',
        37.5794,
        126.9910
    ),
    (
        '덕수궁',
        '조선 후기와 대한제국 시대의 궁궐입니다.',
        'https://example.com/images/deoksugung.jpg',
        'QR_DEOKSUGUNG_003',
        '덕수궁은 조선 후기와 대한제국 시대의 궁궐로, 전통 건축과 서구식 건축이 조화를 이루고 있는 특별한 궁궐입니다.',
        37.5658,
        126.9751
    ),
    (
        '창경궁',
        '조선 성종 때 건립된 이궁입니다.',
        'https://example.com/images/changgyeonggung.jpg',
        'QR_CHANGGYEONGGUNG_004',
        '창경궁은 조선 성종 14년(1483년)에 세조의 왕비인 정희왕후, 덕종의 왕비인 소혜왕후, 예종의 왕비인 안순왕후를 위해 지어진 궁궐입니다.',
        37.5792,
        126.9949
    ),
    (
        '경희궁',
        '조선 광해군 때 건립된 이궁입니다.',
        'https://example.com/images/gyeonghuigung.jpg',
        'QR_GYEONGHUIGUNG_005',
        '경희궁은 조선 광해군 8년(1616년)에 건립된 이궁으로, 서궐이라고도 불렸습니다. 일제강점기에 많은 건물이 훼손되었다가 복원되었습니다.',
        37.5710,
        126.9653
    ),

-- 성문과 성곽
(
    '남대문',
    '서울의 정문 역할을 했던 조선시대 성문입니다.',
    'https://example.com/images/namdaemun.jpg',
    'QR_NAMDAEMUN_006',
    '남대문(숭례문)은 조선 태조 7년(1398년)에 축조된 한양도성의 정문으로, 국보 제1호로 지정되어 있습니다.',
    37.5601,
    126.9754
),
(
    '동대문',
    '한양도성의 동쪽 대문입니다.',
    'https://example.com/images/dongdaemun.jpg',
    'QR_DONGDAEMUN_007',
    '동대문(흥인지문)은 조선 태조 5년(1396년)에 건립된 한양도성의 동쪽 대문으로, 보물 제1호로 지정되어 있습니다.',
    37.5708,
    127.0094
),
(
    '서대문',
    '한양도성의 서쪽 대문이었습니다.',
    'https://example.com/images/seodaemun.jpg',
    'QR_SEODAEMUN_008',
    '서대문(돈의문)은 조선시대 한양도성의 서쪽 대문이었으나 일제강점기에 철거되었고, 현재는 돈의문박물관마을이 조성되어 있습니다.',
    37.5699,
    126.9656
),
(
    '낙산성곽',
    '한양도성의 동북쪽 성곽입니다.',
    'https://example.com/images/naksan.jpg',
    'QR_NAKSAN_009',
    '낙산성곽은 조선시대 한양도성의 동북쪽 구간으로, 현재도 원형이 잘 보존되어 있어 성곽길로 인기가 높습니다.',
    37.5812,
    127.0066
),
(
    '인왕산성곽',
    '한양도성의 서북쪽 성곽입니다.',
    'https://example.com/images/inwangsan.jpg',
    'QR_INWANGSAN_010',
    '인왕산성곽은 조선시대 한양도성의 서북쪽 구간으로, 자연 암반을 이용해 축조된 독특한 구조를 보여줍니다.',
    37.5825,
    126.9567
),

-- 사당과 종교 건축
(
    '종묘',
    '조선 왕조 역대 왕과 왕비의 신주를 모신 사당입니다.',
    'https://example.com/images/jongmyo.jpg',
    'QR_JONGMYO_011',
    '종묘는 조선 왕조 역대 왕과 왕비의 신주를 모신 유교 사당으로, 조선 태조 때 건립되어 500년 이상 제례를 지내고 있습니다.',
    37.5743,
    126.9936
),
(
    '사직단',
    '토지신과 곡물신에게 제사를 지내던 제단입니다.',
    'https://example.com/images/sajikdan.jpg',
    'QR_SAJIKDAN_012',
    '사직단은 조선시대 국가에서 토지신(사신)과 곡물신(직신)에게 제사를 지내던 제단으로, 현재 사직공원 내에 위치합니다.',
    37.5748,
    126.9717
),
(
    '문묘',
    '공자를 모신 유교 사당입니다.',
    'https://example.com/images/munmyo.jpg',
    'QR_MUNMYO_013',
    '문묘는 성균관 내에 있는 공자와 중국 및 우리나라 성현들의 위패를 모신 유교 사당으로, 매년 석전대제가 열립니다.',
    37.5940,
    126.9963
),
(
    '조계사',
    '한국 불교 조계종의 본산입니다.',
    'https://example.com/images/jogyesa.jpg',
    'QR_JOGYESA_014',
    '조계사는 대한불교조계종의 본산으로 1395년 창건되었으며, 서울 시내 중심가에 위치한 대표적인 불교 사찰입니다.',
    37.5719,
    126.9826
),
(
    '봉은사',
    '강남 지역의 대표적인 불교 사찰입니다.',
    'https://example.com/images/bongeunsa.jpg',
    'QR_BONGEUNSA_015',
    '봉은사는 신라 원성왕 10년(794년)에 창건된 사찰로, 현재 강남 코엑스 인근에 위치한 도심 속 사찰입니다.',
    37.5144,
    127.0592
),

-- 왕릉
(
    '선정릉',
    '조선 성종과 중종의 능입니다.',
    'https://example.com/images/seonjeongneung.jpg',
    'QR_SEONJEONGNEUNG_016',
    '선정릉은 조선 제9대 성종과 제11대 중종의 능으로, 강남 중심가에 위치한 유일한 조선왕릉입니다.',
    37.5048,
    127.0471
),
(
    '태릉',
    '조선 문정왕후의 능입니다.',
    'https://example.com/images/taereung.jpg',
    'QR_TAEREUNG_017',
    '태릉은 조선 중종의 계비인 문정왕후 윤씨의 능으로, 현재 태릉선수촌과 함께 위치해 있습니다.',
    37.6219,
    127.0777
),
(
    '의릉',
    '조선 경종의 능입니다.',
    'https://example.com/images/uireung.jpg',
    'QR_UIREUNG_018',
    '의릉은 조선 제20대 경종과 선의왕후의 능으로, 성북구에 위치한 조선왕릉입니다.',
    37.6064,
    127.0190
),

-- 역사적 건물과 유적
(
    '광화문',
    '경복궁의 정문입니다.',
    'https://example.com/images/gwanghwamun.jpg',
    'QR_GWANGHWAMUN_019',
    '광화문은 경복궁의 정문으로, "백성을 가르쳐 깨우친다"는 의미를 담고 있으며 현재 위치는 3번째 자리입니다.',
    37.5759,
    126.9768
),
(
    '보신각',
    '조선시대 종루였던 건물입니다.',
    'https://example.com/images/bosingak.jpg',
    'QR_BOSINGAK_020',
    '보신각은 조선시대 종루로 시간을 알리고 성문 개폐를 알리던 곳이며, 현재도 새해 타종행사가 열립니다.',
    37.5698,
    126.9827
),
(
    '동묘',
    '관우를 모신 사당입니다.',
    'https://example.com/images/dongmyo.jpg',
    'QR_DONGMYO_021',
    '동묘는 중국 촉한의 장군 관우를 모신 사당으로, 조선 선조 34년(1601년)에 건립되었습니다.',
    37.5741,
    127.0172
),
(
    '성균관',
    '조선시대 최고 교육기관이었습니다.',
    'https://example.com/images/sungkyunkwan.jpg',
    'QR_SUNGKYUNKWAN_022',
    '성균관은 조선시대 최고 교육기관으로 유학을 가르쳤으며, 현재 성균관대학교의 전신입니다.',
    37.5937,
    126.9964
),
(
    '탑골공원',
    '조선시대 원각사지입니다.',
    'https://example.com/images/tapgol.jpg',
    'QR_TAPGOL_023',
    '탑골공원은 조선시대 원각사가 있던 자리로, 원각사지 십층석탑과 3.1운동 기념비가 있습니다.',
    37.5707,
    126.9857
),

-- 근현대 역사 유적
(
    '독립문',
    '독립협회가 건립한 기념문입니다.',
    'https://example.com/images/dongnimmun.jpg',
    'QR_DONGNIMMUN_024',
    '독립문은 1896년 독립협회가 자주독립 의지를 상징하기 위해 건립한 기념문입니다.',
    37.5743,
    126.9589
),
(
    '서대문형무소',
    '일제강점기 독립투사들이 수감되었던 감옥입니다.',
    'https://example.com/images/seodaemun_prison.jpg',
    'QR_PRISON_025',
    '서대문형무소는 1908년 일제가 건설한 근대식 감옥으로, 수많은 독립투사들이 수감되었던 역사의 현장입니다.',
    37.5742,
    126.9594
),
(
    '명동성당',
    '한국 천주교의 모교회입니다.',
    'https://example.com/images/myeongdong_cathedral.jpg',
    'QR_CATHEDRAL_026',
    '명동성당은 1898년 완공된 한국 천주교의 모교회로, 고딕 양식의 아름다운 건축물입니다.',
    37.5636,
    126.9874
),
(
    '이화장',
    '이승만 대통령의 사저였습니다.',
    'https://example.com/images/ihwajang.jpg',
    'QR_IHWAJANG_027',
    '이화장은 이승만 대통령이 1948년부터 1960년까지 거주했던 사저로, 현재 기념관으로 운영되고 있습니다.',
    37.5838,
    126.9829
),
(
    '백범김구기념관',
    '백범 김구 선생을 기리는 기념관입니다.',
    'https://example.com/images/kimgu_memorial.jpg',
    'QR_KIMGU_028',
    '백범김구기념관은 독립운동가 김구 선생의 생애와 사상을 기리기 위해 효창공원 내에 건립된 기념관입니다.',
    37.5423,
    126.9615
),

-- 전통 건축과 가옥
(
    '운현궁',
    '흥선대원군의 사저였습니다.',
    'https://example.com/images/unhyeongung.jpg',
    'QR_UNHYEONGUNG_029',
    '운현궁은 조선 말기 흥선대원군이 거주했던 사저로, 고종이 즉위하기 전까지 생활했던 곳입니다.',
    37.5779,
    126.9856
),
(
    '석파정',
    '조선 후기의 별장 건축입니다.',
    'https://example.com/images/seokpajeong.jpg',
    'QR_SEOKPAJEONG_030',
    '석파정은 조선 말기 흥선대원군의 별장으로 사용되었던 곳으로, 아름다운 한국 전통 정원을 보여줍니다.',
    37.5511,
    126.9432
),
(
    '북촌한옥마을',
    '조선시대 양반들이 거주했던 한옥 마을입니다.',
    'https://example.com/images/bukchon.jpg',
    'QR_BUKCHON_031',
    '북촌한옥마을은 조선시대부터 양반들이 거주했던 지역으로, 현재도 600여 채의 전통 한옥이 보존되어 있습니다.',
    37.5816,
    126.9831
),
(
    '인사동',
    '전통 문화의 거리입니다.',
    'https://example.com/images/insadong.jpg',
    'QR_INSADONG_032',
    '인사동은 조선시대부터 문화와 예술의 중심지였으며, 현재도 전통 문화와 현대가 조화를 이루는 문화의 거리입니다.',
    37.5719,
    126.9854
),

-- 공원과 자연 문화재
(
    '남산',
    '서울의 중심에 위치한 역사적 명산입니다.',
    'https://example.com/images/namsan.jpg',
    'QR_NAMSAN_033',
    '남산은 조선시대부터 서울의 진산으로 여겨져 왔으며, 현재는 남산공원과 N서울타워가 있는 서울의 상징입니다.',
    37.5512,
    126.9882
),
(
    '한강',
    '서울을 관통하는 젖줄 같은 강입니다.',
    'https://example.com/images/hangang.jpg',
    'QR_HANGANG_034',
    '한강은 서울을 동서로 관통하는 강으로, 조선시대부터 서울 발전의 젖줄 역할을 해왔습니다.',
    37.5219,
    126.9417
),
(
    '청계천',
    '복원된 서울의 역사적 하천입니다.',
    'https://example.com/images/cheonggyecheon.jpg',
    'QR_CHEONGGYECHEON_035',
    '청계천은 조선시대부터 서울 중심가를 흐르던 하천으로, 2005년 복원되어 현재는 시민들의 휴식 공간이 되었습니다.',
    37.5693,
    126.9777
),

-- 근대 건축물
(
    '구 서울역사',
    '일제강점기에 건설된 경성역사입니다.',
    'https://example.com/images/old_seoul_station.jpg',
    'QR_OLD_STATION_036',
    '구 서울역사는 1925년 경성역으로 건설된 근대 건축물로, 현재는 문화역서울284로 활용되고 있습니다.',
    37.5564,
    126.9706
),
(
    '한국은행 본관',
    '일제강점기 조선은행 본점이었습니다.',
    'https://example.com/images/bank_of_korea.jpg',
    'QR_BOK_037',
    '한국은행 본관은 1912년 조선은행 본점으로 건설된 르네상스 양식의 근대 건축물로, 현재 화폐박물관이 있습니다.',
    37.5608,
    126.9784
),
(
    '서울시청',
    '일제강점기 경성부청사였습니다.',
    'https://example.com/images/seoul_city_hall.jpg',
    'QR_CITY_HALL_038',
    '구 서울시청 건물은 1926년 경성부청사로 건설된 근대 건축물로, 현재는 서울도서관으로 사용되고 있습니다.',
    37.5663,
    126.9779
),

-- 교육 문화재
(
    '이화여대 대강당',
    '근대 건축의 걸작입니다.',
    'https://example.com/images/ewha_auditorium.jpg',
    'QR_EWHA_039',
    '이화여자대학교 대강당은 1935년 건립된 근대 건축물로, 한국 근대 건축사의 중요한 유산입니다.',
    37.5595,
    126.9456
),
(
    '연세대학교 언더우드관',
    '선교사가 설립한 근대 교육기관입니다.',
    'https://example.com/images/underwood_hall.jpg',
    'QR_UNDERWOOD_040',
    '연세대학교 언더우드관은 1924년 건립된 근대 건축물로, 한국 근대 교육의 상징적 공간입니다.',
    37.5666,
    126.9384
);

-- Insert quiz data for all heritage sites
INSERT INTO
    quiz (heritage_id)
VALUES (1),
    (2),
    (3),
    (4),
    (5),
    (6),
    (7),
    (8),
    (9),
    (10),
    (11),
    (12),
    (13),
    (14),
    (15),
    (16),
    (17),
    (18),
    (19),
    (20),
    (21),
    (22),
    (23),
    (24),
    (25),
    (26),
    (27),
    (28),
    (29),
    (30),
    (31),
    (32),
    (33),
    (34),
    (35),
    (36),
    (37),
    (38),
    (39),
    (40);

-- Insert questions for 경복궁
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        1,
        '경복궁이 건립된 연도는?',
        '1392년',
        '1394년',
        '1395년',
        '1396년',
        3
    ),
    (
        1,
        '경복궁의 정전은?',
        '근정전',
        '인정전',
        '중화전',
        '선정전',
        1
    ),
    (
        1,
        '경복궁의 "경복"의 의미는?',
        '큰 복을 누린다',
        '영원한 평화',
        '왕의 덕을 기린다',
        '백성의 행복',
        1
    );

-- Insert questions for 창덕궁
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        2,
        '창덕궁이 건립된 연도는?',
        '1405년',
        '1406년',
        '1407년',
        '1408년',
        1
    ),
    (
        2,
        '창덕궁의 특징은?',
        '서구식 건축',
        '자연과의 조화',
        '화려한 장식',
        '거대한 규모',
        2
    ),
    (
        2,
        '창덕궁은 유네스코 세계문화유산에 언제 등재되었나?',
        '1995년',
        '1996년',
        '1997년',
        '1998년',
        3
    );

-- Insert questions for 덕수궁
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        3,
        '덕수궁의 특징은?',
        '순수 전통 건축',
        '전통과 서구식 건축의 조화',
        '순수 서구식 건축',
        '현대식 건축',
        2
    ),
    (
        3,
        '덕수궁에서 즉위한 황제는?',
        '고종',
        '순종',
        '태조',
        '세종',
        1
    ),
    (
        3,
        '덕수궁의 서구식 건물은?',
        '중화전',
        '석조전',
        '덕홍전',
        '정관헌',
        2
    );

-- Insert questions for 종묘
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        4,
        '종묘의 용도는?',
        '왕의 거처',
        '신하들의 회의장',
        '왕과 왕비의 신주를 모신 사당',
        '군사 훈련장',
        3
    ),
    (
        4,
        '종묘제례는 언제 열리나?',
        '매월',
        '매년 5월',
        '3년마다',
        '5년마다',
        2
    ),
    (
        4,
        '종묘 제례악은 무엇인가?',
        '궁중음악',
        '민속음악',
        '서구음악',
        '현대음악',
        1
    );

-- Insert questions for 남대문
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        6,
        '남대문의 정식 명칭은?',
        '광화문',
        '숭례문',
        '흥인지문',
        '돈의문',
        2
    ),
    (
        6,
        '남대문이 건립된 연도는?',
        '1396년',
        '1397년',
        '1398년',
        '1399년',
        3
    ),
    (
        6,
        '남대문은 국보 몇 호인가?',
        '제1호',
        '제2호',
        '제3호',
        '제4호',
        1
    );

-- Insert questions for 창경궁
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        4,
        '창경궁이 건립된 연도는?',
        '1483년',
        '1484년',
        '1485년',
        '1486년',
        1
    ),
    (
        4,
        '창경궁은 누구를 위해 지어졌나?',
        '왕의 거처',
        '세 명의 대비를 위해',
        '신하들을 위해',
        '백성들을 위해',
        2
    ),
    (
        4,
        '창경궁을 건립한 왕은?',
        '태조',
        '세종',
        '성종',
        '중종',
        3
    );

-- Insert questions for 경희궁
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        5,
        '경희궁이 건립된 연도는?',
        '1615년',
        '1616년',
        '1617년',
        '1618년',
        2
    ),
    (
        5,
        '경희궁의 별칭은?',
        '동궐',
        '서궐',
        '남궐',
        '북궐',
        2
    ),
    (
        5,
        '경희궁을 건립한 왕은?',
        '선조',
        '광해군',
        '인조',
        '효종',
        2
    );

-- Insert questions for 동대문
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        7,
        '동대문의 정식 명칭은?',
        '광화문',
        '숭례문',
        '흥인지문',
        '돈의문',
        3
    ),
    (
        7,
        '동대문은 보물 몇 호인가?',
        '제1호',
        '제2호',
        '제3호',
        '제4호',
        1
    ),
    (
        7,
        '동대문이 건립된 연도는?',
        '1395년',
        '1396년',
        '1397년',
        '1398년',
        2
    );

-- Insert questions for 서대문
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        8,
        '서대문의 정식 명칭은?',
        '광화문',
        '숭례문',
        '흥인지문',
        '돈의문',
        4
    ),
    (
        8,
        '서대문은 언제 철거되었나?',
        '일제강점기',
        '한국전쟁 중',
        '1960년대',
        '1980년대',
        1
    ),
    (
        8,
        '현재 서대문 자리에는 무엇이 있나?',
        '공원',
        '박물관마을',
        '상가',
        '아파트',
        2
    );

-- Insert questions for 낙산성곽
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        9,
        '낙산성곽은 한양도성의 어느 방향인가?',
        '동북쪽',
        '동남쪽',
        '서북쪽',
        '서남쪽',
        1
    ),
    (
        9,
        '낙산성곽의 현재 용도는?',
        '군사시설',
        '성곽길',
        '주거지역',
        '상업지역',
        2
    ),
    (
        9,
        '낙산성곽이 위치한 산은?',
        '남산',
        '인왕산',
        '낙산',
        '북악산',
        3
    );

-- Insert questions for 인왕산성곽
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        10,
        '인왕산성곽은 한양도성의 어느 방향인가?',
        '동북쪽',
        '동남쪽',
        '서북쪽',
        '서남쪽',
        3
    ),
    (
        10,
        '인왕산성곽의 특징은?',
        '평지에 건설',
        '자연 암반 이용',
        '인공 재료만 사용',
        '목재로 건설',
        2
    ),
    (
        10,
        '인왕산성곽이 위치한 산은?',
        '남산',
        '인왕산',
        '낙산',
        '북악산',
        2
    );

-- Insert questions for 사직단
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        12,
        '사직단에서 모시는 신은?',
        '왕의 조상신',
        '토지신과 곡물신',
        '바다의 신',
        '하늘의 신',
        2
    ),
    (
        12,
        '사직단이 현재 위치한 곳은?',
        '종묘공원',
        '사직공원',
        '탑골공원',
        '효창공원',
        2
    ),
    (
        12,
        '사직의 의미는?',
        '사(社)는 토지, 직(稷)은 곡물',
        '사(士)는 선비, 직(直)은 정직',
        '사(四)는 사방, 직(直)은 바름',
        '사(寺)는 절, 직(職)은 직업',
        1
    );

-- Insert questions for 문묘
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        13,
        '문묘에서 모시는 성인은?',
        '석가모니',
        '공자',
        '노자',
        '맹자',
        2
    ),
    (
        13,
        '문묘가 위치한 곳은?',
        '종묘',
        '성균관',
        '경복궁',
        '창덕궁',
        2
    ),
    (
        13,
        '문묘에서 열리는 제례는?',
        '종묘제례',
        '석전대제',
        '사직제례',
        '영녕전제례',
        2
    );

-- Insert questions for 조계사
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        14,
        '조계사는 어느 종파의 본산인가?',
        '천태종',
        '조계종',
        '진각종',
        '원불교',
        2
    ),
    (
        14,
        '조계사가 창건된 연도는?',
        '1394년',
        '1395년',
        '1396년',
        '1397년',
        2
    ),
    (
        14,
        '조계사의 위치는?',
        '인사동',
        '종로',
        '명동',
        '강남',
        2
    );

-- Insert questions for 봉은사
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        15,
        '봉은사가 창건된 시대는?',
        '백제',
        '고구려',
        '신라',
        '고려',
        3
    ),
    (
        15,
        '봉은사가 현재 위치한 지역은?',
        '종로구',
        '중구',
        '강남구',
        '서초구',
        3
    ),
    (
        15,
        '봉은사 근처에 있는 시설은?',
        '롯데월드',
        '코엑스',
        '용산역',
        '홍대',
        2
    );

-- Insert questions for 선정릉
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        16,
        '선정릉에 모셔진 왕은?',
        '태조와 세종',
        '성종과 중종',
        '세조와 예종',
        '숙종과 영조',
        2
    ),
    (
        16,
        '선정릉의 특징은?',
        '산속에 위치',
        '강남 도심에 위치',
        '바닷가에 위치',
        '섬에 위치',
        2
    ),
    (
        16,
        '선정릉은 조선왕릉 중 어떤 특징이 있나?',
        '가장 큰 규모',
        '도심에 위치한 유일한 왕릉',
        '가장 오래된 왕릉',
        '가장 높은 곳에 위치',
        2
    );

-- Insert questions for 태릉
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        17,
        '태릉에 모셔진 분은?',
        '성종',
        '문정왕후',
        '중종',
        '인종',
        2
    ),
    (
        17,
        '문정왕후는 어느 왕의 왕비인가?',
        '성종',
        '중종',
        '인종',
        '명종',
        2
    ),
    (
        17,
        '태릉 근처에 있는 시설은?',
        '올림픽공원',
        '태릉선수촌',
        '잠실종합운동장',
        '상암월드컵경기장',
        2
    );

-- Insert questions for 의릉
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        18,
        '의릉에 모셌진 왕은?',
        '숙종',
        '경종',
        '영조',
        '정조',
        2
    ),
    (
        18,
        '경종은 조선 몇 대 왕인가?',
        '19대',
        '20대',
        '21대',
        '22대',
        2
    ),
    (
        18,
        '의릉이 위치한 구는?',
        '종로구',
        '성북구',
        '강북구',
        '노원구',
        2
    );

-- Insert questions for 광화문
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        19,
        '광화문은 어느 궁궐의 정문인가?',
        '창덕궁',
        '경복궁',
        '덕수궁',
        '창경궁',
        2
    ),
    (
        19,
        '광화문의 의미는?',
        '큰 문',
        '밝은 문',
        '백성을 가르쳐 깨우친다',
        '왕의 위엄',
        3
    ),
    (
        19,
        '현재 광화문의 위치는?',
        '원래 자리',
        '두 번째 자리',
        '세 번째 자리',
        '네 번째 자리',
        3
    );

-- Insert questions for 보신각
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        20,
        '보신각의 원래 용도는?',
        '관청',
        '종루',
        '창고',
        '주거',
        2
    ),
    (
        20,
        '보신각에서 현재 열리는 행사는?',
        '전통혼례식',
        '새해 타종행사',
        '음악회',
        '전시회',
        2
    ),
    (
        20,
        '보신각에서 종을 치는 목적은?',
        '시간 알림과 성문 개폐',
        '종교 의식',
        '음악 연주',
        '장식용',
        1
    );

-- Insert questions for 동묘
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        21,
        '동묘에서 모시는 인물은?',
        '공자',
        '관우',
        '제갈량',
        '조조',
        2
    ),
    (
        21,
        '관우는 어느 나라의 장군인가?',
        '위나라',
        '촉한',
        '오나라',
        '진나라',
        2
    ),
    (
        21,
        '동묘가 건립된 연도는?',
        '1600년',
        '1601년',
        '1602년',
        '1603년',
        2
    );

-- Insert questions for 성균관
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        22,
        '성균관의 용도는?',
        '왕의 거처',
        '최고 교육기관',
        '종교시설',
        '군사시설',
        2
    ),
    (
        22,
        '성균관에서 가르친 학문은?',
        '불교',
        '도교',
        '유학',
        '천주교',
        3
    ),
    (
        22,
        '성균관의 후신은?',
        '서울대학교',
        '연세대학교',
        '고려대학교',
        '성균관대학교',
        4
    );

-- Insert questions for 탑골공원
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        23,
        '탑골공원의 원래 이름은?',
        '종묘공원',
        '원각사지',
        '조계사지',
        '봉은사지',
        2
    ),
    (
        23,
        '탑골공원에 있는 탑은?',
        '다보탑',
        '원각사지 십층석탑',
        '경천사지 십층석탑',
        '불국사 삼층석탑',
        2
    ),
    (
        23,
        '탑골공원과 관련된 역사적 사건은?',
        '임진왜란',
        '정유재란',
        '3.1운동',
        '6.25전쟁',
        3
    );

-- Insert questions for 독립문
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        24,
        '독립문을 건립한 단체는?',
        '독립협회',
        '신민회',
        '대한민국임시정부',
        '광복군',
        1
    ),
    (
        24,
        '독립문이 건립된 연도는?',
        '1895년',
        '1896년',
        '1897년',
        '1898년',
        2
    ),
    (
        24,
        '독립문의 건립 목적은?',
        '일본에 대한 저항',
        '자주독립 의지 표현',
        '중국에 대한 독립',
        '러시아에 대한 저항',
        2
    );

-- Insert questions for 서대문형무소
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        25,
        '서대문형무소가 건설된 연도는?',
        '1907년',
        '1908년',
        '1909년',
        '1910년',
        2
    ),
    (
        25,
        '서대문형무소를 건설한 주체는?',
        '조선 정부',
        '일제',
        '미군정',
        '대한민국 정부',
        2
    ),
    (
        25,
        '서대문형무소에 수감된 인물들은?',
        '일반 범죄자만',
        '독립투사들',
        '일본인들',
        '외국인들',
        2
    );

-- Insert questions for 명동성당
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        26,
        '명동성당이 완공된 연도는?',
        '1897년',
        '1898년',
        '1899년',
        '1900년',
        2
    ),
    (
        26,
        '명동성당의 건축 양식은?',
        '로마네스크',
        '고딕',
        '바로크',
        '르네상스',
        2
    ),
    (
        26,
        '명동성당의 지위는?',
        '일반 성당',
        '한국 천주교 모교회',
        '대성당',
        '수도원',
        2
    );

-- Insert questions for 이화장
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        27,
        '이화장에 거주했던 대통령은?',
        '이승만',
        '윤보선',
        '박정희',
        '전두환',
        1
    ),
    (
        27,
        '이승만이 이화장에 거주한 기간은?',
        '1945-1960년',
        '1948-1960년',
        '1950-1960년',
        '1953-1960년',
        2
    ),
    (
        27,
        '이화장의 현재 용도는?',
        '대통령 관저',
        '기념관',
        '일반 주택',
        '정부청사',
        2
    );

-- Insert questions for 백범김구기념관
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        28,
        '백범김구기념관이 위치한 공원은?',
        '탑골공원',
        '효창공원',
        '사직공원',
        '남산공원',
        2
    ),
    (
        28,
        '김구 선생의 호는?',
        '단재',
        '백범',
        '도산',
        '우당',
        2
    ),
    (
        28,
        '김구 선생의 주요 활동은?',
        '교육사업',
        '독립운동',
        '종교활동',
        '문학활동',
        2
    );

-- Insert questions for 운현궁
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        29,
        '운현궁에 거주했던 인물은?',
        '흥선대원군',
        '명성황후',
        '고종',
        '순종',
        1
    ),
    (
        29,
        '고종이 운현궁에서 생활한 시기는?',
        '즉위 후',
        '즉위 전',
        '퇴위 후',
        '유년기만',
        2
    ),
    (
        29,
        '운현궁의 현재 상태는?',
        '완전 멸실',
        '일부 복원',
        '원형 보존',
        '현대식 재건',
        2
    );

-- Insert questions for 석파정
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        30,
        '석파정을 사용했던 인물은?',
        '고종',
        '흥선대원군',
        '명성황후',
        '순종',
        2
    ),
    (
        30,
        '석파정의 용도는?',
        '본궁',
        '별장',
        '관청',
        '학교',
        2
    ),
    (
        30,
        '석파정의 특징은?',
        '서구식 건축',
        '전통 정원',
        '종교적 공간',
        '군사시설',
        2
    );

-- Insert questions for 북촌한옥마을
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        31,
        '북촌한옥마을에 거주했던 계층은?',
        '평민',
        '양반',
        '상인',
        '수공업자',
        2
    ),
    (
        31,
        '현재 북촌에 보존된 한옥의 수는?',
        '약 400채',
        '약 600채',
        '약 800채',
        '약 1000채',
        2
    ),
    (
        31,
        '북촌한옥마을의 위치는?',
        '경복궁과 창덕궁 사이',
        '덕수궁 근처',
        '창경궁 근처',
        '경희궁 근처',
        1
    );

-- Insert questions for 인사동
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        32,
        '인사동의 전통적 특징은?',
        '상업 중심지',
        '문화와 예술의 중심지',
        '정치 중심지',
        '종교 중심지',
        2
    ),
    (
        32,
        '인사동에서 주로 거래되는 것은?',
        '식료품',
        '전통 문화재와 예술품',
        '의류',
        '전자제품',
        2
    ),
    (
        32,
        '인사동의 대표적인 거리는?',
        '명동길',
        '인사동길',
        '강남대로',
        '세종대로',
        2
    );

-- Insert questions for 남산
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        33,
        '남산이 조선시대에 불린 이름은?',
        '목멱산',
        '인왕산',
        '낙산',
        '북악산',
        1
    ),
    (
        33,
        '남산에 있는 대표적인 시설은?',
        'N서울타워',
        '63빌딩',
        '롯데타워',
        '동대문디자인플라자',
        1
    ),
    (
        33,
        '남산의 서울에서의 지위는?',
        '주산',
        '진산',
        '안산',
        '조산',
        2
    );

-- Insert questions for 한강
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        34,
        '한강이 서울을 흐르는 방향은?',
        '남북',
        '동서',
        '남동-북서',
        '남서-북동',
        2
    ),
    (
        34,
        '조선시대 한강의 역할은?',
        '군사적 요새',
        '교통과 물류의 중심',
        '종교적 성지',
        '왕실 휴양지',
        2
    ),
    (
        34,
        '한강의 서울 발전에서의 의미는?',
        '장벽 역할',
        '발전의 젖줄',
        '경계선',
        '관광자원만',
        2
    );

-- Insert questions for 청계천
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        35,
        '청계천이 복원된 연도는?',
        '2003년',
        '2004년',
        '2005년',
        '2006년',
        3
    ),
    (
        35,
        '청계천의 조선시대 역할은?',
        '식수 공급',
        '시내 하천',
        '해자',
        '관개용수',
        2
    ),
    (
        35,
        '청계천 복원 전 이 자리에 있던 것은?',
        '도로',
        '청계고가도로',
        '지하철',
        '공원',
        2
    );

-- Insert questions for 구 서울역사
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        36,
        '구 서울역사의 원래 이름은?',
        '서울역',
        '경성역',
        '용산역',
        '중앙역',
        2
    ),
    (
        36,
        '구 서울역사가 건설된 연도는?',
        '1923년',
        '1924년',
        '1925년',
        '1926년',
        3
    ),
    (
        36,
        '현재 구 서울역사의 용도는?',
        '기차역',
        '문화역서울284',
        '박물관',
        '쇼핑몰',
        2
    );

-- Insert questions for 한국은행 본관
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        37,
        '한국은행 본관의 원래 이름은?',
        '한국은행',
        '조선은행',
        '식민지은행',
        '경성은행',
        2
    ),
    (
        37,
        '건물이 건설된 연도는?',
        '1910년',
        '1911년',
        '1912년',
        '1913년',
        3
    ),
    (
        37,
        '현재 이 건물에 있는 시설은?',
        '은행 업무만',
        '화폐박물관',
        '쇼핑센터',
        '호텔',
        2
    );

-- Insert questions for 서울시청
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        38,
        '구 서울시청 건물의 원래 용도는?',
        '서울시청',
        '경성부청사',
        '총독부',
        '법원',
        2
    ),
    (
        38,
        '건물이 건설된 연도는?',
        '1924년',
        '1925년',
        '1926년',
        '1927년',
        3
    ),
    (
        38,
        '현재 구 시청 건물의 용도는?',
        '시청 업무',
        '서울도서관',
        '박물관',
        '문화센터',
        2
    );

-- Insert questions for 이화여대 대강당
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        39,
        '이화여대 대강당이 건립된 연도는?',
        '1933년',
        '1934년',
        '1935년',
        '1936년',
        3
    ),
    (
        39,
        '이 건물의 건축사적 의미는?',
        '한국 최초의 서구식 건물',
        '근대 건축의 걸작',
        '가장 큰 규모',
        '가장 오래된 건물',
        2
    ),
    (
        39,
        '이화여자대학교의 설립 배경은?',
        '일제강점기 저항',
        '근대 여성 교육',
        '종교 선교',
        '정치적 목적',
        2
    );

-- Insert questions for 연세대학교 언더우드관
INSERT INTO
    questions (
        quiz_id,
        question_text,
        option1,
        option2,
        option3,
        option4,
        correct_answer
    )
VALUES (
        40,
        '언더우드관이 건립된 연도는?',
        '1922년',
        '1923년',
        '1924년',
        '1925년',
        3
    ),
    (
        40,
        '언더우드는 누구인가?',
        '미국 선교사',
        '영국 외교관',
        '독일 교육자',
        '프랑스 건축가',
        1
    ),
    (
        40,
        '연세대학교의 전신은?',
        '조선대학교',
        '경성대학교',
        '연희전문학교',
        '보성전문학교',
        3
    );

-- Insert badge data (heritage-specific badges)
INSERT INTO
    badges (
        heritage_id,
        name,
        description,
        image_url
    )
VALUES (
        1,
        '경복궁 마스터',
        '경복궁 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/gyeongbokgung.png'
    ),
    (
        2,
        '창덕궁 마스터',
        '창덕궁 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/changdeokgung.png'
    ),
    (
        3,
        '덕수궁 마스터',
        '덕수궁 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/deoksugung.png'
    ),
    (
        4,
        '창경궁 마스터',
        '창경궁 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/changgyeonggung.png'
    ),
    (
        5,
        '경희궁 마스터',
        '경희궁 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/gyeonghuigung.png'
    ),
    (
        6,
        '남대문 지킴이',
        '남대문 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/namdaemun.png'
    ),
    (
        7,
        '동대문 지킴이',
        '동대문 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/dongdaemun.png'
    ),
    (
        8,
        '서대문 역사가',
        '서대문 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/seodaemun.png'
    ),
    (
        9,
        '낙산성곽 탐험가',
        '낙산성곽 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/naksan.png'
    ),
    (
        10,
        '인왕산성곽 탐험가',
        '인왕산성곽 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/inwangsan.png'
    ),
    (
        11,
        '종묘 제례 전문가',
        '종묘 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/jongmyo.png'
    ),
    (
        12,
        '사직단 제례 전문가',
        '사직단 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/sajikdan.png'
    ),
    (
        13,
        '문묘 유학자',
        '문묘 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/munmyo.png'
    ),
    (
        14,
        '조계사 불자',
        '조계사 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/jogyesa.png'
    ),
    (
        15,
        '봉은사 불자',
        '봉은사 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/bongeunsa.png'
    ),
    (
        16,
        '선정릉 지킴이',
        '선정릉 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/seonjeongneung.png'
    ),
    (
        17,
        '태릉 지킴이',
        '태릉 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/taereung.png'
    ),
    (
        18,
        '의릉 지킴이',
        '의릉 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/uireung.png'
    ),
    (
        19,
        '광화문 지킴이',
        '광화문 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/gwanghwamun.png'
    ),
    (
        20,
        '보신각 종지기',
        '보신각 퀴즈를 모두 맞혔습니다!',
        'https://example.com/badges/bosingak.png'
    );