# Heritage Tour API 명세서

## 📋 목차
- [개요](#개요)
- [인증](#인증)
- [Heritage API](#heritage-api)
- [User API](#user-api)
- [Badge API](#badge-api)
- [Auth API](#auth-api)
- [Quiz API](#quiz-api)
- [에러 코드](#에러-코드)

---

## 개요

Heritage Tour API는 문화유산 관광 애플리케이션을 위한 RESTful API입니다.

**Base URL**: `http://localhost:8080`

**응답 형식**: JSON

**공통 응답 구조**:
```json
{
  "success": true,
  "data": { ... }
}
```

---

## 인증

JWT 토큰을 사용한 Bearer Authentication

**Header**:
```
Authorization: Bearer <JWT_TOKEN>
```

**인증 필요 엔드포인트**:
- 모든 `/api/users/me/*` 엔드포인트
- 모든 `/api/heritage/*/visits` 및 `/api/heritage/*/reviews` 엔드포인트
- 모든 `/api/quiz/*` 엔드포인트

---

## Heritage API

### 1. 문화유산 검색

**`GET /api/heritage`**

**Query Parameters**:
- `search` (string, required): 검색 키워드

**Response**:
```json
{
  "success": true,
  "data": {
    "name": "경복궁",
    "latitude": 37.5796,
    "longitude": 126.9770,
    "imageUrl": "https://example.com/image.jpg",
    "description": "조선왕조의 정궁"
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
  "qrCode": "HERITAGE_QR_CODE_123"
}
```

**Response**:
```json
{
  "success": true,
  "data": {
    "heritageId": 1,
    "name": "경복궁",
    "imageUrl": "https://example.com/image.jpg",
    "description": "조선왕조의 정궁",
    "story": "경복궁은 1395년에 창건된...",
    "isFirstVisit": true
  }
}
```

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
    "createdAt": "2024-01-01 10:00:00",
    "updatedAt": "2024-01-01 10:00:00"
  }
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
    "createdAt": "2024-01-01 10:00:00"
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
    "visitCount": 5,
    "visits": [
      {
        "name": "경복궁",
        "visitedAt": "2024-01-01",
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
    "visitCount": 5
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
    "reviewCount": 3,
    "reviews": [
      {
        "reviewId": 1,
        "heritageId": 1,
        "heritageName": "경복궁",
        "rating": 5,
        "reviewText": "정말 아름다운 궁궐이었습니다!",
        "createdAt": "2024-01-01 10:00:00",
        "updatedAt": "2024-01-01 10:00:00"
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
    "reviewCount": 3
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
    "badgeCount": 2,
    "badges": [
      {
        "name": "첫 방문자",
        "imageUrl": "https://example.com/badge.png",
        "earnedAt": "2024-01-01",
        "heritageName": "경복궁"
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
    "badgeCount": 2
  }
}
```

---

## Badge API

### 1. 모든 배지 조회

**`GET /api/badges`**

**Response**:
```json
{
  "success": true,
  "data": {
    "badgeCount": 10,
    "badges": [
      {
        "badgeId": 1,
        "name": "첫 방문자",
        "imageUrl": "https://example.com/badge.png",
        "heritageName": "경복궁",
        "description": "첫 번째 문화유산 방문 시 획득"
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
    "name": "첫 방문자",
    "description": "첫 번째 문화유산 방문 시 획득",
    "imageUrl": "https://example.com/badge.png",
    "heritageName": "경복궁",
    "createdAt": "2024-01-01 10:00:00"
  }
}
```

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

**Response**:
```json
{
  "success": true,
  "data": {
    "userId": 1,
    "username": "user123",
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs..."
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
    "username": "user123",
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

### 3. 토큰 갱신

**`POST /api/auth/refresh`**

**Request Body**:
```json
{
  "refreshToken": "eyJhbGciOiJIUzI1NiIs..."
}
```

**Response**:
```json
{
  "success": true,
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIs...",
    "refreshToken": "eyJhbGciOiJIUzI1NiIs..."
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
    "quizId": 1,
    "question": "경복궁이 창건된 연도는?",
    "options": ["1392년", "1395년", "1398년", "1400년"],
    "correctAnswer": 1
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
  "answer": 1
}
```

**Response**:
```json
{
  "success": true,
  "data": {
    "correct": true,
    "score": 100,
    "explanation": "정답입니다! 경복궁은 1395년에 창건되었습니다."
  }
}
```

---

## 에러 코드

### HTTP 상태 코드

| 상태 코드 | 설명 |
|-----------|------|
| 200 | 성공 |
| 201 | 생성 성공 |
| 400 | 잘못된 요청 |
| 401 | 인증 실패 |
| 403 | 권한 없음 |
| 404 | 리소스 없음 |
| 409 | 충돌 (중복 생성) |
| 500 | 서버 오류 |

### 에러 응답 형식

```json
{
  "success": false,
  "error": {
    "code": "HERITAGE_NOT_FOUND",
    "message": "Heritage not found"
  }
}
```

### 주요 에러 코드

| 에러 코드 | 설명 |
|-----------|------|
| `HERITAGE_NOT_FOUND` | 문화유산을 찾을 수 없음 |
| `INVALID_QR_CODE` | 유효하지 않은 QR 코드 |
| `REVIEW_ALREADY_EXISTS` | 이미 후기가 존재함 |
| `REVIEW_NOT_FOUND` | 후기를 찾을 수 없음 |
| `USER_NOT_AUTHENTICATED` | 사용자 인증 실패 |
| `INVALID_TOKEN` | 유효하지 않은 토큰 |
| `BADGE_NOT_FOUND` | 배지를 찾을 수 없음 |

---

## 최종 API 개수: 12개

### 📍 Heritage API (5개)
- GET /api/heritage?search=keyword
- POST /api/heritage/:heritageId/visits
- POST /api/heritage/:heritageId/reviews
- PUT /api/heritage/:heritageId/reviews
- GET /api/heritage/:heritageId/reviews/me

### 👤 User API (4개)
- GET /api/users/me
- GET /api/users/me/visits
- GET /api/users/me/reviews
- GET /api/users/me/badges

### 🏆 Badge API (2개)
- GET /api/badges
- GET /api/badges/:badgeId

### 🔐 Auth API (3개)
- POST /api/auth/register
- POST /api/auth/login
- POST /api/auth/refresh

### 🎯 Quiz API (2개)
- GET /api/quiz/:heritageId
- POST /api/quiz/:heritageId/submit

---

**마지막 업데이트**: 2024년 12월

🔐 = 인증 필요