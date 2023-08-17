package handlers

import (
	"gobanking/internal/core-layer/ports/accounts"
)

type AccountHandler struct {
	UserService *accounts.AccountService
}
