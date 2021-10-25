# altcoin-backend
Backend for a Bitcoin puchasing system.


## List of features
- User Sign up - When he signups, assign him 500,000 dollars of virtual money
- User Log in
- Validate User token
- Purchase bitcon
- Sell bitcoin

## Tools and technologies:
- Go
- Sqllite

## Models:
- User
  - id
  - created_at
  - updated_at
  - username
  - encrypted_password
  - wallet_amount
  - bitcoin_amount

- Transaction
  - user_id
  - bitcoin_amount
  - bitcoin_price
  - created_at

## API endpoints:
- [POST] /api/signup
    - in
```json
{
    "username": "",
    "password": ""
}
```
    - out
Response: `200 | 422`,
data:
```json
{
    "errors": {
        "msg": ""
    }
}
```
      - Response type: 400, 401 along with error dict
    - Info
      - Checks if the username exists already, If yes return error,
      - If no, create user, and assign the value amount of 500,000

- [POST] /api/login
  - in
```json
{
    "username": "",
    "password": ""
}
```
    - out
Response `200 | 401`
data:
```json
{
    "token": "" // valid jwt token
}
```
    - info
      - Validates username, password and if yes, returns jwt token

- [POST] /api/buy
    - in
Headers:
Authorization: Bearer <token>
```json
{
    "amount": 1.4,
    "price": 8000,
}
    - out
Response type: `422`
```json
{
    "errors": {
        "msg": ""
    }
}
```
OR
Respose type: `200`
    - Info: check if the final price (amount * price) is less than or equal to the wallet amount.

- [POST] /api/sell
  - in
Headers  -
Authorization: Bearer <token>
```json
{
    "amount": 1.4,
    "price": 8222
}
```
  - out
Response type: `422`
```json
{
    "errors": {
        "msg": ""
    }
}
```
OR
Respose type: `200`
  - Info: Check if the amount is less than or equal the bitcoin amount stored in user

- [GET] /api/user
  - in
Headers  -
Authorization: Bearer <token>

  - out
Response Types: `200`
data:

```json
{
    "data": {
        "username": "",
        "wallet_amount": "",
        "bitcoin_amount": "",
        "bitcoin_value": "" // computed field, based on the current value of the bitcoin
    }
}
```

## Caveats:
- Modified the naming convention compared to the given requirements:
  - `amount` - refers to the quanity of bitcoin purchased / sold / present with the user
  - `price` - refers to the price of the bitcoin when it was sold / purchased
  - `wallet_amount` - refers to the amount of money in user's wallet
- We have a separate table called transaction which refers to all the transations (purchased / sold) by the user
- In the user model, I have `bitcoin_value` as the computed field, as it's value will depend on the current market value of bitcoin


# Go commands
```bash
go mod tidy # to add missing modele requirements in go mod file
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .)) # to add the go bin's files to path
```

## Important info:
- The install directory is controlled by the GOPATH and GOBIN environment variables. If GOBIN is set, binaries are installed to that directory. If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go).
- A standalone program (as opposed to a library) is always in package main.
- As you might know, Go requires all exported fields to start with a capitalized letter.




