<div align=center>

# Remote config Auth Server

</div>


## π Tech Stack

- Go
- Gingonic
- MySQL
- Redis
- Docker
- Travis CI
- Deepsource Go


## π API Spec

### Config
#### Token μ‘°ν

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

#### Token λ±λ‘
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

#### Token μλ°μ΄νΈ
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

#### Token μ­μ 
```azure
DELETE /v1/config?store_id=hackingrowth

Status Code: 204
```

#### Script tag μ€μΉ
```azure
GET /v1/scripttag?store_id=hackingrowth?tracker_type=mixpanel

Status Code: 200
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