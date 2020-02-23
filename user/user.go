package user

import (
  "encoding/json"
)

type User struct {
  Name string 'json:"name"'
  Token string 'json:"token"'
}