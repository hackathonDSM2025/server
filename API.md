# Heritage Tour API 명세서

## 📋 목차

- [개요](#개요)
- [인증](#인증)
- [Auth API](#auth-api)
- [Heritage API](#heritage-api)
- [User API](#user-api)
- [Quiz API](#quiz-api)
- [Badge API](#badge-api)
- [에러 코드](#에러-코드)
- [검증 규칙](#검증-규칙)

---

## 개요

Heritage Tour API는 문화유산 관광 애플리케이션을 위한 RESTful API입니다.

**Base URL**: `https://pastport.ijw.app`

**응답 형식**: JSON (UTF-8 인코딩)

**공통 응답 구조**:

```json
{
  "success": true,
  "data": { ... }
}
```

**에러 응답 구조**:

```json
{
  "success": false,
  "message": "에러 메시지"
}
```

---

## 인증

JWT 토큰을 사용한 Bearer Authentication

**Header**:

```
Authorization: Bearer <JWT_TOKEN>
```

**토큰 만료 시간**: 1시간

**인증 필요 엔드포인트**:

- 모든 `/api/users/me/*` 엔드포인트
- 모든 `/api/heritage/*/visits` 및 `/api/heritage/*/reviews` 엔드포인트
- 모든 `/api/quiz/*` 엔드포인트

---

## Auth API

### 1. 회원가입

**`POST /api/auth/register`**

**Request Body**:

```json
{
  "username": "user123",
  "password": "password123"
}
```

**검증 규칙**:
- `username`: 필수, 중복 불가
- `password`: 필수, bcrypt 해시화 저장

**Response**:

```json
{
  "success": true,
  "data": {
    "userId": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 2. 로그인

**`POST /api/auth/login`**

**Request Body**:

```json
{
  "username": "user123",
  "password": "password123"
}
```

**Response**:

```json
{
  "success": true,
  "data": {
    "userId": 1,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 3. 사용자명 중복 확인

**`POST /api/auth/check-username`**

**Request Body**:

```json
{
  "username": "user123"
}
```

**Response**:

```json
{
  "success": true,
  "data": {
    "available": true,
    "message": "Username is available"
  }
}
```

---

## Heritage API

### 1. 문화유산 검색

**`GET /api/heritage?keyword=검색어`**

**Query Parameters**:

- `keyword` (string, required): 검색 키워드 (한글 지원)

**검증 규칙**:
- `keyword`: 필수 파라미터, 빈 문자열 불가

**Response**:

```json
{
  "success": true,
  "data": {
    "name": "경복궁",
    "latitude": 37.5796,
    "longitude": 126.977,
    "imageUrl": "https://example.com/images/gyeongbokgung.jpg",
    "description": "조선 왕조의 정궁으로 1395년에 창건되었습니다."
  }
}
```

### 2. 방문 기록 생성 (QR 코드 스캔)

**`POST /api/heritage/:heritageId/visits`** 🔐

**Path Parameters**:

- `heritageId` (int): 문화유산 ID

**Request Body**:

```json
{
  "qrCode": "QR_GYEONGBOKGUNG_001"
}
```

**검증 규칙**:
- `qrCode`: 필수, 유효한 QR 코드여야 함
- QR 코드와 heritageId가 일치해야 함

**Response**:

```json
{
  "success": true,
  "data": {
    "heritageId": 1,
    "name": "경복궁",
    "imageUrl": "https://example.com/images/gyeongbokgung.jpg",
    "description": "조선 왕조의 정궁으로 1395년에 창건되었습니다.",
    "story": "경복궁은 조선 태조 4년(1395년)에 창건된 조선 왕조의 정궁입니다...",
    "isFirstVisit": true
  }
}
```

**특별 기능**:
- 첫 방문 기록 저장
- 중복 방문 시 `isFirstVisit: false`

### 3. 후기 작성

**`POST /api/heritage/:heritageId/reviews`** 🔐

**Path Parameters**:

- `heritageId` (int): 문화유산 ID

**Request Body**:

```json
{
  "rating": 5,
  "reviewText": "정말 아름다운 궁궐이었습니다!"
}
```

**검증 규칙**:
- `rating`: 필수, 1-5 범위의 정수
- `reviewText`: 필수, 빈 문자열 불가
- 사용자당 문화유산 하나에 대해 하나의 리뷰만 작성 가능

**Response**:

```json
{
  "success": true,
  "message": "Review created successfully"
}
```

### 4. 후기 수정

**`PUT /api/heritage/:heritageId/reviews`** 🔐

**Path Parameters**:

- `heritageId` (int): 문화유산 ID

**Request Body**:

```json
{
  "rating": 4,
  "reviewText": "수정된 후기입니다."
}
```

**검증 규칙**:
- 기존 리뷰가 존재해야 함
- `rating`: 필수, 1-5 범위의 정수
- `reviewText`: 필수, 빈 문자열 불가

**Response**:

```json
{
  "success": true,
  "message": "Review updated successfully"
}
```

### 5. 내 후기 조회

**`GET /api/heritage/:heritageId/reviews/me`** 🔐

**Path Parameters**:

- `heritageId` (int): 문화유산 ID

**Response**:

```json
{
  "success": true,
  "data": {
    "reviewId": 1,
    "rating": 5,
    "reviewText": "정말 아름다운 궁궐이었습니다!",
    "createdAt": "2025-07-16 16:50:08",
    "updatedAt": "2025-07-16 16:50:08"
  }
}
```

**리뷰가 없는 경우**:

```json
{
  "success": true,
  "data": null
}
```

---

## User API

### 1. 사용자 프로필 조회

**`GET /api/users/me`** 🔐

**Response**:

```json
{
  "success": true,
  "data": {
    "userId": 1,
    "username": "user123",
    "createdAt": "2025-07-16 16:41:06"
  }
}
```

### 2. 내 방문 기록 조회

**`GET /api/users/me/visits`** 🔐

**Query Parameters**:

- `count_only` (boolean, optional): true면 개수만 반환

**Response (전체 목록)**:

```json
{
  "success": true,
  "data": {
    "visitCount": 1,
    "visits": [
      {
        "name": "경복궁",
        "visitedAt": "2025-07-16",
        "completed": true
      }
    ]
  }
}
```

**Response (개수만)**:

```json
{
  "success": true,
  "data": {
    "visitCount": 1
  }
}
```

### 3. 내 후기 조회

**`GET /api/users/me/reviews`** 🔐

**Query Parameters**:

- `count_only` (boolean, optional): true면 개수만 반환

**Response (전체 목록)**:

```json
{
  "success": true,
  "data": {
    "reviewCount": 1,
    "reviews": [
      {
        "reviewId": 2,
        "heritageId": 1,
        "heritageName": "경복궁",
        "heritageImageUrl": "https://example.com/images/gyeongbokgung.jpg",
        "rating": 4,
        "reviewText": "Very nice palace with rich history",
        "createdAt": "2025-07-16 16:50:08",
        "updatedAt": "2025-07-16 16:50:25"
      }
    ]
  }
}
```

**Response (개수만)**:

```json
{
  "success": true,
  "data": {
    "reviewCount": 1
  }
}
```

### 4. 내 배지 조회

**`GET /api/users/me/badges`** 🔐

**Query Parameters**:

- `count_only` (boolean, optional): true면 개수만 반환

**Response (전체 목록)**:

```json
{
  "success": true,
  "data": {
    "badgeCount": 1,
    "badges": [
      {
        "name": "경복궁 마스터",
        "imageUrl": "https://example.com/badges/gyeongbokgung.png",
        "earnedAt": "2025-07-16",
        "heritageName": "경복궁",
        "heritageImageUrl": "https://example.com/images/gyeongbokgung.jpg"
      }
    ]
  }
}
```

**Response (개수만)**:

```json
{
  "success": true,
  "data": {
    "badgeCount": 1
  }
}
```

---

## Quiz API

### 1. 퀴즈 조회

**`GET /api/quiz/:heritageId`** 🔐

**Path Parameters**:

- `heritageId` (int): 문화유산 ID

**Response**:

```json
{
  "success": true,
  "data": {
    "questions": [
      {
        "id": 1,
        "question": "경복궁이 건립된 연도는?",
        "options": ["1392년", "1394년", "1395년", "1396년"],
        "correctAnswer": 2
      },
      {
        "id": 2,
        "question": "경복궁의 정전은?",
        "options": ["근정전", "인정전", "중화전", "선정전"],
        "correctAnswer": 0
      }
    ]
  }
}
```

### 2. 퀴즈 답안 제출

**`POST /api/quiz/:heritageId/submit`** 🔐

**Path Parameters**:

- `heritageId` (int): 문화유산 ID

**Request Body**:

```json
{
  "answers": [2, 0, 0]
}
```

**검증 규칙**:
- `answers`: 필수, 0-based index 배열
- 문제 개수와 답안 개수가 일치해야 함

**Response (100점 - 배지 수여)**:

```json
{
  "success": true,
  "data": {
    "score": 100,
    "correctCount": 3,
    "totalQuestions": 3,
    "allCorrect": true,
    "canRetry": false,
    "newBadge": {
      "name": "경복궁 마스터",
      "imageUrl": "https://example.com/badges/gyeongbokgung.png",
      "heritageName": "경복궁",
      "heritageImageUrl": "https://example.com/images/gyeongbokgung.jpg"
    }
  }
}
```

**Response (100점 - 이미 배지 보유)**:

```json
{
  "success": true,
  "data": {
    "score": 100,
    "correctCount": 3,
    "totalQuestions": 3,
    "allCorrect": true,
    "canRetry": false
  }
}
```

**Response (부분 점수)**:

```json
{
  "success": true,
  "data": {
    "score": 67,
    "correctCount": 2,
    "totalQuestions": 3,
    "allCorrect": false,
    "canRetry": true,
    "incorrectAnswers": [
      {
        "questionId": 3,
        "correctAnswer": 1,
        "userAnswer": 2,
        "questionText": "경복궁의 정전은?"
      }
    ]
  }
}
```

**배지 수여 조건**:
- **만점(100점) 달성** 시에만 해당 문화유산 배지 수여
- 이미 보유한 배지는 중복 수여하지 않음
- 배지 수여 시 `newBadge` 필드에 배지 정보 포함

**특별 기능**:
- 재시도 가능 (최신 점수로 업데이트)
- 만점 달성 시 배지 자동 수여

---

## Badge API

### 1. 모든 배지 조회

**`GET /api/badges`**

**Response**:

```json
{
  "success": true,
  "data": {
    "badgeCount": 5,
    "badges": [
      {
        "badgeId": 1,
        "name": "경복궁 마스터",
        "imageUrl": "https://example.com/badges/gyeongbokgung.png",
        "heritageName": "경복궁",
        "heritageImageUrl": "https://example.com/images/gyeongbokgung.jpg",
        "description": "경복궁 퀴즈를 모두 맞혔습니다!"
      },
      {
        "badgeId": 2,
        "name": "창덕궁 마스터",
        "imageUrl": "https://example.com/badges/changdeokgung.png",
        "heritageName": "창덕궁",
        "heritageImageUrl": "https://example.com/images/changdeokgung.jpg",
        "description": "창덕궁 퀴즈를 모두 맞혔습니다!"
      }
    ]
  }
}
```

### 2. 특정 배지 상세 조회

**`GET /api/badges/:badgeId`**

**Path Parameters**:

- `badgeId` (int): 배지 ID

**Response**:

```json
{
  "success": true,
  "data": {
    "badgeId": 1,
    "name": "경복궁 마스터",
    "description": "경복궁 퀴즈를 모두 맞혔습니다!",
    "imageUrl": "https://example.com/badges/gyeongbokgung.png",
    "heritageName": "경복궁",
    "heritageImageUrl": "https://example.com/images/gyeongbokgung.jpg",
    "createdAt": "2025-07-16 16:22:30"
  }
}
```

---

## 에러 코드

### HTTP 상태 코드

| 상태 코드 | 설명             |
| --------- | ---------------- |
| 200       | 성공             |
| 201       | 생성 성공        |
| 400       | 잘못된 요청      |
| 401       | 인증 실패        |
| 403       | 권한 없음        |
| 404       | 리소스 없음      |
| 409       | 충돌 (중복 생성) |
| 500       | 서버 오류        |

### 에러 응답 형식

```json
{
  "success": false,
  "message": "Heritage not found"
}
```

### 주요 에러 메시지

| 에러 메시지                     | 설명                    |
| ------------------------------ | ----------------------- |
| `Heritage not found`           | 문화유산을 찾을 수 없음 |
| `Invalid QR code`              | 유효하지 않은 QR 코드   |
| `Review already exists for this heritage` | 이미 후기가 존재함      |
| `Review not found for this heritage` | 후기를 찾을 수 없음     |
| `User not authenticated`       | 사용자 인증 실패        |
| `Invalid request body`         | 잘못된 요청 본문        |
| `Keyword is required`          | 검색 키워드 필수        |
| `QR code does not match heritage ID` | QR 코드와 문화유산 ID 불일치 |

---

## 검증 규칙

### 인증 관련
- JWT 토큰 필수 (Bearer 방식)
- 토큰 만료 시간: 1시간
- 토큰 검증 실패 시 401 에러

### 문화유산 검색
- `keyword` 파라미터 필수
- 한글 키워드 지원
- 부분 일치 검색 (LIKE 패턴)
- 정확한 일치 우선순위

### 방문 기록
- QR 코드 필수
- QR 코드와 문화유산 ID 일치 검증
- 중복 방문 허용 (방문 시간 업데이트)
- 첫 방문 기록 저장

### 후기 시스템
- 평점: 1-5 범위의 정수
- 리뷰 텍스트: 빈 문자열 불가
- 사용자당 문화유산 하나에 대해 하나의 리뷰만 허용
- 한글 텍스트 지원

### 퀴즈 시스템
- 답안 배열 필수
- 문제 개수와 답안 개수 일치
- 0-based index 사용
- 재시도 가능 (최신 점수 업데이트)

### 배지 시스템
- **퀴즈 만점(100점) 달성 시에만** 해당 문화유산 배지 수여
- 이미 보유한 배지는 중복 수여하지 않음
- 배지 수여 시 응답에 `newBadge` 정보 포함

---

## 총 API 개수: 13개

### 🔐 Auth API (3개)
- `POST /api/auth/register`
- `POST /api/auth/login`
- `POST /api/auth/check-username`

### 📍 Heritage API (5개)
- `GET /api/heritage?keyword=검색어`
- `POST /api/heritage/:heritageId/visits`
- `POST /api/heritage/:heritageId/reviews`
- `PUT /api/heritage/:heritageId/reviews`
- `GET /api/heritage/:heritageId/reviews/me`

### 👤 User API (4개)
- `GET /api/users/me`
- `GET /api/users/me/visits`
- `GET /api/users/me/reviews`
- `GET /api/users/me/badges`

### 🎯 Quiz API (2개)
- `GET /api/quiz/:heritageId`
- `POST /api/quiz/:heritageId/submit`

### 🏆 Badge API (2개)
- `GET /api/badges`
- `GET /api/badges/:badgeId`

---

**최종 검증 완료**: 2025년 7월 16일

🔐 = 인증 필요  
✅ = 테스트 완료  
🇰🇷 = 한글 지원