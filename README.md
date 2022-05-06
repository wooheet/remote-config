<div align=center>

# Remote config Auth Server

</div>


## 📚 Tech Stack

- Go
- Gingonic
- MySQL
- Redis
- Docker
- Travis CI
- Deepsource Go


## 📄 API Spec

### Config
#### API Token 조회

```azure
GET /api/v1/config?store_id=hackingrowth

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

#### API Token 등록
```azure
POST /api/v1/config

Body
{
"token": "ab12cbf123ff2",
"tracker_type": "mixpanel",
"store_id": "hackinggrowth"
}

Status Code: 201
```

#### API Token 업데이트
```azure
[PUT|PATCH] /api/v1/config?store_id=hackingrowth

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

#### API Token 삭제
```azure
DELETE /api/v1/config?store_id=hackingrowth

Status Code: 204
```

### Auth

#### Login
```azure
POST /v1/login

Status Code: 201
```

#### Logout
```azure
DELETE /v1/logout

Status Code: 201
```