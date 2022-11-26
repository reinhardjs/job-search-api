
## Endpoints

API Endpoint Host : http://103.134.154.18:30822/

### Login
`GET` http://103.134.154.18:30822/login

Authentication : `Bearer <Token>`

Example Body Payload:
```
{
    "email": "admin@email.com",
    "password": "password"
}
```

<br> 

Example Response Payload:
```
{
    "status": 200,
    "message": "this token will be valid for the next 3 minutes, login again if it expired",
    "data": {
        "email": "admin@email.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImFkbWluQGVtYWlsLmNvbSIsImV4cCI6MTY2OTQ3MjAxOH0.8KfQnlZHC8tFldiaqbj7DQlW7QwIbpWn16TBDSu_p9w"
    }
}
```

<br>

### Get list of position
`GET` http://103.134.154.18:30822/positions?description={description}&location={location}&page={page}

Authentication : `Bearer <Token>`

<br>

### Get position
`GET` http://103.134.154.18:30822/positions/{position-id}

Authentication : `Bearer <Token>`

## Credentials
```
{
    "email": "admin@email.com",
    "password": "password"
}
```

<br>

Your token will be expiring for 3 minutes. You should request for new token from `/login` endpoint <br>
You can use non expiring token below, if u disturbed by the expired token

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOiI2MzgwYzIxNmE2NjBhOWQ3ZjRmMDZmZDIiLCJFbWFpbCI6ImFkbWluQGVtYWlsLmNvbSIsIlJvbGUiOiJhZG1pbiJ9.kkcnAqajjcx0YmtRnWk-P594v_2wIEObwUzTtuMq_JY
