package api

import (
	"log"
	"net/http"

	"github.com/FadyGamilM/go-banking-v2/common"
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
type GetAccountByIdReqDto struct {
	ID int64 `uri:"id" binding:"required,min=0"`
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

func (h *AccountHandler) HandleGetAccountByID(ctx *gin.Context) {

	var reqUri GetAccountByIdReqDto

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, responseWithError(common.BadRequest))
	}

	log.Println(reqUri.ID)
	// id, err := strconv.ParseInt(accID, 10, 64)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, responseWithError(common.InternalServerError))
	// }

	acc, err := h.srv.GetByID(ctx, int64(reqUri.ID))
	if err != nil {
		if err == common.NotFound {
			ctx.JSON(http.StatusNotFound, responseWithError(err))
		}
		ctx.JSON(http.StatusInternalServerError, responseWithError(err))
	}

	ctx.JSON(http.StatusOK, acc)
}

type PaginationReqDto struct {
	PageID   int32 `form:"page_id"`
	PageSize int32 `form:"page_size"`
}

func (h *AccountHandler) HandleGetAccountsInPages(ctx *gin.Context) {
	var reqDto PaginationReqDto
	if err := ctx.ShouldBindQuery(&reqDto); err != nil {
		ctx.JSON(http.StatusBadRequest, responseWithError(err))
	}

	log.Println("the page number is : ", reqDto.PageID)
	log.Println("the accounts per page are : ", reqDto.PageSize)
	ctx.JSON(http.StatusOK, "[]")
}
