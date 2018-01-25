HTTPKit
--------
[![Go Report Card](https://goreportcard.com/badge/github.com/gokit/httpkit)](https://goreportcard.com/report/github.com/gokit/httpkit)

HTTPKit implements a code generator which automatically generates go packages for a http CRUD implementations for annotated structs.

## Install

```
go get -u github.com/gokit/httpkit
```

## Examples

See [Examples](./examples) for demonstrations of httpkit code generation.

## Usage

Running the following commands instantly generates all necessary files and packages for giving code gen.

```go
> httpkit generate
```

## How It works

You annotate any giving struct with `@httpapi` which marks giving struct has a target for code generation. 

It also respects attributes like `New` and `Update` on the `@httpapi` annotation, these two provide information to the generator for:

1. What struct to be used as representing a new data for struct type.
2. What struct contain information for representing the update data for struct type.

Sample below:

```go
// User is a type defining the given user related fields for a given.
//@httpapi(New => NewUser, Update => UpdatedUser)
type User struct {
	Username      string    `json:"username"`
	PublicID      string    `json:"public_id"`
	PrivateID     string    `json:"private_id"`
	Hash          string    `json:"hash"`
	TwoFactorAuth bool      `json:"two_factor_auth"`
	Created       time.Time `json:"created_at"`
	Updated       time.Time `json:"updated_at"`
}
```
