User HTTP API 
===============================

[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/httpkit/example/userapi)](https://goreportcard.com/report/github.com/gokit/httpkit/example/userapi)

User HTTP API is a auto generated http api for the struct `User`.

The API expects the user to implement and provide the backend interface to provided underline db interaction:

```go
type Backend interface{
    Delete(context.Context, string) error
    Get(context.Context, string) (users.User, error)
    Update(context.Context, string, users.User) error
    GetAll(context.Context, string, string, int, int) ([]users.User, int, error)
    Create(context.Context, users.User) (users.User, error)
}
```

It exposes the following methods for each endpoint:

## Create { POST /User/ }
### Method: `func (api *HTTPAPI) Create(ctx *httputil.Context) error`

Create receives the provided record of the User type which is delieved the 
JSON content to the HTTP API. This will in turn return a respective status code.

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00",

    "public_id":	"c5dfpkxg4twl5cehbuco11g8xr5mph",

    "username":	"NicholasWilliams",

    "email":	"nBrown@Yodel.com"

}
```

- Expected Status Code

```
Failure: 500
Success: 201
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
{


    "updated_at":	"2018-01-25T09:07:49+01:00",

    "public_id":	"eeh4cn165l2b3tzy9n9ijch9cuqyr7",

    "username":	"JohnBaker",

    "email":	"aut@Oba.edu",

    "created_at":	"2018-01-25T09:07:49+01:00"

}
```

## INFO /User/
### Method: `func (api *HTTPAPI) Info(ctx *httputil.Context) error`

Info returns total of records available in api for type users.User.

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
{
    "total": 10,
}
```

## GET /User/:public_id
### Method: `func (api *HTTPAPI) Get(ctx *httputil.Context) error`

Get retrives a giving record of the User type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Request Parameters

```
    :public_id
```

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
{


    "public_id":	"djgg5zt4a2uv1g5zkp2nei5m8xy5lh",

    "username":	"RogerBradley",

    "email":	"voluptatibus_cumque@Dabshots.biz",

    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00"

}
```

## GET /User/
### Method: `func (api *HTTPAPI) GetAll(ctx *httputil.Context) error`

Get retrives all records of the User type from the HTTP API.

- Expected Status Code

```
Failure: 500
Success: 200
```

- Expected Response Body

```http
    Content-Type: application/json
```

```json
[{


    "public_id":	"0u1u6awnxkwubzpz4y77yavytwfl2a",

    "username":	"6Lee",

    "email":	"0West@Yakijo.biz",

    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00"

},{


    "email":	"1Garcia@Aibox.gov",

    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00",

    "public_id":	"18130s3bntsjnexp871jdplvft93er",

    "username":	"AshleyLopez"

}]
```

## PUT /User/:public_id
### Method: `func (api *HTTPAPI) Update(ctx *httputil.Context) error`

Update attempts to update a giving record of the User type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record with the provided JSON request body.

- Expected Request Parameters

```
    :public_id
```

- Expected Request Body

```http
    Content-Type: application/json
```

```json
{


    "public_id":	"qvgpxzibfr33p2f42c0j6j4847vp2i",

    "username":	"VirginiaGreen",

    "email":	"RoseHarper@Yodo.name",

    "created_at":	"2018-01-25T09:07:49+01:00",

    "updated_at":	"2018-01-25T09:07:49+01:00"

}
```

- Expected Status Code

```
Failure: 500
Success: 204
```

## DELETE /User/:public_id
### Method: `func (api *HTTPAPI) Delete(ctx *httputil.Context) error`

Get deletes a giving record of the User type from the HTTP API returning received result as a JSON
response. It uses the provided `:public_id` parameter as the paramter to identify the record.

- Expected Request Parameters

```
    :public_id
```

- Expected Status Code

```
Failure: 500
Success: 204
```

