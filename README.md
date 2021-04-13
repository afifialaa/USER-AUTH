# user-auth

A user authentication api written in GO with MongoDB as main datastore.

## Installation

1. Clone repo
```bash
git clone https://github.com/afifialaa/user-auth.git
```

2. Install dependencies
```bash
go get
```


## Requests

| Action  | Method  | URL |
| -------------  | ------------- | ------------- |
| [Login](#login) | POST  | /api/user/login  |
| [Signup](#signup)  | POST  | /api/user/signup  |

## Login
### Request
```curl
curl -X POST -d 'email=johndoe@example.com&password=jdoe123' http://localhost:8000/api/user/login
```

### Response
```json
HTTP/1.1 200 Ok
Date: Tue, 13 Apr 2021 11:23:38 GMT
Status: 200 Ok
Content-Type: application/json
Content-Length: 130

{"token":"IUzI1NiIsInR5cCI6IkpXVC.IUzI1NiIsInR5cCI6IkpXVC.IUzI1NiIsInR5cCI6IkpXVC"}
```

## Signup
### Request
```curl
curl -X POST -d 'email=johndoe@example.com&password=jdoe123' http://localhost:8000/api/user/signup
```

### Response
```json
HTTP/1.1 201 Created
Date: Tue, 13 Apr 2021 11:23:38 GMT
Status: 201 Created
Content-Type: application/json
Content-Length: 40

{"msg": "User was created successfully"}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
