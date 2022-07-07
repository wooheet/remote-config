<div align=center>

# Remote config Auth Server

</div>


## 📚 Tech Stack

- Go
- Gingonic
- MySQL
- Redis
- Docker


## 📄 API Spec

### Config
#### Token 조회

```azure
GET /v1/config?store_id=hackingrowth

Status Code: 200

Response

{
"store_id": "hackinggrowth",
"token": "ab12cbf123ff2",
"tracker_type": "mixpanel"
"inserted_at": "2022-01-01T00:00:00.00Z",
"updated_at": "2022-01-01T00:00:00.00Z",
}
```

#### Token 등록
```azure
POST /v1/config

Body
{
"token": "ab12cbf123ff2",
"tracker_type": "mixpanel",
"store_id": "hackinggrowth"
}

Status Code: 201
```

#### Token 업데이트
```azure
[PUT|PATCH] /v1/config?store_id=hackingrowth

Body
{
"token": "ab12cbf123ff3",
"tracker_type": "mixpanel"
}

Status Code: 200

Response

{
"store_id": "hackinggrowth",
"token": "ab12cbf123ff3",
"tracker_type": "mixpanel"
"inserted_at": "2022-01-01T00:00:00.00Z",
"updated_at": "2022-01-01T00:00:00.00Z",
}
```

#### Token 삭제
```azure
DELETE /v1/config?store_id=hackingrowth

Status Code: 204
```

#### Script tag 설치
```azure
GET /v1/scripttag?store_id=hackingrowth?tracker_type=mixpanel

Status Code: 200
```

### Auth

#### Signin
```azure
POST /v1/signin

Status Code: 201
```

#### Logout
```azure
DELETE /v1/logout

Status Code: 201
```