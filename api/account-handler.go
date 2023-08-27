package api

import (
	"net/http"

	account "github.com/FadyGamilM/go-banking-v2/internal/account/domain"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	srv account.AccountService
}

func NewAccountHandler(s account.AccountService) *AccountHandler {
	return &AccountHandler{
		srv: s,
	}
}

// ===> DTOs
type CreateAccountReqDto struct {
	OwnerName string `json:"owner_name" binding:"required"`
	Currency  string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (h *AccountHandler) HandleCreateAccount(ctx *gin.Context) {
	// validate the incoming data from the request body
	req := new(CreateAccountReqDto)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, responseWithError(err))
	}

	accDomain := &account.Account{
		OwnerName: req.OwnerName,
		Currency:  req.Currency,
		Balance:   float64(0),
	}

	var acc *account.Account
	var err error
	if acc, err = h.srv.Create(ctx, accDomain); err != nil {
		ctx.JSON(http.StatusInternalServerError, responseWithError(err))
	}

	ctx.JSON(http.StatusCreated, acc)

}
