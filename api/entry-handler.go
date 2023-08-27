package api

import (
	"net/http"

	entry "github.com/FadyGamilM/go-banking-v2/internal/entry/domain"
	"github.com/gin-gonic/gin"
)

type EntryHandler struct {
	srv entry.EntryService
}

func NewEntryHandler(s entry.EntryService) *EntryHandler {
	return &EntryHandler{
		srv: s,
	}
}

func (h *EntryHandler) HandleCreateEntry(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"res": "done"})
}
