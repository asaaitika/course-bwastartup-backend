package handler

import (
	"course-bwastartup-backend/campaign"
	"course-bwastartup-backend/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*1. tangkap parameter di handler
2. handler ke service, di service menentukan apakah repository mana yang dicall
3. repository akses ke db: GetAll (FindAll), GetByUserID(FinfByUserID)*/

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}

// func (h *campaignHandler) GetCampaign(c *gin.Context) {
// 	/*1. butuh repository untuk get campaign by id, karena untuk mendapat satu single data dan pada end pointnya idnya sendiri
// 	2. service untuk inputnya struct => menangkap id di url, memanggil repo
// 	3. handler: mapping id yg url ke struct input yaitu service dan memanggil formatter*/
// }
