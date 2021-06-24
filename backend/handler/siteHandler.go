package handler

import (
	"assessment/site"
	"strconv"

	"github.com/gin-gonic/gin"
)

type siteHandler struct {
	service site.Service
}

func NewSiteHandler(service site.Service) *siteHandler {
	return &siteHandler{service}

}

// Handler

func (h *siteHandler) GetAllSiteByUser(c *gin.Context) {
	userID := int(c.MustGet("CurrentUser").(int))

	userIDstr := strconv.Itoa(userID)

	site, err := h.service.GetAllSiteByUser(userIDstr)

	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	if userID == 0 {
		c.JSON(401, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, site)
}

func (h *siteHandler) CreateSite(c *gin.Context) {
	var input site.CreateSite

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	userID := c.MustGet("CurrentUser").(int)

	site, err := h.service.CreateSite(userID, input)
	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(201, site)
}

func (h *siteHandler) GetSiteByID(c *gin.Context) {
	siteID := c.Params.ByName("site_id")

	site, err := h.service.GetSiteByID(siteID)

	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, site)
}

func (h *siteHandler) UpdateSite(c *gin.Context) {
	siteID := c.Params.ByName("site_id")

	var input site.UpdateSite

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}

	site, err := h.service.UpdateSite(siteID, input)

	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(200, site)
}

func (h *siteHandler) DeleteSite(c *gin.Context) {
	siteID := c.Params.ByName("site_id")

	message, err := h.service.DeleteSite(siteID)

	if err != nil {
		c.JSON(500, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": message})
}
