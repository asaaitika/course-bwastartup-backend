package handler

import (
	"course-bwastartup-backend/helper"
	"course-bwastartup-backend/transaction"
	"course-bwastartup-backend/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*1. handler: parameter di uri, perlu tangkap parameter ke input struct dan memanggil service
2. service: ambil dari campaign id, lalu memanggil repo
3. repo mencari data transaction campaign*/

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionCampaign(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to get campaign's transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

}
