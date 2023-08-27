package api

import (
	"net/http"

	transfer "github.com/FadyGamilM/go-banking-v2/internal/transfer/domain"
	"github.com/gin-gonic/gin"
)

type TransferHandler struct {
	srv transfer.TransferService
}

func NewTransferHandler(s transfer.TransferService) *TransferHandler {
	return &TransferHandler{
		srv: s,
	}
}

func (h TransferHandler) HandleCreateTransfer(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"res": "done"})
}
