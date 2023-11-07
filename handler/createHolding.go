package handler

import (
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
