package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromartinsb/gomerit/schemas"
)

// @BasePath /api/v1

// @Summary Create holding
// @Description Create a new holding
// @Tags Holdings
// @Accept json
// @Produce json
// @Param request body CreateHoldingRequest true "Request body"
// @Success 200 {object} CreateHoldingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /holding [post]
func CreateHoldingHandler(ctx *gin.Context) {
	request := &CreateHoldingRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	holding := schemas.Holding{
		Name: request.Name,
	}

	if err := db.Create(&holding).Error; err != nil {
		logger.Errorf("error creating holding: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-holding", holding)
}

// @BasePath /api/v1

// @Summary Show holding
// @Description Show a holding
// @Tags Holdings
// @Accept json
// @Produce json
// @Param id query string true "Holding identification"
// @Success 200 {object} ShowHoldingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /holding [get]
func ShowHoldingHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	holding := schemas.Holding{}

	if err := db.First(&holding, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "holding not found")
		return
	}

	sendSuccess(ctx, "show-holding", holding)
}

// @BasePath /api/v1

// @Summary List holdings
// @Description List all holdings
// @Tags Holdings
// @Accept json
// @Produce json
// @Success 200 {object} ListHoldingsResponse
// @Failure 500 {object} ErrorResponse
// @Router /holdings [get]
func ListHoldingsHandler(ctx *gin.Context) {
	holdings := []schemas.Holding{}

	if err := db.Find(&holdings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing holdings")
		return
	}

	sendSuccess(ctx, "list-holdings", holdings)
}

// @BasePath /api/v1

// @Summary Update holding
// @Description Update a holding
// @Tags Holdings
// @Accept json
// @Produce json
// @Param id query string true "Holding identification"
// @Param opening body UpdateOpeningRequest true "Holding data to Update"
// @Success 200 {object} UpdateHoldingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /holding [put]
func UpdateHoldingHandler(ctx *gin.Context) {
	request := UpdateHoldingRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	holding := schemas.Holding{}

	if err := db.First(&holding).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("holding with id: %s not found", id))
		return
	}

	// Update holding
	if request.Name != "" {
		holding.Name = request.Name
	}

	// Save holding
	if err := db.Save(&holding).Error; err != nil {
		logger.Errorf("error updating holding: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating holding")
		return
	}
	sendSuccess(ctx, "update-holding", holding)
}

// @BasePath /api/v1

// @Summary Delete holding
// @Description Delete a holding
// @Tags Holdings
// @Accept json
// @Produce json
// @Param id query string true "Holding identification"
// @Success 200 {object} DeleteHoldingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /holding [delete]
func DeleteHoldingHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	holding := schemas.Holding{}

	// Find Holding
	if err := db.First(&holding).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("holding with id: %s not found", id))
		return
	}

	// Delete Holding
	if err := db.Delete(&holding).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting holding with id: %s", id))
		return
	}

	sendSuccess(ctx, "delete-holding", holding)
}
